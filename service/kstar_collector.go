package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/HavvokLab/true-solar-monitoring/config"
	"github.com/HavvokLab/true-solar-monitoring/constant"
	"github.com/HavvokLab/true-solar-monitoring/inverter"
	"github.com/HavvokLab/true-solar-monitoring/inverter/kstar"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/HavvokLab/true-solar-monitoring/repo"
	"github.com/HavvokLab/true-solar-monitoring/util"
	"go.openly.dev/pointy"
)

type KStarCollectorService interface {
	Run(credential *model.KStarCredential) error
}

type kstarCollectorService struct {
	vendorType     string
	siteRegionRepo repo.SiteRegionMappingRepo
	solarRepo      repo.SolarRepo
	siteRegions    []model.SiteRegionMapping
	elasticConfig  config.ElasticsearchConfig
	logger         logger.Logger
}

func NewKStarCollectorService(
	solarRepo repo.SolarRepo, siteRegionRepo repo.SiteRegionMappingRepo, logger logger.Logger,
) (KStarCollectorService, error) {
	return &kstarCollectorService{
		vendorType:     strings.ToUpper(constant.VENDOR_TYPE_KSTAR),
		siteRegionRepo: siteRegionRepo,
		solarRepo:      solarRepo,
		logger:         logger,
		siteRegions:    make([]model.SiteRegionMapping, 0),
		elasticConfig:  config.GetConfig().Elastic,
	}, nil
}

func (s *kstarCollectorService) Run(credential *model.KStarCredential) error {
	defer func() {
		if r := recover(); r != nil {
			s.logger.Warnf("[%v] - KStarCollectorService.Run(): %v", credential.Username, r)
		}
	}()

	siteRegions, err := s.siteRegionRepo.GetSiteRegionMappings()
	if err != nil {
		s.logger.Errorf("[%v] - Failed to get site region mappings: %v", credential.Username, err)
		return err
	}
	s.siteRegions = siteRegions

	documents := make([]interface{}, 0)
	siteDocuments := make([]model.SiteItem, 0)
	doneCh := make(chan bool)
	errorCh := make(chan error)
	documentCh := make(chan interface{})
	go s.run(credential, documentCh, errorCh, doneCh)

DONE:
	for {
		select {
		case <-doneCh:
			break DONE
		case err := <-errorCh:
			s.logger.Errorf("[%v] - Failed to run KStarCollectorService: %v", credential.Username, err)
			return err
		case document := <-documentCh:
			documents = append(documents, document)
			if plantItemDoc, ok := document.(model.PlantItem); ok {
				siteItemDoc := model.SiteItem{
					Timestamp:   plantItemDoc.Timestamp,
					VendorType:  plantItemDoc.VendorType,
					Area:        plantItemDoc.Area,
					SiteID:      plantItemDoc.SiteID,
					NodeType:    plantItemDoc.NodeType,
					Name:        plantItemDoc.Name,
					Location:    plantItemDoc.Location,
					PlantStatus: plantItemDoc.PlantStatus,
				}
				siteDocuments = append(siteDocuments, siteItemDoc)
			}
		}
	}

	collectorIndex := fmt.Sprintf("%s-%s", s.elasticConfig.SolarIndex, time.Now().Format("2006.01.02"))
	if err := s.solarRepo.BulkIndex(collectorIndex, documents); err != nil {
		s.logger.Errorf("[%v] - KStarCollectorService.Run(): %v", credential.Username, err)
		return err
	}
	s.logger.Infof("[%v] - KStarCollectorService.Run(): %v documents indexed", credential.Username, len(documents))

	if err := s.solarRepo.UpsertSiteStation(siteDocuments); err != nil {
		s.logger.Errorf("[%v] - KStarCollectorService.Run(): %v", credential.Username, err)
		return err
	}
	s.logger.Infof("[%v] - KStarCollectorService.Run(): %v site stations upserted", credential.Username, len(siteDocuments))

	close(doneCh)
	close(errorCh)
	close(documentCh)
	return nil
}

func (s *kstarCollectorService) run(credential *model.KStarCredential, documentCh chan interface{}, errorCh chan error, doneCh chan bool) {
	now := time.Now()
	client, err := kstar.NewKStarClient(&kstar.KStarCredential{
		Username: credential.Username,
		Password: credential.Password,
	})

	if err != nil {
		s.logger.Errorf("[%v] - Failed to create KStar client: %v", credential.Username, err)
		errorCh <- err
		return
	}

	mapPlantIDToDeviceList := make(map[string][]kstar.DeviceItem)
	deviceList, err := client.GetDeviceList()
	if err != nil {
		s.logger.Errorf("[%v] - Failed to get device list: %v", credential.Username, err)
		errorCh <- err
		return
	}

	for _, device := range deviceList {
		mapPlantIDToDeviceList[device.GetPlantID()] = append(mapPlantIDToDeviceList[device.GetPlantID()], *device)
	}

	plantListResp, err := client.GetPlantList()
	if err != nil {
		s.logger.Errorf("[%v] - Failed to get plant list: %v", credential.Username, err)
		errorCh <- err
		return
	}

	stationCount := 1
	stationSize := len(plantListResp.Data)
	for _, station := range plantListResp.Data {
		stationID := station.GetID()
		stationName := station.GetName()
		plantNameInfo, _ := inverter.ParsePlantID(stationName)
		cityName, cityCode, cityArea := inverter.ParseSiteID(s.siteRegions, plantNameInfo.SiteID)

		s.logger.Infof("[%v] - Plant[%v/%v] - %v", credential.Username, stationCount, stationSize, stationName)
		stationCount++

		var plantStatus string
		var currentPower float64
		var totalProduction float64
		var dailyProduction float64
		var monthlyProduction float64
		var yearlyProduction float64
		var location *string

		if station.Latitude != nil && station.Longitude != nil {
			location = pointy.String(fmt.Sprintf("%f,%f", station.GetLatitude(), station.GetLongitude()))
		}

		deviceCount := 1
		deviceSize := len(mapPlantIDToDeviceList[stationID])
		for _, device := range mapPlantIDToDeviceList[stationID] {
			deviceID := device.GetID()

			s.logger.Infof("[%v] - Device[%v/%v] - %v", credential.Username, deviceCount, deviceSize, device.GetName())
			deviceCount++

			realtimeAlarmResp, err := client.GetRealtimeAlarmListOfDevice(deviceID)
			if err != nil {
				s.logger.Errorf("[%v] - Failed to get realtime alarm list of device: %v", credential.Username, err)
				errorCh <- err
				return
			}

			deviceStatus := device.Status
			if len(realtimeAlarmResp.Data) > 0 {
				deviceStatus = pointy.Int(2)
				alarmCount := 1
				alarmSize := len(realtimeAlarmResp.Data)
				for _, alarm := range realtimeAlarmResp.Data {
					s.logger.Infof("[%v] - Alarm[%v/%v] - %v", credential.Username, alarmCount, alarmSize, alarm.GetMessage())
					alarmCount++

					alarmItem := model.AlarmItem{
						Timestamp:    now,
						Month:        now.Format("01"),
						Year:         now.Format("2006"),
						MonthYear:    now.Format("01-2006"),
						VendorType:   s.vendorType,
						DataType:     constant.DATA_TYPE_ALARM,
						Area:         cityArea,
						SiteID:       plantNameInfo.SiteID,
						SiteCityCode: cityCode,
						SiteCityName: cityName,
						NodeType:     plantNameInfo.NodeType,
						ACPhase:      plantNameInfo.ACPhase,
						PlantID:      alarm.PlantID,
						PlantName:    alarm.PlantName,
						Latitude:     station.Latitude,
						Longitude:    station.Longitude,
						Location:     location,
						DeviceID:     alarm.DeviceID,
						DeviceSN:     device.SN,
						DeviceName:   alarm.DeviceName,
						DeviceStatus: pointy.String(kstar.KSTAR_DEVICE_STATUS_ALARM),
						ID:           nil,
						Message:      alarm.Message,
					}

					if alarm.SaveTime != nil {
						if alarmTime, err := time.Parse("2006-01-02 15:04:05", *alarm.SaveTime); err == nil {
							alarmItem.AlarmTime = &alarmTime
						}
					}

					documentCh <- alarmItem
				}
			}

			deviceItem := model.DeviceItem{
				Timestamp:    now,
				Month:        now.Format("01"),
				Year:         now.Format("2006"),
				MonthYear:    now.Format("01-2006"),
				VendorType:   s.vendorType,
				DataType:     constant.DATA_TYPE_DEVICE,
				Area:         cityArea,
				SiteID:       plantNameInfo.SiteID,
				SiteCityCode: cityCode,
				SiteCityName: cityName,
				NodeType:     plantNameInfo.NodeType,
				ACPhase:      plantNameInfo.ACPhase,
				PlantID:      device.PlantID,
				PlantName:    device.PlantName,
				Latitude:     station.Latitude,
				Longitude:    station.Longitude,
				Location:     location,
				ID:           device.ID,
				SN:           device.SN,
				Name:         device.Name,
				DeviceType:   pointy.String(kstar.KSTAR_DEVICE_TYPE_INVERTER),
			}

			deviceInfoResp, err := client.GetRealtimeDeviceData(deviceID)
			if err != nil {
				s.logger.Errorf("[%v] - Failed to get realtime device data: %v", credential.Username, err)
				errorCh <- err
				return
			}

			if deviceInfoResp.Data != nil {
				if deviceInfoResp.Data.SaveTime != nil {
					if saveTime, err := time.Parse("2006-01-02 15:04:05", *deviceInfoResp.Data.SaveTime); err == nil {
						deviceItem.LastUpdateTime = &saveTime
					}
				}

				deviceItem.TotalPowerGeneration = deviceInfoResp.Data.TotalGeneration
				deviceItem.DailyPowerGeneration = deviceInfoResp.Data.DayGeneration
				deviceItem.MonthlyPowerGeneration = deviceInfoResp.Data.MonthGeneration
				deviceItem.YearlyPowerGeneration = deviceInfoResp.Data.YearGeneration

				currentPower += deviceInfoResp.Data.GetPowerInter()
				totalProduction += deviceInfoResp.Data.GetTotalGeneration()
				dailyProduction += deviceInfoResp.Data.GetDayGeneration()
				monthlyProduction += deviceInfoResp.Data.GetMonthGeneration()
				yearlyProduction += deviceInfoResp.Data.GetYearGeneration()
			}

			if deviceStatus != nil {
				switch *deviceStatus {
				case 0: // OFF
					deviceItem.Status = pointy.String(kstar.KSTAR_DEVICE_STATUS_OFF)
					if plantStatus != kstar.KSTAR_DEVICE_STATUS_ALARM {
						plantStatus = kstar.KSTAR_DEVICE_STATUS_OFF
					}
				case 1: // ON
					deviceItem.Status = pointy.String(kstar.KSTAR_DEVICE_STATUS_ON)
					if plantStatus != kstar.KSTAR_DEVICE_STATUS_OFF && plantStatus != kstar.KSTAR_DEVICE_STATUS_ALARM {
						plantStatus = kstar.KSTAR_DEVICE_STATUS_ON
					}
				case 2: // ALARM
					deviceItem.Status = pointy.String(kstar.KSTAR_DEVICE_STATUS_ALARM)
					plantStatus = kstar.KSTAR_DEVICE_STATUS_ALARM
				default:
				}
			}

			documentCh <- deviceItem
		}

		if util.EmptyString(plantStatus) {
			plantStatus = kstar.KSTAR_DEVICE_STATUS_OFF
		}

		plantItem := model.PlantItem{
			Timestamp:         now,
			Month:             now.Format("01"),
			Year:              now.Format("2006"),
			MonthYear:         now.Format("01-2006"),
			VendorType:        s.vendorType,
			DataType:          constant.DATA_TYPE_PLANT,
			Area:              cityArea,
			SiteID:            plantNameInfo.SiteID,
			SiteCityCode:      cityCode,
			SiteCityName:      cityName,
			NodeType:          plantNameInfo.NodeType,
			ACPhase:           plantNameInfo.ACPhase,
			ID:                station.ID,
			Name:              station.Name,
			Latitude:          station.Latitude,
			Longitude:         station.Longitude,
			Location:          location,
			LocationAddress:   station.Address,
			InstalledCapacity: station.InstalledCapacity,
			TotalCO2:          nil,
			MonthlyCO2:        nil,
			TotalSavingPrice:  pointy.Float64(totalProduction * station.GetElectricPrice()),
			Currency:          station.ElectricUnit,
			CurrentPower:      pointy.Float64(currentPower / 1000), // W to kW
			TotalProduction:   pointy.Float64(totalProduction),
			DailyProduction:   pointy.Float64(dailyProduction),
			MonthlyProduction: pointy.Float64(monthlyProduction),
			YearlyProduction:  pointy.Float64(yearlyProduction),
			PlantStatus:       pointy.String(plantStatus),
		}

		if station.CreatedTime != nil {
			if createdTime, err := time.Parse("2006-01-02 15:04:05", *station.CreatedTime); err == nil {
				plantItem.CreatedDate = &createdTime
			}
		}

		documentCh <- plantItem
	}

	doneCh <- true
}

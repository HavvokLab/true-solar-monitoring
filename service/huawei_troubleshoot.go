package service

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/HavvokLab/true-solar-monitoring/config"
	"github.com/HavvokLab/true-solar-monitoring/constant"
	"github.com/HavvokLab/true-solar-monitoring/inverter"
	"github.com/HavvokLab/true-solar-monitoring/inverter/huawei"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/HavvokLab/true-solar-monitoring/repo"
	"github.com/HavvokLab/true-solar-monitoring/util"
	"github.com/schollz/progressbar/v3"
	"go.openly.dev/pointy"
)

type HuaweiTroubleShootService interface {
	Run(*model.HuaweiCredential) error
}

type huaweiTroubleShootService struct {
	vendorType     string
	siteRegionRepo repo.SiteRegionMappingRepo
	siteRegions    []model.SiteRegionMapping
	solarRepo      repo.SolarRepo
	elasticConfig  config.ElasticsearchConfig
	logger         logger.Logger
}

func NewHuaweiTroubleShootService(solarRepo repo.SolarRepo, siteRegionRepo repo.SiteRegionMappingRepo, logger logger.Logger) HuaweiTroubleShootService {
	return &huaweiTroubleShootService{
		vendorType:     strings.ToUpper(constant.VENDOR_TYPE_HUAWEI),
		siteRegionRepo: siteRegionRepo,
		solarRepo:      solarRepo,
		logger:         logger,
		siteRegions:    make([]model.SiteRegionMapping, 0),
		elasticConfig:  config.GetConfig().Elastic,
	}
}

func (s *huaweiTroubleShootService) Run(credential *model.HuaweiCredential) error {
	defer func() {
		if r := recover(); r != nil {
			s.logger.Warnf("[%v] - HuaweiTroubleShootService.Run(): %v", credential.Username, r)
		}
	}()

	siteRegions, err := s.siteRegionRepo.GetSiteRegionMappings()
	if err != nil {
		s.logger.Errorf("[%v] - HuaweiTroubleShootService.Run(): %v", credential.Username, err)
		return err
	}
	s.siteRegions = siteRegions

	documents := make([]interface{}, 0)
	siteDocuments := make([]model.SiteItem, 0)
	doneCh := make(chan bool)
	errorCh := make(chan error)
	documentCh := make(chan interface{})
	date := time.Date(2023, time.November, 20, 0, 0, 0, 0, time.Local)
	go s.run(credential, date, documentCh, doneCh, errorCh)

DONE:
	for {
		select {
		case <-doneCh:
			break DONE
		case err := <-errorCh:
			s.logger.Errorf("[%v] - HuaweiTroubleShootService.Run(): %v", credential.Username, err)
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
					Owner:       credential.Owner,
				}
				siteDocuments = append(siteDocuments, siteItemDoc)
			}
		}
	}

	// collectorIndex := fmt.Sprintf("%s-%s", s.elasticConfig.SolarIndex, time.Now().Format("2006.01.02"))
	// if err := s.solarRepo.BulkIndex(collectorIndex, documents); err != nil {
	// 	s.logger.Errorf("[%v] - HuaweiTroubleShootService.Run(): %v", credential.Username, err)
	// 	return err
	// }
	util.PrintJSON(map[string]any{"result": documents})
	s.logger.Infof("[%v] - HuaweiTroubleShootService.Run(): %v documents indexed", credential.Username, len(documents))

	// if err := s.solarRepo.UpsertSiteStation(siteDocuments); err != nil {
	// 	s.logger.Errorf("[%v] - HuaweiTroubleShootService.Run(): %v", credential.Username, err)
	// 	return err
	// }
	util.PrintJSON(map[string]any{"result": siteDocuments})
	s.logger.Infof("[%v] - HuaweiTroubleShootService.Run(): %v site stations upserted", credential.Username, len(siteDocuments))

	close(doneCh)
	close(errorCh)
	close(documentCh)
	return nil
}

// monthly co2
// monthly production
// daily production

func (s *huaweiTroubleShootService) run(credential *model.HuaweiCredential, date time.Time, documentCh chan interface{}, doneCh chan bool, errorCh chan error) {
	beginTime := time.Date(date.Year(), date.Month(), date.Day(), 6, 0, 0, 0, date.Location()).UnixNano() / 1e6
	collectTime := date.UnixNano() / 1e6

	client, err := huawei.NewHuaweiClient(&huawei.HuaweiCredential{
		Username: credential.Username,
		Password: credential.Password,
	})

	if err != nil {
		errorCh <- err
		return
	}

	plantListResp, err := client.GetPlantList()
	if err != nil {
		s.logger.Error(err)
		errorCh <- err
		return
	}
	s.logger.Infof("[%v] - HuaweiCollectorService.run(): %v plants found", credential.Username, len(plantListResp.Data))
	s.logger.Infof("[%v] - HuaweiCollectorService.run(): start preparing plant data", credential.Username)

	var stationCodeList []string
	var stationCodeListString []string
	for _, station := range plantListResp.Data {
		if len(stationCodeList) == 100 {
			stationCodeListString = append(stationCodeListString, strings.Join(stationCodeList, ","))
			stationCodeList = make([]string, 0)
		}

		if station.Code != nil {
			stationCodeList = append(stationCodeList, station.GetCode())
		}
	}
	stationCodeListString = append(stationCodeListString, strings.Join(stationCodeList, ","))
	fmt.Println(stationCodeListString)

	var inverterList []huawei.DeviceItem
	// mapPlantCodeToRealtimeData := make(map[string]huawei.RealtimePlantData)
	mapPlantCodeToDailyData := make(map[string]huawei.HistoricalPlantData)
	mapPlantCodeToMonthlyData := make(map[string]huawei.HistoricalPlantData)
	mapPlantCodeToYearlyPower := make(map[string]float64)
	mapPlantCodeToTotalPower := make(map[string]float64)
	mapPlantCodeToTotalCO2 := make(map[string]float64)
	mapPlantCodeToDevice := make(map[string][]huawei.DeviceItem)
	mapDeviceSNToAlarm := make(map[string][]huawei.DeviceAlarmItem)
	bar := progressbar.Default(int64(len(stationCodeListString)))
	for _, stationCode := range stationCodeListString {
		// realtimePlantDataResp, err := client.GetRealtimePlantData(stationCode)
		// if err != nil {
		// 	s.logger.Error(err)
		// 	errorCh <- err
		// 	return
		// }

		// for _, item := range realtimePlantDataResp.Data {
		// 	if item.Code != nil {
		// 		mapPlantCodeToRealtimeData[item.GetCode()] = *item
		// 	}
		// }

		dailyPlantDataResp, err := client.GetDailyPlantData(stationCode, collectTime)
		if err != nil {
			s.logger.Error(err)
			errorCh <- err
			return
		}

		for _, item := range dailyPlantDataResp.Data {
			if item.Code != nil {
				if date.Format("2006-01-02") == time.Unix(item.GetCollectTime()/1e3, 0).Format("2006-01-02") {
					mapPlantCodeToDailyData[item.GetCode()] = *item
				}
			}
		}

		monthlyPlantDataResp, err := client.GetMonthlyPlantData(stationCode, collectTime)
		if err != nil {
			s.logger.Error(err)
			errorCh <- err
			return
		}

		for _, item := range monthlyPlantDataResp.Data {
			if item.Code != nil {
				if date.Format("2006-01") == time.Unix(item.GetCollectTime()/1e3, 0).Format("2006-01") {
					mapPlantCodeToMonthlyData[item.GetCode()] = *item
				}
				mapPlantCodeToYearlyPower[item.GetCode()] = mapPlantCodeToYearlyPower[item.GetCode()] + item.DataItemMap.GetInverterPower()
			}
		}

		yearlyPlantDataResp, err := client.GetYearlyPlantData(stationCode, collectTime)
		if err != nil {
			s.logger.Error(err)
			errorCh <- err
			return
		}

		for _, item := range yearlyPlantDataResp.Data {
			if item.Code != nil {
				mapPlantCodeToTotalPower[item.GetCode()] = mapPlantCodeToTotalPower[item.GetCode()] + item.DataItemMap.GetInverterPower()
				mapPlantCodeToTotalCO2[item.GetCode()] = mapPlantCodeToTotalCO2[item.GetCode()] + item.DataItemMap.GetReductionTotalCO2()
			}
		}

		deviceListResp, err := client.GetDeviceList(stationCode)
		if err != nil {
			s.logger.Error(err)
			errorCh <- err
			return
		}

		for _, item := range deviceListResp.Data {
			if item.PlantCode != nil {
				mapPlantCodeToDevice[item.GetPlantCode()] = append(mapPlantCodeToDevice[item.GetPlantCode()], *item)
			}

			if item.GetTypeID() == 1 {
				inverterList = append(inverterList, *item)
			}
		}

		deviceAlarmListResp, err := client.GetDeviceAlarmList(stationCode, beginTime, collectTime)
		if err != nil {
			s.logger.Error(err)
			errorCh <- err
			return
		}

		for _, item := range deviceAlarmListResp.Data {
			doubleAlarm := false

			if item.DeviceSN != nil {
				for i, alarm := range mapDeviceSNToAlarm[item.GetDeviceSN()] {
					if alarm.GetAlarmName() == item.GetAlarmName() {
						doubleAlarm = true

						if alarm.GetRaiseTime() < item.GetRaiseTime() {
							mapDeviceSNToAlarm[item.GetDeviceSN()][i] = *item
							break
						}
					}
				}

				if !doubleAlarm {
					mapDeviceSNToAlarm[item.GetDeviceSN()] = append(mapDeviceSNToAlarm[item.GetDeviceSN()], *item)
				}
			}
		}

		bar.Add(1)
	}

	s.logger.Infof("[%v] - HuaweiCollectorService.run(): start preparing device data", credential.Username)
	var inverterIDList []string
	var inverterIDListString []string
	for _, device := range inverterList {
		if len(inverterIDList) == 100 {
			inverterIDListString = append(inverterIDListString, strings.Join(inverterIDList, ","))
			inverterIDList = make([]string, 0)
		}

		if device.ID != nil {
			inverterIDList = append(inverterIDList, strconv.Itoa(device.GetID()))
		}
	}
	inverterIDListString = append(inverterIDListString, strings.Join(inverterIDList, ","))

	// mapDeviceToRealtimeData := make(map[int]huawei.RealtimeDeviceData)
	mapDeviceToDailyData := make(map[int]huawei.HistoricalDeviceData)
	mapDeviceToMonthlyData := make(map[int]huawei.HistoricalDeviceData)
	mapDeviceToYearlyPower := make(map[int]float64)
	for _, deviceID := range inverterIDListString {
		// realtimeDeviceDataResp, err := client.GetRealtimeDeviceData(deviceID, "1")
		// if err != nil {
		// 	s.logger.Error(err)
		// 	errorCh <- err
		// 	return
		// }

		// for _, item := range realtimeDeviceDataResp.Data {
		// 	if item.ID != nil {
		// 		mapDeviceToRealtimeData[item.GetID()] = *item
		// 	}
		// }

		dailyDeviceDataResp, err := client.GetDailyDeviceData(deviceID, "1", collectTime)
		if err != nil {
			s.logger.Error(err)
			errorCh <- err
			return
		}

		for _, item := range dailyDeviceDataResp.Data {
			if item.ID != nil {
				if date.Format("2006-01-02") == time.Unix(item.GetCollectTime()/1e3, 0).Format("2006-01-02") {
					deviceID := item.GetID()
					switch deviceID := deviceID.(type) {
					case float64:
						parsedDeviceID := int(deviceID)
						mapDeviceToDailyData[parsedDeviceID] = *item
					default:
					}
				}
			}
		}

		monthlyDeviceDataResp, err := client.GetMonthlyDeviceData(deviceID, "1", collectTime)
		if err != nil {
			s.logger.Error(err)
			errorCh <- err
			return
		}

		for _, item := range monthlyDeviceDataResp.Data {
			if item.ID != nil {
				deviceID := item.GetID()
				switch deviceID := deviceID.(type) {
				case float64:
					parsedDeviceID := int(deviceID)
					mapDeviceToYearlyPower[parsedDeviceID] = mapDeviceToYearlyPower[parsedDeviceID] + item.DataItemMap.GetProductPower()
					if date.Format("2006-01") == time.Unix(item.GetCollectTime()/1e3, 0).Format("2006-01") {
						mapDeviceToMonthlyData[parsedDeviceID] = *item
					}
				default:
				}
			}
		}
	}

	s.logger.Infof("[%v] - HuaweiCollectorService.run(): start preparing documents", credential.Username)
	plantCount := 1
	plantSize := len(plantListResp.Data)
	for _, station := range plantListResp.Data {
		stationCode := station.GetCode()
		stationName := station.GetName()
		plantNameInfo, _ := inverter.ParsePlantID(stationName)
		cityName, cityCode, cityArea := inverter.ParseSiteID(s.siteRegions, plantNameInfo.SiteID)

		s.logger.Infof("[%v] - HuaweiCollectorService.run(): preparing document %v/%v", credential.Username, plantCount, plantSize)
		plantCount++

		var latitude *float64
		var longitude *float64
		var location *string
		currentPower := 0.0
		// plantStatus := huawei.HuaweiMapDeviceStatus[mapPlantCodeToRealtimeData[stationCode].DataItemMap.GetRealHealthState()]
		var plantStatus string = "UNKNOWN"

		deviceCount := 1
		deviceSize := len(mapPlantCodeToDevice[stationCode])
		for _, device := range mapPlantCodeToDevice[stationCode] {
			deviceID := device.GetID()
			deviceSN := device.GetSN()

			s.logger.Infof("[%v] - HuaweiCollectorService.run(): plant: %v, device: %v, preparing document %v/%v - %v/%v", credential.Username, stationCode, deviceID, plantCount, plantSize, deviceCount, deviceSize)
			deviceCount++

			if device.Latitude != nil && device.Longitude != nil {
				location = pointy.String(fmt.Sprintf("%f,%f", device.GetLatitude(), device.GetLongitude()))
			}

			var deviceStatus *int
			// if mapDeviceToRealtimeData[deviceID].DataItemMap != nil {
			// 	if mapDeviceToRealtimeData[deviceID].DataItemMap.Status != nil {
			// 		deviceStatus = mapDeviceToRealtimeData[deviceID].DataItemMap.Status
			// 	}
			// }

			if len(mapDeviceSNToAlarm[deviceSN]) > 0 {
				deviceStatus = pointy.Int(2)

				for _, deviceAlarm := range mapDeviceSNToAlarm[deviceSN] {
					alarmItem := model.AlarmItem{
						Timestamp:    date,
						Month:        date.Format("01"),
						Year:         date.Format("2006"),
						MonthYear:    date.Format("01-2006"),
						VendorType:   s.vendorType,
						DataType:     constant.DATA_TYPE_ALARM,
						Area:         cityArea,
						SiteID:       plantNameInfo.SiteID,
						SiteCityCode: cityCode,
						SiteCityName: cityName,
						NodeType:     plantNameInfo.NodeType,
						ACPhase:      plantNameInfo.ACPhase,
						PlantID:      station.Code,
						PlantName:    station.Name,
						Latitude:     latitude,
						Longitude:    longitude,
						Location:     location,
						DeviceID:     pointy.String(strconv.Itoa(deviceID)),
						DeviceSN:     deviceAlarm.DeviceSN,
						DeviceName:   deviceAlarm.DeviceName,
						DeviceStatus: pointy.String(huawei.HUAWEI_STATUS_ALARM),
						ID:           pointy.String(strconv.Itoa(deviceAlarm.GetAlarmID())),
						Message:      deviceAlarm.AlarmName,
						Owner:        credential.Owner,
					}

					if deviceAlarm.RaiseTime != nil {
						alarmTime := time.Unix(deviceAlarm.GetRaiseTime()/1e3, 0)
						alarmItem.AlarmTime = &alarmTime
					}

					documentCh <- alarmItem
				}
			}

			deviceItem := model.DeviceItem{
				Timestamp:      date,
				Month:          date.Format("01"),
				Year:           date.Format("2006"),
				MonthYear:      date.Format("01-2006"),
				VendorType:     s.vendorType,
				DataType:       constant.DATA_TYPE_DEVICE,
				Area:           cityArea,
				SiteID:         plantNameInfo.SiteID,
				SiteCityCode:   cityCode,
				SiteCityName:   cityName,
				NodeType:       plantNameInfo.NodeType,
				ACPhase:        plantNameInfo.ACPhase,
				PlantID:        station.Code,
				PlantName:      station.Name,
				Latitude:       device.Latitude,
				Longitude:      device.Longitude,
				Location:       location,
				ID:             pointy.String(strconv.Itoa(deviceID)),
				SN:             device.SN,
				Name:           device.Name,
				LastUpdateTime: nil,
				Owner:          credential.Owner,
			}

			if deviceStatus != nil {
				switch *deviceStatus {
				case 0:
					deviceItem.Status = pointy.String(huawei.HUAWEI_STATUS_OFF)
					if plantStatus != huawei.HUAWEI_STATUS_ALARM {
						plantStatus = huawei.HUAWEI_STATUS_OFF
					}
				case 1:
					deviceItem.Status = pointy.String(huawei.HUAWEI_STATUS_ON)
					if plantStatus != huawei.HUAWEI_STATUS_OFF && plantStatus != huawei.HUAWEI_STATUS_ALARM {
						plantStatus = huawei.HUAWEI_STATUS_ON
					}
				case 2:
					deviceItem.Status = pointy.String(huawei.HUAWEI_STATUS_ALARM)
					plantStatus = huawei.HUAWEI_STATUS_ALARM
				}
			}

			if device.TypeID != nil {
				deviceItem.DeviceType = pointy.String(huawei.HuaweiMapDeviceType[*device.TypeID])

				if device.GetTypeID() == 1 {
					// if mapDeviceToRealtimeData[deviceID].DataItemMap != nil {
					// 	deviceItem.TotalPowerGeneration = pointy.Float64(mapDeviceToRealtimeData[deviceID].DataItemMap.GetTotalEnergy())
					// }

					if mapDeviceToDailyData[deviceID].DataItemMap != nil {
						deviceItem.DailyPowerGeneration = pointy.Float64(mapDeviceToDailyData[deviceID].DataItemMap.GetProductPower())
					}

					if mapDeviceToMonthlyData[deviceID].DataItemMap != nil {
						deviceItem.MonthlyPowerGeneration = pointy.Float64(mapDeviceToMonthlyData[deviceID].DataItemMap.GetProductPower())
					}

					deviceItem.YearlyPowerGeneration = pointy.Float64(mapDeviceToYearlyPower[deviceID])

					// currentPower += mapDeviceToRealtimeData[deviceID].DataItemMap.GetActivePower()
					latitude = deviceItem.Latitude
					longitude = deviceItem.Longitude
				}
			}

			documentCh <- deviceItem
		}

		var dailyProduction float64
		if mapPlantCodeToDailyData[stationCode].DataItemMap != nil {
			dailyProduction = mapPlantCodeToDailyData[stationCode].DataItemMap.GetInverterPower()
		}

		var monthlyProduction float64
		if mapPlantCodeToMonthlyData[stationCode].DataItemMap != nil {
			monthlyProduction = mapPlantCodeToMonthlyData[stationCode].DataItemMap.GetInverterPower()
		}

		var monthlyCO2 float64 = 0.0
		if data, ok := mapPlantCodeToMonthlyData[stationCode]; ok {
			monthlyCO2 = data.DataItemMap.GetReductionTotalCO2() * 1000
		}

		plantItem := model.PlantItem{
			Timestamp:         date,
			Month:             date.Format("01"),
			Year:              date.Format("2006"),
			MonthYear:         date.Format("01-2006"),
			VendorType:        s.vendorType,
			DataType:          constant.DATA_TYPE_PLANT,
			Area:              cityArea,
			SiteID:            plantNameInfo.SiteID,
			SiteCityCode:      cityCode,
			SiteCityName:      cityName,
			NodeType:          plantNameInfo.NodeType,
			ACPhase:           plantNameInfo.ACPhase,
			ID:                station.Code,
			Name:              station.Name,
			Latitude:          latitude,
			Longitude:         longitude,
			Location:          location,
			LocationAddress:   station.Address,
			CreatedDate:       nil,
			InstalledCapacity: pointy.Float64(station.GetCapacity() * 1000),
			TotalCO2:          pointy.Float64(mapPlantCodeToTotalCO2[stationCode] * 1000),
			MonthlyCO2:        &monthlyCO2,
			TotalSavingPrice:  nil,
			Currency:          pointy.String(huawei.HUAWEI_CURRENCY_USD),
			CurrentPower:      pointy.Float64(currentPower),
			DailyProduction:   &dailyProduction,
			MonthlyProduction: &monthlyProduction,
			YearlyProduction:  pointy.Float64(mapPlantCodeToYearlyPower[stationCode]),
			PlantStatus:       pointy.String(plantStatus),
			Owner:             credential.Owner,
			TotalProduction:   pointy.Float64(mapPlantCodeToTotalPower[stationCode]),
		}

		// plantItem.TotalProduction = pointy.Float64(mapPlantCodeToTotalPower[stationCode])
		// plantItem.TotalProduction = mapPlantCodeToRealtimeData[stationCode].DataItemMap.TotalPower
		// if pointy.Float64Value(plantItem.TotalProduction, 0.0) < pointy.Float64Value(plantItem.YearlyProduction, 0.0) {
		// 	plantItem.TotalProduction = pointy.Float64(mapPlantCodeToTotalPower[stationCode])
		// }

		documentCh <- plantItem
	}

	doneCh <- true

}

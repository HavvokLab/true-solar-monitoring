package service

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/HavvokLab/true-solar-monitoring/config"
	"github.com/HavvokLab/true-solar-monitoring/constant"
	"github.com/HavvokLab/true-solar-monitoring/inverter"
	"github.com/HavvokLab/true-solar-monitoring/inverter/huawei/v2"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/HavvokLab/true-solar-monitoring/repo"
	"go.openly.dev/pointy"
)

type HuaweiCollectorV3Service interface {
	Run(*model.HuaweiCredential) error
}

type huaweiCollectorV3Service struct {
	vendorType         string
	huaweiAltervimRepo repo.HuaweiAltervimRepo
	siteRegionRepo     repo.SiteRegionMappingRepo
	siteRegions        []model.SiteRegionMapping
	solarRepo          repo.SolarRepo
	elasticConfig      config.ElasticsearchConfig
	logger             logger.Logger
}

func NewHuaweiCollectorV3Service(
	huaweiAltervimRepo repo.HuaweiAltervimRepo,
	solarRepo repo.SolarRepo,
	siteRegionRepo repo.SiteRegionMappingRepo,
	logger logger.Logger,
) *huaweiCollectorV3Service {
	return &huaweiCollectorV3Service{
		vendorType:         strings.ToUpper(constant.VENDOR_TYPE_HUAWEI),
		huaweiAltervimRepo: huaweiAltervimRepo,
		siteRegionRepo:     siteRegionRepo,
		solarRepo:          solarRepo,
		logger:             logger,
		siteRegions:        make([]model.SiteRegionMapping, 0),
		elasticConfig:      config.GetConfig().Elastic,
	}
}

func (s *huaweiCollectorV3Service) Run(credential *model.HuaweiCredential) error {
	defer func() {
		if r := recover(); r != nil {
			s.logger.Warnf("[%v] - HuaweiCollectorService.Run(): %v", credential.Username, r)
		}
	}()

	s.PreparePlantAndDevice(credential)
	siteRegions, err := s.siteRegionRepo.GetSiteRegionMappings()
	if err != nil {
		s.logger.Errorf("[%v] - HuaweiCollectorService.Run(): %v", credential.Username, err)
		return err
	}
	s.siteRegions = siteRegions

	documents := make([]interface{}, 0)
	siteDocuments := make([]model.SiteItem, 0)
	doneCh := make(chan bool)
	errorCh := make(chan error)
	documentCh := make(chan interface{})
	go s.run(credential, documentCh, doneCh, errorCh)

DONE:
	for {
		select {
		case <-doneCh:
			break DONE
		case err := <-errorCh:
			s.logger.Errorf("[%v] - HuaweiCollectorService.Run(): %v", credential.Username, err)
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

	collectorIndex := fmt.Sprintf("%s-%s", s.elasticConfig.SolarIndex, time.Now().Format("2006.01.02"))
	if err := s.solarRepo.BulkIndex(collectorIndex, documents); err != nil {
		s.logger.Errorf("[%v] - HuaweiCollectorService.Run(): %v", credential.Username, err)
		return err
	}
	s.logger.Infof("[%v] - HuaweiCollectorService.Run(): %v documents indexed", credential.Username, len(documents))

	if err := s.solarRepo.UpsertSiteStation(siteDocuments); err != nil {
		s.logger.Errorf("[%v] - HuaweiCollectorService.Run(): %v", credential.Username, err)
		return err
	}
	s.logger.Infof("[%v] - HuaweiCollectorService.Run(): %v site stations upserted", credential.Username, len(siteDocuments))

	close(doneCh)
	close(errorCh)
	close(documentCh)
	return nil
}

func (s *huaweiCollectorV3Service) run(
	credential *model.HuaweiCredential,
	documentCh chan interface{},
	doneCh chan bool,
	errorCh chan error,
) {
	now := time.Now()
	beginTime := time.Date(now.Year(), now.Month(), now.Day(), 6, 0, 0, 0, now.Location()).UnixNano() / 1e6
	collectTime := now.UnixNano() / 1e6

	client, err := huawei.NewHuaweiClient(&huawei.HuaweiCredential{
		Username: credential.Username,
		Password: credential.Password,
	})

	if err != nil {
		errorCh <- err
		return
	}

	plantList, err := s.huaweiAltervimRepo.GetPlants()
	if err != nil {
		s.logger.Error(err)
		errorCh <- err
		return
	}
	s.logger.Infof("[%v] - HuaweiCollectorService.run(): collection %v plants...", credential.Username, len(plantList))

	var stationCodeList []string
	var stationCodeListString []string
	for _, station := range plantList {
		if len(stationCodeList) == 100 {
			stationCodeListString = append(stationCodeListString, strings.Join(stationCodeList, ","))
			stationCodeList = make([]string, 0)
		}

		stationCodeList = append(stationCodeList, station.Code)
	}
	stationCodeListString = append(stationCodeListString, strings.Join(stationCodeList, ","))

	var inverterList []model.HuaweiAltervimDevice
	mapPlantCodeToRealtimeData := make(map[string]huawei.RealtimePlantData)
	mapPlantCodeToDailyData := make(map[string]huawei.HistoricalPlantData)
	mapPlantCodeToMonthlyData := make(map[string]huawei.HistoricalPlantData)
	mapPlantCodeToYearlyPower := make(map[string]float64)
	mapPlantCodeToTotalPower := make(map[string]float64)
	mapPlantCodeToTotalCO2 := make(map[string]float64)
	mapPlantCodeToDevice := make(map[string][]model.HuaweiAltervimDevice)
	mapDeviceSNToAlarm := make(map[string][]huawei.DeviceAlarmItem)
	for _, stationCode := range stationCodeListString {
		realtimePlantDataResp, err := client.GetRealtimePlantData(stationCode)
		if err != nil {
			s.logger.Error(err)
			errorCh <- err
			return
		}

		for _, item := range realtimePlantDataResp.Data {
			if item.Code != nil {
				mapPlantCodeToRealtimeData[item.GetCode()] = *item
			}
		}

		dailyPlantDataResp, err := client.GetDailyPlantData(stationCode, collectTime)
		if err != nil {
			s.logger.Error(err)
			errorCh <- err
			return
		}

		for _, item := range dailyPlantDataResp.Data {
			if item.Code != nil {
				if now.Format("2006-01-02") == time.Unix(item.GetCollectTime()/1e3, 0).Format("2006-01-02") {
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
				if now.Format("2006-01") == time.Unix(item.GetCollectTime()/1e3, 0).Format("2006-01") {
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

		deviceList, err := s.huaweiAltervimRepo.GetDeviceByPlantCode(stationCode)
		if err != nil {
			s.logger.Error(err)
			errorCh <- err
			return
		}

		for _, item := range deviceList {
			if item.PlantCode != nil {
				mapPlantCodeToDevice[*item.PlantCode] = append(mapPlantCodeToDevice[*item.PlantCode], item)
			}

			if *item.TypeID == 1 {
				inverterList = append(inverterList, item)
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
	}

	s.logger.Infof("[%v] - HuaweiCollectorService.run(): start preparing device data", credential.Username)

	var inverterIDList []string
	var inverterIDListString []string
	for _, device := range inverterList {
		if len(inverterIDList) == 100 {
			inverterIDListString = append(inverterIDListString, strings.Join(inverterIDList, ","))
			inverterIDList = make([]string, 0)
		}

		inverterIDList = append(inverterIDList, strconv.Itoa(device.ID))
	}
	inverterIDListString = append(inverterIDListString, strings.Join(inverterIDList, ","))

	mapDeviceToRealtimeData := make(map[int]huawei.RealtimeDeviceData)
	mapDeviceToDailyData := make(map[int]huawei.HistoricalDeviceData)
	mapDeviceToMonthlyData := make(map[int]huawei.HistoricalDeviceData)
	mapDeviceToYearlyPower := make(map[int]float64)
	for _, deviceID := range inverterIDListString {
		realtimeDeviceDataResp, err := client.GetRealtimeDeviceData(deviceID, "1")
		if err != nil {
			s.logger.Error(err)
			errorCh <- err
			return
		}

		for _, item := range realtimeDeviceDataResp.Data {
			if item.ID != nil {
				mapDeviceToRealtimeData[item.GetID()] = *item
			}
		}

		dailyDeviceDataResp, err := client.GetDailyDeviceData(deviceID, "1", collectTime)
		if err != nil {
			s.logger.Error(err)
			errorCh <- err
			return
		}

		for _, item := range dailyDeviceDataResp.Data {
			if item.ID != nil {
				if now.Format("2006-01-02") == time.Unix(item.GetCollectTime()/1e3, 0).Format("2006-01-02") {
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
					if now.Format("2006-01") == time.Unix(item.GetCollectTime()/1e3, 0).Format("2006-01") {
						mapDeviceToMonthlyData[parsedDeviceID] = *item
					}
				default:
				}
			}
		}
	}

	s.logger.Infof("[%v] - HuaweiCollectorService.run(): start preparing documents", credential.Username)
	plantCount := 1
	plantSize := len(plantList)
	for _, station := range plantList {
		stationCode := station.Code
		stationName := station.GetPlantName()
		plantNameInfo, _ := inverter.ParsePlantID(stationName)
		cityName, cityCode, cityArea := inverter.ParseSiteID(s.siteRegions, plantNameInfo.SiteID)

		s.logger.Infof("[%v] - HuaweiCollectorService.run(): preparing document %v/%v", credential.Username, plantCount, plantSize)
		plantCount++

		var latitude *float64
		var longitude *float64
		var location *string
		currentPower := 0.0
		plantStatus := huawei.HuaweiMapDeviceStatus[mapPlantCodeToRealtimeData[stationCode].DataItemMap.GetRealHealthState()]

		deviceCount := 1
		deviceSize := len(mapPlantCodeToDevice[stationCode])
		for _, device := range mapPlantCodeToDevice[stationCode] {
			deviceID := device.ID
			deviceSN := device.GetSN()

			s.logger.Infof("[%v] - HuaweiCollectorService.run(): plant: %v, device: %v, preparing document %v/%v - %v/%v", credential.Username, stationCode, deviceID, plantCount, plantSize, deviceCount, deviceSize)
			deviceCount++

			if device.Latitude != nil && device.Longitude != nil {
				location = pointy.String(fmt.Sprintf("%f,%f", device.GetLatitude(), device.GetLongitude()))
			}

			var deviceStatus *int
			if mapDeviceToRealtimeData[deviceID].DataItemMap != nil {
				if mapDeviceToRealtimeData[deviceID].DataItemMap.Status != nil {
					deviceStatus = mapDeviceToRealtimeData[deviceID].DataItemMap.Status
				}
			}

			if len(mapDeviceSNToAlarm[deviceSN]) > 0 {
				deviceStatus = pointy.Int(2)

				for _, deviceAlarm := range mapDeviceSNToAlarm[deviceSN] {
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
						PlantID:      &station.Code,
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
				Timestamp:      now,
				Month:          now.Format("01"),
				Year:           now.Format("2006"),
				MonthYear:      now.Format("01-2006"),
				VendorType:     s.vendorType,
				DataType:       constant.DATA_TYPE_DEVICE,
				Area:           cityArea,
				SiteID:         plantNameInfo.SiteID,
				SiteCityCode:   cityCode,
				SiteCityName:   cityName,
				NodeType:       plantNameInfo.NodeType,
				ACPhase:        plantNameInfo.ACPhase,
				PlantID:        &station.Code,
				PlantName:      station.Name,
				Latitude:       device.Latitude,
				Longitude:      device.Longitude,
				Location:       location,
				ID:             pointy.String(strconv.Itoa(deviceID)),
				SN:             device.SerialNumber,
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
					if mapDeviceToRealtimeData[deviceID].DataItemMap != nil {
						deviceItem.TotalPowerGeneration = pointy.Float64(mapDeviceToRealtimeData[deviceID].DataItemMap.GetTotalEnergy())
					}

					if mapDeviceToDailyData[deviceID].DataItemMap != nil {
						deviceItem.DailyPowerGeneration = pointy.Float64(mapDeviceToDailyData[deviceID].DataItemMap.GetProductPower())
					}

					if mapDeviceToMonthlyData[deviceID].DataItemMap != nil {
						deviceItem.MonthlyPowerGeneration = pointy.Float64(mapDeviceToMonthlyData[deviceID].DataItemMap.GetProductPower())
					}

					deviceItem.YearlyPowerGeneration = pointy.Float64(mapDeviceToYearlyPower[deviceID])

					currentPower += mapDeviceToRealtimeData[deviceID].DataItemMap.GetActivePower()
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
			ID:                &station.Code,
			Name:              station.Name,
			Latitude:          latitude,
			Longitude:         longitude,
			Location:          location,
			LocationAddress:   station.Address,
			CreatedDate:       nil,
			InstalledCapacity: pointy.Float64(station.GetCapacity() * 1000),
			TotalCO2:          pointy.Float64(mapPlantCodeToTotalCO2[stationCode] * 1000),
			MonthlyCO2:        &monthlyCO2,
			TotalSavingPrice:  mapPlantCodeToRealtimeData[stationCode].DataItemMap.TotalIncome,
			Currency:          pointy.String(huawei.HUAWEI_CURRENCY_USD),
			CurrentPower:      pointy.Float64(currentPower),
			DailyProduction:   &dailyProduction,
			MonthlyProduction: &monthlyProduction,
			YearlyProduction:  pointy.Float64(mapPlantCodeToYearlyPower[stationCode]),
			PlantStatus:       pointy.String(plantStatus),
			Owner:             credential.Owner,
		}

		plantItem.TotalProduction = mapPlantCodeToRealtimeData[stationCode].DataItemMap.TotalPower
		if pointy.Float64Value(plantItem.TotalProduction, 0.0) < pointy.Float64Value(plantItem.YearlyProduction, 0.0) {
			plantItem.TotalProduction = pointy.Float64(mapPlantCodeToTotalPower[stationCode])
		}

		documentCh <- plantItem

	}

	doneCh <- true
}

func (s *huaweiCollectorV3Service) PreparePlantAndDevice(credential *model.HuaweiCredential) error {
	// check latest plant update
	s.logger.Info("check latest plant")
	plant, err := s.huaweiAltervimRepo.GetLatestPlant()
	if err != nil {
		s.logger.Error(err)
		return err
	}

	if plant.Name != nil {
		nowStr := time.Now().Format("2006-01-02")
		plantUpdateStr := plant.UpdatedAt.Format("2006-01-02")
		if nowStr == plantUpdateStr {
			return nil
		}
	}

	// preparing
	client, err := huawei.NewHuaweiClient(&huawei.HuaweiCredential{
		Username: credential.Username,
		Password: credential.Password,
	})

	if err != nil {
		return err
	}

	plantData := []model.HuaweiAltervimPlant{}
	deviceData := []model.HuaweiAltervimDevice{}
	plantCodes := []string{}
	plantCodeStr := []string{}
	deviceIds := []int{}

	s.logger.Info("getting plant list...")
	plants, err := client.GetPlantList()
	if err != nil {
		s.logger.Error(err)
		return err
	}
	s.logger.Infof("got %v plant", len(plants))

	count := 1
	s.logger.Info("start preparing")
	for _, plant := range plants {
		s.logger.Infof("%v/%v", count, len(plants))
		count += 1

		if len(plantCodes) >= 100 {
			plantCodeStr = append(plantCodeStr, strings.Join(plantCodes, ","))
			plantCodes = []string{}
		}
		plantCodes = append(plantCodes, *plant.PlantCode)

		plantData = append(plantData, model.HuaweiAltervimPlant{
			Code:               plant.GetPlantCode(),
			Name:               plant.PlantName,
			Address:            plant.PlantAddress,
			Longitude:          plant.Longitude,
			Latitude:           plant.Latitude,
			Capacity:           plant.Capacity,
			ContactPerson:      plant.ContactPerson,
			ContactMethod:      plant.ContactMethod,
			GridConnectionData: plant.GridConnectionDate,
		})
	}
	plantCodeStr = append(plantCodeStr, strings.Join(plantCodes, ","))

	for _, code := range plantCodeStr {
		if resp, err := client.GetDeviceList(code); err != nil {
			s.logger.Error(err)
			return err
		} else if len(resp.Data) > 0 {
			for _, device := range resp.Data {
				deviceIds = append(deviceIds, *device.ID)
				deviceData = append(deviceData, model.HuaweiAltervimDevice{
					ID:              device.GetID(),
					SerialNumber:    device.SN,
					Name:            device.Name,
					TypeID:          device.TypeID,
					InverterModel:   device.InverterModel,
					Latitude:        device.Latitude,
					Longitude:       device.Longitude,
					SoftwareVersion: device.SoftwareVersion,
					PlantCode:       device.PlantCode,
				})
			}
		}
	}

	if err := s.huaweiAltervimRepo.BatchInsertDevices(deviceData); err != nil {
		s.logger.Error(err)
		return err
	}

	if err := s.huaweiAltervimRepo.BatchInsertPlants(plantData); err != nil {
		s.logger.Error(err)
		return err
	}

	// delete disappear plant and device
	if err := s.huaweiAltervimRepo.DeletePlantNotIn(plantCodes); err != nil {
		s.logger.Error(err)
		return err
	}

	if err := s.huaweiAltervimRepo.DeleteDeviceNotIn(deviceIds); err != nil {
		s.logger.Error(err)
		return err
	}

	return nil
}

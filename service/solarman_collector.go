package service

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/HavvokLab/true-solar-monitoring/config"
	"github.com/HavvokLab/true-solar-monitoring/constant"
	"github.com/HavvokLab/true-solar-monitoring/inverter"
	"github.com/HavvokLab/true-solar-monitoring/inverter/solarman"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/HavvokLab/true-solar-monitoring/repo"
	"github.com/sourcegraph/conc"
	"go.openly.dev/pointy"
)

type SolarmanCollectorService interface {
	Run(*model.SolarmanCredential) error
}

type solarmanCollectorService struct {
	vendorType     string
	siteRegionRepo repo.SiteRegionMappingRepo
	siteRegions    []model.SiteRegionMapping
	solarRepo      repo.SolarRepo
	elasticConfig  config.ElasticsearchConfig
	logger         logger.Logger
}

func NewSolarmanCollectorService(solarRepo repo.SolarRepo, siteRegionRepo repo.SiteRegionMappingRepo, logger logger.Logger) (SolarmanCollectorService, error) {
	return &solarmanCollectorService{
		vendorType:     strings.ToUpper(constant.VENDOR_TYPE_INVT),
		siteRegions:    make([]model.SiteRegionMapping, 0),
		siteRegionRepo: siteRegionRepo,
		solarRepo:      solarRepo,
		elasticConfig:  config.GetConfig().Elastic,
		logger:         logger,
	}, nil
}

func (s *solarmanCollectorService) Run(credential *model.SolarmanCredential) error {
	defer func() {
		if r := recover(); r != nil {
			s.logger.Warnf("[%v] - SolarmanCollectorService.Run(): %v", credential.Username, r)
		}
	}()

	siteRegions, err := s.siteRegionRepo.GetSiteRegionMappings()
	if err != nil {
		s.logger.Errorf("[%v] - SolarmanCollectorService.Run(): %v", credential.Username, err)
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
			s.logger.Errorf("[%v] - SolarmanCollectorService.Run(): %v", credential.Username, err)
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
		s.logger.Errorf("[%v] - SolarmanCollectorService.Run(): %v", credential.Username, err)
		return err
	}
	s.logger.Infof("[%v] - SolarmanCollectorService.Run(): indexed %v documents to %v", credential.Username, len(documents), collectorIndex)

	if err := s.solarRepo.UpsertSiteStation(siteDocuments); err != nil {
		s.logger.Errorf("[%v] - SolarmanCollectorService.Run(): %v", credential.Username, err)
		return err
	}
	s.logger.Infof("[%v] - SolarmanCollectorService.Run(): upserted %v site-documents", credential.Username, len(siteDocuments))

	close(doneCh)
	close(errorCh)
	close(documentCh)
	return nil
}

func (s *solarmanCollectorService) run(credential *model.SolarmanCredential, documentCh chan interface{}, doneCh chan bool, errorCh chan error) {
	now := time.Now()
	beginningOfToday := time.Date(now.Year(), now.Month(), now.Day(), 6, 0, 0, 0, time.Local)
	client, err := solarman.NewSolarmanClient(&solarman.SolarmanCredential{
		Username:  credential.Username,
		Password:  credential.Password,
		AppID:     credential.AppID,
		AppSecret: credential.AppSecret,
	})

	if err != nil {
		s.logger.Errorf("[%v] - SolarmanCollectorService.run(): %v", credential.Username, err)
		errorCh <- err
		return
	}

	userInfoResp, err := client.GetUserInfo()
	if err != nil {
		s.logger.Errorf("[%v] - SolarmanCollectorService.run(): %v", credential.Username, err)
		errorCh <- err
		return
	}

	wg := conc.NewWaitGroup()
	for _, company := range userInfoResp.OrgInfoList {
		company := company
		credential := credential

		producer := func() {
			client, err := solarman.NewSolarmanClient(&solarman.SolarmanCredential{
				Username:  credential.Username,
				Password:  credential.Password,
				AppID:     credential.AppID,
				AppSecret: credential.AppSecret,
			})
			if err != nil {
				s.logger.Errorf("[%v] - SolarmanCollectorService.run(): %v", credential.Username, err)
				errorCh <- err
				return
			}

			tokenResp, err := client.GetBusinessToken(company.GetCompanyID())
			if err != nil {
				s.logger.Errorf("[%v] - SolarmanCollectorService.run(): %v", credential.Username, err)
				errorCh <- err
				return
			}

			if tokenResp.AccessToken == nil {
				err := fmt.Errorf("access token should not be empty")
				s.logger.Errorf("[%v] - SolarmanCollectorService.run(): %v", credential.Username, err)
				errorCh <- err
				return
			}

			token := tokenResp.GetAccessToken()
			plantList, err := client.GetPlantList(token)
			if err != nil {
				s.logger.Errorf("[%v] - SolarmanCollectorService.run(): %v", credential.Username, err)
				errorCh <- err
				return
			}

			plantCount := 1
			plantSize := len(plantList)
			for _, station := range plantList {
				if station == nil {
					s.logger.Warnf("[%v] - SolarmanCollectorService.run(): some plant of company(%v) is nil", credential.Username, company.GetCompanyID())
					plantCount++
					continue
				}

				s.logger.Infof("[%v] - collecting plant(%v) of company(%v): %v/%v", credential.Username, station.GetID(), company.GetCompanyID(), plantCount, plantSize)
				plantCount++

				stationID := station.GetID()
				plantID, _ := inverter.ParsePlantID(station.GetName())
				cityName, cityCode, cityArea := inverter.ParseSiteID(s.siteRegions, plantID.SiteID)

				plantItemDoc := model.PlantItem{
					Timestamp:         now,
					Month:             now.Format("01"),
					Year:              now.Format("2006"),
					MonthYear:         now.Format("01-2006"),
					VendorType:        s.vendorType,
					DataType:          constant.DATA_TYPE_PLANT,
					Area:              cityArea,
					SiteID:            plantID.SiteID,
					SiteCityName:      cityName,
					SiteCityCode:      cityCode,
					NodeType:          plantID.NodeType,
					ACPhase:           plantID.ACPhase,
					ID:                pointy.String(strconv.Itoa(stationID)),
					Name:              station.Name,
					Latitude:          station.LocationLat,
					Longitude:         station.LocationLng,
					LocationAddress:   station.LocationAddress,
					InstalledCapacity: station.InstalledCapacity,
					Owner:             credential.Owner,
				}

				var (
					mergedElectricPrice         *float64
					totalPowerGenerationKWh     *float64
					sumYearlyPowerGenerationKWh *float64
				)

				if plantItemDoc.Latitude != nil && plantItemDoc.Longitude != nil {
					plantItemDoc.Location = pointy.String(fmt.Sprintf("%f,%f", *plantItemDoc.Latitude, *plantItemDoc.Longitude))
				}

				if station.CreatedDate != nil {
					parsed := time.Unix(int64(*station.CreatedDate), 0)
					plantItemDoc.CreatedDate = &parsed
				}

				if plantInfoResp, err := client.GetPlantBaseInfo(token, stationID); err == nil {
					plantItemDoc.Currency = plantInfoResp.Currency
					mergedElectricPrice = plantInfoResp.MergeElectricPrice
				}

				if realtimeDataResp, err := client.GetPlantRealtimeData(token, stationID); err == nil {
					plantItemDoc.CurrentPower = pointy.Float64(realtimeDataResp.GetGenerationPower() / 1000.0)
				}

				if resp, err := client.GetHistoricalPlantData(
					token,
					stationID,
					solarman.TIME_TYPE_DAY,
					now.Unix(),
					now.Unix(),
				); err == nil && resp != nil {
					if resp.StationDataItems != nil {
						if len(resp.StationDataItems) > 0 {
							plantItemDoc.DailyProduction = resp.StationDataItems[0].GenerationValue
						}
					}
				}

				if resp, err := client.GetHistoricalPlantData(
					token,
					stationID,
					solarman.TIME_TYPE_MONTH,
					now.Unix(),
					now.Unix(),
				); err == nil && len(resp.StationDataItems) > 0 {
					plantItemDoc.MonthlyProduction = resp.StationDataItems[0].GenerationValue
				}

				startTime := time.Date(2015, now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
				if resp, err := client.GetHistoricalPlantData(
					token,
					stationID,
					solarman.TIME_TYPE_YEAR,
					startTime.Unix(),
					now.Unix(),
				); err == nil && len(resp.StationDataItems) > 0 {
					for _, item := range resp.StationDataItems {
						if item.GetYear() == now.Year() {
							plantItemDoc.YearlyProduction = item.GenerationValue
						}

						sumYearlyPowerGenerationKWh = pointy.Float64(
							pointy.Float64Value(sumYearlyPowerGenerationKWh, 0.0) + item.GetGenerationValue(),
						)
					}
				}

				deviceList, err := client.GetPlantDeviceList(token, stationID)
				if err != nil {
					s.logger.Errorf("[%v] - SolarmanCollectorService.run(): %v", credential.Username, err)
					errorCh <- err
					return
				}

				deviceCount := 1
				deviceSize := len(deviceList)
				deviceStatusArray := make([]string, 0)
				for _, device := range deviceList {
					s.logger.Infof("[%v] - collecting device(%v) of plant(%v): %v/%v", credential.Username, device.GetDeviceSN(), stationID, deviceCount, deviceSize)
					deviceCount++

					if device == nil {
						s.logger.Warnf("[%v] - SolarmanCollectorService.run(): some device of plant(%v) is nil", credential.Username, stationID)
						continue
					}

					deviceSN := device.GetDeviceSN()
					deviceID := device.GetDeviceID()

					deviceItemDoc := model.DeviceItem{
						Timestamp:    now,
						Month:        now.Format("01"),
						Year:         now.Format("2006"),
						MonthYear:    now.Format("01-2006"),
						VendorType:   s.vendorType,
						DataType:     constant.DATA_TYPE_DEVICE,
						Area:         cityArea,
						SiteID:       plantID.SiteID,
						SiteCityName: cityName,
						SiteCityCode: cityCode,
						NodeType:     plantID.NodeType,
						ACPhase:      plantID.ACPhase,
						PlantID:      pointy.String(strconv.Itoa(stationID)),
						PlantName:    station.Name,
						Latitude:     plantItemDoc.Latitude,
						Longitude:    plantItemDoc.Longitude,
						Location:     plantItemDoc.Location,
						ID:           pointy.String(strconv.Itoa(deviceID)),
						SN:           device.DeviceSN,
						Name:         device.DeviceSN,
						DeviceType:   device.DeviceType,
						Owner:        credential.Owner,
					}

					if resp, err := client.GetDeviceRealtimeData(token, deviceSN); err == nil {
						if len(resp.DataList) > 0 {
							for _, data := range resp.DataList {
								if data.GetKey() == solarman.DATA_LIST_KEY_CUMULATIVE_PRODUCTION {
									if generation, err := strconv.ParseFloat(data.GetValue(), 64); err == nil {
										totalPowerGenerationKWh = pointy.Float64(
											pointy.Float64Value(totalPowerGenerationKWh, 0.0) + generation,
										)
									}
								}
							}
						}
					}

					if resp, err := client.GetHistoricalDeviceData(token, deviceSN, solarman.TIME_TYPE_DAY, now.Unix(), now.Unix()); err == nil && len(resp.ParamDataList) > 0 {
						for _, param := range resp.ParamDataList {
							if param.DataList != nil {
								for _, data := range param.DataList {
									if data.GetKey() == solarman.DATA_LIST_KEY_GENERATION {
										if generation, err := strconv.ParseFloat(data.GetValue(), 64); err == nil {
											deviceItemDoc.DailyPowerGeneration = pointy.Float64(
												pointy.Float64Value(deviceItemDoc.DailyPowerGeneration, 0.0) + generation,
											)
										}
									}
								}
							}
						}
					}

					if resp, err := client.GetHistoricalDeviceData(token, deviceSN, solarman.TIME_TYPE_MONTH, now.Unix(), now.Unix()); err == nil && len(resp.ParamDataList) > 0 {
						for _, param := range resp.ParamDataList {
							if param.DataList != nil {
								for _, data := range param.DataList {
									if data.GetKey() == solarman.DATA_LIST_KEY_GENERATION {
										if generation, err := strconv.ParseFloat(data.GetValue(), 64); err == nil {
											deviceItemDoc.MonthlyPowerGeneration = pointy.Float64(
												pointy.Float64Value(deviceItemDoc.MonthlyPowerGeneration, 0.0) + generation,
											)
										}
									}
								}
							}
						}
					}

					if resp, err := client.GetHistoricalDeviceData(token, deviceSN, solarman.TIME_TYPE_YEAR, now.Unix(), now.Unix()); err == nil && len(resp.ParamDataList) > 0 {
						for _, param := range resp.ParamDataList {
							if param.DataList != nil {
								for _, data := range param.DataList {
									if data.GetKey() == solarman.DATA_LIST_KEY_GENERATION {
										if generation, err := strconv.ParseFloat(data.GetValue(), 64); err == nil {
											deviceItemDoc.YearlyPowerGeneration = pointy.Float64(
												pointy.Float64Value(deviceItemDoc.YearlyPowerGeneration, 0.0) + generation,
											)
										}
									}
								}
							}
						}
					}

					if device.CollectionTime != nil {
						parsed := time.Unix(device.GetCollectionTime(), 0)
						deviceItemDoc.LastUpdateTime = &parsed
					}

					if device.ConnectStatus != nil {
						switch device.GetConnectStatus() {
						case 0:
							deviceItemDoc.Status = pointy.String(solarman.DEVICE_STATUS_OFF)
						case 1:
							deviceItemDoc.Status = pointy.String(solarman.DEVICE_STATUS_ON)
						case 2:
							deviceItemDoc.Status = pointy.String(solarman.DEVICE_STATUS_FAILURE)

							if alertList, err := client.GetDeviceAlertList(token, deviceSN, beginningOfToday.Unix(), now.Unix()); err == nil {
								alertCount := 0
								alertSize := len(alertList)
								for _, alert := range alertList {
									s.logger.Infof("[%v] - collecting alert(%v) of device(%v): %v/%v", credential.Username, alert.GetAlertID(), deviceSN, alertCount, alertSize)
									alertCount++

									alarmItemDoc := model.AlarmItem{
										Timestamp:    now,
										Month:        now.Format("01"),
										Year:         now.Format("2006"),
										MonthYear:    now.Format("01-2006"),
										VendorType:   s.vendorType,
										DataType:     constant.DATA_TYPE_ALARM,
										Area:         cityArea,
										SiteID:       plantID.SiteID,
										SiteCityName: cityName,
										SiteCityCode: cityCode,
										NodeType:     plantID.NodeType,
										ACPhase:      plantID.ACPhase,
										PlantID:      pointy.String(strconv.Itoa(stationID)),
										PlantName:    station.Name,
										Latitude:     plantItemDoc.Latitude,
										Longitude:    plantItemDoc.Longitude,
										Location:     plantItemDoc.Location,
										DeviceID:     pointy.String(strconv.Itoa(deviceID)),
										DeviceSN:     device.DeviceSN,
										DeviceName:   device.DeviceSN,
										DeviceType:   device.DeviceType,
										DeviceStatus: deviceItemDoc.Status,
										ID:           pointy.String(strconv.Itoa(alert.GetAlertID())),
										Message:      alert.AlertNameInPAAS,
										Owner:        credential.Owner,
									}

									if alert.AlertTime != nil {
										alertTime := time.Unix(alert.GetAlertTime(), 0)
										alarmItemDoc.AlarmTime = &alertTime
									}

									documentCh <- alarmItemDoc
								}
							}
						default:
						}
					}

					if deviceItemDoc.Status != nil {
						deviceStatusArray = append(deviceStatusArray, *deviceItemDoc.Status)
					}

					documentCh <- deviceItemDoc
				}

				plantStatus := solarman.SOLARMAN_PLANT_STATUS_ON
				if len(deviceStatusArray) > 0 {
					var offlineCount int
					var alertingCount int

					for _, status := range deviceStatusArray {
						switch status {
						case solarman.DEVICE_STATUS_OFF:
							offlineCount++
						case solarman.DEVICE_STATUS_ON:
						default:
							alertingCount++
						}
					}

					if alertingCount > 0 {
						plantStatus = solarman.SOLARMAN_PLANT_STATUS_ALARM
					} else if offlineCount > 0 {
						plantStatus = solarman.SOLARMAN_PLANT_STATUS_OFF
					}
				} else {
					plantStatus = solarman.SOLARMAN_PLANT_STATUS_OFF
				}

				plantItemDoc.TotalProduction = totalPowerGenerationKWh
				if pointy.Float64Value(plantItemDoc.TotalProduction, 0.0) < pointy.Float64Value(plantItemDoc.YearlyProduction, 0.0) {
					plantItemDoc.TotalProduction = plantItemDoc.YearlyProduction
				}

				plantItemDoc.PlantStatus = &plantStatus
				plantItemDoc.TotalSavingPrice = pointy.Float64(
					pointy.Float64Value(mergedElectricPrice, 0.0) * pointy.Float64Value(totalPowerGenerationKWh, 0.0),
				)
				documentCh <- plantItemDoc
			}
		}

		wg.Go(producer)
	}

	if r := wg.WaitAndRecover(); r != nil {
		s.logger.Warnf("[%v] - SolarmanCollectorService.run(): \ncaller: %v, \nvalue: %v \nstring: %v", credential.Username, r.Callers, r.Value, r.String())
		return
	}

	doneCh <- true
}

package service

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/HavvokLab/true-solar-monitoring/config"
	"github.com/HavvokLab/true-solar-monitoring/constant"
	"github.com/HavvokLab/true-solar-monitoring/inverter"
	"github.com/HavvokLab/true-solar-monitoring/inverter/growatt"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/HavvokLab/true-solar-monitoring/repo"
	"github.com/mitchellh/mapstructure"
	"github.com/sourcegraph/conc"
	"go.openly.dev/pointy"
)

type GrowattCollectorService interface {
	Run(*model.GrowattCredential) error
}

type growattCollectorService struct {
	vendorType     string
	siteRegionRepo repo.SiteRegionMappingRepo
	siteRegions    []model.SiteRegionMapping
	solarRepo      repo.SolarRepo
	elasticConfig  config.ElasticsearchConfig
	logger         logger.Logger
}

func NewGrowattCollectorService(solarRepo repo.SolarRepo, siteRegionRepo repo.SiteRegionMappingRepo, logger logger.Logger) GrowattCollectorService {
	return &growattCollectorService{
		vendorType:     strings.ToUpper(constant.VENDOR_TYPE_GROWATT),
		siteRegionRepo: siteRegionRepo,
		solarRepo:      solarRepo,
		logger:         logger,
		siteRegions:    make([]model.SiteRegionMapping, 0),
		elasticConfig:  config.GetConfig().Elastic,
	}
}

func (s *growattCollectorService) Run(credential *model.GrowattCredential) error {
	defer func() {
		if r := recover(); r != nil {
			s.logger.Errorf("[%v] Recovered from panic: %v", credential.Username, r)
		}
	}()

	siteRegions, err := s.siteRegionRepo.GetSiteRegionMappings()
	if err != nil {
		s.logger.Errorf("[%v] Error while getting site region mappings: %v", credential.Username, err)
		return err
	}
	s.siteRegions = siteRegions

	plantDeviceStatusMap := make(map[string]string)
	inverterArray := make([]string, 0)
	documents := make([]interface{}, 0)
	siteDocuments := make([]model.SiteItem, 0)

	documentCh := make(chan interface{})
	inverterCh := make(chan string)
	plantDeviceStatusCh := make(chan map[string]string)
	doneCh := make(chan bool)
	errorCh := make(chan error)
	go s.run(credential, documentCh, inverterCh, plantDeviceStatusCh, doneCh, errorCh)

DONE:
	for {
		select {
		case <-doneCh:
			break DONE
		case err := <-errorCh:
			s.logger.Errorf("[%v] Error while running growatt collector: %v", credential.Username, err)
			return err
		case doc := <-documentCh:
			documents = append(documents, doc)
		case plantDeviceStatus := <-plantDeviceStatusCh:
			if err := mapstructure.Decode(&plantDeviceStatus, &plantDeviceStatusMap); err != nil {
				s.logger.Errorf("[%v] Error while decoding plant device status: %v", credential.Username, err)
				return err
			}
		case sn := <-inverterCh:
			inverterArray = append(inverterArray, sn)
		}
	}

	realtimeDeviceMap, err := growatt.CalculateInverterProductions(
		&growatt.GrowattCredential{
			Username: credential.Username,
			Token:    credential.Token,
		},
		inverterArray,
	)

	if err != nil {
		s.logger.Errorf("[%v] Error while calculating inverter productions: %v", credential.Username, err)
	}

	for i, doc := range documents {
		if plantItem, ok := doc.(model.PlantItem); ok {
			if plantItem.ID != nil {
				if plantStatus, found := plantDeviceStatusMap[*plantItem.ID]; found {
					plantItem.PlantStatus = &plantStatus
					documents[i] = plantItem
				}
			}

			siteItem := model.SiteItem{
				Timestamp:   plantItem.Timestamp,
				VendorType:  plantItem.VendorType,
				Area:        plantItem.Area,
				SiteID:      plantItem.SiteID,
				NodeType:    plantItem.NodeType,
				Name:        plantItem.Name,
				Location:    plantItem.Location,
				PlantStatus: plantItem.PlantStatus,
				Owner:       plantItem.Owner,
			}

			siteDocuments = append(siteDocuments, siteItem)
		} else if deviceItem, ok := doc.(model.DeviceItem); ok {
			if deviceItem.SN != nil {
				if data, ok := realtimeDeviceMap[*deviceItem.SN]; ok {
					deviceItem.TotalPowerGeneration = data.Total
					deviceItem.DailyPowerGeneration = data.Today
					documents[i] = deviceItem
				}
			}
		}
	}

	collectorIndex := fmt.Sprintf("%s-%s", s.elasticConfig.SolarIndex, time.Now().Format("2006.01.02"))
	if err := s.solarRepo.BulkIndex(collectorIndex, documents); err != nil {
		s.logger.Errorf("[%v] - GrowattCollectorService.Run(): %v", credential.Username, err)
		return err
	}
	s.logger.Infof("[%v] - GrowattCollectorService.Run(): %v documents indexed", credential.Username, len(documents))

	if err := s.solarRepo.UpsertSiteStation(siteDocuments); err != nil {
		s.logger.Errorf("[%v] - GrowattCollectorService.Run(): %v", credential.Username, err)
		return err
	}
	s.logger.Infof("[%v] - GrowattCollectorService.Run(): %v site stations upserted", credential.Username, len(siteDocuments))

	close(documentCh)
	close(doneCh)
	close(errorCh)
	close(inverterCh)
	close(plantDeviceStatusCh)
	return nil
}

func (s *growattCollectorService) run(credential *model.GrowattCredential, documentCh chan interface{}, inverterCh chan string, plantDeviceStatusCh chan map[string]string, doneCh chan bool, errorCh chan error) {
	now := time.Now()
	wg := conc.NewWaitGroup()
	client, err := growatt.NewGrowattClient(&growatt.GrowattCredential{
		Username: credential.Username,
		Token:    credential.Token,
	})

	if err != nil {
		s.logger.Errorf("[%v] Error while creating growatt client: %v", credential.Username, err)
		errorCh <- err
		return
	}

	plantList, err := client.GetPlantList()
	if err != nil {
		s.logger.Errorf("[%v] Error while getting plant list: %v", credential.Username, err)
		errorCh <- err
		return
	}

	plantCount := 1
	plantSize := len(plantList)
	for _, station := range plantList {
		s.logger.Infof("[%v] Processing plant %v/%v", credential.Username, plantCount, plantSize)
		plantCount++

		station := station
		producer := func() {
			stationID := station.GetPlantID()
			stationIDStr := strconv.Itoa(stationID)
			plantID, _ := inverter.ParsePlantID(station.GetName())
			cityName, cityCode, cityArea := inverter.ParseSiteID(s.siteRegions, plantID.SiteID)

			plantItem := model.PlantItem{
				Timestamp:    now,
				Month:        now.Format("01"),
				Year:         now.Format("2006"),
				MonthYear:    now.Format("01-2006"),
				VendorType:   s.vendorType,
				DataType:     constant.DATA_TYPE_PLANT,
				Area:         cityArea,
				SiteID:       plantID.SiteID,
				SiteCityName: cityName,
				SiteCityCode: cityCode,
				NodeType:     plantID.NodeType,
				ACPhase:      plantID.ACPhase,
				ID:           pointy.String(stationIDStr),
				Name:         station.Name,
				PlantStatus:  pointy.String(growatt.GROWATT_PLANT_STATUS_OFF),
				Owner:        credential.Owner,
			}

			var electricPricePerKWh *float64
			var co2WeightPerKWh *float64

			if station.Latitude != nil {
				if parsed, err := strconv.ParseFloat(station.GetLatitude(), 64); err == nil {
					plantItem.Latitude = &parsed
				}
			}

			if station.Longitude != nil {
				if parsed, err := strconv.ParseFloat(station.GetLongitude(), 64); err == nil {
					plantItem.Longitude = &parsed
				}
			}

			if plantItem.Latitude != nil && plantItem.Longitude != nil {
				plantItem.Location = pointy.String(fmt.Sprintf("%f,%f", *plantItem.Latitude, *plantItem.Longitude))
			}

			if station.City != nil {
				if *station.City != "" {
					plantItem.LocationAddress = station.City
				}
			}

			if station.Country != nil {
				if *station.Country != "" {
					if plantItem.LocationAddress != nil {
						plantItem.LocationAddress = pointy.String(fmt.Sprintf("%s, %s", *plantItem.LocationAddress, *station.Country))
					} else {
						plantItem.LocationAddress = station.Country
					}
				}
			}

			if dataLoggerResp, err := client.GetPlantDataLoggerInfo(stationID); err == nil {
				if dataLoggerResp.Data != nil {
					if dataLoggerResp.Data.PeakPowerActual != nil {
						actualData := dataLoggerResp.Data.PeakPowerActual
						electricPricePerKWh = actualData.FormulaMoney
						co2WeightPerKWh = actualData.FormulaCo2

						if actualData.NominalPower != nil {
							plantItem.InstalledCapacity = pointy.Float64(actualData.GetNominalPower() / 1000.0)
						} else if plantID.Capacity != 0 {
							plantItem.InstalledCapacity = pointy.Float64(plantID.Capacity)
						}

						if actualData.FormulaMoneyUnitID != nil {
							plantItem.Currency = pointy.String(strings.ToUpper(actualData.GetFormulaMoneyUnitID()))
						}
					}
				}
			}

			if overviewInfoResp, err := client.GetPlantOverviewInfo(stationID); err == nil {
				if overviewInfoResp.Data != nil {
					plantItem.CurrentPower = overviewInfoResp.Data.CurrentPower

					if overviewInfoResp.Data.TodayEnergy != nil {
						if parsed, err := strconv.ParseFloat(overviewInfoResp.Data.GetTodayEnergy(), 64); err == nil {
							plantItem.DailyProduction = &parsed
						}
					}

					if overviewInfoResp.Data.MonthlyEnergy != nil {
						if parsed, err := strconv.ParseFloat(overviewInfoResp.Data.GetMonthlyEnergy(), 64); err == nil {
							plantItem.MonthlyProduction = &parsed

							if co2WeightPerKWh != nil {
								plantItem.MonthlyCO2 = pointy.Float64(parsed * pointy.Float64Value(co2WeightPerKWh, 0.0))
							}
						}
					}

					if overviewInfoResp.Data.YearlyEnergy != nil {
						if parsed, err := strconv.ParseFloat(overviewInfoResp.Data.GetYearlyEnergy(), 64); err == nil {
							plantItem.YearlyProduction = &parsed
						}
					}

					if overviewInfoResp.Data.TotalEnergy != nil {
						if parsed, err := strconv.ParseFloat(overviewInfoResp.Data.GetTotalEnergy(), 64); err == nil {
							plantItem.TotalProduction = &parsed

							if electricPricePerKWh != nil {
								plantItem.TotalSavingPrice = pointy.Float64(parsed * pointy.Float64Value(electricPricePerKWh, 0.0))
							}

							if co2WeightPerKWh != nil {
								plantItem.TotalCO2 = pointy.Float64(parsed * pointy.Float64Value(co2WeightPerKWh, 0.0))
							}
						}
					}
				}
			}

			documentCh <- plantItem

			deviceList, err := client.GetPlantDeviceList(stationID)
			if err != nil {
				s.logger.Errorf("[%v] Error while getting plant device list: %v", credential.Username, err)
				errorCh <- err
				return
			}

			deviceCount := 1
			deviceSize := len(deviceList)
			deviceStatusArray := make([]string, 0)
			for _, device := range deviceList {
				deviceSN := device.GetDeviceSN()
				deviceID := device.GetDeviceID()
				deviceTypeRaw := device.GetType()
				deviceType := growatt.ParseGrowattDeviceType(deviceTypeRaw)

				s.logger.Infof("[%v] Processing device %v/%v", credential.Username, deviceCount, deviceSize)
				deviceCount++

				deviceItem := model.DeviceItem{
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
					PlantID:      pointy.String(stationIDStr),
					PlantName:    station.Name,
					Latitude:     plantItem.Latitude,
					Longitude:    plantItem.Longitude,
					Location:     plantItem.Location,
					ID:           pointy.String(strconv.Itoa(deviceID)),
					SN:           device.DeviceSN,
					Name:         device.DeviceSN,
					DeviceType:   &deviceType,
					Owner:        credential.Owner,
				}

				if device.LastUpdateTime != nil {
					if parsed, err := time.Parse("2006-01-02 15:04:05", device.GetLastUpdateTime()); err == nil {
						deviceItem.LastUpdateTime = &parsed
					}
				}

				switch deviceTypeRaw {
				case growatt.GROWATT_DEVICE_TYPE_INVERTER:
					if device.Status != nil {
						switch *device.Status {
						case 0: // Offline
							deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_OFFLINE)
						case 1: // Online
							deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_ONLINE)
						default: // Others
							if *device.Status == 2 { // Stand by
								deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_STAND_BY)
							} else if *device.Status == 3 { // Failure
								deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_FAILURE)
							} else {
								deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_OFFLINE)
							}

							if alarms, err := client.GetInverterAlertList(deviceSN); err == nil {
								if len(alarms) > 0 {
									if latestAlert := alarms[0]; latestAlert != nil {
										if startTime := latestAlert.StartTime; startTime != nil {
											if pointy.StringValue(device.LastUpdateTime, "0000-00-00")[0:10] == (*startTime)[0:10] {
												alarmItem := model.AlarmItem{
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
													PlantID:      pointy.String(stationIDStr),
													PlantName:    station.Name,
													Latitude:     plantItem.Latitude,
													Longitude:    plantItem.Longitude,
													Location:     plantItem.Location,
													DeviceID:     pointy.String(strconv.Itoa(deviceID)),
													DeviceSN:     device.DeviceSN,
													DeviceName:   device.DeviceSN,
													DeviceType:   pointy.String(growatt.ParseGrowattDeviceType(growatt.GROWATT_DEVICE_TYPE_INVERTER)),
													DeviceStatus: deviceItem.Status,
													ID:           pointy.String(strconv.Itoa(pointy.IntValue(latestAlert.AlarmCode, 0))),
													Message:      latestAlert.AlarmMessage,
													Owner:        credential.Owner,
												}

												if latestAlert.StartTime != nil {
													if parsed, err := time.Parse("2006-01-02 15:04:05.0", *latestAlert.StartTime); err == nil {
														utcParsed := parsed.UTC()
														alarmItem.AlarmTime = &utcParsed
													}
												}

												documentCh <- alarmItem
											}
										}
									}

								}
							}
						}

					}
				case growatt.GROWATT_DEVICE_TYPE_ENERGY_STORAGE_MACHINE:
					if device.Status != nil {
						if *device.Status == 0 { // Stand by
							deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_STAND_BY)
						} else if *device.Status == 1 { // Charging
							deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_CHARGING)
						} else if *device.Status == 2 { // Discharging
							deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_DISCHARGING)
						} else if *device.Status == 3 { // Failure
							deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_FAILURE)
						} else if *device.Status == 4 { // Burning
							deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_BURNING)
						} else {
							deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_OFFLINE)
						}

						if alertListResp, err := client.GetEnergyStorageMachineAlertList(deviceSN, now.Unix()); err == nil {
							if alertListResp.Data != nil && device.LastUpdateTime != nil {
								if len(alertListResp.Data.Alarms) > 0 {
									if latestAlert := alertListResp.Data.Alarms[0]; latestAlert != nil {
										if startTime := latestAlert.StartTime; startTime != nil {
											if pointy.StringValue(device.LastUpdateTime, "0000-00-00")[0:10] == (*startTime)[0:10] {
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
													PlantID:      pointy.String(stationIDStr),
													PlantName:    station.Name,
													Latitude:     plantItem.Latitude,
													Longitude:    plantItem.Longitude,
													Location:     plantItem.Location,
													DeviceID:     pointy.String(strconv.Itoa(deviceID)),
													DeviceSN:     device.DeviceSN,
													DeviceName:   device.DeviceSN,
													DeviceType:   pointy.String(growatt.ParseGrowattDeviceType(growatt.GROWATT_DEVICE_TYPE_ENERGY_STORAGE_MACHINE)),
													DeviceStatus: deviceItem.Status,
													ID:           pointy.String(strconv.Itoa(pointy.IntValue(latestAlert.AlarmCode, 0))),
													Message:      latestAlert.AlarmMessage,
													Owner:        credential.Owner,
												}

												if latestAlert.StartTime != nil {
													if parsed, err := time.Parse("2006-01-02 15:04:05.0", *latestAlert.StartTime); err == nil {
														utcParsed := parsed.UTC()
														alarmItemDoc.AlarmTime = &utcParsed
													}
												}

												documentCh <- alarmItemDoc
											}
										}
									}
								}
							}

						}
					}
				case growatt.GROWATT_DEVICE_TYPE_OTHER_EQUIPMENT:
				case growatt.GROWATT_DEVICE_TYPE_MAX:
					if device.Status != nil {
						switch *device.Status {
						case 1: // Online
							deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_ONLINE)
						default: // Others
							if *device.Status == 0 || *device.Status == 2 { // Stand by
								deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_STAND_BY)
							} else if *device.Status == 3 { // Failure
								deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_FAILURE)
							} else {
								deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_OFFLINE)
							}

							if alarms, err := client.GetMaxAlertList(deviceSN, now.Unix()); err == nil {
								if len(alarms) > 0 {
									if latestAlert := alarms[0]; latestAlert != nil {
										if startTime := latestAlert.StartTime; startTime != nil {
											if pointy.StringValue(device.LastUpdateTime, "0000-00-00")[0:10] == (*startTime)[0:10] {
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
													PlantID:      pointy.String(stationIDStr),
													PlantName:    station.Name,
													Latitude:     plantItem.Latitude,
													Longitude:    plantItem.Longitude,
													Location:     plantItem.Location,
													DeviceID:     pointy.String(strconv.Itoa(deviceID)),
													DeviceSN:     device.DeviceSN,
													DeviceName:   device.DeviceSN,
													DeviceType:   pointy.String(growatt.ParseGrowattDeviceType(growatt.GROWATT_DEVICE_TYPE_MAX)),
													DeviceStatus: deviceItem.Status,
													ID:           pointy.String(strconv.Itoa(pointy.IntValue(latestAlert.AlarmCode, 0))),
													Message:      latestAlert.AlarmMessage,
													Owner:        credential.Owner,
												}

												if latestAlert.StartTime != nil {
													if parsed, err := time.Parse("2006-01-02 15:04:05.0", *latestAlert.StartTime); err == nil {
														utcParsed := parsed.UTC()
														alarmItemDoc.AlarmTime = &utcParsed
													}
												}

												documentCh <- alarmItemDoc
											}
										}
									}
								}
							}
						}
					}
				case growatt.GROWATT_DEVICE_TYPE_MIX:
					if device.Status != nil {
						switch *device.Status {
						case 5, 6, 7, 8: // Normal
							deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_ONLINE)
						default: // Others
							if *device.Status == 0 { // Waiting
								deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_WAITING)
							} else if *device.Status == 1 { // Self-check
								deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_SELF_CHECK)
							} else if *device.Status == 3 { // Failure
								deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_FAILURE)
							} else if *device.Status == 4 { // Upgrading
								deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_UPGRADING)
							} else {
								deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_OFFLINE)
							}

							if alarms, err := client.GetMixAlertList(deviceSN, now.Unix()); err == nil {
								if len(alarms) > 0 {
									if latestAlert := alarms[0]; latestAlert != nil {
										if startTime := latestAlert.StartTime; startTime != nil {
											if pointy.StringValue(device.LastUpdateTime, "0000-00-00")[0:10] == (*startTime)[0:10] {
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
													PlantID:      pointy.String(stationIDStr),
													PlantName:    station.Name,
													Latitude:     plantItem.Latitude,
													Longitude:    plantItem.Longitude,
													Location:     plantItem.Location,
													DeviceID:     pointy.String(strconv.Itoa(deviceID)),
													DeviceSN:     device.DeviceSN,
													DeviceName:   device.DeviceSN,
													DeviceType:   pointy.String(growatt.ParseGrowattDeviceType(growatt.GROWATT_DEVICE_TYPE_MIX)),
													DeviceStatus: deviceItem.Status,
													ID:           pointy.String(strconv.Itoa(pointy.IntValue(latestAlert.AlarmCode, 0))),
													Message:      latestAlert.AlarmMessage,
													Owner:        credential.Owner,
												}

												if latestAlert.StartTime != nil {
													if parsed, err := time.Parse("2006-01-02 15:04:05.0", *latestAlert.StartTime); err == nil {
														utcParsed := parsed.UTC()
														alarmItemDoc.AlarmTime = &utcParsed
													}
												}

												documentCh <- alarmItemDoc
											}
										}
									}
								}
							}
						}

					}
				case growatt.GROWATT_DEVICE_TYPE_SPA:
					if device.Status != nil {
						switch *device.Status {
						case 5, 6, 7, 8: // Normal
							deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_ONLINE)
						default: // Others
							if *device.Status == 0 { // Waiting
								deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_WAITING)
							} else if *device.Status == 1 { // Self-check
								deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_SELF_CHECK)
							} else if *device.Status == 3 { // Failure
								deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_FAILURE)
							} else if *device.Status == 4 { // Upgrading
								deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_UPGRADING)
							} else {
								deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_OFFLINE)
							}

							if alarms, err := client.GetSpaAlertList(deviceSN, now.Unix()); err == nil {
								if len(alarms) > 0 {
									if latestAlert := alarms[0]; latestAlert != nil {
										if startTime := latestAlert.StartTime; startTime != nil {
											if pointy.StringValue(device.LastUpdateTime, "0000-00-00")[0:10] == (*startTime)[0:10] {
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
													PlantID:      pointy.String(stationIDStr),
													PlantName:    station.Name,
													Latitude:     plantItem.Latitude,
													Longitude:    plantItem.Longitude,
													Location:     plantItem.Location,
													DeviceID:     pointy.String(strconv.Itoa(deviceID)),
													DeviceSN:     device.DeviceSN,
													DeviceName:   device.DeviceSN,
													DeviceType:   pointy.String(growatt.ParseGrowattDeviceType(growatt.GROWATT_DEVICE_TYPE_SPA)),
													DeviceStatus: deviceItem.Status,
													ID:           pointy.String(strconv.Itoa(pointy.IntValue(latestAlert.AlarmCode, 0))),
													Message:      latestAlert.AlarmMessage,
													Owner:        credential.Owner,
												}

												if latestAlert.StartTime != nil {
													if parsed, err := time.Parse("2006-01-02 15:04:05.0", *latestAlert.StartTime); err == nil {
														utcParsed := parsed.UTC()
														alarmItemDoc.AlarmTime = &utcParsed
													}
												}

												documentCh <- alarmItemDoc
											}
										}
									}
								}
							}
						}
					}

				case growatt.GROWATT_DEVICE_TYPE_MIN:
					if device.Status != nil {
						switch *device.Status {
						case 0: // Offline
							deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_OFFLINE)
						case 1: // Online
							deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_ONLINE)
						default: // Others
							if *device.Status == 2 { // Stand by
								deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_STAND_BY)
							} else if *device.Status == 3 { // Failure
								deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_FAILURE)
							} else {
								deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_OFFLINE)
							}

							if alarms, err := client.GetMinAlertList(deviceSN, now.Unix()); err == nil {
								if len(alarms) > 0 {
									if latestAlert := alarms[0]; latestAlert != nil {
										if startTime := latestAlert.StartTime; startTime != nil {
											if pointy.StringValue(device.LastUpdateTime, "0000-00-00")[0:10] == (*startTime)[0:10] {
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
													PlantID:      pointy.String(stationIDStr),
													PlantName:    station.Name,
													Latitude:     plantItem.Latitude,
													Longitude:    plantItem.Longitude,
													Location:     plantItem.Location,
													DeviceID:     pointy.String(strconv.Itoa(deviceID)),
													DeviceSN:     device.DeviceSN,
													DeviceName:   device.DeviceSN,
													DeviceType:   pointy.String(growatt.ParseGrowattDeviceType(growatt.GROWATT_DEVICE_TYPE_MIN)),
													DeviceStatus: deviceItem.Status,
													ID:           pointy.String(strconv.Itoa(pointy.IntValue(latestAlert.AlarmCode, 0))),
													Message:      latestAlert.AlarmMessage,
													Owner:        credential.Owner,
												}

												if latestAlert.StartTime != nil {
													if parsed, err := time.Parse("2006-01-02 15:04:05.0", *latestAlert.StartTime); err == nil {
														utcParsed := parsed.UTC()
														alarmItemDoc.AlarmTime = &utcParsed
													}
												}

												documentCh <- alarmItemDoc
											}
										}
									}
								}
							}
						}
					}
				case growatt.GROWATT_DEVICE_TYPE_PCS:
					if device.Status != nil {
						switch *device.Status {
						case 0: // Offline
							deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_OFFLINE)
						case 1: // Online
							deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_ONLINE)
						default: // Others
							if *device.Status == 2 { // Stand by
								deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_STAND_BY)
							} else if *device.Status == 3 { // Failure
								deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_FAILURE)
							} else {
								deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_OFFLINE)
							}

							if alarms, err := client.GetPcsAlertList(deviceSN, now.Unix()); err == nil {
								if len(alarms) > 0 {
									if latestAlert := alarms[0]; latestAlert != nil {
										if startTime := latestAlert.StartTime; startTime != nil {
											if pointy.StringValue(device.LastUpdateTime, "0000-00-00")[0:10] == (*startTime)[0:10] {
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
													PlantID:      pointy.String(stationIDStr),
													PlantName:    station.Name,
													Latitude:     plantItem.Latitude,
													Longitude:    plantItem.Longitude,
													Location:     plantItem.Location,
													DeviceID:     pointy.String(strconv.Itoa(deviceID)),
													DeviceSN:     device.DeviceSN,
													DeviceName:   device.DeviceSN,
													DeviceType:   pointy.String(growatt.ParseGrowattDeviceType(growatt.GROWATT_DEVICE_TYPE_PCS)),
													DeviceStatus: deviceItem.Status,
													ID:           pointy.String(strconv.Itoa(pointy.IntValue(latestAlert.AlarmCode, 0))),
													Message:      latestAlert.AlarmMessage,
													Owner:        credential.Owner,
												}

												if latestAlert.StartTime != nil {
													if parsed, err := time.Parse("2006-01-02 15:04:05.0", *latestAlert.StartTime); err == nil {
														utcParsed := parsed.UTC()
														alarmItemDoc.AlarmTime = &utcParsed
													}
												}

												documentCh <- alarmItemDoc
											}
										}
									}
								}
							}
						}

					}
				case growatt.GROWATT_DEVICE_TYPE_PBD:
					if device.Status != nil {
						switch *device.Status {
						case 0: // Offline
							deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_OFFLINE)
						case 1: // Online
							deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_ONLINE)
						default: // Others
							if *device.Status == 2 { // Stand by
								deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_STAND_BY)
							} else if *device.Status == 3 { // Failure
								deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_FAILURE)
							} else {
								deviceItem.Status = pointy.String(growatt.GROWATT_DEVICE_STATUS_OFFLINE)
							}

							if alarms, err := client.GetPbdAlertList(deviceSN, now.Unix()); err == nil {
								if len(alarms) > 0 {
									if latestAlert := alarms[0]; latestAlert != nil {
										if startTime := latestAlert.StartTime; startTime != nil {
											if pointy.StringValue(device.LastUpdateTime, "0000-00-00")[0:10] == (*startTime)[0:10] {
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
													PlantID:      pointy.String(stationIDStr),
													PlantName:    station.Name,
													Latitude:     plantItem.Latitude,
													Longitude:    plantItem.Longitude,
													Location:     plantItem.Location,
													DeviceID:     pointy.String(strconv.Itoa(deviceID)),
													DeviceSN:     device.DeviceSN,
													DeviceName:   device.DeviceSN,
													DeviceType:   pointy.String(growatt.ParseGrowattDeviceType(growatt.GROWATT_DEVICE_TYPE_PBD)),
													DeviceStatus: deviceItem.Status,
													ID:           pointy.String(strconv.Itoa(pointy.IntValue(latestAlert.AlarmCode, 0))),
													Message:      latestAlert.AlarmMessage,
													Owner:        credential.Owner,
												}

												if latestAlert.StartTime != nil {
													if parsed, err := time.Parse("2006-01-02 15:04:05.0", *latestAlert.StartTime); err == nil {
														utcParsed := parsed.UTC()
														alarmItemDoc.AlarmTime = &utcParsed
													}
												}

												documentCh <- alarmItemDoc
											}
										}
									}
								}
							}
						}

					}
				default:
				}

				if deviceItem.Status != nil {
					deviceStatusArray = append(deviceStatusArray, *deviceItem.Status)
				}

				documentCh <- deviceItem

				if deviceTypeRaw == growatt.GROWATT_DEVICE_TYPE_INVERTER {
					inverterCh <- deviceSN
				}
			}

			plantStatus := growatt.GROWATT_PLANT_STATUS_ON
			if len(deviceStatusArray) > 0 {
				var offlineCount int
				var alertingCount int

				for _, status := range deviceStatusArray {
					switch status {
					case growatt.GROWATT_DEVICE_STATUS_OFFLINE:
						offlineCount++
					case growatt.GROWATT_DEVICE_STATUS_ONLINE:
					default:
						alertingCount++
					}
				}

				if alertingCount > 0 {
					plantStatus = growatt.GROWATT_PLANT_STATUS_ALARM
				} else if offlineCount > 0 {
					plantStatus = growatt.GROWATT_PLANT_STATUS_OFF
				}
			} else {
				plantStatus = growatt.GROWATT_PLANT_STATUS_OFF
			}

			plantDeviceStatusCh <- map[string]string{stationIDStr: plantStatus}
		}

		wg.Go(producer)
	}

	if r := wg.WaitAndRecover(); r != nil {
		s.logger.Errorf("[%v] - GrowattCollectorService.run(): \ncaller: %v, \nvalue: %v \nstring: %v", credential.Username, r.Callers, r.Value, r.String())
		return
	}

	doneCh <- true
}

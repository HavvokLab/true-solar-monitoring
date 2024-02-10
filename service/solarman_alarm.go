package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/HavvokLab/true-solar-monitoring/config"
	"github.com/HavvokLab/true-solar-monitoring/constant"
	"github.com/HavvokLab/true-solar-monitoring/inverter/solarman"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/HavvokLab/true-solar-monitoring/repo"
	"github.com/HavvokLab/true-solar-monitoring/util"
	"github.com/go-redis/redis/v8"
)

type SolarmanAlarmService interface {
	Run(*model.SolarmanCredential) error
}

type solarmanAlarmService struct {
	brand       string
	solarRepo   repo.SolarRepo
	snmpRepo    repo.SnmpRepo
	redisClient *redis.Client
	logger      logger.Logger
}

func NewSolarmanAlarmService(solarRepo repo.SolarRepo, snmpRepo repo.SnmpRepo, redisClient *redis.Client, logger logger.Logger) SolarmanAlarmService {
	const brand = "INVT-Ipanda"
	return &solarmanAlarmService{
		brand:       brand,
		solarRepo:   solarRepo,
		snmpRepo:    snmpRepo,
		redisClient: redisClient,
		logger:      logger,
	}
}

func (s *solarmanAlarmService) Run(credential *model.SolarmanCredential) error {
	defer func() {
		if r := recover(); r != nil {
			s.logger.Warnf("[%v] - SolarmanAlarmService.Run(): %v", credential.Username, r)
		}
	}()

	ctx := context.Background()
	now := time.Now()
	beginningOfDay := time.Date(now.Year(), now.Month(), now.Day(), 6, 0, 0, 0, time.Local)
	documents := make([]interface{}, 0)

	if credential == nil {
		s.logger.Errorf("credential should not be empty")
		return errors.New("credential should not be empty")
	}

	client, err := solarman.NewSolarmanClient(&solarman.SolarmanCredential{
		Username:  credential.Username,
		Password:  credential.Password,
		AppID:     credential.AppID,
		AppSecret: credential.AppSecret,
	})

	if err != nil {
		s.logger.Error(err)
		return err
	}

	userInfoResp, err := client.GetUserInfo()
	if err != nil {
		s.logger.Error(err)
		return err
	}

	if userInfoResp.OrgInfoList == nil {
		s.logger.Errorf("organization should not be empty")
		return errors.New("organization should not be empty")
	}

	companyCount := 1
	companyTotal := len(userInfoResp.OrgInfoList)
	for _, company := range userInfoResp.OrgInfoList {
		s.logger.Infof("%v-COMPANY:[%v/%v]", credential.Username, companyCount, companyTotal)
		companyCount++

		tokenResp, err := client.GetBusinessToken(company.GetCompanyID())
		if err != nil {
			s.logger.Error(err)
			return err
		}

		if tokenResp.AccessToken == nil {
			s.logger.Errorf("accesstoken should not be empty")
			return errors.New("accesstoken should not be empty")
		}
		token := tokenResp.GetAccessToken()

		plantList, err := client.GetPlantList(token)
		if err != nil {
			s.logger.Error(err)
			return err
		}

		plantCount := 1
		plantTotal := len(plantList)
		for _, plant := range plantList {
			s.logger.Infof("%v-PLANT:[%v/%v]", credential.Username, plantCount, plantTotal)
			plantCount++

			stationID := plant.GetID()
			stationName := plant.GetName()

			deviceList, err := client.GetPlantDeviceList(token, stationID)
			if err != nil {
				s.logger.Error(err)
				return err
			}

			deviceCount := 1
			deviceTotal := len(deviceList)
			for _, device := range deviceList {
				s.logger.Infof("%v-DEVICE:[%v/%v]", credential.Username, deviceCount, deviceTotal)
				deviceCount++

				deviceID := device.GetDeviceID()
				deviceSN := device.GetDeviceSN()
				deviceType := device.GetDeviceType()
				deviceCollectionTime := device.GetCollectionTime()
				deviceCollectionTimeStr := strconv.FormatInt(deviceCollectionTime, 10)

				if device.ConnectStatus != nil {
					var document interface{}
					switch device.GetConnectStatus() {
					case 0:
						rkey := fmt.Sprintf("%s,%d,%s,%s,%d,%s", s.brand, stationID, deviceType, deviceSN, deviceID, "Disconnect")
						val := fmt.Sprintf("%s,%s", stationName, deviceCollectionTimeStr)

						err := s.redisClient.Set(ctx, rkey, val, 0).Err()
						if err != nil {
							s.logger.Error(err)
							return err
						}

						name := fmt.Sprintf("%s-%s", stationName, deviceSN)
						alert := strings.ReplaceAll(fmt.Sprintf("%s-%s", deviceType, "Disconnect"), " ", "-")
						description := fmt.Sprintf("%s,%d,%s,%d", s.brand, stationID, deviceSN, deviceID)
						document = model.NewSnmpAlarmItem(s.brand, name, alert, description, constant.MAJOR_SEVERITY, deviceCollectionTimeStr)
						err = s.snmpRepo.SendAlarmTrap(name, alert, description, constant.MAJOR_SEVERITY, deviceCollectionTimeStr)
						if err != nil {
							s.logger.Error(err)
							return err
						}

						s.logger.Infof("SendAlarmTrap: %s, %s, %s, %s", name, alert, description, constant.MAJOR_SEVERITY)
					case 1:
						var keys []string
						var cursor uint64

						for {
							var scanKeys []string
							match := fmt.Sprintf("%s,%d,%s,%s,%d,*", s.brand, stationID, deviceType, deviceSN, deviceID)
							scanKeys, cursor, err = s.redisClient.Scan(ctx, cursor, match, 10).Result()
							if err != nil {
								s.logger.Error(err)
								return err
							}

							keys = append(keys, scanKeys...)
							if cursor == 0 {
								break
							}
						}

						for _, key := range keys {
							val, err := s.redisClient.Get(ctx, key).Result()
							if err == redis.Nil {
								s.logger.Warnf("%v", err)
								continue
							}

							if err != nil {
								s.logger.Error(err)
								return err
							}

							if !util.EmptyString(val) {
								splitKey := strings.Split(key, ",")
								splitVal := strings.Split(val, ",")

								name := fmt.Sprintf("%s-%s", stationName, deviceSN)
								alert := strings.ReplaceAll(fmt.Sprintf("%s-%s", deviceType, splitKey[5]), " ", "-")
								description := fmt.Sprintf("%s,%d,%s,%d", s.brand, stationID, deviceSN, deviceID)
								document = model.NewSnmpAlarmItem(s.brand, name, alert, description, constant.CLEAR_SEVERITY, splitVal[1])
								if err := s.snmpRepo.SendAlarmTrap(name, alert, description, constant.CLEAR_SEVERITY, splitVal[1]); err != nil {
									s.logger.Error(err)
									return err
								}

								s.logger.Infof("SendAlarmTrap: %s, %s, %s, %s", name, alert, description, constant.CLEAR_SEVERITY)
							}

							if err := s.redisClient.Del(ctx, key).Err(); err != nil {
								s.logger.Error(err)
								return err
							}
						}
					case 2:
						alertList, err := client.GetDeviceAlertList(token, deviceSN, beginningOfDay.Unix(), now.Unix())
						if err != nil {
							s.logger.Error(err)
							return err
						}

						for _, alert := range alertList {
							alertName := alert.GetAlertNameInPAAS()
							alertTime := alert.GetAlertTime()
							alertTimeStr := strconv.FormatInt(alertTime, 10)

							if alert.AlertNameInPAAS != nil && alert.AlertTime != nil {
								rkey := fmt.Sprintf("%s,%d,%s,%s,%d,%s", s.brand, stationID, deviceType, deviceSN, deviceID, alertName)
								val := fmt.Sprintf("%s,%s", stationName, alertTimeStr)

								if err := s.redisClient.Set(ctx, rkey, val, 0).Err(); err != nil {
									s.logger.Error(err)
									return err
								}

								name := fmt.Sprintf("%s-%s", stationName, deviceSN)
								alert := strings.ReplaceAll(fmt.Sprintf("%s-%s", deviceType, alertName), " ", "-")
								description := fmt.Sprintf("%s,%d,%s,%d", s.brand, stationID, deviceSN, deviceID)
								document = model.NewSnmpAlarmItem(s.brand, name, alert, description, constant.MAJOR_SEVERITY, alertTimeStr)
								if err := s.snmpRepo.SendAlarmTrap(name, alert, description, constant.MAJOR_SEVERITY, alertTimeStr); err != nil {
									s.logger.Error(err)
									return err
								}

								// ILeek Fail
								s.logger.Infof("SendAlarmTrap: %s, %s, %s, %s", name, alert, description, constant.MAJOR_SEVERITY)
							}
						}
					default:
					}
					documents = append(documents, document)
				}
			}
		}
	}

	elasticCfg := config.GetConfig().Elastic
	index := fmt.Sprintf("%s-%s", elasticCfg.AlarmIndex, time.Now().Format("2006.01.02"))
	if err := s.solarRepo.BulkIndex(index, documents); err != nil {
		s.logger.Error(err)
		return err
	}
	s.logger.Infof("SolarmanAlarm(): saved %v alarms", len(documents))

	return nil
}

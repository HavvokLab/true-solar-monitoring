package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/HavvokLab/true-solar-monitoring/config"
	"github.com/HavvokLab/true-solar-monitoring/constant"
	"github.com/HavvokLab/true-solar-monitoring/inverter/kstar"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/HavvokLab/true-solar-monitoring/repo"
	"github.com/HavvokLab/true-solar-monitoring/util"
	"github.com/go-redis/redis/v8"
)

type KStarAlarmService interface {
	Run(*model.KStarCredential) error
}

type kstarAlarmService struct {
	vendorType  string
	solarRepo   repo.SolarRepo
	snmpRepo    repo.SnmpRepo
	redisClient *redis.Client
	logger      logger.Logger
}

func NewKStarAlarmService(
	solarRepo repo.SolarRepo,
	snmpRepo repo.SnmpRepo,
	redisClient *redis.Client,
	logger logger.Logger,
) KStarAlarmService {
	return &kstarAlarmService{
		vendorType:  strings.ToUpper(constant.VENDOR_TYPE_KSTAR),
		solarRepo:   solarRepo,
		snmpRepo:    snmpRepo,
		redisClient: redisClient,
		logger:      logger,
	}
}

func (s *kstarAlarmService) Run(credential *model.KStarCredential) error {
	defer func() {
		if r := recover(); r != nil {
			s.logger.Warnf("[%v] - KStarAlarmService.Run(): %v", credential.Username, r)
		}
	}()

	ctx := context.Background()
	client, err := kstar.NewKStarClient(&kstar.KStarCredential{
		Username: credential.Username,
		Password: credential.Password,
	})

	if err != nil {
		s.logger.Errorf("[%v] - KStarAlarmService.Run(): %v", credential.Username, err)
		return err
	}

	deviceList, err := client.GetDeviceList()
	if err != nil {
		s.logger.Errorf("[%v] - KStarAlarmService.Run(): %v", credential.Username, err)
		return err
	}

	if deviceList == nil {
		s.logger.Errorf("[%v] - KStarAlarmService.Run(): deviceList is nil", credential.Username)
		return errors.New("empty deviceList")
	}

	deviceCount := 1
	deviceSize := len(deviceList)
	documents := make([]interface{}, 0)
	for _, device := range deviceList {
		deviceID := device.GetID()
		deviceName := device.GetName()
		plantID := device.GetPlantID()
		plantName := device.GetPlantName()
		saveTime := device.GetSaveTime()

		s.logger.Infof("[%v] - KStarAlarmService.Run(): device %v/%v, deviceID: %v, deviceName: %v, plantID: %v, plantName: %v, saveTime: %v", credential.Username, deviceCount, deviceSize, deviceID, deviceName, plantID, plantName, saveTime)
		deviceCount++

		realtimeDeviceDataResp, err := client.GetRealtimeDeviceData(deviceID)
		if err != nil {
			s.logger.Errorf("[%v] - KStarAlarmService.Run(): %v", credential.Username, err)
			return err
		}

		if realtimeDeviceDataResp == nil {
			s.logger.Warnf("[%v] - KStarAlarmService.Run(): realtimeDeviceDataResp is nil", credential.Username)
			continue
		}

		if realtimeDeviceDataResp.Data != nil {
			saveTime = realtimeDeviceDataResp.Data.GetSaveTime()
		}

		if device.Status != nil {
			var document interface{}
			switch *device.Status {
			case 0:
				key := fmt.Sprintf("Kstar,%s,%s,%s,%s", plantID, deviceID, deviceName, "Kstar-Disconnect")
				val := fmt.Sprintf("%s,%s", plantName, saveTime)
				if err := s.redisClient.Set(ctx, key, val, 0).Err(); err != nil {
					s.logger.Errorf("[%v] - KStarAlarmService.Run(): %v", credential.Username, err)
					return err
				}

				alarmName := "Kstar-Disconnect"
				payload := fmt.Sprintf("Kstar,%s,%s,%s", plantID, deviceID, deviceName)
				document = model.NewSnmpAlarmItem(s.vendorType, plantName, alarmName, payload, constant.MAJOR_SEVERITY, saveTime)
				if err := s.snmpRepo.SendAlarmTrap(plantName, alarmName, payload, constant.MAJOR_SEVERITY, saveTime); err != nil {
					s.logger.Errorf("[%v] - KStarAlarmService.Run(): %v", credential.Username, err)
					return err
				}

				documents = append(documents, document)
				s.logger.Infof("[%v] - SendAlarmTrap(): plant: %v, alarm: %v, payload: %v, severity: %v, lastedUpdatedTime: %v", credential.Username, plantName, alarmName, payload, constant.MAJOR_SEVERITY, saveTime)
			case 1:
				realtimeAlarmResp, err := client.GetRealtimeAlarmListOfDevice(deviceID)
				if err != nil {
					s.logger.Errorf("[%v] - KStarAlarmService.Run(): %v", credential.Username, err)
					return err
				}

				if len(realtimeAlarmResp.Data) > 0 {
					alarmName := "Kstar-Disconnect"
					payload := fmt.Sprintf("Kstar,%s,%s,%s", plantID, deviceID, deviceName)
					document = model.NewSnmpAlarmItem(s.vendorType, plantName, alarmName, payload, constant.MAJOR_SEVERITY, saveTime)
					if err := s.snmpRepo.SendAlarmTrap(plantName, alarmName, payload, constant.CLEAR_SEVERITY, saveTime); err != nil {
						s.logger.Errorf("[%v] - KStarAlarmService.Run(): %v", credential.Username, err)
						return err
					}
					documents = append(documents, document)

					if err := s.redisClient.Del(ctx, alarmName).Err(); err != nil {
						s.logger.Errorf("[%v] - KStarAlarmService.Run(): %v", credential.Username, err)
						return err
					}

					for _, alarm := range realtimeAlarmResp.Data {
						alarmTime := alarm.GetSaveTime()
						alarmMessage := strings.ReplaceAll(alarm.GetMessage(), " ", "-")

						key := fmt.Sprintf("Kstar,%s,%s,%s,%s", plantID, deviceID, deviceName, alarmMessage)
						val := fmt.Sprintf("%s,%s", plantName, alarmTime)
						if err := s.redisClient.Set(ctx, key, val, 0).Err(); err != nil {
							s.logger.Errorf("[%v] - KStarAlarmService.Run(): %v", credential.Username, err)
							return err
						}

						payload := fmt.Sprintf("Kstar,%s,%s,%s", plantID, deviceID, deviceName)
						document = model.NewSnmpAlarmItem(s.vendorType, plantName, alarmName, payload, constant.MAJOR_SEVERITY, saveTime)
						if err := s.snmpRepo.SendAlarmTrap(plantName, alarmMessage, payload, constant.MAJOR_SEVERITY, alarmTime); err != nil {
							s.logger.Errorf("[%v] - KStarAlarmService.Run(): %v", credential.Username, err)
							return err
						}

						documents = append(documents, document)
						s.logger.Infof("[%v] - SendAlarmTrap(): plant: %v, alarm: %v, payload: %v, severity: %v, lastedUpdatedTime: %v", credential.Username, plantName, alarmMessage, payload, constant.MAJOR_SEVERITY, alarmTime)
					}
					continue
				}

				var keys []string
				var cursor uint64
				for {
					var scanKeys []string
					match := fmt.Sprintf("Kstar,%s,%s,%s,*", plantID, deviceID, deviceName)
					scanKeys, cursor, err = s.redisClient.Scan(ctx, cursor, match, 10).Result()
					if err != nil {
						s.logger.Errorf("[%v] - KStarAlarmService.Run(): %v", credential.Username, err)
						return err
					}

					keys = append(keys, scanKeys...)
					if cursor == 0 {
						break
					}
				}

				for _, key := range keys {
					val, err := s.redisClient.Get(ctx, key).Result()
					if err != nil {
						s.logger.Errorf("[%v] - KStarAlarmService.Run(): %v", credential.Username, err)
						if err != redis.Nil {
							return err
						}
						continue
					}

					if !util.EmptyString(val) {
						splitKey := strings.Split(key, ",")
						splitVal := strings.Split(val, ",")

						plantName := splitVal[0]
						alarmName := strings.ReplaceAll(splitKey[4], " ", "-")
						payload := fmt.Sprintf("Kstar,%s,%s,%s", plantID, deviceID, deviceName)
						document = model.NewSnmpAlarmItem(s.vendorType, plantName, alarmName, payload, constant.CLEAR_SEVERITY, splitVal[1])
						if err := s.snmpRepo.SendAlarmTrap(plantName, alarmName, payload, constant.CLEAR_SEVERITY, splitVal[1]); err != nil {
							s.logger.Errorf("[%v] - KStarAlarmService.Run(): %v", credential.Username, err)
							return err
						}

						s.logger.Infof("[%v] - SendAlarmTrap(): plant: %v, alarm: %v, payload: %v, severity: %v, lastedUpdatedTime: %v", credential.Username, plantName, alarmName, payload, constant.CLEAR_SEVERITY, splitVal[1])

						if err := s.redisClient.Del(ctx, key).Err(); err != nil {
							s.logger.Errorf("[%v] - KStarAlarmService.Run(): %v", credential.Username, err)
							return err
						}
					}
				}
			case 2:
				realtimeAlarmResp, err := client.GetRealtimeAlarmListOfDevice(deviceID)
				if err != nil {
					s.logger.Errorf("[%v] - KStarAlarmService.Run(): %v", credential.Username, err)
					return err
				}

				if len(realtimeAlarmResp.Data) > 0 {
					for _, alarm := range realtimeAlarmResp.Data {
						alarmTime := alarm.GetSaveTime()
						alarmMessage := strings.ReplaceAll(alarm.GetMessage(), " ", "-")

						key := fmt.Sprintf("Kstar,%s,%s,%s,%s", plantID, deviceID, deviceName, alarmMessage)
						val := fmt.Sprintf("%s,%s", plantName, alarmTime)
						if err := s.redisClient.Set(ctx, key, val, 0).Err(); err != nil {
							s.logger.Errorf("[%v] - KStarAlarmService.Run(): %v", credential.Username, err)
							return err
						}

						payload := fmt.Sprintf("Kstar,%s,%s,%s", plantID, deviceID, deviceName)
						document = model.NewSnmpAlarmItem(s.vendorType, plantName, alarmMessage, payload, constant.MAJOR_SEVERITY, alarmTime)
						if err := s.snmpRepo.SendAlarmTrap(plantName, alarmMessage, payload, constant.MAJOR_SEVERITY, alarmTime); err != nil {
							s.logger.Errorf("[%v] - KStarAlarmService.Run(): %v", credential.Username, err)
							return err
						}
						documents = append(documents, document)
					}
				}
			default:
			}
		}
	}

	elasticCfg := config.GetConfig().Elastic
	index := fmt.Sprintf("%s-%s", elasticCfg.AlarmIndex, time.Now().Format("2006.01.02"))
	if err := s.solarRepo.BulkIndex(index, documents); err != nil {
		s.logger.Error(err)
		return err
	}
	s.logger.Infof("GrowattAlarm(): saved %v alarms", len(documents))

	return nil
}

package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/HavvokLab/true-solar-monitoring/constant"
	"github.com/HavvokLab/true-solar-monitoring/inverter/growatt"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/HavvokLab/true-solar-monitoring/repo"
	"github.com/HavvokLab/true-solar-monitoring/util"
	"github.com/go-redis/redis/v8"
)

type GrowattAlarmService interface {
	Run(*model.GrowattCredential) error
}

type growattAlarmService struct {
	snmpRepo    repo.SnmpRepo
	redisClient *redis.Client
	logger      logger.Logger
}

func NewGrowattAlarmService(
	snmpRepo repo.SnmpRepo,
	redisClient *redis.Client,
	logger logger.Logger,
) GrowattAlarmService {
	return &growattAlarmService{
		snmpRepo:    snmpRepo,
		redisClient: redisClient,
		logger:      logger,
	}
}

func (s *growattAlarmService) Run(credential *model.GrowattCredential) error {
	now := time.Now()
	ctx := context.Background()
	client, err := growatt.NewGrowattClient(&growatt.GrowattCredential{Username: credential.Username, Token: credential.Token})
	if err != nil {
		s.logger.Errorf("[%v] Failed to create growatt client: %v", credential.Username, err)
		return err
	}

	plants, err := client.GetPlantList()
	if err != nil {
		s.logger.Errorf("[%v] Failed to get plant list: %v", credential.Username, err)
		return err
	}

	for _, plant := range plants {
		plantID := plant.GetPlantID()
		plantName := plant.GetName()

		devices, err := client.GetPlantDeviceList(plant.GetPlantID())
		if err != nil {
			s.logger.Errorf("[%v] Failed to get plant device list: %v", credential.Username, err)
			continue
		}

		for _, device := range devices {
			deviceSN := device.GetDeviceSN()
			deviceModel := device.GetModel()
			deviceLastUpdateTime := device.GetLastUpdateTime()
			deviceStatus := growatt.GROWATT_INVERTER_STATUS_MAPPER[device.GetStatus()]
			deviceType := growatt.GROWATT_EQUIP_TYPE_MAPPER[device.GetType()]
			deviceName := fmt.Sprintf("%s_%d_%s", plantName, plantID, deviceSN)

			switch deviceStatus {
			case "Online":
				key := fmt.Sprintf("%d,%s,%s,%s", plantID, plantName, deviceType, deviceSN)
				val, err := s.redisClient.Get(ctx, key).Result()
				if err != nil {
					s.logger.Errorf("[%v] Failed to get redis key: %v", credential.Username, err)
					continue
				}

				if !util.EmptyString(val) {
					vals := strings.Split(val, ",")
					alarmName := fmt.Sprintf("Growatt,%s,%s", vals[1], deviceModel)
					payload := fmt.Sprintf("%s-Error-%s", alarmName, vals[0])
					severity := constant.CLEAR_SEVERITY
					if err := s.snmpRepo.SendAlarmTrap(deviceName, payload, alarmName, severity, deviceLastUpdateTime); err != nil {
						s.logger.Errorf("[%v] Failed to send alarm trap: %v", credential.Username, err)
						continue
					}

					s.logger.Infof("[%v] - SendAlarmTrap(): plant: %v, alarm: %v, payload: %v, severity: %v, lastedUpdatedTime: %v", credential.Username, deviceName, payload, alarmName, severity, deviceLastUpdateTime)
				}

				if err := s.redisClient.Del(ctx, key).Err(); err != nil {
					s.logger.Errorf("[%v] Failed to delete redis key: %v", credential.Username, err)
					continue
				}

			case "Disconnect":
				rkey := fmt.Sprintf("%d,%s,%s,%s", plantID, plantName, deviceType, deviceSN)
				val := "0,Disconnect"
				if err := s.redisClient.Set(ctx, rkey, val, 0).Err(); err != nil {
					s.logger.Error(err)
					continue
				}

				alarmName := fmt.Sprintf("Growatt,Disconnect,%s", deviceModel)
				payload := fmt.Sprintf("%s-Error-0", deviceType)
				severity := "4"
				if err := s.snmpRepo.SendAlarmTrap(deviceName, payload, alarmName, severity, deviceLastUpdateTime); err != nil {
					s.logger.Errorf("snmp.SendAlarmTrap(%v,%v,%v,%v,%v): %v", deviceName, payload, alarmName, severity, deviceLastUpdateTime, err)
					continue
				}

				s.logger.Infof("[%v] - SendAlarmTrap(): plant: %v, alarm: %v, payload: %v, severity: %v, lastedUpdatedTime: %v", credential.Username, deviceName, payload, alarmName, severity, deviceLastUpdateTime)
			default:
				date := now.AddDate(0, 0, -1).Format("2006-01-02")
				alarms, err := client.GetInverterAlertList(deviceSN)
				if err != nil {
					s.logger.Errorf("[%v] Failed to get inverter alert list: %v", credential.Username, err)
					continue
				}

				if len(alarms) > 0 {
					alarm := alarms[0]
					rkey := fmt.Sprintf("%d,%s,%s,%s", plantID, plantName, deviceType, deviceSN)
					val := fmt.Sprintf("%d,%s", alarm.GetAlarmCode(), alarm.GetAlarmMessage())
					if err := s.redisClient.Set(ctx, rkey, val, 0).Err(); err != nil {
						s.logger.Errorf("[%v] Failed to set redis key: %v", credential.Username, err)
						continue
					}

					alarmName := fmt.Sprintf("Growatt,%s,%s", alarm.GetAlarmMessage(), deviceModel)
					payload := fmt.Sprintf("%s-Error-%d", deviceType, alarm.GetAlarmCode())
					severity := constant.MAJOR_SEVERITY
					if err := s.snmpRepo.SendAlarmTrap(deviceName, payload, alarmName, severity, date); err != nil {
						s.logger.Errorf("[%v] Failed to send alarm trap: %v", credential.Username, err)
						continue
					}

					s.logger.Infof("[%v] - SendAlarmTrap(): plant: %v, alarm: %v, payload: %v, severity: %v, lastedUpdatedTime: %v", credential.Username, deviceName, alarmName, payload, severity, date)
				}
			}
		}

		time.Sleep(10 * time.Second)
	}

	return nil
}

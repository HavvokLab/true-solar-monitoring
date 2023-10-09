package service

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/HavvokLab/true-solar-monitoring/constant"
	"github.com/HavvokLab/true-solar-monitoring/inverter/huawei"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/HavvokLab/true-solar-monitoring/repo"
	"github.com/HavvokLab/true-solar-monitoring/util"
	"github.com/go-redis/redis/v8"
)

type HuaweiAlarmService interface {
	Run(*model.HuaweiCredential) error
}

type huaweiAlarmService struct {
	snmpRepo    repo.SnmpRepo
	redisClient *redis.Client
	logger      logger.Logger
}

func NewHuaweiAlarmService(
	snmpRepo repo.SnmpRepo,
	redisClient *redis.Client,
	logger logger.Logger,
) HuaweiAlarmService {
	return &huaweiAlarmService{
		snmpRepo:    snmpRepo,
		redisClient: redisClient,
		logger:      logger,
	}
}

func (s *huaweiAlarmService) Run(credential *model.HuaweiCredential) error {
	now := time.Now()
	ctx := context.Background()
	beginTime := time.Date(now.Year(), now.Month(), now.Day(), 6, 0, 0, 0, time.Local).UnixNano() / 1e6
	endTime := now.UnixNano() / 1e6

	client, err := huawei.NewHuaweiClient(&huawei.HuaweiCredential{
		Username: credential.Username,
		Password: credential.Password,
	})

	if err != nil {
		s.logger.Error(err)
		return err
	}

	plantListResp, err := client.GetPlantList()
	if err != nil {
		s.logger.Error(err)
		return err
	}
	s.logger.Infof("[%v] - huaweiAlarmService.Run(): plant size: %v", credential.Username, len(plantListResp.Data))
	s.logger.Infof("[%v] - huaweiAlarmService.Run(): start plant preparation", credential.Username)

	var stationCodeList []string
	var stationCodeListString []string
	for _, plant := range plantListResp.Data {
		if len(stationCodeList) == 100 {
			stationCodeListString = append(stationCodeListString, strings.Join(stationCodeList, ","))
			stationCodeList = []string{}
		}

		if plant.Code != nil {
			stationCodeList = append(stationCodeList, plant.GetCode())
		}
	}
	stationCodeListString = append(stationCodeListString, strings.Join(stationCodeList, ","))

	var inverterList []huawei.DeviceItem
	mapPlantCodeToDevice := make(map[string][]huawei.DeviceItem)
	mapDeviceSNToAlarm := make(map[string][]huawei.DeviceAlarmItem)
	mapInverterIDToRealtimeData := make(map[int]huawei.RealtimeDeviceData)
	for _, stationCode := range stationCodeListString {
		deviceListResp, err := client.GetDeviceList(stationCode)
		if err != nil {
			s.logger.Error(err)
			return err
		}

		for _, device := range deviceListResp.Data {
			if device.PlantCode != nil {
				mapPlantCodeToDevice[device.GetPlantCode()] = append(mapPlantCodeToDevice[device.GetPlantCode()], *device)
			}

			if device.GetTypeID() == 1 {
				inverterList = append(inverterList, *device)
			}
		}

		deviceAlarmListResp, err := client.GetDeviceAlarmList(stationCode, beginTime, endTime)
		if err != nil {
			s.logger.Error(err)
			return err
		}

		for _, alarm := range deviceAlarmListResp.Data {
			doubleAlarm := false

			if alarm.DeviceSN != nil {
				for i, deviceAlarm := range mapDeviceSNToAlarm[alarm.GetDeviceSN()] {
					if deviceAlarm.GetAlarmName() == alarm.GetAlarmName() {
						doubleAlarm = true

						if deviceAlarm.GetRaiseTime() < alarm.GetRaiseTime() {
							mapDeviceSNToAlarm[alarm.GetDeviceSN()][i] = *alarm
							break
						}
					}
				}

				if !doubleAlarm {
					mapDeviceSNToAlarm[alarm.GetDeviceSN()] = append(mapDeviceSNToAlarm[alarm.GetDeviceSN()], *alarm)
				}
			}
		}
	}

	s.logger.Infof("[%v] - huaweiAlarmService.Run(): start device preparation", credential.Username)
	var inverterIDList []string
	var inverterIDListString []string
	for _, device := range inverterList {
		if len(inverterIDList) == 100 {
			inverterIDListString = append(inverterIDListString, strings.Join(inverterIDList, ","))
			inverterIDList = []string{}
		}

		if device.ID != nil {
			inverterIDList = append(inverterIDList, strconv.Itoa(device.GetID()))
		}
	}
	inverterIDListString = append(inverterIDListString, strings.Join(inverterIDList, ","))

	for _, inverterID := range inverterIDListString {
		realtimeDeviceResp, err := client.GetRealtimeDeviceData(inverterID, "1")
		if err != nil {
			s.logger.Error(err)
			return err
		}

		for _, realtimeDevice := range realtimeDeviceResp.Data {
			if realtimeDevice.ID != nil {
				mapInverterIDToRealtimeData[realtimeDevice.GetID()] = *realtimeDevice
			}
		}
	}

	plantCount := 1
	plantSize := len(plantListResp.Data)
	for _, plant := range plantListResp.Data {
		plantCode := plant.GetCode()
		plantName := plant.GetName()

		s.logger.Infof("[%v] - huaweiAlarmService.Run(): start plant %v/%v", credential.Username, plantCount, plantSize)
		plantCount++

		deviceCount := 1
		deviceSize := len(mapPlantCodeToDevice[plantCode])
		for _, device := range mapPlantCodeToDevice[plantCode] {
			deviceID := device.GetID()
			deviceSN := device.GetSN()
			deviceName := device.GetName()

			s.logger.Infof("[%v] - huaweiAlarmService.Run(): start device %v/%v", credential.Username, deviceCount, deviceSize)
			deviceCount++

			if device.GetTypeID() == 1 {
				realtimeDevice := mapInverterIDToRealtimeData[deviceID].DataItemMap
				if realtimeDevice == nil {
					s.logger.Warnf("realtimeDevice is nil, deviceID: %d", deviceID)
					continue
				}

				if realtimeDevice.GetStatus(10) == 0 {
					shutdownTime := strconv.Itoa(int(endTime))
					if mapInverterIDToRealtimeData[deviceID].DataItemMap != nil {

						if mapInverterIDToRealtimeData[deviceID].DataItemMap.InverterShutdown != nil {
							inverterShutdown, ok := realtimeDevice.GetInverterShutdown().(float64)
							if ok {
								shutdownTime = strconv.Itoa(int(inverterShutdown))
							}
						}
					}

					key := fmt.Sprintf("Huawei,%s,%s,%s,%s", plantCode, deviceSN, deviceName, "Disconnect")
					val := fmt.Sprintf("%s,%s,%s", plantName, "Disconnect", shutdownTime)
					err := s.redisClient.Set(ctx, key, val, 0).Err()
					if err != nil {
						s.logger.Error(err)
						return err
					}

					alarmName := fmt.Sprintf("HUW-%s", "Disconnect")
					payload := fmt.Sprintf("Huawei,%s,%s", deviceName, "Disconnect")
					if err := s.snmpRepo.SendAlarmTrap(plantName, alarmName, payload, constant.MAJOR_SEVERITY, shutdownTime); err != nil {
						s.logger.Error(err)
						return err
					}
					s.logger.Infof("send alarm trap, plantName: %s, alarmName: %s, payload: %s, severity: %s, shutdownTime: %s", plantName, alarmName, payload, constant.MAJOR_SEVERITY, shutdownTime)
					continue
				}
			}

			if len(mapDeviceSNToAlarm[deviceSN]) > 0 {
				for _, alarm := range mapDeviceSNToAlarm[deviceSN] {
					alarmName := alarm.GetAlarmName()
					alarmCause := alarm.GetAlarmCause()
					alarmTime := strconv.Itoa(int(alarm.GetRaiseTime()))

					key := fmt.Sprintf("Huawei,%s,%s,%s,%s", plantCode, deviceSN, deviceName, alarmName)
					val := fmt.Sprintf("%s,%s,%s", plantName, alarmCause, alarmTime)
					err := s.redisClient.Set(ctx, key, val, 0).Err()
					if err != nil {
						s.logger.Error(err)
						return err
					}

					alarmName = strings.ReplaceAll(fmt.Sprintf("HUW-%s", alarmName), " ", "-")
					payload := fmt.Sprintf("Huawei,%s,%s", deviceName, alarmCause)
					if err := s.snmpRepo.SendAlarmTrap(plantName, alarmName, payload, constant.MAJOR_SEVERITY, alarmTime); err != nil {
						s.logger.Error(err)
						return err
					}
					s.logger.Infof("send alarm trap, plantName: %s, alarmName: %s, payload: %s, severity: %s, alarmTime: %s", plantName, alarmName, payload, constant.MAJOR_SEVERITY, alarmTime)
				}

				continue
			}

			var keys []string
			var cursor uint64
			for {
				var scanKeys []string
				match := fmt.Sprintf("Huawei,%s,%s,%s,*", plantCode, deviceSN, deviceName)
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
				if err != nil {
					if err != redis.Nil {
						s.logger.Error(err)
						return err
					}
					continue
				}

				if !util.EmptyString(val) {
					splitKey := strings.Split(key, ",")
					splitVal := strings.Split(val, ",")

					alarmName := strings.ReplaceAll(fmt.Sprintf("HUW-%s", splitKey[4]), " ", "-")
					payload := fmt.Sprintf("Huawei,%s,%s", deviceName, splitVal[1])
					if err := s.snmpRepo.SendAlarmTrap(
						splitVal[0],
						alarmName,
						payload,
						constant.CLEAR_SEVERITY,
						splitVal[2],
					); err != nil {
						s.logger.Error(err)
						return err
					}

					s.logger.Infof("send alarm trap, plantName: %s, alarmName: %s, payload: %s, severity: %s, alarmTime: %s", splitVal[0], alarmName, payload, constant.CLEAR_SEVERITY, splitVal[2])

					if err := s.redisClient.Del(ctx, key).Err(); err != nil {
						s.logger.Error(err)
						return err
					}
				}
			}
		}
	}

	return nil
}

package service

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/HavvokLab/true-solar-monitoring/constant"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/HavvokLab/true-solar-monitoring/repo"
	"github.com/HavvokLab/true-solar-monitoring/util"
	"go.openly.dev/pointy"
)

type LowPerformanceAlarmService interface {
	Run() error
}

type lowPerformanceAlarmService struct {
	solarRepo                  repo.SolarRepo
	installedCapacityRepo      repo.InstalledCapacityRepo
	performanceAlarmConfigRepo repo.PerformanceAlarmConfigRepo
	snmpRepo                   repo.SnmpRepo
	logger                     logger.Logger
}

func NewLowPerformanceAlarmService(solarRepo repo.SolarRepo, installedCapacityRepo repo.InstalledCapacityRepo, performanceAlarmConfigRepo repo.PerformanceAlarmConfigRepo, snmpRepo repo.SnmpRepo, logger logger.Logger) LowPerformanceAlarmService {
	return &lowPerformanceAlarmService{
		solarRepo:                  solarRepo,
		installedCapacityRepo:      installedCapacityRepo,
		performanceAlarmConfigRepo: performanceAlarmConfigRepo,
		snmpRepo:                   snmpRepo,
		logger:                     logger,
	}
}

func (s *lowPerformanceAlarmService) Run() error {
	now := time.Now()
	installedCapacityConfig, err := s.getInstalledCapacity()
	if err != nil {
		return err
	}

	config, err := s.getConfig()
	if err != nil {
		return err
	}

	efficiencyFactor := installedCapacityConfig.EfficiencyFactor
	focusHour := installedCapacityConfig.FocusHour
	hitDay := *config.HitDay
	duration := *config.Duration
	percentage := config.Percentage / 100.0

	s.logger.Infof("Retrieving low performance alarm service with duration: %d, hit day: %d, percentage: %.2f%%, efficiency factor: %.2f, focus hour: %v", duration, hitDay, percentage*100.0, efficiencyFactor, focusHour)
	buckets, err := s.solarRepo.GetPerformanceLow(duration, efficiencyFactor, focusHour, percentage)
	if err != nil {
		s.logger.Error(err)
		return err
	}
	s.logger.Infof("Retrieved %d buckets", len(buckets))

	period := fmt.Sprintf("%s - %s", now.AddDate(0, 0, -duration).Format("02Jan2006"), now.AddDate(0, 0, -1).Format("02Jan2006"))
	filteredBuckets := make(map[string]map[string]interface{})
	for _, bucketPtr := range buckets {
		if bucketPtr != nil {
			bucket := *bucketPtr

			if len(bucket.Key) == 0 {
				continue
			}

			var plantItem *model.PlantItem
			var key string
			var installedCapacity float64

			if len(bucket.Key) > 0 {
				if vendorType, ok := bucket.Key["vendor_type"]; ok {
					if id, ok := bucket.Key["id"]; ok {
						key = fmt.Sprintf("%s_%s", vendorType, id)
					}
				}
			}

			if avgCapacity, ok := bucket.ValueCount("avg_capacity"); ok {
				installedCapacity = pointy.Float64Value(avgCapacity.Value, 0.0)
			}

			if topHits, found := bucket.Aggregations.TopHits("hits"); found {
				if topHits.Hits != nil {
					if len(topHits.Hits.Hits) == 1 {
						searchHitPtr := topHits.Hits.Hits[0]
						if searchHitPtr != nil {
							if err := util.Recast(searchHitPtr.Source, &plantItem); err != nil {
								s.logger.Warn(err.Error())
								continue
							}
						}
					}
				}
			}

			if !util.EmptyString(key) {
				if item, found := filteredBuckets[key]; found {
					if count, ok := item["count"].(int); ok {
						item["count"] = count + 1
					}
					filteredBuckets[key] = item
				} else {
					filteredBuckets[key] = map[string]interface{}{
						"count":             1,
						"installedCapacity": installedCapacity,
						"plantItem":         plantItem,
						"period":            period,
					}
				}
			}
		}
	}

	s.logger.Infof("start sending low performance alarm with %d plants", len(filteredBuckets))

	var alarmCount int
	var failedAlarmCount int
	if len(filteredBuckets) > 0 {
		bucketBatches := s.chunkBy(filteredBuckets, constant.PERFORMANCE_ALARM_SNMP_BATCH_SIZE)

		var batchAlarmCount int
		var failedBatchAlarmCount int
		for i, batches := range bucketBatches {
			batchAlarmCount = 0
			failedBatchAlarmCount = 0

			for _, batch := range batches {
				for _, data := range batch {
					if count, ok := data["count"].(int); ok {
						if count >= hitDay {
							plantName, alarmName, description, severity, err := s.getSNMPPayload(constant.PERFORMANCE_ALARM_TYPE_PERFORMANCE_LOW, *config, *installedCapacityConfig, data)
							if err != nil {
								s.logger.Error(err)
								continue
							}

							if err := s.snmpRepo.SendAlarmTrap(plantName, alarmName, description, constant.MAJOR_SEVERITY, now.Format(time.RFC3339Nano)); err != nil {
								failedAlarmCount++
								failedBatchAlarmCount++
								s.logger.Error(err)
								continue
							}

							s.logger.Infof("SendAlarmTrap: plantName: %s, alarmName: %s, description: %s, severity: %s", plantName, alarmName, description, severity)
							alarmCount++
							batchAlarmCount++
						}
					}
				}
			}

			s.logger.Infof("batch %v completed to send %d alarms", i+1, batchAlarmCount)
			s.logger.Infof("batch %v failed to send %d alarms", i+1, failedBatchAlarmCount)
			s.logger.Infof("batch %v sleeping for %.2fs", i+1, constant.PERFORMANCE_ALARM_SNMP_BATCH_DELAY.Seconds())
			time.Sleep(constant.PERFORMANCE_ALARM_SNMP_BATCH_DELAY)
		}

		s.logger.Infof("completed to send %d alarms", alarmCount)
		s.logger.Infof("failed to send %d alarms", failedAlarmCount)
		s.logger.Infof("polling finished in %v", time.Since(now).String())
	}

	return nil
}

func (s *lowPerformanceAlarmService) getConfig() (*model.PerformanceAlarmConfig, error) {
	config, err := s.performanceAlarmConfigRepo.GetLowPerformanceAlarmConfig()
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	if config == nil {
		err := errors.New("performance alarm config not found")
		s.logger.Error(err)
		return nil, err
	}

	if pointy.IntValue(config.HitDay, 0) == 0 {
		err := errors.New("hit day must not be zero value")
		s.logger.Error(err)
		return nil, err
	}

	if pointy.IntValue(config.Duration, 0) == 0 {
		err := errors.New("duration must not be zero value")
		s.logger.Error(err)
		return nil, err
	}

	return config, nil
}

func (s *lowPerformanceAlarmService) getInstalledCapacity() (*model.InstalledCapacity, error) {
	installedCapacity, err := s.installedCapacityRepo.FindOne()
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	if installedCapacity == nil {
		err := errors.New("installed capacity not found")
		s.logger.Error(err)
		return nil, err
	}

	return installedCapacity, nil
}

func (p *lowPerformanceAlarmService) getSNMPPayload(alarmType int, alarmConfig model.PerformanceAlarmConfig, capacityConfig model.InstalledCapacity, data map[string]interface{}) (string, string, string, string, error) {
	if alarmType != constant.PERFORMANCE_ALARM_TYPE_PERFORMANCE_LOW && alarmType != constant.PERFORMANCE_ALARM_TYPE_SUM_PERFORMANCE_LOW {
		return "", "", "", "", errors.New("alarm type must be 1 (PERFORMANCE_LOW) or 2 (SUM_PERFORMANCE_LOW)")
	}

	var capacity float64
	if cap, ok := data["installedCapacity"].(float64); ok {
		capacity = cap
	}

	var plantItem model.PlantItem
	if item, ok := data["plantItem"].(*model.PlantItem); ok {
		if item != nil {
			plantItem = *item
		}
	}

	var period string
	if p, ok := data["period"].(string); ok {
		period = p
	}

	var vendorName string
	switch strings.ToLower(plantItem.VendorType) {
	case constant.VENDOR_TYPE_GROWATT:
		vendorName = "Growatt"
	case constant.VENDOR_TYPE_HUAWEI:
		vendorName = "HUA"
	case constant.VENDOR_TYPE_KSTAR:
		vendorName = "Kstar"
	case constant.VENDOR_TYPE_INVT:
		vendorName = "INVT-Ipanda"
	case constant.VENDOR_TYPE_SOLARMAN: // Todo: Remove after change SOLARMAN to INVT in elasticsearch
		vendorName = "INVT-Ipanda"
	default:
		// no-op
	}

	if vendorName == "" {
		return "", "", "", "", fmt.Errorf("vendor type (%s) not supported", plantItem.VendorType)
	}

	plantName := pointy.StringValue(plantItem.Name, "")
	alarmName := fmt.Sprintf("SolarCell-%s", strings.ReplaceAll(alarmConfig.Name, " ", ""))
	alarmNameInDescription := util.AddSpace(alarmConfig.Name)
	severity := "5"
	duration := pointy.IntValue(alarmConfig.Duration, 0)
	hitDay := pointy.IntValue(alarmConfig.HitDay, 0)
	multipliedCapacity := capacity * capacityConfig.EfficiencyFactor * float64(capacityConfig.FocusHour)

	// PerformanceLow
	if alarmType == constant.PERFORMANCE_ALARM_TYPE_PERFORMANCE_LOW {
		severity := "3"
		payload := fmt.Sprintf("%s, %s, Less than or equal %.2f%%, Expected Daily Production:%.2f KWH, Actual Production less than:%.2f KWH, Duration:%d days, Period:%s",
			vendorName, alarmNameInDescription, alarmConfig.Percentage, multipliedCapacity, multipliedCapacity*(alarmConfig.Percentage/100.0), hitDay, period)
		return plantName, alarmName, payload, severity, nil
	}

	// SumPerformanceLow
	var totalProduction float64
	if x, ok := data["totalProduction"].(float64); ok {
		totalProduction = x
	}

	payload := fmt.Sprintf("%s, %s, Less than or equal %.2f%%, Expected Production:%.2f KWH, Actual Production:%.2f KWH (less than %.2f KWH), Duration:%d days, Period:%s",
		vendorName, alarmNameInDescription, alarmConfig.Percentage, multipliedCapacity*float64(duration), totalProduction, (multipliedCapacity*float64(duration))*(alarmConfig.Percentage/100.0), duration, period)
	return plantName, alarmName, payload, severity, nil
}

func (s *lowPerformanceAlarmService) chunkBy(items map[string]map[string]interface{}, chunkSize int) (chunks [][]map[string]map[string]interface{}) {
	slice := make([]map[string]map[string]interface{}, 0)

	for k, v := range items {
		slice = append(slice, map[string]map[string]interface{}{k: v})
	}

	for chunkSize < len(slice) {
		slice, chunks = slice[chunkSize:], append(chunks, slice[0:chunkSize:chunkSize])
	}

	return append(chunks, slice)
}

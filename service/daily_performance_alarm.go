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

var ErrDailyProductionUnderThreshold = errors.New("daily production under threshold")

type DailyPerformanceAlarmService interface {
	Run() error
}

type dailyPerformanceAlarmService struct {
	solarRepo                  repo.SolarRepo
	installedCapacityRepo      repo.InstalledCapacityRepo
	performanceAlarmConfigRepo repo.PerformanceAlarmConfigRepo
	snmpRepo                   repo.SnmpRepo
	logger                     logger.Logger
}

func NewDailyPerformanceAlarmService(
	solarRepo repo.SolarRepo,
	installedCapacityRepo repo.InstalledCapacityRepo,
	performanceAlarmConfigRepo repo.PerformanceAlarmConfigRepo,
	snmpRepo repo.SnmpRepo,
	logger logger.Logger,
) DailyPerformanceAlarmService {
	return &dailyPerformanceAlarmService{
		solarRepo:                  solarRepo,
		installedCapacityRepo:      installedCapacityRepo,
		performanceAlarmConfigRepo: performanceAlarmConfigRepo,
		snmpRepo:                   snmpRepo,
		logger:                     logger,
	}
}

func (s *dailyPerformanceAlarmService) Run() error {
	defer func() {
		if err := recover(); err != nil {
			s.logger.Errorf("DailyPerformanceAlarm.Run(): %v", err)
		}
	}()

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
	duration := *config.Duration
	percentage := config.Percentage / 100.0
	s.logger.Infof("Retrieving daily performance alarm service with duration: %d, percentage: %.2f%%, efficiency factor: %.2f, focus hour: %v", duration, percentage*100.0, efficiencyFactor, focusHour)

	date := now.AddDate(0, 0, -1)
	buckets, err := s.solarRepo.GetPerformanceByDate(&date, efficiencyFactor, focusHour, percentage)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	s.logger.Infof("Found %d buckets", len(buckets))
	var failAlarmCount int
	var clearAlarmCount int
	var count int = 1
	size := len(buckets)
	for _, bucketPtr := range buckets {
		s.logger.Infof("Processing bucket %d/%v", count, size)
		count++

		if bucketPtr == nil {
			continue
		}

		bucket := *bucketPtr
		if len(bucket.Key) > 0 {
			var dailyProduction float64
			dailyProductionValue, ok := bucket.ValueCount("daily")
			if !ok {
				continue
			} else {
				dailyProduction = pointy.Float64Value(dailyProductionValue.Value, 0.0)
			}

			var threshold float64
			thresholdValue, ok := bucket.ValueCount("threshold")
			if !ok {
				continue
			} else {
				threshold = pointy.Float64Value(thresholdValue.Value, 0.0)
			}

			var plantItem *model.PlantItem
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

			plantName, alarmName, payload, severity, err := s.getPayload(&date, dailyProduction, threshold, plantItem, *installedCapacityConfig, *config)
			if err != nil {
				if errors.Is(err, ErrDailyProductionUnderThreshold) {
					continue
				}

				s.logger.Error(err)
				continue
			}

			if err := s.snmpRepo.SendAlarmTrap(plantName, alarmName, payload, severity, now.Format(time.RFC3339Nano)); err != nil {
				s.logger.Error(err)
				continue
			}

			s.logger.Infof("SendAlarmTrap: %s, %s, %s, %s", plantName, alarmName, payload, severity)
			if severity == constant.MAJOR_SEVERITY {
				failAlarmCount++
			} else {
				clearAlarmCount++
			}
		}

		s.logger.Infof("Sending alarm fail: %v, clear: %v", failAlarmCount, clearAlarmCount)
	}

	return nil
}

func (s *dailyPerformanceAlarmService) getConfig() (*model.PerformanceAlarmConfig, error) {
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

	if pointy.IntValue(config.Duration, 0) == 0 {
		err := errors.New("duration must not be zero value")
		s.logger.Error(err)
		return nil, err
	}

	return config, nil
}

func (s *dailyPerformanceAlarmService) getInstalledCapacity() (*model.InstalledCapacity, error) {
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

func (p *dailyPerformanceAlarmService) getPayload(date *time.Time, dailyProduction, threshold float64, plantItem *model.PlantItem, capacityConfig model.InstalledCapacity, alarmConfig model.PerformanceAlarmConfig) (plantName string, alarmName string, payload string, severity string, err error) {
	var vendorType string
	switch strings.ToLower(plantItem.VendorType) {
	case constant.VENDOR_TYPE_GROWATT:
		vendorType = "Growatt"
	case constant.VENDOR_TYPE_HUAWEI:
		vendorType = "HUA"
	case constant.VENDOR_TYPE_KSTAR:
		vendorType = "Kstar"
	case constant.VENDOR_TYPE_INVT:
		vendorType = "INVT-Ipanda"
	case constant.VENDOR_TYPE_SOLARMAN:
		vendorType = "INVT-Ipanda"
	default:
	}

	if util.EmptyString(vendorType) {
		err = fmt.Errorf("vendor type (%s) not supported", plantItem.VendorType)
		return
	}

	plantName = pointy.StringValue(plantItem.Name, "")
	alarmName = fmt.Sprintf("SolarCell-%s", strings.ReplaceAll(alarmConfig.Name, " ", ""))
	alarmNameInPayload := util.AddSpace(alarmConfig.Name)
	if dailyProduction > threshold {
		severity = constant.CLEAR_SEVERITY
		payload = fmt.Sprintf("%s, %s, More than or equal %.2f%%, Expected Daily Production:%.2f KWH, Actual Production:%.2f KWH, Date:%s", vendorType, alarmNameInPayload, alarmConfig.Percentage, threshold, dailyProduction, date.Format("2006-01-02"))
	} else {
		err = ErrDailyProductionUnderThreshold
		return
	}

	return
}

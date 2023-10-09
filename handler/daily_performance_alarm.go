package handler

import (
	"time"

	"github.com/HavvokLab/true-solar-monitoring/constant"
	"github.com/HavvokLab/true-solar-monitoring/infra"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/repo"
	"github.com/HavvokLab/true-solar-monitoring/service"
)

type DailyPerformanceAlarmHandler struct {
	logger logger.Logger
}

func NewDailyPerformanceAlarmHandler() *DailyPerformanceAlarmHandler {
	logger := logger.NewLogger(
		&logger.LoggerOption{
			LogName:     constant.GetLogName(constant.DAILY_PERFORMANCE_ALARM_LOG_NAME),
			LogSize:     1024,
			LogAge:      90,
			LogBackup:   1,
			LogCompress: false,
			LogLevel:    logger.LOG_LEVEL_DEBUG,
			SkipCaller:  1,
		},
	)

	return &DailyPerformanceAlarmHandler{logger: logger}
}

func (h *DailyPerformanceAlarmHandler) Run() {
	now := time.Now()
	h.logger = logger.NewLogger(
		&logger.LoggerOption{
			LogName:     constant.GetLogName(constant.DAILY_PERFORMANCE_ALARM_LOG_NAME),
			LogSize:     1024,
			LogAge:      90,
			LogBackup:   1,
			LogCompress: false,
			LogLevel:    logger.LOG_LEVEL_DEBUG,
			SkipCaller:  1,
		},
	)
	defer h.logger.Close()

	elastic, err := infra.NewElasticsearch()
	if err != nil {
		h.logger.Errorf("Failed to connect to elasticsearch: %v", err)
		return
	}

	snmp, err := infra.NewSnmp()
	if err != nil {
		h.logger.Errorf("Failed to connect to snmp: %v", err)
		return
	}

	db, err := infra.NewGormDB()
	if err != nil {
		h.logger.Errorf("Failed to connect to gorm: %v", err)
		return
	}

	solarRepo := repo.NewSolarRepo(elastic)
	snmpRepo := repo.NewSnmpRepo(snmp)
	performanceAlarmConfigRepo := repo.NewPerformanceAlarmConfigRepo(db)
	installedCapacityRepo := repo.NewInstalledCapacityRepo(db)
	serv := service.NewDailyPerformanceAlarmService(solarRepo, installedCapacityRepo, performanceAlarmConfigRepo, snmpRepo, h.logger)
	if err := serv.Run(); err != nil {
		h.logger.Errorf("Failed to run daily performance alarm: %v", err)
		return
	}
	h.logger.Infof("DailyPerformanceAlarmHandler finished in %v", time.Since(now))
}

func (h *DailyPerformanceAlarmHandler) Mock() {
	h.logger = logger.NewLogger(
		&logger.LoggerOption{
			LogName:     constant.GetLogName(constant.DAILY_PERFORMANCE_ALARM_LOG_NAME),
			LogSize:     1024,
			LogAge:      90,
			LogBackup:   1,
			LogCompress: false,
			LogLevel:    logger.LOG_LEVEL_DEBUG,
			SkipCaller:  1,
		},
	)
	defer h.logger.Close()

	elastic, err := infra.NewElasticsearch()
	if err != nil {
		h.logger.Errorf("Failed to connect to elasticsearch: %v", err)
		return
	}

	db, err := infra.NewGormDB()
	if err != nil {
		h.logger.Errorf("Failed to connect to gorm: %v", err)
		return
	}

	solarRepo := repo.NewSolarRepo(elastic)
	snmpRepo := repo.NewMockSnmpRepo()
	performanceAlarmConfigRepo := repo.NewPerformanceAlarmConfigRepo(db)
	installedCapacityRepo := repo.NewInstalledCapacityRepo(db)
	serv := service.NewDailyPerformanceAlarmService(solarRepo, installedCapacityRepo, performanceAlarmConfigRepo, snmpRepo, h.logger)
	serv.Run()
}

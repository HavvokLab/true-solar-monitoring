package main

import (
	"github.com/HavvokLab/true-solar-monitoring/config"
	"github.com/HavvokLab/true-solar-monitoring/infra"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/HavvokLab/true-solar-monitoring/repo"
	"github.com/HavvokLab/true-solar-monitoring/service"
	"github.com/HavvokLab/true-solar-monitoring/util"
)

func init() {
	config.InitConfig()
}

func init() {
	util.SetTimezone()
}

func main() {
	credential := model.HuaweiCredential{Username: "trueapi", Password: "Trueapi12@"}
	logger := logger.NewLogger(&logger.LoggerOption{
		LogName:     "huawei_troubleshoot",
		LogSize:     1024,
		LogAge:      90,
		LogBackup:   1,
		LogCompress: false,
		LogLevel:    logger.LOG_LEVEL_DEBUG,
		SkipCaller:  1,
	})

	elastic, err := infra.NewElasticsearch()
	if err != nil {
		logger.Errorf("[%v]Failed to connect to elasticsearch", credential.Username)
		return
	}
	solarRepo := repo.NewSolarRepo(elastic)

	db, err := infra.NewGormDB()
	if err != nil {
		logger.Errorf("[%v]Failed to connect to gorm", credential.Username)
		return
	}
	siteRegionRepo := repo.NewSiteRegionMappingRepo(db)
	serv := service.NewHuaweiTroubleShootService(solarRepo, siteRegionRepo, logger)
	if err := serv.Run(&credential); err != nil {
		panic(err)
	}
}

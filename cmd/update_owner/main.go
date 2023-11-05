package main

import (
	"fmt"
	"time"

	"github.com/HavvokLab/true-solar-monitoring/config"
	"github.com/HavvokLab/true-solar-monitoring/constant"
	"github.com/HavvokLab/true-solar-monitoring/infra"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/repo"
	"github.com/HavvokLab/true-solar-monitoring/util"
	"github.com/sourcegraph/conc/pool"
)

func init() {
	util.SetTimezone()
	config.InitConfig()
}

func main() {
	logger := logger.NewLogger(
		&logger.LoggerOption{
			LogLevel:   logger.LOG_LEVEL_DEBUG,
			LogName:    "logs/update_owner/main.log",
			SkipCaller: 1,
			LogSize:    1024,
			LogBackup:  3,
		},
	)

	start := time.Date(2021, time.September, 23, 0, 0, 0, 0, time.Local)
	now := time.Now()
	p := pool.New().WithMaxGoroutines(15)
	for {
		curr := start
		index := fmt.Sprintf("solarcell-%v", curr.Format("2006.01.02"))
		p.Go(run(index, logger))
		if curr.Day() == now.Day() && curr.Month() == now.Month() && curr.Year() == now.Year() {
			break
		}
		start = start.AddDate(0, 0, 1)
	}
	p.Wait()
}

func run(index string, logger logger.Logger) func() {
	return func() {
		logger.Infof("Start update owner to index %v", index)
		elastic, err := infra.NewElasticsearch()
		if err != nil {
			logger.Errorf("Update owner to index %v failed: %v", index, err)
			return
		}

		repo := repo.NewSolarRepo(elastic)
		err = repo.UpdateOwnerToIndex(index, string(constant.TRUE_OWNER))
		if err != nil {
			logger.Errorf("Update owner to index %v failed: %v", index, err)
			return
		}
		logger.Infof("Update owner to index %v success", index)
	}
}

package handler

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/HavvokLab/true-solar-monitoring/constant"
	"github.com/HavvokLab/true-solar-monitoring/infra"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/HavvokLab/true-solar-monitoring/repo"
	"github.com/HavvokLab/true-solar-monitoring/service"
	"github.com/HavvokLab/true-solar-monitoring/util"
	"github.com/gammazero/workerpool"
	"github.com/gocarina/gocsv"
)

type GrowattTroubleShootHandler struct {
	logger logger.Logger
}

func NewGrowattTroubleShootHandler() *GrowattTroubleShootHandler {
	return &GrowattTroubleShootHandler{}
}

type GrowattMissingSite struct {
	ID   int    `csv:"id"`
	Name string `csv:"name"`
	Date string `csv:"date"`
	Area string `csv:"area"`
}

func (s GrowattMissingSite) Unix() (int64, error) {
	t, err := time.Parse("2006/01/02", s.Date)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

func (h *GrowattTroubleShootHandler) Run() {
	h.logger = logger.NewLogger(
		&logger.LoggerOption{
			LogName:     util.GetLogName(constant.GROWATT_TROUBLESHOT_LOG_NAME),
			LogSize:     1024,
			LogAge:      90,
			LogBackup:   1,
			LogCompress: false,
			LogLevel:    logger.LOG_LEVEL_DEBUG,
			SkipCaller:  1,
		},
	)
	defer h.logger.Close()

	db, err := infra.NewGormDB()
	if err != nil {
		h.logger.Error(err)
		return
	}

	credentialRepo := repo.NewGrowattCredentialRepo(db)
	credentials, err := credentialRepo.FindAll()
	if err != nil {
		h.logger.Error(err)
		return
	}

	pool := workerpool.New(len(credentials))
	for _, credential := range credentials {
		clone := credential
		pool.Submit(h.run(&clone))
	}
	pool.StopWait()
}

func (h *GrowattTroubleShootHandler) run(credential *model.GrowattCredential) func() {
	h.logger.Infof("[%v] Start Troubleshooting", credential.Username)
	getSites := func() ([]GrowattMissingSite, error) {
		file, err := os.OpenFile("growatt_troubleshoot.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		var data []GrowattMissingSite
		if err := gocsv.UnmarshalFile(file, &data); err != nil {
			log.Fatal(err)
		}

		result := make([]GrowattMissingSite, 0)
		for _, item := range data {
			if !util.EmptyString(item.Area) {
				parts := strings.Split(item.Area, "-")
				num, _ := strconv.Atoi(parts[1])
				area := strings.ToLower(fmt.Sprintf("%v%v", parts[0], num))
				match, _ := regexp.MatchString(area, credential.Username)
				if match {
					result = append(result, item)
				}
			}
		}

		return result, nil
	}

	getSiteIDByDate := func(sites []GrowattMissingSite, date string) ([]int, error) {
		result := make([]int, 0)
		for _, site := range sites {
			if site.Date == date {
				result = append(result, site.ID)
			}
		}
		return result, nil
	}

	return func() {
		elastic, err := infra.NewElasticsearch()
		if err != nil {
			h.logger.Errorf("[%v]Failed to connect to elasticsearch", credential.Username)
			return
		}

		db, err := infra.NewGormDB()
		if err != nil {
			h.logger.Errorf("[%v]Failed to connect to gorm", credential.Username)
			return
		}

		sites, err := getSites()
		if err != nil {
			h.logger.Errorf("[%v]Failed to get sites: %v", credential.Username, err)
			return
		}

		uniqueDateSet := util.NewSet()
		for _, site := range sites {
			uniqueDateSet.Add(site.Date)
		}
		uniqueDate := uniqueDateSet.Keys()

		for _, dateStr := range uniqueDate {
			sites, err := getSiteIDByDate(sites, dateStr)
			if err != nil {
				h.logger.Errorf("[%v]Failed to get site by date: %v", credential.Username, err)
				return
			}

			date, err := time.Parse("2006/01/02", dateStr)
			if err != nil {
				h.logger.Errorf("[%v]Failed to parse date: %v", credential.Username, err)
				return
			}

			now := time.Now()
			solarRepo := repo.NewSolarRepo(elastic)

			siteRegionRepo := repo.NewSiteRegionMappingRepo(db)
			plantRepo := repo.NewPlantRepo(db)

			serv := service.NewGrowattTroubleshootService(solarRepo, plantRepo, siteRegionRepo, h.logger)
			if err := serv.RunWithPlantIdList(credential, sites, date.Unix(), date.Unix()); err != nil {
				h.logger.Errorf("[%v]Failed to run service: %v", credential.Username, err)
				return
			}

			h.logger.Infof("[%v] Finished in %v", credential.Username, time.Since(now).String())
		}
	}
}

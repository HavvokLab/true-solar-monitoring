package service

import (
	"fmt"
	"time"

	"github.com/HavvokLab/true-solar-monitoring/config"
	"github.com/HavvokLab/true-solar-monitoring/constant"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/HavvokLab/true-solar-monitoring/repo"
	"go.openly.dev/pointy"
)

type MonthlyProductionService interface {
	Run(start, end *time.Time) error
	DailyRun() error
}

type monthlyProductionService struct {
	solarRepo      repo.SolarRepo
	masterSiteRepo repo.MasterSiteRepo
	logger         logger.Logger
}

func NewMonthlyProductionService(solarRepo repo.SolarRepo, masterSiteRepo repo.MasterSiteRepo, logger logger.Logger) MonthlyProductionService {
	return &monthlyProductionService{
		solarRepo:      solarRepo,
		masterSiteRepo: masterSiteRepo,
		logger:         logger,
	}
}

func (s monthlyProductionService) DailyRun() error {
	now := time.Now()
	var end time.Time = now.AddDate(0, 0, -1)
	var start time.Time = time.Date(end.Year(), end.Month(), 1, 0, 0, 0, 0, time.Local)

	defer func() {
		if r := recover(); r != nil {
			s.logger.Warnf("[%v]MonthlyProduction.Run(): %v", end.Format(constant.YEAR_MONTH), r)
		}
	}()

	documents, err := s.generateDocuments(&start, &end)
	if err != nil {
		s.logger.Errorf("[%v]MonthlyProduction.Run(): %v", start.Format(constant.YEAR_MONTH), err)
		return err
	}
	s.logger.Infof("[%v]MonthlyProduction.Run(): generated %v documents", start.Format(constant.YEAR_MONTH), len(documents))

	if len(documents) == 0 {
		s.logger.Errorf("[%v]MonthlyProduction.Run(): %v", start.Format(constant.YEAR_MONTH), "documents is empty")
		return nil
	}

	conf := config.GetConfig().Elastic
	index := fmt.Sprintf("%v_%v", conf.MonthlyProductionIndex, start.Format("200601"))

	if end.Day() > 1 {
		if err := s.solarRepo.DeleteIndex(index); err != nil {
			s.logger.Errorf("[%v]MonthlyProduction.Run(): %v", start.Format(constant.YEAR_MONTH), err)
			return err
		}
		s.logger.Infof("[%v]MonthlyProduction.Run(): index %q deleted", index)
	}

	if err := s.solarRepo.BulkIndex(index, documents); err != nil {
		s.logger.Errorf("[%v]MonthlyProduction.Run(): %v", start.Format(constant.YEAR_MONTH), err)
		return err
	}

	s.logger.Infof("MonthlyProduction.Run(): bulked new index %q", index)
	return nil
}

func (s monthlyProductionService) Run(start, end *time.Time) error {
	defer func() {
		if r := recover(); r != nil {
			s.logger.Warnf("[%v]MonthlyProduction.Run(): %v", start.Format(constant.YEAR_MONTH), r)
		}
	}()

	documents, err := s.generateDocuments(start, end)
	if err != nil {
		s.logger.Errorf("[%v]MonthlyProduction.Run(): %v", start.Format(constant.YEAR_MONTH), err)
		return err
	}
	s.logger.Infof("[%v]MonthlyProduction.Run(): generated %v documents", start.Format(constant.YEAR_MONTH), len(documents))

	if len(documents) == 0 {
		s.logger.Errorf("[%v]MonthlyProduction.Run(): %v", start.Format(constant.YEAR_MONTH), "documents is empty")
		return nil
	}

	conf := config.GetConfig().Elastic
	index := fmt.Sprintf("%v_%v", conf.MonthlyProductionIndex, start.Format("200601"))
	if err := s.solarRepo.BulkIndex(index, documents); err != nil {
		s.logger.Errorf("[%v]MonthlyProduction.Run(): %v", start.Format(constant.YEAR_MONTH), err)
		return err
	}

	s.logger.Infof("MonthlyProduction.Run(): bulked new index %q", index)
	return nil
}

func (s monthlyProductionService) generateDocuments(start, end *time.Time) ([]interface{}, error) {
	data, err := s.solarRepo.GetPlantMonthlyProduction(start, end)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	if data == nil {
		err := fmt.Errorf("[%v]MonthlyProduction.Run(): %v", start.Format(constant.YEAR_MONTH), "aggregate productions are empty")
		s.logger.Error(err)
		return nil, err
	}

	masterSiteMap := s.masterSiteRepo.ExportToMap()
	var count int
	var size = len(data)
	documents := make([]interface{}, 0)
	for _, item := range data {
		// |=> generated exist site
		if item == nil {
			continue
		}

		if len(item.Key) == 0 {
			continue
		}

		doc := model.MonthlyProductionDocument{}

		// took data from key
		if val, ok := item.Key["date"].(string); ok {
			date, _ := time.Parse(constant.YEAR_MONTH_DAY, val)
			doc.SetDate(&date)
		}

		if val, ok := item.Key["vendor_type"].(string); ok {
			doc.SetVendorType(val)
		}

		if val, ok := item.Key["area"].(string); ok {
			doc.SetArea(val)
		}

		if val, ok := item.Key["name"].(string); ok {
			doc.SetSiteName(val)
		}

		if val, ok := item.Key["owner"].(string); ok {
			doc.SetOwner(val)
		}

		// took data from max_aggregation
		if val, ok := item.Aggregations.Max("lat"); ok {
			doc.SetLatitude(val.Value)
		}

		if val, ok := item.Aggregations.Max("long"); ok {
			doc.SetLongitude(val.Value)
		}
		doc.SetLocation(doc.Latitude, doc.Longitude)

		if val, ok := item.Aggregations.Max("installed_capacity"); ok {
			doc.SetInstalledCapacity(val.Value)
		}

		if val, ok := item.Aggregations.Max("monthly_production"); ok {
			doc.SetMonthlyProduction(val.Value)
		}

		// took data from bucket script
		if val, ok := item.BucketScript("target"); ok {
			doc.SetTarget(val.Value)
		}

		if val, ok := item.BucketScript("production_to_target"); ok {
			doc.SetProductionToTarget(val.Value)
		}
		doc.SetCriteria(doc.ProductionToTarget)
		doc.ClearZeroValue()

		s.logger.Infof("[%v/%v] generateDocument vendor_type: %v, name: %v, monthly_production: %v, target: %v, product2target: %v, criteria: %v",
			count,
			size,
			*doc.VendorType,
			*doc.SiteName,
			pointy.Float64Value(doc.MonthlyProduction, 0.0),
			pointy.Float64Value(doc.Target, 0.0),
			pointy.Float64Value(doc.ProductionToTarget, 0.0),
			*doc.Criteria,
		)
		count += 1
		s.logger.Infof("[%v/%v] generateDocument of %v", count, size, start.Format(constant.YEAR_MONTH))
		documents = append(documents, doc)

		// |=> generated non-exist site
		masterSite := model.MasterSite{
			Vendor:   doc.VendorType,
			Area:     doc.Area,
			SiteName: doc.SiteName,
		}
		delete(masterSiteMap, masterSite.GetKey())
	}

	count = 0
	size = len(masterSiteMap)
	for _, site := range masterSiteMap {
		doc := model.MonthlyProductionDocument{
			Date:               start,
			VendorType:         site.Vendor,
			Area:               site.Area,
			SiteName:           site.SiteName,
			InstalledCapacity:  site.InstalledCapacity,
			MonthlyProduction:  nil,
			Latitude:           site.Latitude,
			Longitude:          site.Longitude,
			Target:             nil,
			ProductionToTarget: nil,
			Criteria:           nil,
		}
		doc.SetLocation(doc.Latitude, doc.Longitude)
		doc.SetCriteria(doc.ProductionToTarget)
		doc.ClearZeroValue()

		count += 1
		s.logger.Infof("[%v/%v] non-exist generateDocument of %v", count, size, start.Format(constant.YEAR_MONTH_DAY))
		documents = append(documents, doc)
	}

	return documents, nil
}

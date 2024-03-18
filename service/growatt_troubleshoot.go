package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/HavvokLab/true-solar-monitoring/config"
	"github.com/HavvokLab/true-solar-monitoring/constant"
	"github.com/HavvokLab/true-solar-monitoring/inverter"
	"github.com/HavvokLab/true-solar-monitoring/inverter/growatt"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/HavvokLab/true-solar-monitoring/repo"
	"github.com/HavvokLab/true-solar-monitoring/util"
	"go.openly.dev/pointy"
)

type GrowattTroubleshootService interface {
}

type growattTroubleshootService struct {
	vendorType     string
	siteRegionRepo repo.SiteRegionMappingRepo
	siteRegions    []model.SiteRegionMapping
	solarRepo      repo.SolarRepo
	plantRepo      repo.PlantRepo
	elasticConfig  config.ElasticsearchConfig
	logger         logger.Logger
}

func NewGrowattTroubleshootService(solarRepo repo.SolarRepo, plantRepo repo.PlantRepo, siteRegionRepo repo.SiteRegionMappingRepo, logger logger.Logger) *growattTroubleshootService {
	return &growattTroubleshootService{
		vendorType:     strings.ToUpper(constant.VENDOR_TYPE_GROWATT),
		siteRegionRepo: siteRegionRepo,
		solarRepo:      solarRepo,
		plantRepo:      plantRepo,
		logger:         logger,
		siteRegions:    make([]model.SiteRegionMapping, 0),
		elasticConfig:  config.GetConfig().Elastic,
	}
}
func (s growattTroubleshootService) RunWithPlantIdList(credential *model.GrowattCredential, plantIdList []int, start, end int64) error {
	documents := make([]interface{}, 0)
	documentCh := make(chan interface{})
	errCh := make(chan error)
	doneCh := make(chan bool)

	siteRegions, err := s.siteRegionRepo.GetSiteRegionMappings()
	if err != nil {
		s.logger.Errorf("[%v] Error while getting site region mappings: %v", credential.Username, err)
		return err
	}
	s.siteRegions = siteRegions

	go s.runWithPlantIdList(credential, plantIdList, start, end, documentCh, errCh, doneCh)

DONE:
	for {
		select {
		case document := <-documentCh:
			documents = append(documents, document)
			util.PrintJSON(document)
		case err := <-errCh:
			s.logger.Errorf("error: %v", err)
		case <-doneCh:
			s.logger.Infof("done")
			break DONE
		}
	}

	collectorIndex := fmt.Sprintf("%s-%s", s.elasticConfig.SolarIndex, time.Unix(start, 0).Format("2006.01.02"))
	if err := s.solarRepo.BulkIndex(collectorIndex, documents); err != nil {
		s.logger.Errorf("[%v] - GrowattTroubleShootService.Run(): %v", credential.Username, err)
		return err
	}

	return nil
}

func (s *growattTroubleshootService) runWithPlantIdList(credential *model.GrowattCredential, plantIdList []int, start, end int64, documentCh chan interface{}, errCh chan error, doneCh chan bool) {
	client, err := growatt.NewGrowattClient(&growatt.GrowattCredential{
		Username: credential.Username,
		Token:    credential.Token,
	})

	if err != nil {
		s.logger.Errorf("[%v] %v", credential.Username, err)
		errCh <- err
		return
	}

	startDate := time.Unix(start, 0)
	endDate := time.Unix(end, 0)

	if startDate.Month() != endDate.Month() || startDate.Year() != endDate.Year() {
		s.logger.Errorf("the range of start date and end date should be in the same period")
		return
	}

	for _, plantId := range plantIdList {
		plant, err := client.GetPlantBasicInfo(plantId)
		if err != nil {
			s.logger.Errorf("[%v] %v", credential.Username, err)
			errCh <- err
			continue
		}

		if plant == nil || plant.Data == nil {
			continue
		}

		plantData := plant.Data
		plantIdStr := fmt.Sprintf("%d", plantId)
		plantIdentity, _ := inverter.ParsePlantID(plantData.GetName())
		cityName, cityCode, cityArea := inverter.ParseSiteID(s.siteRegions, plantIdentity.SiteID)

		var monthlyProduction *float64
		var yearlyProduction *float64
		yearlyEnergies, err := client.GetHistoricalPlantPowerGeneration(plantId, start, end, "year")
		if err != nil {
			s.logger.Error(err)
			errCh <- err
			continue
		}

		if len(yearlyEnergies) > 0 {
			yearlyProduction = pointy.Float64(yearlyEnergies[0].GetEnergy())
		}

		monthlyEnergies, err := client.GetHistoricalPlantPowerGeneration(plantId, start, end, "month")
		if err != nil {
			s.logger.Error(err)
			errCh <- err
			continue
		}

		if len(monthlyEnergies) > 0 {
			monthlyProduction = pointy.Float64(monthlyEnergies[0].GetEnergy())
		}

		dailyEnergies, err := client.GetHistoricalPlantPowerGeneration(plantId, start, end, "day")
		if err != nil {
			s.logger.Error(err)
			errCh <- err
			continue
		}

		if len(dailyEnergies) == 0 {
			continue
		}

		masterPlant, err := s.plantRepo.FindOneByName(plantData.GetName())
		if err != nil {
			errCh <- err
			continue
		} else if util.EmptyString(masterPlant.Name) {
			errCh <- fmt.Errorf("")
			continue
		}

		var location *string
		if masterPlant.Latitude != nil && masterPlant.Longitude != nil {
			location = pointy.String(fmt.Sprintf("%f,%f", *masterPlant.Latitude, *masterPlant.Longitude))
		}

		for _, daily := range dailyEnergies {
			date := time.Unix(start, 0)

			plantItem := model.PlantItem{
				Timestamp:         date,
				Month:             date.Format("01"),
				Year:              date.Format("2006"),
				MonthYear:         date.Format("01-2006"),
				VendorType:        s.vendorType,
				DataType:          constant.DATA_TYPE_PLANT,
				Area:              cityArea,
				SiteID:            plantIdentity.SiteID,
				SiteCityName:      cityName,
				SiteCityCode:      cityCode,
				NodeType:          plantIdentity.NodeType,
				ACPhase:           plantIdentity.ACPhase,
				ID:                pointy.String(plantIdStr),
				Name:              plantData.Name,
				PlantStatus:       pointy.String("UNKNOWN"),
				Owner:             credential.Owner,
				Latitude:          masterPlant.Latitude,
				Longitude:         masterPlant.Longitude,
				Location:          location,
				LocationAddress:   plantData.City,
				YearlyProduction:  yearlyProduction,
				MonthlyProduction: monthlyProduction,
				DailyProduction:   pointy.Float64(daily.GetEnergy()),
				InstalledCapacity: &masterPlant.InstalledCapacity,
				MonthlyCO2:        pointy.Float64(pointy.Float64Value(monthlyProduction, 0) * 2.079),
			}

			documentCh <- plantItem
			break
		}
	}

	doneCh <- true
}

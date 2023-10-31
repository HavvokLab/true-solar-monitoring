package service

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/HavvokLab/true-solar-monitoring/config"
	"github.com/HavvokLab/true-solar-monitoring/constant"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/HavvokLab/true-solar-monitoring/repo"
	"github.com/HavvokLab/true-solar-monitoring/util"
	"go.openly.dev/pointy"
)

type PlantAggregateService interface {
	UpdatePlantByDate(date *time.Time) error
}

type plantAggregateService struct {
	plantRepo repo.PlantRepo
	solarRepo repo.SolarRepo
	logger    logger.Logger
}

func NewPlantAggregateService(plantRepo repo.PlantRepo, solarRepo repo.SolarRepo, logger logger.Logger) PlantAggregateService {
	return &plantAggregateService{
		plantRepo: plantRepo,
		solarRepo: solarRepo,
		logger:    logger,
	}
}

func (s *plantAggregateService) UpdatePlantByDate(date *time.Time) error {
	index := fmt.Sprintf("%s-%s", config.GetConfig().Elastic.SolarIndex, date.Format("2006.01.*"))
	data, err := s.solarRepo.GetUniquePlantByIndex(index)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	plants := make([]*model.Plant, 0)
	for _, item := range data {
		if item == nil {
			continue
		}

		mapData := map[string]interface{}{}
		if topHits, found := item.Aggregations.TopHits("data"); found {
			if topHits.Hits != nil {
				hit := topHits.Hits.Hits[0]
				if hit != nil {
					if err := util.Recast(hit.Source, &mapData); err != nil {
						s.logger.Error(err)
						return err
					}
				}
			}
		}

		plant := new(model.Plant)
		if val, ok := mapData["name"]; ok {
			if parsed, ok := val.(string); ok {
				plant.Name = parsed
			} else {
				continue
			}
		} else {
			continue
		}

		if val, ok := mapData["area"]; ok {
			if parsed, ok := val.(string); ok {
				plant.Area = &parsed
			}
		}

		if val, ok := mapData["vendor_type"]; ok {
			if parsed, ok := val.(string); ok {
				plant.VendorType = parsed
			} else {
				continue
			}
		} else {
			continue
		}

		if val, ok := mapData["installed_capacity"]; ok {
			if parsed, ok := val.(float64); ok {
				plant.InstalledCapacity = parsed
			}
		}

		if val, ok := mapData["owner"]; ok {
			if parsed, ok := val.(string); ok {
				plant.Owner = &parsed
			} else {
				plant.Owner = pointy.String(string(constant.TRUE_OWNER))
			}
		} else {
			plant.Owner = pointy.String(string(constant.TRUE_OWNER))
		}

		if val, ok := mapData["location"]; ok {
			if parsed, ok := val.(string); ok {
				if len(parsed) > 0 {
					parts := strings.Split(parsed, ",")
					if len(parts) == 2 {
						lat, err := strconv.ParseFloat(parts[0], 64)
						if err == nil {
							plant.Latitude = &lat
						}

						long, err := strconv.ParseFloat(parts[0], 64)
						if err == nil {
							plant.Longitude = &long
						}
					}
				}
			}
		}

		plants = append(plants, plant)
	}

	if err := s.plantRepo.BatchUpsertAvailable(plants); err != nil {
		s.logger.Error(err)
		return err
	}

	return nil
}

package service

import (
	"fmt"
	"strings"

	"github.com/HavvokLab/true-solar-monitoring/domain"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/HavvokLab/true-solar-monitoring/repo"
	"github.com/HavvokLab/true-solar-monitoring/util"
	"github.com/google/uuid"
)

type SiteRegionMappingService interface {
	FindAll(limit, offset int) (*domain.FindAllSiteRegionMappingsResponse, error)
	FindRegion() (*model.Regions, error)
}

type siteRegionMappingService struct {
	siteRegionMappingRepo repo.SiteRegionMappingRepo
	logger                logger.Logger
}

func NewSiteRegionMappingService(siteRegionMappingRepo repo.SiteRegionMappingRepo, logger logger.Logger) SiteRegionMappingService {
	return &siteRegionMappingService{
		siteRegionMappingRepo: siteRegionMappingRepo,
		logger:                logger,
	}
}

func (s *siteRegionMappingService) FindAll(limit, offset int) (*domain.FindAllSiteRegionMappingsResponse, error) {
	count, err := s.siteRegionMappingRepo.Count()
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	if offset < 0 {
		data, err := s.siteRegionMappingRepo.GetSiteRegionMappings()
		if err != nil {
			s.logger.Error(err)
			return nil, err
		}

		return &domain.FindAllSiteRegionMappingsResponse{
			Count:              count,
			SiteRegionMappings: data,
		}, nil
	}

	if limit < 12 {
		limit = 12
	}

	data, err := s.siteRegionMappingRepo.GetSiteRegionMappingsWithPagination(limit, offset)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	return &domain.FindAllSiteRegionMappingsResponse{
		Count:              count,
		SiteRegionMappings: data,
	}, nil
}

func (s *siteRegionMappingService) FindRegion() (*model.Regions, error) {
	siteRegionMapping, err := s.siteRegionMappingRepo.GetSiteRegionMappings()
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	mapAreaCities := make(map[string][]model.SiteRegionMapping)
	for _, site := range siteRegionMapping {
		if site.Area != nil {
			mapAreaCities[*site.Area] = append(mapAreaCities[*site.Area], site)
		}
	}

	var region model.Regions
	for area, cities := range mapAreaCities {
		if len(cities) == 0 {
			cities = make([]model.SiteRegionMapping, 0)
		}

		region.Regions = append(region.Regions, model.AreaWithCity{
			Area:   area,
			Cities: cities,
		})
	}

	return &region, nil
}

func (s *siteRegionMappingService) UpdateRegion(req *domain.UpdateRegionRequest) error {
	if err := util.ValidateStruct(req); err != nil {
		return err
	}

	err := s.siteRegionMappingRepo.UpdateCityToNullArea(strings.ToUpper(req.Area))
	if err != nil {
		s.logger.Error(err)
		return err
	}

	if len(req.Cities) == 0 {
		id := uuid.New()
		cityCode := strings.ToUpper(fmt.Sprintf("EMPTY-%s", id.String()))

		err = s.siteRegionMappingRepo.CreateCity(&model.SiteRegionMapping{
			Code: cityCode,
		})

		if err != nil {
			s.logger.Error(err)
			return err
		}

		req.Cities = append(req.Cities, cityCode)
	}

	codeListString := strings.ToUpper(fmt.Sprintf("'%s'", strings.Join(req.Cities, "','")))
	err = s.siteRegionMappingRepo.UpdateSiteRegionMapping(codeListString, req.Area)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	return nil
}

func (s *siteRegionMappingService) DeleteCity(id int64) error {
	err := s.siteRegionMappingRepo.DeleteCity(id)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	return nil
}

func (s *siteRegionMappingService) DeleteArea(area string) error {
	err := s.siteRegionMappingRepo.UpdateCityToNullArea(strings.ToUpper(area))
	if err != nil {
		s.logger.Error(err)
		return err
	}

	return nil
}

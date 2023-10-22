package service

import (
	"github.com/HavvokLab/true-solar-monitoring/domain"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/HavvokLab/true-solar-monitoring/repo"
	"github.com/HavvokLab/true-solar-monitoring/util"
)

type InstalledCapacityService interface {
	FindOne(*domain.UserContext) (*model.InstalledCapacity, error)
	Update(*domain.UserContext, int64, *domain.UpdateInstalledCapacityRequest) error
}

type installedCapacityService struct {
	repo   repo.InstalledCapacityRepo
	logger logger.Logger
}

func NewInstalledCapacityService(repo repo.InstalledCapacityRepo, logger logger.Logger) InstalledCapacityService {
	return &installedCapacityService{repo: repo, logger: logger}
}

func (s *installedCapacityService) FindOne(utx *domain.UserContext) (*model.InstalledCapacity, error) {
	result, err := s.repo.FindOne()
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	return result, nil
}

func (s *installedCapacityService) Update(utx *domain.UserContext, id int64, request *domain.UpdateInstalledCapacityRequest) error {
	if err := util.ValidateStruct(request); err != nil {
		s.logger.Error(err)
		return err
	}

	installedCapacity := new(model.InstalledCapacity)
	installedCapacity.EfficiencyFactor = request.EfficiencyFactor
	installedCapacity.FocusHour = request.FocusHour

	if err := s.repo.Update(id, installedCapacity); err != nil {
		s.logger.Error(err)
		return err
	}

	return nil
}

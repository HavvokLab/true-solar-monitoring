package service

import (
	"github.com/HavvokLab/true-solar-monitoring/domain"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/HavvokLab/true-solar-monitoring/repo"
	"github.com/HavvokLab/true-solar-monitoring/util"
)

type PerformanceAlarmConfigService interface {
	GetLowPerformanceAlarmConfig(utx *domain.UserContext) (*model.PerformanceAlarmConfig, error)
	UpdateLowPerformanceAlarmConfig(*domain.UserContext, *domain.UpdatePerformanceAlarmConfigRequest) error
}

type performanceAlarmConfigService struct {
	performanceAlarmConfigRepo repo.PerformanceAlarmConfigRepo
	logger                     logger.Logger
}

func NewPerformanceAlarmConfigService(
	performanceAlarmConfigRepo repo.PerformanceAlarmConfigRepo,
	logger logger.Logger,
) PerformanceAlarmConfigService {
	return &performanceAlarmConfigService{
		performanceAlarmConfigRepo: performanceAlarmConfigRepo,
		logger:                     logger,
	}
}

func (s *performanceAlarmConfigService) GetLowPerformanceAlarmConfig(utx *domain.UserContext) (*model.PerformanceAlarmConfig, error) {
	result, err := s.performanceAlarmConfigRepo.GetLowPerformanceAlarmConfig()
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	return result, nil
}

func (s *performanceAlarmConfigService) UpdateLowPerformanceAlarmConfig(utx *domain.UserContext, request *domain.UpdatePerformanceAlarmConfigRequest) error {
	if err := util.ValidateStruct(request); err != nil {
		return err
	}

	data := new(model.PerformanceAlarmConfig)
	data.Interval = request.Interval
	data.Percentage = request.Percentage
	data.HitDay = request.HitDay
	data.Duration = request.Duration

	if err := s.performanceAlarmConfigRepo.UpdateLowPerformanceAlarmConfig(data); err != nil {
		s.logger.Error(err)
		return err
	}

	return nil
}

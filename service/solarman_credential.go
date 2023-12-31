package service

import (
	"github.com/HavvokLab/true-solar-monitoring/domain"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/HavvokLab/true-solar-monitoring/repo"
	"github.com/HavvokLab/true-solar-monitoring/util"
)

type SolarmanCredentialService interface {
	FindAll(utx *domain.UserContext) ([]model.SolarmanCredential, error)
	Create(utx *domain.UserContext, request *domain.CreateSolarmanCredentialRequest) error
	Update(utx *domain.UserContext, id int64, request *domain.UpdateSolarmanCredentialRequest) error
	Delete(utx *domain.UserContext, id int64) error
}

type solarmanCredentialService struct {
	repo   repo.SolarmanCredentialRepo
	logger logger.Logger
}

func NewSolarmanCredentialService(repo repo.SolarmanCredentialRepo, logger logger.Logger) SolarmanCredentialService {
	return &solarmanCredentialService{
		repo:   repo,
		logger: logger,
	}
}

func (s *solarmanCredentialService) FindAll(utx *domain.UserContext) ([]model.SolarmanCredential, error) {
	result, err := s.repo.FindAll()
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	return result, nil
}

func (s *solarmanCredentialService) Create(utx *domain.UserContext, request *domain.CreateSolarmanCredentialRequest) error {
	if err := util.ValidateStruct(request); err != nil {
		return err
	}

	credential := new(model.SolarmanCredential)
	credential.Username = request.Username
	credential.Password = request.Password
	credential.AppID = request.AppID
	credential.AppSecret = request.AppSecret
	credential.Owner = request.Owner

	if err := s.repo.Create(credential); err != nil {
		s.logger.Error(err)
		return err
	}

	return nil
}

func (s *solarmanCredentialService) Update(utx *domain.UserContext, id int64, request *domain.UpdateSolarmanCredentialRequest) error {
	if err := util.ValidateStruct(request); err != nil {
		return err
	}

	credential := new(model.SolarmanCredential)
	credential.Username = request.Username
	credential.Password = request.Password
	credential.AppID = request.AppID
	credential.AppSecret = request.AppSecret
	credential.Owner = request.Owner

	if err := s.repo.Update(id, credential); err != nil {
		s.logger.Error(err)
		return err
	}

	return nil
}

func (s *solarmanCredentialService) Delete(utx *domain.UserContext, id int64) error {
	if err := s.repo.Delete(id); err != nil {
		s.logger.Error(err)
		return err
	}

	return nil
}

package service

import (
	"github.com/HavvokLab/true-solar-monitoring/domain"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/HavvokLab/true-solar-monitoring/repo"
	"github.com/HavvokLab/true-solar-monitoring/util"
)

type GrowattCredentialService interface {
	FindAll(utx *domain.UserContext) ([]model.GrowattCredential, error)
	Create(utx *domain.UserContext, request *domain.CreateGrowattCredentialRequest) error
	Update(utx *domain.UserContext, id int64, request *domain.UpdateGrowattCredentialRequest) error
	Delete(utx *domain.UserContext, id int64) error
}

type growattCredentialService struct {
	repo   repo.GrowattCredentialRepo
	logger logger.Logger
}

func NewGrowattCredentialService(repo repo.GrowattCredentialRepo, logger logger.Logger) GrowattCredentialService {
	return &growattCredentialService{
		repo:   repo,
		logger: logger,
	}
}

func (s *growattCredentialService) FindAll(utx *domain.UserContext) ([]model.GrowattCredential, error) {
	result, err := s.repo.FindAll()
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	return result, nil
}

func (s *growattCredentialService) Create(utx *domain.UserContext, request *domain.CreateGrowattCredentialRequest) error {
	if err := util.ValidateStruct(request); err != nil {
		return err
	}

	credential := new(model.GrowattCredential)
	credential.Username = request.Username
	credential.Password = request.Password
	credential.Token = request.Token
	credential.Owner = request.Owner

	if err := s.repo.Create(credential); err != nil {
		s.logger.Error(err)
		return err
	}

	return nil
}

func (s *growattCredentialService) Update(utx *domain.UserContext, id int64, request *domain.UpdateGrowattCredentialRequest) error {
	if err := util.ValidateStruct(request); err != nil {
		return err
	}

	credential := new(model.GrowattCredential)
	credential.Username = request.Username
	credential.Password = request.Password
	credential.Token = request.Token
	credential.Owner = request.Owner

	if err := s.repo.Update(id, credential); err != nil {
		s.logger.Error(err)
		return err
	}

	return nil
}

func (s *growattCredentialService) Delete(utx *domain.UserContext, id int64) error {
	if err := s.repo.Delete(id); err != nil {
		s.logger.Error(err)
		return err
	}

	return nil
}

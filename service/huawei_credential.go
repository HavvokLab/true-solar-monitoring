package service

import (
	"github.com/HavvokLab/true-solar-monitoring/domain"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/HavvokLab/true-solar-monitoring/repo"
	"github.com/HavvokLab/true-solar-monitoring/util"
)

type HuaweiCredentialService interface {
	FindAll() ([]model.HuaweiCredential, error)
	Create(request *domain.CreateHuaweiCredentialRequest) error
	Update(id int64, request *domain.UpdateHuaweiCredentialRequest) error
	Delete(id int64) error
}

type huaweiCredentialService struct {
	repo   repo.HuaweiCredentialRepo
	logger logger.Logger
}

func NewHuaweiCredentialService(repo repo.HuaweiCredentialRepo, logger logger.Logger) HuaweiCredentialService {
	return &huaweiCredentialService{
		repo:   repo,
		logger: logger,
	}
}

func (s *huaweiCredentialService) FindAll() ([]model.HuaweiCredential, error) {
	result, err := s.repo.FindAll()
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	return result, nil
}

func (s *huaweiCredentialService) Create(request *domain.CreateHuaweiCredentialRequest) error {
	if err := util.ValidateStruct(request); err != nil {
		return err
	}

	credential := new(model.HuaweiCredential)
	credential.Username = request.Username
	credential.Password = request.Password
	credential.Owner = request.Owner

	if err := s.repo.Create(credential); err != nil {
		s.logger.Error(err)
		return err
	}

	return nil
}

func (s *huaweiCredentialService) Update(id int64, request *domain.UpdateHuaweiCredentialRequest) error {
	if err := util.ValidateStruct(request); err != nil {
		return err
	}

	credential := new(model.HuaweiCredential)
	credential.Username = request.Username
	credential.Password = request.Password
	credential.Owner = request.Owner

	if err := s.repo.Update(id, credential); err != nil {
		s.logger.Error(err)
		return err
	}

	return nil
}

func (s *huaweiCredentialService) Delete(id int64) error {
	if err := s.repo.Delete(id); err != nil {
		s.logger.Error(err)
		return err
	}

	return nil
}

package service

import (
	"github.com/HavvokLab/true-solar-monitoring/domain"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/HavvokLab/true-solar-monitoring/repo"
	"github.com/HavvokLab/true-solar-monitoring/util"
)

type KStarCredentialService interface{}

type kstarCredentialService struct {
	repo   repo.KStarCredentialRepo
	logger logger.Logger
}

func NewKStarCredentialService(repo repo.KStarCredentialRepo, logger logger.Logger) KStarCredentialService {
	return &kstarCredentialService{
		repo:   repo,
		logger: logger,
	}
}

func (s *kstarCredentialService) FindAll() ([]model.KStarCredential, error) {
	result, err := s.repo.FindAll()
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	return result, nil
}

func (s *kstarCredentialService) Create(request *domain.CreateKStarCredentialRequest) error {
	if err := util.ValidateStruct(request); err != nil {
		return err
	}

	credential := new(model.KStarCredential)
	credential.Username = request.Username
	credential.Password = request.Password
	credential.Owner = request.Owner

	if err := s.repo.Create(credential); err != nil {
		s.logger.Error(err)
		return err
	}

	return nil
}

func (s *kstarCredentialService) Update(id int64, request *domain.UpdateKStarCredentialRequest) error {
	if err := util.ValidateStruct(request); err != nil {
		return err
	}

	credential := new(model.KStarCredential)
	credential.Username = request.Username
	credential.Password = request.Password
	credential.Owner = request.Owner

	if err := s.repo.Update(id, credential); err != nil {
		s.logger.Error(err)
		return err
	}

	return nil
}

func (s *kstarCredentialService) Delete(id int64) error {
	if err := s.repo.Delete(id); err != nil {
		s.logger.Error(err)
		return err
	}

	return nil
}

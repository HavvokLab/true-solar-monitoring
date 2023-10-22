package service

import (
	"github.com/HavvokLab/true-solar-monitoring/domain"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/HavvokLab/true-solar-monitoring/repo"
)

type KibanaCredentialService interface {
	FindOne(*domain.UserContext) (*model.KibanaCredential, error)
}

type kibanaCredentialService struct {
	repo   repo.KibanaCredentialRepo
	logger logger.Logger
}

func NewKibanaCredentialService(repo repo.KibanaCredentialRepo, logger logger.Logger) KibanaCredentialService {
	return &kibanaCredentialService{repo: repo, logger: logger}
}

func (s *kibanaCredentialService) FindOne(utx *domain.UserContext) (*model.KibanaCredential, error) {
	result, err := s.repo.FindOne()
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	return result, nil
}

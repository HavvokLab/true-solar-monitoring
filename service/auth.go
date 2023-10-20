package service

import (
	"net/http"

	"github.com/HavvokLab/true-solar-monitoring/domain"
	"github.com/HavvokLab/true-solar-monitoring/errors"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/repo"
	"github.com/HavvokLab/true-solar-monitoring/util"
)

type AuthService interface {
	Login(*domain.LoginRequest) (*domain.LoginResponse, error)
}

type authService struct {
	userRepo repo.UserRepo
	logger   logger.Logger
}

func NewAuthService(userRepo repo.UserRepo, logger logger.Logger) AuthService {
	return &authService{
		userRepo: userRepo,
		logger:   logger,
	}
}

func (s *authService) Login(req *domain.LoginRequest) (*domain.LoginResponse, error) {
	if err := util.ValidateStruct(req); err != nil {
		return nil, err
	}

	user, err := s.userRepo.FindByUsername(req.Username)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	if !util.CompareHash(req.Password, user.HashedPassword) {
		err := errors.NewServerError(http.StatusBadRequest, "invalid username or password")
		s.logger.Error(err)
		return nil, err
	}

	return &domain.LoginResponse{
		AccessToken: "token",
	}, nil
}

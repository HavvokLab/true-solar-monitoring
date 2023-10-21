package service

import (
	"net/http"
	"time"

	"github.com/HavvokLab/true-solar-monitoring/config"
	"github.com/HavvokLab/true-solar-monitoring/domain"
	"github.com/HavvokLab/true-solar-monitoring/errors"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/HavvokLab/true-solar-monitoring/repo"
	"github.com/HavvokLab/true-solar-monitoring/util"
	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	Login(*domain.LoginRequest) (*domain.LoginResponse, error)
	Register(*domain.RegisterRequest) error
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

	accessToken, err := createAccessToken(user)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	return &domain.LoginResponse{
		AccessToken: accessToken,
	}, nil
}

func createAccessToken(user *model.User) (string, error) {
	conf := config.GetConfig().Authentication
	claims := new(domain.AccessToken)
	claims.ID = user.ID
	claims.DisplayName = user.Username
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(conf.Expired)))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(conf.Secret))
}

func (s *authService) Register(req *domain.RegisterRequest) error {
	if err := util.ValidateStruct(req); err != nil {
		s.logger.Error(err)
		return err
	}

	hashed, err := util.GenerateHash(req.Password)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	user := new(model.User)
	user.Username = req.Username
	user.HashedPassword = hashed
	if err := s.userRepo.Create(user); err != nil {
		s.logger.Error(err)
		return err
	}

	return nil
}

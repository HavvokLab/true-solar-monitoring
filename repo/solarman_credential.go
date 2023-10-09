package repo

import (
	"github.com/HavvokLab/true-solar-monitoring/model"
	"gorm.io/gorm"
)

type SolarmanCredentialRepo interface {
	GetCredentials() ([]model.SolarmanCredential, error)
}

type solarmanCredentialRepo struct {
	db *gorm.DB
}

func NewSolarmanCredentialRepo(db *gorm.DB) SolarmanCredentialRepo {
	return &solarmanCredentialRepo{db: db}
}

func (r *solarmanCredentialRepo) GetCredentials() ([]model.SolarmanCredential, error) {
	var credentials []model.SolarmanCredential
	tx := r.db.Session(&gorm.Session{})
	if err := tx.Find(&credentials).Error; err != nil {
		return nil, err
	}

	return credentials, nil
}

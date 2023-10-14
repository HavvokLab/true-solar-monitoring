package repo

import (
	"github.com/HavvokLab/true-solar-monitoring/model"
	"gorm.io/gorm"
)

type KStarCredentialRepo interface {
	GetCredentials() ([]model.KStarCredential, error)
}

type kStarCredentialRepo struct {
	db *gorm.DB
}

func NewKStarCredentialRepo(db *gorm.DB) KStarCredentialRepo {
	return &kStarCredentialRepo{db: db}
}

func (r *kStarCredentialRepo) GetCredentials() ([]model.KStarCredential, error) {
	var credentials []model.KStarCredential
	tx := r.db.Session(&gorm.Session{})
	if err := tx.Find(&credentials).Error; err != nil {
		return nil, err
	}

	return credentials, nil
}

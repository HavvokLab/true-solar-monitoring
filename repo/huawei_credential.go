package repo

import (
	"github.com/HavvokLab/true-solar-monitoring/constant"
	"github.com/HavvokLab/true-solar-monitoring/model"
	"gorm.io/gorm"
)

type HuaweiCredentialRepo interface {
	GetCredentialsByOwner(constant.Owner) ([]model.HuaweiCredential, error)
	GetCredentials() ([]model.HuaweiCredential, error)
}

type huaweiCredentialRepo struct {
	db *gorm.DB
}

func NewHuaweiCredentialRepo(db *gorm.DB) HuaweiCredentialRepo {
	return &huaweiCredentialRepo{db: db}
}

func (r *huaweiCredentialRepo) GetCredentialsByOwner(owner constant.Owner) ([]model.HuaweiCredential, error) {
	var credentials []model.HuaweiCredential
	tx := r.db.Session(&gorm.Session{})
	if err := tx.Find(&credentials, "owner = ?", owner).Error; err != nil {
		return nil, err
	}

	return credentials, nil
}

func (r *huaweiCredentialRepo) GetCredentials() ([]model.HuaweiCredential, error) {
	var credentials []model.HuaweiCredential
	tx := r.db.Session(&gorm.Session{})
	if err := tx.Find(&credentials).Error; err != nil {
		return nil, err
	}

	return credentials, nil
}

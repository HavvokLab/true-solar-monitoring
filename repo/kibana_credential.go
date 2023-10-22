package repo

import (
	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/HavvokLab/true-solar-monitoring/util"
	"gorm.io/gorm"
)

type KibanaCredentialRepo interface {
	FindOne() (*model.KibanaCredential, error)
}

type kibanaCredentialRepo struct {
	db *gorm.DB
}

func NewKibanaCredentialRepo(db *gorm.DB) KibanaCredentialRepo {
	return &kibanaCredentialRepo{db: db}
}

func (r *kibanaCredentialRepo) FindOne() (*model.KibanaCredential, error) {
	var credential model.KibanaCredential
	tx := r.db.Session(&gorm.Session{})
	if err := tx.First(&credential).Error; err != nil {
		return nil, util.TranslateSqliteError(err)
	}

	return &credential, nil
}

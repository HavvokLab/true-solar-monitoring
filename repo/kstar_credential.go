package repo

import (
	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/HavvokLab/true-solar-monitoring/util"
	"gorm.io/gorm"
)

type KStarCredentialRepo interface {
	FindAll() ([]model.KStarCredential, error)
	Create(credential *model.KStarCredential) error
	Update(id int64, credential *model.KStarCredential) error
	Delete(id int64) error
}

type kStarCredentialRepo struct {
	db *gorm.DB
}

func NewKStarCredentialRepo(db *gorm.DB) KStarCredentialRepo {
	return &kStarCredentialRepo{db: db}
}

func (r *kStarCredentialRepo) FindAll() ([]model.KStarCredential, error) {
	var credentials []model.KStarCredential
	tx := r.db.Session(&gorm.Session{})
	if err := tx.Find(&credentials).Error; err != nil {
		return nil, err
	}

	return credentials, nil
}

func (r *kStarCredentialRepo) Create(credential *model.KStarCredential) error {
	tx := r.db.Session(&gorm.Session{})
	if err := tx.Create(credential).Error; err != nil {
		return util.TranslateSqliteError(err)
	}

	return nil
}

func (r *kStarCredentialRepo) Update(id int64, credential *model.KStarCredential) error {
	tx := r.db.Session(&gorm.Session{})
	if err := tx.Where("id = ?", id).Updates(credential).Error; err != nil {
		return util.TranslateSqliteError(err)
	}

	return nil
}

func (r *kStarCredentialRepo) Delete(id int64) error {
	tx := r.db.Session(&gorm.Session{})
	if err := tx.Where("id = ?", id).Delete(&model.KStarCredential{}).Error; err != nil {
		return util.TranslateSqliteError(err)
	}

	return nil
}

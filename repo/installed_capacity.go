package repo

import (
	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/HavvokLab/true-solar-monitoring/util"
	"gorm.io/gorm"
)

type InstalledCapacityRepo interface {
	FindOne() (*model.InstalledCapacity, error)
	Update(id int64, installedCapacity *model.InstalledCapacity) error
}

type installedCapacityRepo struct {
	db *gorm.DB
}

func NewInstalledCapacityRepo(db *gorm.DB) InstalledCapacityRepo {
	return &installedCapacityRepo{db: db}
}

func (r *installedCapacityRepo) FindOne() (*model.InstalledCapacity, error) {
	tx := r.db.Session(&gorm.Session{})
	var installedCapacity model.InstalledCapacity
	err := tx.First(&installedCapacity).Error
	if err != nil {
		return nil, util.TranslateSqliteError(err)
	}

	return &installedCapacity, nil
}

func (r *installedCapacityRepo) Update(id int64, installedCapacity *model.InstalledCapacity) error {
	tx := r.db.Session(&gorm.Session{})
	err := tx.Where("id = ?", id).Updates(installedCapacity).Error
	if err != nil {
		return util.TranslateSqliteError(err)
	}

	return nil
}

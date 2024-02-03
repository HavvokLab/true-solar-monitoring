package repo

import (
	"github.com/HavvokLab/true-solar-monitoring/model"
	"gorm.io/gorm"
)

type HuaweiAltervimRepo interface {
	BatchInsertPlants([]model.HuaweiAltervimPlant) error
	BatchInsertDevices([]model.HuaweiAltervimDevice) error
	GetPlants() ([]model.HuaweiAltervimPlant, error)
	GetDevices() ([]model.HuaweiAltervimDevice, error)
	DeletePlantNotIn(codes []string) error
	DeleteDeviceNotIn(ids []int) error
}

type huaweiAltervimRepo struct {
	db *gorm.DB
}

func NewHuaweiAltervimRepo(db *gorm.DB) *huaweiAltervimRepo {
	return &huaweiAltervimRepo{
		db: db,
	}
}

func (r *huaweiAltervimRepo) BatchInsertPlants(data []model.HuaweiAltervimPlant) error {
	tx := r.db.Session(&gorm.Session{})
	if err := tx.CreateInBatches(data, 100).Error; err != nil {
		return err
	}

	return nil
}

func (r *huaweiAltervimRepo) BatchInsertDevices(data []model.HuaweiAltervimDevice) error {
	tx := r.db.Session(&gorm.Session{})
	if err := tx.CreateInBatches(data, 100).Error; err != nil {
		return err
	}

	return nil
}

func (r *huaweiAltervimRepo) GetPlants() ([]model.HuaweiAltervimPlant, error) {
	tx := r.db.Session(&gorm.Session{})
	data := []model.HuaweiAltervimPlant{}
	if err := tx.Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (r *huaweiAltervimRepo) GetDevices() ([]model.HuaweiAltervimDevice, error) {
	tx := r.db.Session(&gorm.Session{})
	data := []model.HuaweiAltervimDevice{}
	if err := tx.Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (r *huaweiAltervimRepo) DeletePlantNotIn(codes []string) error {
	tx := r.db.Session(&gorm.Session{})
	if err := tx.
		Where("code NOT IN ?", codes).
		Delete(&model.HuaweiAltervimPlant{}).
		Error; err != nil {
		return err
	}

	return nil
}

func (r *huaweiAltervimRepo) DeleteDeviceNotIn(ids []int) error {
	tx := r.db.Session(&gorm.Session{})
	if err := tx.
		Where("id NOT IN ?", ids).
		Delete(&model.HuaweiAltervimPlant{}).
		Error; err != nil {
		return err
	}

	return nil
}

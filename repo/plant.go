package repo

import (
	"github.com/HavvokLab/true-solar-monitoring/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PlantRepo interface {
	FindAll() ([]*model.Plant, error)
	FindOneByName(name string) (*model.Plant, error)
	FindAllWithPagination(offset, limit int) ([]*model.Plant, error)
	Create(*model.Plant) error
	BatchCreate([]*model.Plant) error
	BatchUpsertAvailable([]*model.Plant) error
	Count() (int64, error)
	Delete(id int64) error
}

type plantRepo struct {
	db *gorm.DB
}

func NewPlantRepo(db *gorm.DB) PlantRepo {
	return &plantRepo{db}
}

func (r *plantRepo) Create(plant *model.Plant) error {
	tx := r.db.Session(&gorm.Session{})
	conflictUpdateData := map[string]interface{}{
		"available": true,
	}

	onConflict := clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}},
		DoUpdates: clause.Assignments(conflictUpdateData),
	}

	return tx.Clauses(onConflict).Create(plant).Error
}

func (r *plantRepo) BatchCreate(plants []*model.Plant) error {
	tx := r.db.Session(&gorm.Session{})
	conflictUpdateData := map[string]interface{}{
		"available": true,
	}

	onConflict := clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}},
		DoUpdates: clause.Assignments(conflictUpdateData),
	}

	return tx.Clauses(onConflict).CreateInBatches(plants, 100).Error
}

func (r *plantRepo) Count() (int64, error) {
	tx := r.db.Session(&gorm.Session{})
	var count int64
	if err := tx.Model(&model.Plant{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (r *plantRepo) FindAll() ([]*model.Plant, error) {
	tx := r.db.Session(&gorm.Session{})
	var plants []*model.Plant
	if err := tx.Find(&plants).Error; err != nil {
		return nil, err
	}

	return plants, nil
}

func (r *plantRepo) FindAllWithPagination(offset, limit int) ([]*model.Plant, error) {
	tx := r.db.Session(&gorm.Session{})
	var plants []*model.Plant
	if err := tx.Offset(offset).Limit(limit).Find(&plants).Error; err != nil {
		return nil, err
	}

	return plants, nil
}

func (r *plantRepo) Delete(id int64) error {
	tx := r.db.Session(&gorm.Session{})
	if err := tx.Delete(&model.Plant{}, "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}

func (r *plantRepo) BatchUpsertAvailable(plants []*model.Plant) error {
	tx := r.db.Session(&gorm.Session{})

	conflictUpdateData := map[string]interface{}{
		"available": true,
	}

	onConflict := clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}},
		DoUpdates: clause.Assignments(conflictUpdateData),
	}

	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.Raw("UPDATE tbl_plants SET available = false").Error; err != nil {
			return err
		}

		return tx.Clauses(onConflict).CreateInBatches(plants, 100).Error
	})

	return err
}

func (r plantRepo) FindOneByName(name string) (*model.Plant, error) {
	tx := r.db.Session(&gorm.Session{})
	data := model.Plant{}
	if err := tx.Find(&data, "name=?", name).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

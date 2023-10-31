package repo

import (
	"github.com/HavvokLab/true-solar-monitoring/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PlantRepo interface {
	FindAll() ([]*model.Plant, error)
	Create(*model.Plant) error
	BulkCreate([]*model.Plant) error
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

func (r *plantRepo) BulkCreate(plants []*model.Plant) error {
	tx := r.db.Session(&gorm.Session{})
	conflictUpdateData := map[string]interface{}{
		"available": true,
	}

	onConflict := clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}},
		DoUpdates: clause.Assignments(conflictUpdateData),
	}

	return tx.Clauses(onConflict).Create(plants).Error
}

func (r *plantRepo) FindAll() ([]*model.Plant, error) {
	tx := r.db.Session(&gorm.Session{})
	var plants []*model.Plant
	if err := tx.Find(&plants).Error; err != nil {
		return nil, err
	}

	return plants, nil
}

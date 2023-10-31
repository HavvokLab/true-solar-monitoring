package repo

import (
	"github.com/HavvokLab/true-solar-monitoring/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PlantRepo interface {
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

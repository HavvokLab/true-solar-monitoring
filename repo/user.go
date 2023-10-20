package repo

import (
	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/HavvokLab/true-solar-monitoring/util"
	"gorm.io/gorm"
)

type UserRepo interface {
	FindByUsername(username string) (*model.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db}
}

func (r *userRepo) FindByUsername(username string) (*model.User, error) {
	tx := r.db.Session(&gorm.Session{})
	var user model.User
	if err := tx.Where("username = ?", username).Take(&user).Error; err != nil {
		return nil, util.TranslateSqliteError(err)
	}

	return &user, nil
}

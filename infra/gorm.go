package infra

import (
	"path/filepath"

	"github.com/HavvokLab/true-solar-monitoring/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewGormDB(paths ...string) (*gorm.DB, error) {
	var path string = "database.db"
	if len(paths) > 0 {
		path = paths[0]
	}

	db, err := gorm.Open(sqlite.Open(filepath.Join(config.OriginPath, path)), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

var GormDB *gorm.DB

func InitGormDB() {
	var err error
	GormDB, err = NewGormDB()
	if err != nil {
		panic(err)
	}
}

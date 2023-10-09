package infra

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewGormDB(paths ...string) (*gorm.DB, error) {
	var path string = "database.db"
	if len(paths) > 0 {
		path = paths[0]
	}

	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

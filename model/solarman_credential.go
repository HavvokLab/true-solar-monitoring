package model

import (
	"time"
)

type SolarmanCredential struct {
	ID        int64      `gorm:"column:id"`
	Username  string     `gorm:"column:username"`
	Password  string     `gorm:"column:password"`
	AppSecret string     `gorm:"column:app_secret"`
	AppID     string     `gorm:"column:app_id"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

func (*SolarmanCredential) TableName() string {
	return "tbl_solarman_credentials"
}

package model

import (
	"time"
)

type SolarmanCredential struct {
	ID        int64      `gorm:"column:id" json:"id"`
	Username  string     `gorm:"column:username" json:"username"`
	Password  string     `gorm:"column:password" json:"password"`
	AppSecret string     `gorm:"column:app_secret" json:"app_secret"`
	AppID     string     `gorm:"column:app_id" json:"app_id"`
	Owner     string     `gorm:"column:owner" json:"owner"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (*SolarmanCredential) TableName() string {
	return "tbl_solarman_credentials"
}

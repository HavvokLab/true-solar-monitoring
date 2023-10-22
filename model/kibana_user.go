package model

import (
	"time"
)

type KibanaCredential struct {
	ID        int64      `gorm:"column:id"`
	Username  string     `gorm:"column:username"`
	Password  string     `gorm:"column:password"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

func (*KibanaCredential) TableName() string {
	return "tbl_kibana_credentials"
}

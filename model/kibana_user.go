package model

import (
	"time"
)

type KibanaCredential struct {
	ID        int64      `gorm:"column:id" json:"id"`
	Username  string     `gorm:"column:username" json:"username"`
	Password  string     `gorm:"column:password" json:"password"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (*KibanaCredential) TableName() string {
	return "tbl_kibana_credentials"
}

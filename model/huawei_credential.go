package model

import (
	"time"
)

type HuaweiCredential struct {
	ID        int64      `gorm:"column:id"`
	Username  string     `gorm:"column:username"`
	Password  string     `gorm:"column:password"`
	Owner     string     `gorm:"column:owner"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

func (*HuaweiCredential) TableName() string {
	return "tbl_huawei_credentials"
}

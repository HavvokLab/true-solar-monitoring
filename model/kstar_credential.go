package model

import "time"

type KStarCredential struct {
	ID        int64      `gorm:"column:id"`
	Username  string     `gorm:"column:username"`
	Password  string     `gorm:"column:password"`
	Owner     string     `gorm:"column:owner"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

func (*KStarCredential) TableName() string {
	return "tbl_kstar_credentials"
}

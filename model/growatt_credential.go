package model

import "time"

type GrowattCredential struct {
	ID        int64      `gorm:"column:id" json:"id"`
	Username  string     `gorm:"column:username" json:"username"`
	Password  string     `gorm:"column:password" json:"password"`
	Token     string     `gorm:"column:token" json:"token"`
	Owner     string     `gorm:"column:owner" json:"owner"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (*GrowattCredential) TableName() string {
	return "tbl_growatt_credentials"
}

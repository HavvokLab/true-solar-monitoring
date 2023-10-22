package model

import (
	"time"

	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

type User struct {
	ID             string     `gorm:"column:id" json:"id"`
	Username       string     `gorm:"column:username" json:"username"`
	HashedPassword string     `gorm:"column:hashed_password" json:"hashed_password"`
	CreatedAt      *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      *time.Time `gorm:"column:created_at" json:"created_at"`
}

func (*User) TableName() string {
	return "tbl_users"
}

func (u *User) BeforeCreate(*gorm.DB) error {
	id := ksuid.New()
	u.ID = id.String()
	return nil
}

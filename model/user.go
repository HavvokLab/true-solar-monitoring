package model

import (
	"time"

	"github.com/segmentio/ksuid"
)

type User struct {
	ID             string     `gorm:"column:id"`
	Username       string     `gorm:"column:username"`
	HashedPassword string     `gorm:"column:hashed_password"`
	CreatedAt      *time.Time `gorm:"column:created_at"`
	UpdatedAt      *time.Time `gorm:"column:created_at"`
}

func (*User) TableName() string {
	return "tbl_users"
}

func (u *User) BeforeCreate() error {
	id := ksuid.New()
	u.ID = id.String()
	return nil
}

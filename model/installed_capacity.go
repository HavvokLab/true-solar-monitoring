package model

import (
	"time"
)

type InstalledCapacity struct {
	ID               int64      `gorm:"column:id"`
	EfficiencyFactor float64    `gorm:"column:efficiency_factor"`
	FocusHour        int        `gorm:"column:focus_hour"`
	CreatedAt        *time.Time `gorm:"column:created_at"`
	UpdatedAt        *time.Time `gorm:"column:updated_at"`
}

func (*InstalledCapacity) TableName() string {
	return "tbl_installed_capacity"
}

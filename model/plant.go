package model

import "time"

type Plant struct {
	ID                int64      `gorm:"column:id" json:"id"`
	Name              string     `gorm:"column:name" json:"name"`
	Area              *string    `gorm:"column:area" json:"area"`
	VendorType        string     `gorm:"column:vendor_type" json:"vendor_type"`
	InstalledCapacity float64    `gorm:"column:installed_capacity" json:"installed_capacity"`
	Latitude          *float64   `gorm:"column:lat" json:"lat"`
	Longitude         *float64   `gorm:"column:long" json:"long"`
	Owner             *string    `gorm:"column:owner" json:"owner"`
	Available         bool       `gorm:"column:available" json:"available"`
	CreatedAt         *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt         *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (p *Plant) TableName() string {
	return "tbl_plants"
}

package model

import (
	"fmt"
	"time"

	"github.com/HavvokLab/true-solar-monitoring/constant"
)

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

func (p Plant) CsvRow() []string {
	result := make([]string, 0)
	result = append(result, p.Name)
	if p.Owner != nil {
		result = append(result, *p.Owner)
	} else {
		result = append(result, string(constant.TRUE_OWNER))
	}

	result = append(result, p.VendorType)

	if p.Area != nil {
		result = append(result, *p.Area)
	} else {
		result = append(result, "-")
	}

	if p.Available {
		result = append(result, "true")
	} else {
		result = append(result, "false")
	}

	result = append(result, fmt.Sprintf("%f", p.InstalledCapacity))

	if p.Latitude != nil {
		result = append(result, fmt.Sprintf("%f", *p.Latitude))
	} else {
		result = append(result, "-")
	}

	if p.Longitude != nil {
		result = append(result, fmt.Sprintf("%f", *p.Longitude))
	} else {
		result = append(result, "-")
	}

	return result
}

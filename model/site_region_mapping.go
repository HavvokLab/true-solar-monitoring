package model

import (
	"time"
)

type SiteRegionMapping struct {
	ID        int64      `gorm:"column:id"`
	Code      string     `gorm:"column:code"`
	Name      string     `gorm:"column:name"`
	Area      *string    `gorm:"column:area"`
	CreatedAt *time.Time `gorm:"created_at"`
	UpdatedAt *time.Time `gorm:"updated_at"`
}

func (*SiteRegionMapping) TableName() string {
	return "tbl_site_region_mapping"
}

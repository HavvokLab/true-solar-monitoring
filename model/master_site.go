package model

import (
	"fmt"

	"go.openly.dev/pointy"
)

type MasterSite struct {
	Vendor            *string  `csv:"vendor,omitempty" json:"vendor,omitempty" mapstructure:"vendor"`
	Area              *string  `csv:"area,omitempty" json:"area,omitempty" mapstructure:"area"`
	SiteName          *string  `csv:"site_name,omitempty" json:"site_name,omitempty" mapstructure:"site_name"`
	InstalledCapacity *float64 `csv:"installed_capacity,omitempty" json:"installed_capacity,omitempty" mapstructure:"installed_capacity"`
	Latitude          *float64 `csv:"latitude,omitempty" json:"latitude,omitempty" mapstructure:"latitude"`
	Longitude         *float64 `csv:"longitude,omitempty" json:"longitude,omitempty" mapstructure:"longitude"`
}

func (m *MasterSite) GetVendor(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(m.Vendor, value)
}

func (m *MasterSite) GetArea(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(m.Area, value)
}

func (m *MasterSite) GetSiteName(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(m.SiteName, value)
}

func (m *MasterSite) GetInstalledCapacity(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.InstalledCapacity, value)
}

func (m *MasterSite) GetLatitude(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.Latitude, value)
}

func (m *MasterSite) GetLongitude(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.Longitude, value)
}

func (m *MasterSite) GetKey() string {
	return fmt.Sprintf("%v,%v,%v", m.GetVendor(), m.GetArea(), m.GetSiteName())
}

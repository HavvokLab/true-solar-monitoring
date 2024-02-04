package model

import (
	"time"

	"go.openly.dev/pointy"
)

type HuaweiAltervimPlant struct {
	Code               string     `gorm:"column:code" json:"code"`
	Name               *string    `gorm:"column:name" json:"name"`
	Address            *string    `gorm:"column:address" json:"address"`
	Longitude          *string    `gorm:"column:longitude" json:"longitude"`
	Latitude           *string    `gorm:"column:latitude" json:"latitude"`
	Capacity           *float64   `gorm:"column:capacity" json:"capacity"`
	ContactPerson      *string    `gorm:"column:contact_person" json:"contact_person"`
	ContactMethod      *string    `gorm:"column:contact_method" json:"contact_method"`
	GridConnectionData *string    `gorm:"column:grid_connection_data" json:"grid_connection_data"`
	CreatedAt          *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt          *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (*HuaweiAltervimPlant) TableName() string {
	return "tbl_huawei_altervim_plants"
}

func (p *HuaweiAltervimPlant) GetPlantName(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.Name, value)
}

func (p *HuaweiAltervimPlant) GetPlantAddress(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.Address, value)
}

func (p *HuaweiAltervimPlant) GetLongitude(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.Longitude, value)
}

func (p *HuaweiAltervimPlant) GetLatitude(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.Latitude, value)
}

func (p *HuaweiAltervimPlant) GetCapacity(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.Capacity, value)
}

func (p *HuaweiAltervimPlant) GetContactPerson(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.ContactPerson, value)
}

func (p *HuaweiAltervimPlant) GetContactMethod(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.ContactMethod, value)
}

func (p *HuaweiAltervimPlant) GetGridConnectionDate(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.GridConnectionData, value)
}

type HuaweiAltervimDevice struct {
	ID              int        `gorm:"column:id" json:"id"`
	SerialNumber    *string    `gorm:"column:serial_number" json:"serial_number"`
	Name            *string    `gorm:"column:name" json:"name"`
	TypeID          *int       `gorm:"column:type_id" json:"type_id"`
	InverterModel   *string    `gorm:"column:inverter_model" json:"inverter_model"`
	Latitude        *float64   `gorm:"column:latitude" json:"latitude"`
	Longitude       *float64   `gorm:"column:longitude" json:"longitude"`
	SoftwareVersion *string    `gorm:"column:software_version" json:"software_version"`
	PlantCode       *string    `gorm:"column:plant_code" json:"plant_code"`
	CreatedAt       *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (*HuaweiAltervimDevice) TableName() string {
	return "tbl_huawei_altervim_devices"
}

func (d *HuaweiAltervimDevice) GetSN(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.SerialNumber, value)
}

func (d *HuaweiAltervimDevice) GetName(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.Name, value)
}

func (d *HuaweiAltervimDevice) GetTypeID(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(d.TypeID, value)
}

func (d *HuaweiAltervimDevice) GetInverterModel(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.InverterModel, value)
}

func (d *HuaweiAltervimDevice) GetLatitude(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(d.Latitude, value)
}

func (d *HuaweiAltervimDevice) GetLongitude(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(d.Longitude, value)
}

func (d *HuaweiAltervimDevice) GetSoftwareVersion(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.SoftwareVersion, value)
}

func (d *HuaweiAltervimDevice) GetPlantCode(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.PlantCode, value)
}

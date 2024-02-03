package model

import "time"

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
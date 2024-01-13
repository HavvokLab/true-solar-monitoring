package model

import (
	"fmt"
	"time"

	"github.com/HavvokLab/true-solar-monitoring/util"
	"go.openly.dev/pointy"
)

type MonthlyProductionDocument struct {
	Date               *time.Time `json:"date"`
	VendorType         *string    `json:"vendor_type"`
	Area               *string    `json:"area"`
	SiteName           *string    `json:"name"`
	Owner              *string    `json:"owner"`
	InstalledCapacity  *float64   `json:"installed_capacity"`
	MonthlyProduction  *float64   `json:"monthly_production"`
	Latitude           *float64   `json:"lat"`
	Longitude          *float64   `json:"lng"`
	Location           *string    `json:"location"`
	Target             *float64   `json:"target"`
	ProductionToTarget *float64   `json:"production_to_target"`
	Criteria           *string    `json:"criteria"`
}

func (d *MonthlyProductionDocument) parseString(data string) *string {
	if util.EmptyString(data) {
		return nil
	}

	return pointy.String(data)
}

func (d *MonthlyProductionDocument) SetDate(data *time.Time) {
	d.Date = data
}

func (d *MonthlyProductionDocument) SetVendorType(data string) {
	d.VendorType = d.parseString(data)
}

func (d *MonthlyProductionDocument) SetArea(data string) {
	d.Area = d.parseString(data)
}

func (d *MonthlyProductionDocument) SetSiteName(data string) {
	d.SiteName = d.parseString(data)
}

func (d *MonthlyProductionDocument) SetOwner(data string) {
	d.Owner = d.parseString(data)
}

func (d *MonthlyProductionDocument) SetInstalledCapacity(data *float64) {
	d.InstalledCapacity = data
}

func (d *MonthlyProductionDocument) SetMonthlyProduction(data *float64) {
	d.MonthlyProduction = data
}

func (d *MonthlyProductionDocument) SetLatitude(data *float64) {
	d.Latitude = data
}

func (d *MonthlyProductionDocument) SetLongitude(data *float64) {
	d.Longitude = data
}

func (d *MonthlyProductionDocument) SetLocation(lat, long *float64) {
	if lat != nil && long != nil {
		d.Location = pointy.String(fmt.Sprintf("%f,%f", *lat, *long))
	}
}

func (d *MonthlyProductionDocument) SetProductionToTarget(data *float64) {
	d.ProductionToTarget = data
}

func (d *MonthlyProductionDocument) SetTarget(data *float64) {
	d.Target = data
}

func (d *MonthlyProductionDocument) SetCriteria(data *float64) {
	if data == nil {
		d.Criteria = pointy.String("-")
		return
	}

	value := pointy.Float64Value(data, 0)
	if value >= 100 {
		d.Criteria = pointy.String(">=100%")
	} else if value >= 80 {
		d.Criteria = pointy.String(">=80%")
	} else if value >= 60 {
		d.Criteria = pointy.String(">=60%")
	} else if value >= 50 {
		d.Criteria = pointy.String(">=50%")
	} else if value >= 30 {
		d.Criteria = pointy.String(">=30%")
	} else if value > 0 {
		d.Criteria = pointy.String("<30%")
	} else {
		d.Criteria = pointy.String("=0%")
	}
}

func (d *MonthlyProductionDocument) ClearZeroValue() {
	if d.MonthlyProduction == nil {
		d.MonthlyProduction = pointy.Float64(0)
	}

	if d.Target == nil {
		d.Target = pointy.Float64(0)
	}

	if d.ProductionToTarget == nil {
		d.ProductionToTarget = pointy.Float64(0)
	}
}

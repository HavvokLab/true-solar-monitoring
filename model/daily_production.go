package model

import (
	"time"

	"github.com/HavvokLab/true-solar-monitoring/util"
	"go.openly.dev/pointy"
)

type DailyProductionDocument struct {
	Date               *time.Time `json:"date"`
	VendorType         *string    `json:"vendor_type"`
	Area               *string    `json:"area"`
	SiteName           *string    `json:"site_name"`
	InstalledCapacity  *float64   `json:"installed_capacity"`
	DailyProduction    *float64   `json:"daily_production"`
	Latitude           *float64   `json:"lat"`
	Longitude          *float64   `json:"lng"`
	Target             *float64   `json:"target"`
	ProductionToTarget *float64   `json:"production_to_target"`
	Criteria           *string    `json:"criteria"`
}

func (d *DailyProductionDocument) parseString(data string) *string {
	if util.EmptyString(data) {
		return nil
	}

	return pointy.String(data)
}

func (d *DailyProductionDocument) SetDate(data *time.Time) {
	d.Date = data
}

func (d *DailyProductionDocument) SetVendorType(data string) {
	d.VendorType = d.parseString(data)
}

func (d *DailyProductionDocument) SetArea(data string) {
	d.Area = d.parseString(data)
}

func (d *DailyProductionDocument) SetSiteName(data string) {
	d.SiteName = d.parseString(data)
}

func (d *DailyProductionDocument) SetInstalledCapacity(data *float64) {
	d.InstalledCapacity = data
}

func (d *DailyProductionDocument) SetDailyProduction(data *float64) {
	d.DailyProduction = data
}

func (d *DailyProductionDocument) SetLatitude(data *float64) {
	d.Latitude = data
}

func (d *DailyProductionDocument) SetLongitude(data *float64) {
	d.Longitude = data
}

func (d *DailyProductionDocument) SetProductionToTarget(data *float64) {
	d.ProductionToTarget = data
}

func (d *DailyProductionDocument) SetTarget(data *float64) {
	d.Target = data
}

func (d *DailyProductionDocument) SetCriteria(data *float64) {
	if data == nil {
		d.Criteria = nil
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
	} else {
		d.Criteria = pointy.String("-")
	}
}

func (d *DailyProductionDocument) ClearZeroValue() {
	if d.DailyProduction == nil {
		d.DailyProduction = pointy.Float64(0)
	}

	if d.Target == nil {
		d.Target = pointy.Float64(0)
	}

	if d.ProductionToTarget == nil {
		d.ProductionToTarget = pointy.Float64(0)
	}
}

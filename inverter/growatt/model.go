package growatt

import (
	"strconv"
	"time"

	"go.openly.dev/pointy"
)

type DefaultResponse struct {
	ErrorCode *int    `json:"error_code,omitempty"`
	ErrorMsg  *string `json:"error_msg,omitempty"`
}

func (r *DefaultResponse) GetErrorCode(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(r.ErrorCode, value)
}

func (r *DefaultResponse) GetErrorMsg(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(r.ErrorMsg, value)
}

type ErrorResponse struct {
	DefaultResponse
	Data *string `json:"data,omitempty"`
}

func (r *ErrorResponse) GetData(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(r.Data, value)
}

// |=> GetPlantList
type PlantItem struct {
	Status       *int         `json:"status,omitempty"`
	Locale       *string      `json:"locale,omitempty"`
	TotalEnergy  *string      `json:"total_energy,omitempty"`
	Operator     *string      `json:"operator,omitempty"`
	Country      *string      `json:"country,omitempty"`
	City         *string      `json:"city,omitempty"`
	CurrentPower *string      `json:"current_power,omitempty"`
	CreateDate   *string      `json:"create_date,omitempty"`
	ImageURL     *string      `json:"image_url,omitempty"`
	PlantID      *int         `json:"plant_id,omitempty"`
	Name         *string      `json:"name,omitempty"`
	Installer    *string      `json:"installer,omitempty"`
	UserID       *int         `json:"user_id,omitempty"`
	Longitude    *string      `json:"longitude,omitempty"`
	Latitude     *string      `json:"latitude,omitempty"`
	PeakPower    *float64     `json:"peak_power,omitempty"`
	LatitudeD    *interface{} `json:"latitude_d,omitempty"`
	LatitudeF    *interface{} `json:"latitude_f,omitempty"`
}

func (p *PlantItem) GetStatus(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	return pointy.IntValue(p.Status, value)
}

func (p *PlantItem) GetLocale(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	return pointy.StringValue(p.Locale, value)
}

func (p *PlantItem) GetTotalEnergy(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	return pointy.StringValue(p.TotalEnergy, value)
}

func (p *PlantItem) GetOperator(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	return pointy.StringValue(p.Operator, value)
}

func (p *PlantItem) GetCountry(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	return pointy.StringValue(p.Country, value)
}

func (p *PlantItem) GetCity(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	return pointy.StringValue(p.City, value)
}

func (p *PlantItem) GetCurrentPower(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	return pointy.StringValue(p.CurrentPower, value)
}

func (p *PlantItem) GetCreateDate(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	return pointy.StringValue(p.CreateDate, value)
}

func (p *PlantItem) GetImageURL(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	return pointy.StringValue(p.ImageURL, value)
}

func (p *PlantItem) GetPlantID(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	return pointy.IntValue(p.PlantID, value)
}

func (p *PlantItem) GetName(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	return pointy.StringValue(p.Name, value)
}

func (p *PlantItem) GetInstaller(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	return pointy.StringValue(p.Installer, value)
}

func (p *PlantItem) GetUserID(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	return pointy.IntValue(p.UserID, value)
}

func (p *PlantItem) GetLongitude(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	return pointy.StringValue(p.Longitude, value)
}

func (p *PlantItem) GetLatitude(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	return pointy.StringValue(p.Latitude, value)
}

func (p *PlantItem) GetPeakPower(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	return pointy.Float64Value(p.PeakPower, value)
}

type PlantData struct {
	Plants []*PlantItem `json:"plants,omitempty"`
	Count  *int         `json:"count,omitempty"`
}

func (r *PlantData) GetCount(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(r.Count, value)
}

type GetPlantListResponse struct {
	DefaultResponse
	Data *PlantData `json:"data,omitempty"`
}

// |=> GetPlantOverviewInfo
type PlantOverviewInfo struct {
	PeakPowerActual *float64 `json:"peak_power_actual,omitempty"`
	MonthlyEnergy   *string  `json:"monthly_energy,omitempty"`
	LastUpdateTime  *string  `json:"last_update_time,omitempty"`
	CurrentPower    *float64 `json:"current_power,omitempty"`
	Timezone        *string  `json:"timezone,omitempty"`
	YearlyEnergy    *string  `json:"yearly_energy,omitempty"`
	TodayEnergy     *string  `json:"today_energy,omitempty"`
	CarbonOffset    *string  `json:"carbon_offset,omitempty"`
	Efficiency      *string  `json:"efficiency,omitempty"`
	TotalEnergy     *string  `json:"total_energy,omitempty"`
}

func (p *PlantOverviewInfo) GetPeakPowerActual(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.PeakPowerActual, value)
}

func (p *PlantOverviewInfo) GetMonthlyEnergy(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.MonthlyEnergy, value)
}

func (p *PlantOverviewInfo) GetLastUpdateTime(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.LastUpdateTime, value)
}

func (p *PlantOverviewInfo) GetCurrentPower(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.CurrentPower, value)
}

func (p *PlantOverviewInfo) GetTimezone(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.Timezone, value)
}

func (p *PlantOverviewInfo) GetYearlyEnergy(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.YearlyEnergy, value)
}

func (p *PlantOverviewInfo) GetTodayEnergy(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.TodayEnergy, value)
}

func (p *PlantOverviewInfo) GetCarbonOffset(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.CarbonOffset, value)
}

func (p *PlantOverviewInfo) GetEfficiency(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.Efficiency, value)
}

func (p *PlantOverviewInfo) GetTotalEnergy(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.TotalEnergy, value)
}

type GetPlantOverviewInfoResponse struct {
	DefaultResponse
	Data *PlantOverviewInfo `json:"data,omitempty"`
}

// |=> GetPlantDataLoggerInfo
type Date struct {
	Time           *int64 `json:"time,omitempty"`
	Minutes        *int   `json:"minutes,omitempty"`
	Seconds        *int   `json:"seconds,omitempty"`
	Hours          *int   `json:"hours,omitempty"`
	Month          *int   `json:"month,omitempty"`
	TimezoneOffset *int   `json:"timezoneOffset,omitempty"`
	Year           *int   `json:"year,omitempty"`
	Day            *int   `json:"day,omitempty"`
	Date           *int   `json:"date,omitempty"`
}

func (d *Date) GetTime(defaultValue ...int64) int64 {
	value := int64(0)
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Int64Value(d.Time, value)
}

func (d *Date) GetMinutes(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(d.Minutes, value)
}

func (d *Date) GetSeconds(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(d.Seconds, value)
}

func (d *Date) GetHours(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(d.Hours, value)
}

func (d *Date) GetMonth(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(d.Month, value)
}

func (d *Date) GetTimezoneOffset(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(d.TimezoneOffset, value)
}

func (d *Date) GetYear(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(d.Year, value)
}

func (d *Date) GetDay(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(d.Day, value)
}

func (d *Date) GetDate(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(d.Date, value)
}

type PeakPowerActual struct {
	MapCityID                *int           `json:"map_cityId,omitempty"`
	MapCountryID             *int           `json:"map_countryId,omitempty"`
	StorageBattoryPercentage *float64       `json:"storage_BattoryPercentage,omitempty"`
	DefaultPlant             *bool          `json:"defaultPlant,omitempty"`
	PeakPeriodPrice          *float64       `json:"peakPeriodPrice,omitempty"`
	City                     *string        `json:"city,omitempty"`
	NominalPower             *float64       `json:"nominalPower,omitempty"`
	AlarmValue               *float64       `json:"alarmValue,omitempty"`
	CurrentPacTxt            *string        `json:"currentPacTxt,omitempty"`
	FixedPowerPrice          *float64       `json:"fixedPowerPrice,omitempty"`
	PlantFromBean            *interface{}   `json:"plantFromBean,omitempty"`
	DeviceCount              *int           `json:"deviceCount,omitempty"`
	PlantImgName             *string        `json:"plantImgName,omitempty"`
	EtodaySo2Text            *string        `json:"etodaySo2Text,omitempty"`
	CompanyName              *string        `json:"companyName,omitempty"`
	EmonthMoneyText          *string        `json:"emonthMoneyText,omitempty"`
	FormulaMoney             *float64       `json:"formulaMoney,omitempty"`
	UserAccount              *string        `json:"userAccount,omitempty"`
	MapLat                   *string        `json:"mapLat,omitempty"`
	CreateDateTextA          *string        `json:"createDateTextA,omitempty"`
	MapLng                   *string        `json:"mapLng,omitempty"`
	OnLineEnvCount           *int           `json:"onLineEnvCount,omitempty"`
	EventMessBeanList        []*interface{} `json:"eventMessBeanList,omitempty"`
	LatitudeText             *string        `json:"latitudeText,omitempty"`
	PlantAddress             *string        `json:"plantAddress,omitempty"`
	HasDeviceOnLine          *int           `json:"hasDeviceOnLine,omitempty"`
	FormulaMoneyStr          *string        `json:"formulaMoneyStr,omitempty"`
	EtodayMoney              *float64       `json:"etodayMoney,omitempty"`
	CreateDate               *Date          `json:"createDate,omitempty"`
	MapCity                  *string        `json:"mapCity,omitempty"`
	PrMonth                  *string        `json:"prMonth,omitempty"`
	StorageTodayToGrid       *float64       `json:"storage_TodayToGrid,omitempty"`
	FormulaCo2               *float64       `json:"formulaCo2,omitempty"`
	ETotal                   *float64       `json:"eTotal,omitempty"`
	EmonthSo2Text            *string        `json:"emonthSo2Text,omitempty"`
	WindAngle                *float64       `json:"windAngle,omitempty"`
	EtotalCoalText           *string        `json:"etotalCoalText,omitempty"`
	WindSpeed                *float64       `json:"windSpeed,omitempty"`
	EmonthCoalText           *string        `json:"emonthCoalText,omitempty"`
	EtodayMoneyText          *string        `json:"etodayMoneyText,omitempty"`
	EYearMoneyText           *string        `json:"EYearMoneyText,omitempty"`
	PlantLng                 *string        `json:"plant_lng,omitempty"`
	LatitudeM                *string        `json:"latitude_m,omitempty"`
	PairViewUserAccount      *string        `json:"pairViewUserAccount,omitempty"`
	StorageTotalToUser       *float64       `json:"storage_TotalToUser,omitempty"`
	LatitudeD                *string        `json:"latitude_d,omitempty"`
	LatitudeF                *string        `json:"latitude_f,omitempty"`
	Remark                   *string        `json:"remark,omitempty"`
	TreeID                   *string        `json:"treeID,omitempty"`
	FlatPeriodPrice          *float64       `json:"flatPeriodPrice,omitempty"`
	LongitudeText            *string        `json:"longitudeText,omitempty"`
	StorageEChargeToday      *float64       `json:"storage_eChargeToday,omitempty"`
	DataLogList              []*interface{} `json:"dataLogList,omitempty"`
	DesignCompany            *string        `json:"designCompany,omitempty"`
	TimezoneText             *string        `json:"timezoneText,omitempty"`
	FormulaCoal              *float64       `json:"formulaCoal,omitempty"`
	StorageEDisChargeToday   *float64       `json:"storage_eDisChargeToday,omitempty"`
	UnitMap                  *interface{}   `json:"unitMap,omitempty"`
	Timezone                 *int           `json:"timezone,omitempty"`
	PhoneNum                 *string        `json:"phoneNum,omitempty"`
	Level                    *int           `json:"level,omitempty"`
	FormulaMoneyUnitID       *string        `json:"formulaMoneyUnitId,omitempty"`
	ImgPath                  *string        `json:"imgPath,omitempty"`
	PanelTemp                *float64       `json:"panelTemp,omitempty"`
	LocationImgName          *string        `json:"locationImgName,omitempty"`
	MoneyUnitText            *string        `json:"moneyUnitText,omitempty"`
	StorageTotalToGrid       *float64       `json:"storage_TotalToGrid,omitempty"`
	PrToday                  *string        `json:"prToday,omitempty"`
	EnergyMonth              *float64       `json:"energyMonth,omitempty"`
	PlantName                *string        `json:"plantName,omitempty"`
	EToday                   *float64       `json:"eToday,omitempty"`
	Status                   *int           `json:"status,omitempty"`
	PlantType                *int           `json:"plantType,omitempty"`
	Country                  *string        `json:"country,omitempty"`
	LongitudeD               *string        `json:"longitude_d,omitempty"`
	MapAreaID                *int           `json:"map_areaId,omitempty"`
	LongitudeF               *string        `json:"longitude_f,omitempty"`
	CreateDateText           *string        `json:"createDateText,omitempty"`
	LongitudeM               *string        `json:"longitude_m,omitempty"`
	FormulaSo2               *float64       `json:"formulaSo2,omitempty"`
	ValleyPeriodPrice        *float64       `json:"valleyPeriodPrice,omitempty"`
	EnergyYear               *float64       `json:"energyYear,omitempty"`
	TreeName                 *string        `json:"treeName,omitempty"`
	PlantLat                 *string        `json:"plant_lat,omitempty"`
	EtodayCo2Text            *string        `json:"etodayCo2Text,omitempty"`
	NominalPowerStr          *string        `json:"nominalPowerStr,omitempty"`
	FormulaTree              *float64       `json:"formulaTree,omitempty"`
	EtotalSo2Text            *string        `json:"etotalSo2Text,omitempty"`
	Children                 []*interface{} `json:"children,omitempty"`
	ID                       *int           `json:"id,omitempty"`
	EtodayCoalText           *string        `json:"etodayCoalText,omitempty"`
	ParamBean                *interface{}   `json:"paramBean,omitempty"`
	EtotalMoney              *int           `json:"etotalMoney,omitempty"`
	EnvTemp                  *int           `json:"envTemp,omitempty"`
	LogoImgName              *string        `json:"logoImgName,omitempty"`
	Alias                    *string        `json:"alias,omitempty"`
	EtotalCo2Text            *string        `json:"etotalCo2Text,omitempty"`
	CurrentPacStr            *string        `json:"currentPacStr,omitempty"`
	MapProvinceID            *int           `json:"map_provinceId,omitempty"`
	EtotalMoneyText          *string        `json:"etotalMoneyText,omitempty"`
	EmonthCo2Text            *string        `json:"emonthCo2Text,omitempty"`
	Irradiance               *int           `json:"irradiance,omitempty"`
	HasStorage               *int           `json:"hasStorage,omitempty"`
	ParentID                 *string        `json:"parentID,omitempty"`
	UserBean                 *interface{}   `json:"userBean,omitempty"`
	StorageTodayToUser       *int           `json:"storage_TodayToUser,omitempty"`
	IsShare                  *bool          `json:"isShare,omitempty"`
	CurrentPac               *int           `json:"currentPac,omitempty"`
}

func (p *PeakPowerActual) GetMapCountryID(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(p.MapCountryID, value)
}

func (p *PeakPowerActual) GetStorageBattoryPercentage(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.StorageBattoryPercentage, value)
}

func (p *PeakPowerActual) GetDefaultPlant(defaultValue ...bool) bool {
	value := false
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.BoolValue(p.DefaultPlant, value)
}

func (p *PeakPowerActual) GetPeakPeriodPrice(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.PeakPeriodPrice, value)
}

func (p *PeakPowerActual) GetCity(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.City, value)
}

func (p *PeakPowerActual) GetNominalPower(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.NominalPower, value)
}

func (p *PeakPowerActual) GetAlarmValue(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.AlarmValue, value)
}

func (p *PeakPowerActual) GetCurrentPacTxt(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.CurrentPacTxt, value)
}

func (p *PeakPowerActual) GetFixedPowerPrice(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.FixedPowerPrice, value)
}

func (p *PeakPowerActual) GetDeviceCount(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(p.DeviceCount, value)
}

func (p *PeakPowerActual) GetPlantImgName(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.PlantImgName, value)
}

func (p *PeakPowerActual) GetEtodaySo2Text(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.EtodaySo2Text, value)
}

func (p *PeakPowerActual) GetCompanyName(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.CompanyName, value)
}

func (p *PeakPowerActual) GetEmonthMoneyText(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.EmonthMoneyText, value)
}

func (p *PeakPowerActual) GetFormulaMoney(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.FormulaMoney, value)
}

func (p *PeakPowerActual) GetUserAccount(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.UserAccount, value)
}

func (p *PeakPowerActual) GetMapLat(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.MapLat, value)
}

func (p *PeakPowerActual) GetCreateDateTextA(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.CreateDateTextA, value)
}

func (p *PeakPowerActual) GetMapLng(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.MapLng, value)
}

func (p *PeakPowerActual) GetOnLineEnvCount(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(p.OnLineEnvCount, value)
}

func (p *PeakPowerActual) GetLatitudeText(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.LatitudeText, value)
}

func (p *PeakPowerActual) GetPlantAddress(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.PlantAddress, value)
}

func (p *PeakPowerActual) GetHasDeviceOnLine(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(p.HasDeviceOnLine, value)
}

func (p *PeakPowerActual) GetFormulaMoneyStr(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.FormulaMoneyStr, value)
}

func (p *PeakPowerActual) GetEtodayMoney(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.EtodayMoney, value)
}

func (p *PeakPowerActual) GetMapCity(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.MapCity, value)
}

func (p *PeakPowerActual) GetPrMonth(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.PrMonth, value)
}

func (p *PeakPowerActual) GetStorageTodayToGrid(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.StorageTodayToGrid, value)
}

func (p *PeakPowerActual) GetFormulaCo2(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.FormulaCo2, value)
}

func (p *PeakPowerActual) GetETotal(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.ETotal, value)
}

func (p *PeakPowerActual) GetEmonthSo2Text(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.EmonthSo2Text, value)
}

func (p *PeakPowerActual) GetWindAngle(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.WindAngle, value)
}

func (p *PeakPowerActual) GetEtotalCoalText(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.EtotalCoalText, value)
}

func (p *PeakPowerActual) GetWindSpeed(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.WindSpeed, value)
}

func (p *PeakPowerActual) GetEmonthCoalText(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.EmonthCoalText, value)
}

func (p *PeakPowerActual) GetEtodayMoneyText(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.EtodayMoneyText, value)
}

func (p *PeakPowerActual) GetEYearMoneyText(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.EYearMoneyText, value)
}

func (p *PeakPowerActual) GetPlantLng(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.PlantLng, value)
}

func (p *PeakPowerActual) GetLatitudeM(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.LatitudeM, value)
}

func (p *PeakPowerActual) GetPairViewUserAccount(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.PairViewUserAccount, value)
}

func (p *PeakPowerActual) GetStorageTotalToUser(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.StorageTotalToUser, value)
}

func (p *PeakPowerActual) GetLatitudeD(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.LatitudeD, value)
}

func (p *PeakPowerActual) GetLatitudeF(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.LatitudeF, value)
}

func (p *PeakPowerActual) GetRemark(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.Remark, value)
}

func (p *PeakPowerActual) GetTreeID(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.TreeID, value)
}

func (p *PeakPowerActual) GetFlatPeriodPrice(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.FlatPeriodPrice, value)
}

func (p *PeakPowerActual) GetLongitudeText(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.LongitudeText, value)
}

func (p *PeakPowerActual) GetStorageEChargeToday(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.StorageEChargeToday, value)
}

func (p *PeakPowerActual) GetDesignCompany(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.DesignCompany, value)
}

func (p *PeakPowerActual) GetTimezoneText(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.TimezoneText, value)
}

func (p *PeakPowerActual) GetFormulaCoal(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.FormulaCoal, value)
}

func (p *PeakPowerActual) GetStorageEDisChargeToday(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.StorageEDisChargeToday, value)
}

func (p *PeakPowerActual) GetTimezone(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(p.Timezone, value)
}

func (p *PeakPowerActual) GetPhoneNum(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.PhoneNum, value)
}

func (p *PeakPowerActual) GetLevel(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(p.Level, value)
}

func (p *PeakPowerActual) GetFormulaMoneyUnitID(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.FormulaMoneyUnitID, value)
}

func (p *PeakPowerActual) GetImgPath(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.ImgPath, value)
}

func (p *PeakPowerActual) GetPanelTemp(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.PanelTemp, value)
}

func (p *PeakPowerActual) GetLocationImgName(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.LocationImgName, value)
}

func (p *PeakPowerActual) GetMoneyUnitText(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.MoneyUnitText, value)
}

func (p *PeakPowerActual) GetStorageTotalToGrid(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.StorageTotalToGrid, value)
}

func (p *PeakPowerActual) GetPrToday(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.PrToday, value)
}

func (p *PeakPowerActual) GetEnergyMonth(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.EnergyMonth, value)
}

func (p *PeakPowerActual) GetPlantName(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.PlantName, value)
}

func (p *PeakPowerActual) GetEToday(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.EToday, value)
}

func (p *PeakPowerActual) GetStatus(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(p.Status, value)
}

func (p *PeakPowerActual) GetPlantType(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(p.PlantType, value)
}

func (p *PeakPowerActual) GetCountry(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.Country, value)
}

func (p *PeakPowerActual) GetLongitudeD(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.LongitudeD, value)
}

func (p *PeakPowerActual) GetMapAreaID(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(p.MapAreaID, value)
}

func (p *PeakPowerActual) GetLongitudeF(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.LongitudeF, value)
}

func (p *PeakPowerActual) GetCreateDateText(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.CreateDateText, value)
}

func (p *PeakPowerActual) GetLongitudeM(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.LongitudeM, value)
}

func (p *PeakPowerActual) GetFormulaSo2(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.FormulaSo2, value)
}

func (p *PeakPowerActual) GetValleyPeriodPrice(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.ValleyPeriodPrice, value)
}

func (p *PeakPowerActual) GetEnergyYear(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.EnergyYear, value)
}

func (p *PeakPowerActual) GetTreeName(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.TreeName, value)
}

func (p *PeakPowerActual) GetPlantLat(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.PlantLat, value)
}

func (p *PeakPowerActual) GetEtodayCo2Text(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.EtodayCo2Text, value)
}

func (p *PeakPowerActual) GetNominalPowerStr(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.NominalPowerStr, value)
}

func (p *PeakPowerActual) GetFormulaTree(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.FormulaTree, value)
}

func (p *PeakPowerActual) GetEtotalSo2Text(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.EtotalSo2Text, value)
}

func (p *PeakPowerActual) GetID(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(p.ID, value)
}

func (p *PeakPowerActual) GetEtodayCoalText(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.EtodayCoalText, value)
}

func (p *PeakPowerActual) GetEtotalMoney(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(p.EtotalMoney, value)
}

func (p *PeakPowerActual) GetEnvTemp(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(p.EnvTemp, value)
}

func (p *PeakPowerActual) GetLogoImgName(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.LogoImgName, value)
}

func (p *PeakPowerActual) GetAlias(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.Alias, value)
}

func (p *PeakPowerActual) GetEtotalCo2Text(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.EtotalCo2Text, value)
}

func (p *PeakPowerActual) GetCurrentPacStr(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.CurrentPacStr, value)
}

func (p *PeakPowerActual) GetMapProvinceID(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(p.MapProvinceID, value)
}

func (p *PeakPowerActual) GetEtotalMoneyText(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.EtotalMoneyText, value)
}

func (p *PeakPowerActual) GetEmonthCo2Text(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.EmonthCo2Text, value)
}

func (p *PeakPowerActual) GetIrradiance(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(p.Irradiance, value)
}

func (p *PeakPowerActual) GetHasStorage(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(p.HasStorage, value)
}

func (p *PeakPowerActual) GetParentID(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.ParentID, value)
}

func (p *PeakPowerActual) GetStorageTodayToUser(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(p.StorageTodayToUser, value)
}

func (p *PeakPowerActual) GetIsShare(defaultValue ...bool) bool {
	value := false
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.BoolValue(p.IsShare, value)
}

func (p *PeakPowerActual) GetCurrentPac(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(p.CurrentPac, value)
}

type DataLogger struct {
	LastUpdateTime *Date   `json:"last_update_time,omitempty"`
	Model          *string `json:"model,omitempty"`
	SN             *string `json:"sn,omitempty"`
	Lost           *bool   `json:"lost,omitempty"`
	Manufacturer   *string `json:"manufacturer,omitempty"`
	Type           *int    `json:"type,omitempty"`
}

func (d *DataLogger) GetModel(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.Model, value)
}

func (d *DataLogger) GetSN(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.SN, value)
}

func (d *DataLogger) GetLost(defaultValue ...bool) bool {
	value := false
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.BoolValue(d.Lost, value)
}

func (d *DataLogger) GetManufacturer(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.Manufacturer, value)
}

func (d *DataLogger) GetType(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(d.Type, value)
}

type PlantDataLoggerInfo struct {
	Count           *int             `json:"count,omitempty"`
	PeakPowerActual *PeakPowerActual `json:"peak_power_actual,omitempty"`
	DataLoggers     []*DataLogger
}

func (r *PlantDataLoggerInfo) GetCount(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(r.Count, value)
}

type GetPlantDataLoggerInfoResponse struct {
	DefaultResponse
	Data *PlantDataLoggerInfo `json:"data,omitempty"`
}

// |=> GetPlantDeviceList
type DeviceItem struct {
	DeviceSN       *string `json:"device_sn,omitempty"`
	LastUpdateTime *string `json:"last_update_time,omitempty"`
	Model          *string `json:"model,omitempty"`
	Lost           *bool   `json:"lost,omitempty"`
	Status         *int    `json:"status,omitempty"`
	Manufacturer   *string `json:"manufacturer,omitempty"`
	DeviceID       *int    `json:"device_id,omitempty"`
	DataLoggerSN   *string `json:"datalogger_sn,omitempty"`
	Type           *int    `json:"type,omitempty"`
}

func (d *DeviceItem) GetDeviceSN(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.DeviceSN, value)
}

func (d *DeviceItem) GetLastUpdateTime(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.LastUpdateTime, value)
}

func (d *DeviceItem) GetModel(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.Model, value)
}

func (d *DeviceItem) GetLost(defaultValue ...bool) bool {
	value := false
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.BoolValue(d.Lost, value)
}

func (d *DeviceItem) GetStatus(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(d.Status, value)
}

func (d *DeviceItem) GetManufacturer(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.Manufacturer, value)
}

func (d *DeviceItem) GetDeviceID(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(d.DeviceID, value)
}

func (d *DeviceItem) GetDataLoggerSN(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.DataLoggerSN, value)
}

func (d *DeviceItem) GetType(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(d.Type, value)
}

type DeviceData struct {
	Count   *int          `json:"count,omitempty"`
	Devices []*DeviceItem `json:"devices,omitempty"`
}

func (d *DeviceData) GetCount(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(d.Count, value)
}

type GetPlantDeviceListResponse struct {
	DefaultResponse
	Data *DeviceData `json:"data,omitempty"`
}

// GetRealtimeDeviceData
type TimeZone struct {
	LastRuleInstance *interface{} `json:"lastRuleInstance,omitempty"`
	RawOffset        *int         `json:"rawOffset,omitempty"`
	DSTSavings       *int         `json:"DSTSavings,omitempty"`
	Dirty            *bool        `json:"dirty,omitempty"`
	ID               *string      `json:"ID,omitempty"`
	DisplayName      *string      `json:"displayName,omitempty"`
}

func (t *TimeZone) GetRawOffset(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(t.RawOffset, value)
}

func (t *TimeZone) GetDSTSavings(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(t.DSTSavings, value)
}

func (t *TimeZone) GetDirty(defaultValue ...bool) bool {
	value := false
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.BoolValue(t.Dirty, value)
}

func (t *TimeZone) GetID(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(t.ID, value)
}

func (t *TimeZone) GetDisplayName(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(t.DisplayName, value)
}

type TimeCalendar struct {
	MinimalDaysInFirstWeek *int      `json:"minimalDaysInFirstWeek,omitempty"`
	WeekYear               *int      `json:"weekYear,omitempty"`
	Time                   *Date     `json:"time,omitempty"`
	WeeksInWeekYear        *int      `json:"weeksInWeekYear,omitempty"`
	GregorianChange        *Date     `json:"gregorianChange,omitempty"`
	TimeZone               *TimeZone `json:"timeZone,omitempty"`
	TimeInMillis           *int64    `json:"timeInMillis,omitempty"`
	Lenient                *bool     `json:"lenient,omitempty"`
	FirstDayOfWeek         *int      `json:"firstDayOfWeek,omitempty"`
	WeekDateSupported      *bool     `json:"weekDateSupported,omitempty"`
}

func (t *TimeCalendar) GetMinimalDaysInFirstWeek(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(t.MinimalDaysInFirstWeek, value)
}

func (t *TimeCalendar) GetWeekYear(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(t.WeekYear, value)
}

func (t *TimeCalendar) GetWeeksInWeekYear(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(t.WeeksInWeekYear, value)
}

func (t *TimeCalendar) GetTimeInMillis(defaultValue ...int64) int64 {
	value := int64(0)
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Int64Value(t.TimeInMillis, value)
}

func (t *TimeCalendar) GetLenient(defaultValue ...bool) bool {
	value := false
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.BoolValue(t.Lenient, value)
}

func (t *TimeCalendar) GetFirstDayOfWeek(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(t.FirstDayOfWeek, value)
}

func (t *TimeCalendar) GetWeekDateSupported(defaultValue ...bool) bool {
	value := false
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.BoolValue(t.WeekDateSupported, value)
}

type RealtimeDeviceData struct {
	IPidPvcpe             *float64      `json:"iPidPvcpe,omitempty"`
	Epv4Total             *float64      `json:"epv4Total,omitempty"`
	RealOPPercent         *float64      `json:"realOPPercent,omitempty"`
	PidBus                *float64      `json:"pidBus,omitempty"`
	Ppv7                  *float64      `json:"ppv7,omitempty"`
	Ctharis               *float64      `json:"ctharis,omitempty"`
	Ctir                  *float64      `json:"ctir,omitempty"`
	VacTr                 *float64      `json:"vacTr,omitempty"`
	Ppv6                  *float64      `json:"ppv6,omitempty"`
	ERacTotal             *float64      `json:"eRacTotal,omitempty"`
	Ctharit               *float64      `json:"ctharit,omitempty"`
	Ctis                  *float64      `json:"ctis,omitempty"`
	Ppv9                  *float64      `json:"ppv9,omitempty"`
	Epv1Total             *float64      `json:"epv1Total,omitempty"`
	WStringStatusValue    *float64      `json:"wStringStatusValue,omitempty"`
	Ppv8                  *float64      `json:"ppv8,omitempty"`
	Ctharir               *float64      `json:"ctharir,omitempty"`
	WarningValue3         *float64      `json:"warningValue3,omitempty"`
	VPidPvape             *float64      `json:"vPidPvape,omitempty"`
	WarningValue1         *float64      `json:"warningValue1,omitempty"`
	Ctit                  *float64      `json:"ctit,omitempty"`
	FaultCode1            *float64      `json:"faultCode1,omitempty"`
	WarningValue2         *float64      `json:"warningValue2,omitempty"`
	Temperature           *float64      `json:"temperature,omitempty"`
	FaultCode2            *float64      `json:"faultCode2,omitempty"`
	Time                  *string       `json:"time,omitempty"`
	IPidPvbpe             *float64      `json:"iPidPvbpe,omitempty"`
	IPidPvdpe             *float64      `json:"iPidPvdpe,omitempty"`
	Epv2Total             *float64      `json:"epv2Total,omitempty"`
	WarnBit               *float64      `json:"warnBit,omitempty"`
	IPidPvepe             *float64      `json:"iPidPvepe,omitempty"`
	VacSt                 *float64      `json:"vacSt,omitempty"`
	VPidPvcpe             *float64      `json:"vPidPvcpe,omitempty"`
	Epv8Total             *float64      `json:"epv8Total,omitempty"`
	Ipv10                 *float64      `json:"ipv10,omitempty"`
	Again                 *bool         `json:"again,omitempty"`
	Vpv10                 *float64      `json:"vpv10,omitempty"`
	StrBreak              *float64      `json:"strBreak,omitempty"`
	Compqt                *float64      `json:"compqt,omitempty"`
	IpmTemperature        *float64      `json:"ipmTemperature,omitempty"`
	Compqs                *float64      `json:"compqs,omitempty"`
	Ppv                   *float64      `json:"ppv,omitempty"`
	Compqr                *float64      `json:"compqr,omitempty"`
	Ctqt                  *float64      `json:"ctqt,omitempty"`
	Pf                    *float64      `json:"pf,omitempty"`
	Epv7Total             *float64      `json:"epv7Total,omitempty"`
	Vpv1                  *float64      `json:"vpv1,omitempty"`
	Epv10Today            *float64      `json:"epv10Today,omitempty"`
	IPidPvape             *float64      `json:"iPidPvape,omitempty"`
	Vpv3                  *float64      `json:"vpv3,omitempty"`
	Ctqr                  *float64      `json:"ctqr,omitempty"`
	Vpv2                  *float64      `json:"vpv2,omitempty"`
	Ctqs                  *float64      `json:"ctqs,omitempty"`
	Vpv5                  *float64      `json:"vpv5,omitempty"`
	Vpv4                  *float64      `json:"vpv4,omitempty"`
	Vpv7                  *float64      `json:"vpv7,omitempty"`
	Vpv6                  *float64      `json:"vpv6,omitempty"`
	PowerTotal            *float64      `json:"powerTotal,omitempty"`
	Vpv9                  *float64      `json:"vpv9,omitempty"`
	TDci                  *float64      `json:"tDci,omitempty"`
	Vpv8                  *float64      `json:"vpv8,omitempty"`
	Epv2Today             *float64      `json:"epv2Today,omitempty"`
	TimeTotal             *float64      `json:"timeTotal,omitempty"`
	Epv1Today             *float64      `json:"epv1Today,omitempty"`
	Epv6Today             *float64      `json:"epv6Today,omitempty"`
	Epv9Today             *float64      `json:"epv9Today,omitempty"`
	TimeTotalText         *string       `json:"timeTotalText,omitempty"`
	DwStringWarningValue1 *float64      `json:"dwStringWarningValue1,omitempty"`
	VPidPvepe             *float64      `json:"vPidPvepe,omitempty"`
	EpvTotal              *float64      `json:"epvTotal,omitempty"`
	VPidPvgpe             *float64      `json:"vPidPvgpe,omitempty"`
	FaultType             *float64      `json:"faultType,omitempty"`
	CurrentString12       *float64      `json:"currentString12,omitempty"`
	CurrentString11       *float64      `json:"currentString11,omitempty"`
	CurrentString10       *float64      `json:"currentString10,omitempty"`
	ERacToday             *float64      `json:"eRacToday,omitempty"`
	CurrentString16       *float64      `json:"currentString16,omitempty"`
	Epv5Today             *float64      `json:"epv5Today,omitempty"`
	CurrentString15       *float64      `json:"currentString15,omitempty"`
	CurrentString14       *float64      `json:"currentString14,omitempty"`
	CurrentString13       *float64      `json:"currentString13,omitempty"`
	WPIDFaultValue        *float64      `json:"wPIDFaultValue,omitempty"`
	VString11             *float64      `json:"vString11,omitempty"`
	VString10             *float64      `json:"vString10,omitempty"`
	PowerToday            *float64      `json:"powerToday,omitempty"`
	VString16             *float64      `json:"vString16,omitempty"`
	VString13             *float64      `json:"vString13,omitempty"`
	VString12             *float64      `json:"vString12,omitempty"`
	VString15             *float64      `json:"vString15,omitempty"`
	VString14             *float64      `json:"vString14,omitempty"`
	BigDevice             *bool         `json:"bigDevice,omitempty"`
	Epv9Total             *float64      `json:"epv9Total,omitempty"`
	WarnCode              *float64      `json:"warnCode,omitempty"`
	PvIso                 *float64      `json:"pvIso,omitempty"`
	Epv6Total             *float64      `json:"epv6Total,omitempty"`
	InverterID            *string       `json:"inverterId,omitempty"`
	Temperature3          *float64      `json:"temperature3,omitempty"`
	Temperature2          *float64      `json:"temperature2,omitempty"`
	TimeCalendar          *TimeCalendar `json:"timeCalendar,omitempty"`
	PBusVoltage           *float64      `json:"pBusVoltage,omitempty"`
	CurrentString5        *float64      `json:"currentString5,omitempty"`
	StrFault              *float64      `json:"strFault,omitempty"`
	CurrentString4        *float64      `json:"currentString4,omitempty"`
	VPidPvdpe             *float64      `json:"vPidPvdpe,omitempty"`
	CurrentString3        *float64      `json:"currentString3,omitempty"`
	Epv3Today             *float64      `json:"epv3Today,omitempty"`
	CurrentString2        *float64      `json:"currentString2,omitempty"`
	CurrentString9        *float64      `json:"currentString9,omitempty"`
	Status                *float64      `json:"status,omitempty"`
	CurrentString8        *float64      `json:"currentString8,omitempty"`
	CurrentString7        *float64      `json:"currentString7,omitempty"`
	CurrentString6        *float64      `json:"currentString6,omitempty"`
	NBusVoltage           *float64      `json:"nBusVoltage,omitempty"`
	CurrentString1        *float64      `json:"currentString1,omitempty"`
	Pacs                  *float64      `json:"pacs,omitempty"`
	Pacr                  *float64      `json:"pacr,omitempty"`
	StrUnblance           *float64      `json:"strUnblance,omitempty"`
	StrUnmatch            *float64      `json:"strUnmatch,omitempty"`
	SDci                  *float64      `json:"sDci,omitempty"`
	Pact                  *float64      `json:"pact,omitempty"`
	Fac                   *float64      `json:"fac,omitempty"`
	VPidPvbpe             *float64      `json:"vPidPvbpe,omitempty"`
	FaultValue            *float64      `json:"faultValue,omitempty"`
	Epv5Total             *float64      `json:"epv5Total,omitempty"`
	Ipv6                  *float64      `json:"ipv6,omitempty"`
	Ipv5                  *float64      `json:"ipv5,omitempty"`
	Ipv4                  *float64      `json:"ipv4,omitempty"`
	Epv4Today             *float64      `json:"epv4Today,omitempty"`
	Ipv3                  *float64      `json:"ipv3,omitempty"`
	Ipv2                  *float64      `json:"ipv2,omitempty"`
	Ipv1                  *float64      `json:"ipv1,omitempty"`
	IPidPvfpe             *float64      `json:"iPidPvfpe,omitempty"`
	StatusText            *string       `json:"statusText,omitempty"`
	VacRs                 *float64      `json:"vacRs,omitempty"`
	IPidPvgpe             *float64      `json:"iPidPvgpe,omitempty"`
	Ipv9                  *float64      `json:"ipv9,omitempty"`
	Ipv8                  *float64      `json:"ipv8,omitempty"`
	Ipv7                  *float64      `json:"ipv7,omitempty"`
	ID                    *float64      `json:"id,omitempty"`
	Epv8Today             *float64      `json:"epv8Today,omitempty"`
	Gfci                  *float64      `json:"gfci,omitempty"`
	IPidPvhpe             *float64      `json:"iPidPvhpe,omitempty"`
	Ppv10                 *float64      `json:"ppv10,omitempty"`
	Epv3Total             *float64      `json:"epv3Total,omitempty"`
	ApfStatus             *float64      `json:"apfStatus,omitempty"`
	Temperature4          *float64      `json:"temperature4,omitempty"`
	RDci                  *float64      `json:"rDci,omitempty"`
	Pac                   *float64      `json:"pac,omitempty"`
	Temperature5          *float64      `json:"temperature5,omitempty"`
	Vact                  *float64      `json:"vact,omitempty"`
	Compharir             *float64      `json:"compharir,omitempty"`
	Vacr                  *float64      `json:"vacr,omitempty"`
	Compharis             *float64      `json:"compharis,omitempty"`
	Vacs                  *float64      `json:"vacs,omitempty"`
	PidFaultCode          *float64      `json:"pidFaultCode,omitempty"`
	Compharit             *float64      `json:"compharit,omitempty"`
	DeratingMode          *float64      `json:"deratingMode,omitempty"`
	VString1              *float64      `json:"vString1,omitempty"`
	Epv7Today             *float64      `json:"epv7Today,omitempty"`
	VString2              *float64      `json:"vString2,omitempty"`
	VString3              *float64      `json:"vString3,omitempty"`
	VPidPvhpe             *float64      `json:"vPidPvhpe,omitempty"`
	VString4              *float64      `json:"vString4,omitempty"`
	VString5              *float64      `json:"vString5,omitempty"`
	VString6              *float64      `json:"vString6,omitempty"`
	VString8              *float64      `json:"vString8,omitempty"`
	PidStatus             *float64      `json:"pidStatus,omitempty"`
	Iacs                  *float64      `json:"iacs,omitempty"`
	OpFullwatt            *float64      `json:"opFullwatt,omitempty"`
	VString7              *float64      `json:"vString7,omitempty"`
	Iact                  *float64      `json:"iact,omitempty"`
	VString9              *float64      `json:"vString9,omitempty"`
	Epv10Total            *float64      `json:"epv10Total,omitempty"`
	VPidPvfpe             *float64      `json:"vPidPvfpe,omitempty"`
	Ppv5                  *float64      `json:"ppv5,omitempty"`
	Debug1                *string       `json:"debug1,omitempty"`
	Ppv4                  *float64      `json:"ppv4,omitempty"`
	Debug2                *string       `json:"debug2,omitempty"`
	Ppv3                  *float64      `json:"ppv3,omitempty"`
	Ppv2                  *float64      `json:"ppv2,omitempty"`
	Ppv1                  *float64      `json:"ppv1,omitempty"`
	Rac                   *float64      `json:"rac,omitempty"`
	Iacr                  *float64      `json:"iacr,omitempty"`
}

func (m *RealtimeDeviceData) GetIPidPvcpe(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.IPidPvcpe, value)
}

func (m *RealtimeDeviceData) GetEpv4Total(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.Epv4Total, value)
}

func (m *RealtimeDeviceData) GetRealOPPercent(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.RealOPPercent, value)
}

func (m *RealtimeDeviceData) GetPidBus(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.PidBus, value)
}

func (m *RealtimeDeviceData) GetPpv7(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.Ppv7, value)
}

func (m *RealtimeDeviceData) GetCtharis(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.Ctharis, value)
}

func (m *RealtimeDeviceData) GetCtir(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.Ctir, value)
}

func (m *RealtimeDeviceData) GetVacTr(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.VacTr, value)
}

func (m *RealtimeDeviceData) GetPpv6(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.Ppv6, value)
}

func (m *RealtimeDeviceData) GetERacTotal(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.ERacTotal, value)
}

func (m *RealtimeDeviceData) GetCtharit(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.Ctharit, value)
}

func (m *RealtimeDeviceData) GetCtis(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.Ctis, value)
}

func (m *RealtimeDeviceData) GetPpv9(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.Ppv9, value)
}

func (m *RealtimeDeviceData) GetEpv1Total(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.Epv1Total, value)
}

func (m *RealtimeDeviceData) GetWStringStatusValue(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.WStringStatusValue, value)
}

func (m *RealtimeDeviceData) GetPpv8(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.Ppv8, value)
}

func (m *RealtimeDeviceData) GetCtharir(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.Ctharir, value)
}

func (m *RealtimeDeviceData) GetWarningValue3(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.WarningValue3, value)
}

func (m *RealtimeDeviceData) GetVPidPvape(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.VPidPvape, value)
}

func (m *RealtimeDeviceData) GetWarningValue1(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.WarningValue1, value)
}

func (m *RealtimeDeviceData) GetCtit(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.Ctit, value)
}

func (m *RealtimeDeviceData) GetFaultCode1(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.FaultCode1, value)
}

func (m *RealtimeDeviceData) GetWarningValue2(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.WarningValue2, value)
}

func (m *RealtimeDeviceData) GetTemperature(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.Temperature, value)
}

func (m *RealtimeDeviceData) GetFaultCode2(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.FaultCode2, value)
}

func (m *RealtimeDeviceData) GetTime(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(m.Time, value)
}

func (m *RealtimeDeviceData) GetIPidPvbpe(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.IPidPvbpe, value)
}

func (m *RealtimeDeviceData) GetIPidPvdpe(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.IPidPvdpe, value)
}

func (m *RealtimeDeviceData) GetEpv2Total(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.Epv2Total, value)
}

func (m *RealtimeDeviceData) GetWarnBit(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.WarnBit, value)
}

func (m *RealtimeDeviceData) GetIPidPvepe(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.IPidPvepe, value)
}

func (m *RealtimeDeviceData) GetVacSt(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.VacSt, value)
}

func (m *RealtimeDeviceData) GetVPidPvcpe(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.VPidPvcpe, value)
}

func (m *RealtimeDeviceData) GetEpv8Total(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.Epv8Total, value)
}

func (m *RealtimeDeviceData) GetIpv10(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.Ipv10, value)
}

func (m *RealtimeDeviceData) GetAgain(defaultValue ...bool) bool {
	value := false
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.BoolValue(m.Again, value)
}

func (m *RealtimeDeviceData) GetVpv10(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.Vpv10, value)
}

func (m *RealtimeDeviceData) GetStrBreak(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.StrBreak, value)
}

func (m *RealtimeDeviceData) GetCompqt(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.Compqt, value)
}

func (m *RealtimeDeviceData) GetIpmTemperature(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.IpmTemperature, value)
}

func (m *RealtimeDeviceData) GetCompqs(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.Compqs, value)
}

func (m *RealtimeDeviceData) GetPpv(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.Ppv, value)
}

func (m *RealtimeDeviceData) GetCompqr(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.Compqr, value)
}

func (m *RealtimeDeviceData) GetCtqt(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(m.Ctqt, value)
}

func (r *RealtimeDeviceData) GetPf(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Pf, value)
}

func (r *RealtimeDeviceData) GetEpv7Total(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Epv7Total, value)
}

func (r *RealtimeDeviceData) GetVpv1(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Vpv1, value)
}

func (r *RealtimeDeviceData) GetEpv10Today(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Epv10Today, value)
}

func (r *RealtimeDeviceData) GetIPidPvape(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.IPidPvape, value)
}

func (r *RealtimeDeviceData) GetVpv3(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Vpv3, value)
}

func (r *RealtimeDeviceData) GetCtqr(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Ctqr, value)
}

func (r *RealtimeDeviceData) GetVpv2(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Vpv2, value)
}

func (r *RealtimeDeviceData) GetCtqs(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Ctqs, value)
}

func (r *RealtimeDeviceData) GetVpv5(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Vpv5, value)
}

func (r *RealtimeDeviceData) GetVpv4(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Vpv4, value)
}

func (r *RealtimeDeviceData) GetVpv7(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Vpv7, value)
}

func (r *RealtimeDeviceData) GetVpv6(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Vpv6, value)
}

func (r *RealtimeDeviceData) GetPowerTotal(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PowerTotal, value)
}

func (r *RealtimeDeviceData) GetVpv9(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Vpv9, value)
}

func (r *RealtimeDeviceData) GetTDci(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.TDci, value)
}

func (r *RealtimeDeviceData) GetVpv8(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Vpv8, value)
}

func (r *RealtimeDeviceData) GetEpv2Today(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Epv2Today, value)
}

func (r *RealtimeDeviceData) GetTimeTotal(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.TimeTotal, value)
}

func (r *RealtimeDeviceData) GetEpv1Today(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Epv1Today, value)
}

func (r *RealtimeDeviceData) GetEpv6Today(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Epv6Today, value)
}

func (r *RealtimeDeviceData) GetEpv9Today(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Epv9Today, value)
}

func (r *RealtimeDeviceData) GetTimeTotalText(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(r.TimeTotalText, value)
}

func (r *RealtimeDeviceData) GetDwStringWarningValue1(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.DwStringWarningValue1, value)
}

func (r *RealtimeDeviceData) GetVPidPvepe(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.VPidPvepe, value)
}

func (r *RealtimeDeviceData) GetEpvTotal(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.EpvTotal, value)
}

func (r *RealtimeDeviceData) GetVPidPvgpe(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.VPidPvgpe, value)
}

func (r *RealtimeDeviceData) GetFaultType(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.FaultType, value)
}

func (r *RealtimeDeviceData) GetCurrentString12(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.CurrentString12, value)
}

func (r *RealtimeDeviceData) GetCurrentString11(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.CurrentString11, value)
}

func (r *RealtimeDeviceData) GetCurrentString10(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.CurrentString10, value)
}

func (r *RealtimeDeviceData) GetERacToday(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.ERacToday, value)
}

func (r *RealtimeDeviceData) GetCurrentString16(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.CurrentString16, value)
}

func (r *RealtimeDeviceData) GetEpv5Today(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Epv5Today, value)
}

func (r *RealtimeDeviceData) GetCurrentString15(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.CurrentString15, value)
}

func (r *RealtimeDeviceData) GetCurrentString14(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.CurrentString14, value)
}

func (r *RealtimeDeviceData) GetCurrentString13(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.CurrentString13, value)
}

func (r *RealtimeDeviceData) GetWPIDFaultValue(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.WPIDFaultValue, value)
}

func (r *RealtimeDeviceData) GetVString11(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.VString11, value)
}

func (r *RealtimeDeviceData) GetVString10(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.VString10, value)
}

func (r *RealtimeDeviceData) GetPowerToday(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PowerToday, value)
}

func (r *RealtimeDeviceData) GetVString16(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.VString16, value)
}

func (r *RealtimeDeviceData) GetVString13(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.VString13, value)
}

func (r *RealtimeDeviceData) GetVString12(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.VString12, value)
}

func (r *RealtimeDeviceData) GetVString15(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.VString15, value)
}

func (r *RealtimeDeviceData) GetVString14(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.VString14, value)
}

func (r *RealtimeDeviceData) GetBigDevice(defaultValue ...bool) bool {
	value := false
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.BoolValue(r.BigDevice, value)
}

func (r *RealtimeDeviceData) GetEpv9Total(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Epv9Total, value)
}

func (r *RealtimeDeviceData) GetWarnCode(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.WarnCode, value)
}

func (r *RealtimeDeviceData) GetPvIso(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PvIso, value)
}

func (r *RealtimeDeviceData) GetEpv6Total(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Epv6Total, value)
}

func (r *RealtimeDeviceData) GetInverterID(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(r.InverterID, value)
}

func (r *RealtimeDeviceData) GetTemperature3(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Temperature3, value)
}

func (r *RealtimeDeviceData) GetTemperature2(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Temperature2, value)
}

func (r *RealtimeDeviceData) GetPBusVoltage(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PBusVoltage, value)
}

func (r *RealtimeDeviceData) GetCurrentString5(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.CurrentString5, value)
}

func (r *RealtimeDeviceData) GetStrFault(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.StrFault, value)
}

func (r *RealtimeDeviceData) GetCurrentString4(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.CurrentString4, value)
}

func (r *RealtimeDeviceData) GetVPidPvdpe(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.VPidPvdpe, value)
}

func (r *RealtimeDeviceData) GetCurrentString3(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.CurrentString3, value)
}

func (r *RealtimeDeviceData) GetEpv3Today(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Epv3Today, value)
}

func (r *RealtimeDeviceData) GetCurrentString2(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.CurrentString2, value)
}

func (r *RealtimeDeviceData) GetCurrentString9(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.CurrentString9, value)
}

func (r *RealtimeDeviceData) GetStatus(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Status, value)
}

func (r *RealtimeDeviceData) GetCurrentString8(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.CurrentString8, value)
}

func (r *RealtimeDeviceData) GetCurrentString7(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.CurrentString7, value)
}

func (r *RealtimeDeviceData) GetCurrentString6(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.CurrentString6, value)
}

func (r *RealtimeDeviceData) GetNBusVoltage(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.NBusVoltage, value)
}

func (r *RealtimeDeviceData) GetCurrentString1(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.CurrentString1, value)
}

func (r *RealtimeDeviceData) GetPacs(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Pacs, value)
}

func (r *RealtimeDeviceData) GetPacr(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Pacr, value)
}

func (r *RealtimeDeviceData) GetStrUnblance(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.StrUnblance, value)
}

func (r *RealtimeDeviceData) GetStrUnmatch(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.StrUnmatch, value)
}

func (r *RealtimeDeviceData) GetSDci(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.SDci, value)
}

func (r *RealtimeDeviceData) GetPact(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Pact, value)
}

func (r *RealtimeDeviceData) GetFac(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Fac, value)
}

func (r *RealtimeDeviceData) GetVPidPvbpe(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.VPidPvbpe, value)
}

func (r *RealtimeDeviceData) GetFaultValue(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.FaultValue, value)
}

func (r *RealtimeDeviceData) GetEpv5Total(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Epv5Total, value)
}

func (r *RealtimeDeviceData) GetIpv6(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Ipv6, value)
}

func (r *RealtimeDeviceData) GetIpv5(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Ipv5, value)
}

func (r *RealtimeDeviceData) GetIpv4(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Ipv4, value)
}

func (r *RealtimeDeviceData) GetEpv4Today(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Epv4Today, value)
}

func (r *RealtimeDeviceData) GetIpv3(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Ipv3, value)
}

func (r *RealtimeDeviceData) GetIpv2(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Ipv2, value)
}

func (r *RealtimeDeviceData) GetIpv1(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Ipv1, value)
}

func (r *RealtimeDeviceData) GetIPidPvfpe(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.IPidPvfpe, value)
}

func (r *RealtimeDeviceData) GetStatusText(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(r.StatusText, value)
}

func (r *RealtimeDeviceData) GetVacRs(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.VacRs, value)
}

func (r *RealtimeDeviceData) GetIPidPvgpe(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.IPidPvgpe, value)
}

func (r *RealtimeDeviceData) GetIpv9(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Ipv9, value)
}

func (r *RealtimeDeviceData) GetIpv8(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Ipv8, value)
}

func (r *RealtimeDeviceData) GetIpv7(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Ipv7, value)
}

func (r *RealtimeDeviceData) GetID(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.ID, value)
}

func (r *RealtimeDeviceData) GetEpv8Today(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Epv8Today, value)
}

func (r *RealtimeDeviceData) GetGfci(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Gfci, value)
}

func (r *RealtimeDeviceData) GetIPidPvhpe(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.IPidPvhpe, value)
}

func (r *RealtimeDeviceData) GetPpv10(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Ppv10, value)
}

func (r *RealtimeDeviceData) GetEpv3Total(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Epv3Total, value)
}

func (r *RealtimeDeviceData) GetApfStatus(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.ApfStatus, value)
}

func (r *RealtimeDeviceData) GetTemperature4(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Temperature4, value)
}

func (r *RealtimeDeviceData) GetRDci(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.RDci, value)
}

func (r *RealtimeDeviceData) GetPac(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Pac, value)
}

func (r *RealtimeDeviceData) GetTemperature5(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Temperature5, value)
}

func (r *RealtimeDeviceData) GetVact(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Vact, value)
}

func (r *RealtimeDeviceData) GetCompharir(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Compharir, value)
}

func (r *RealtimeDeviceData) GetVacr(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Vacr, value)
}

func (r *RealtimeDeviceData) GetCompharis(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Compharis, value)
}

func (r *RealtimeDeviceData) GetVacs(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Vacs, value)
}

func (r *RealtimeDeviceData) GetPidFaultCode(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PidFaultCode, value)
}

func (r *RealtimeDeviceData) GetCompharit(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Compharit, value)
}

func (r *RealtimeDeviceData) GetDeratingMode(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.DeratingMode, value)
}

func (r *RealtimeDeviceData) GetVString1(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.VString1, value)
}

func (r *RealtimeDeviceData) GetEpv7Today(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Epv7Today, value)
}

func (r *RealtimeDeviceData) GetVString2(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.VString2, value)
}

func (r *RealtimeDeviceData) GetVString3(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.VString3, value)
}

func (r *RealtimeDeviceData) GetVPidPvhpe(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.VPidPvhpe, value)
}

func (r *RealtimeDeviceData) GetVString4(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.VString4, value)
}

func (r *RealtimeDeviceData) GetVString5(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.VString5, value)
}

func (r *RealtimeDeviceData) GetVString6(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.VString6, value)
}

func (r *RealtimeDeviceData) GetVString8(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.VString8, value)
}

func (r *RealtimeDeviceData) GetPidStatus(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PidStatus, value)
}

func (r *RealtimeDeviceData) GetIacs(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Iacs, value)
}

func (r *RealtimeDeviceData) GetOpFullwatt(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.OpFullwatt, value)
}

func (r *RealtimeDeviceData) GetVString7(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.VString7, value)
}

func (r *RealtimeDeviceData) GetIact(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Iact, value)
}

func (r *RealtimeDeviceData) GetVString9(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.VString9, value)
}

func (r *RealtimeDeviceData) GetEpv10Total(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Epv10Total, value)
}

func (r *RealtimeDeviceData) GetVPidPvfpe(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.VPidPvfpe, value)
}

func (r *RealtimeDeviceData) GetPpv5(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Ppv5, value)
}

func (r *RealtimeDeviceData) GetDebug1(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(r.Debug1, value)
}

func (r *RealtimeDeviceData) GetPpv4(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Ppv4, value)
}

func (r *RealtimeDeviceData) GetDebug2(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(r.Debug2, value)
}

func (r *RealtimeDeviceData) GetPpv3(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Ppv3, value)
}

func (r *RealtimeDeviceData) GetPpv2(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Ppv2, value)
}

func (r *RealtimeDeviceData) GetPpv1(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Ppv1, value)
}

func (r *RealtimeDeviceData) GetRac(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Rac, value)
}

func (r *RealtimeDeviceData) GetIacr(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Iacr, value)
}

type GetRealtimeDeviceDataResponse struct {
	DefaultResponse
	DeviceSN     *string             `json:"device_sn,omitempty"`
	DataLoggerSN *string             `json:"dataloggerSn,omitempty"`
	Data         *RealtimeDeviceData `json:"data,omitempty"`
}

func (d *GetRealtimeDeviceDataResponse) GetDeviceSN(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.DeviceSN, value)
}

func (d *GetRealtimeDeviceDataResponse) GetDataLoggerSN(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.DataLoggerSN, value)
}

// GetRealtimeDeviceBatchesData
type GetRealtimeDeviceBatchesDataResponse struct {
	DefaultResponse
	Inverters []*string                         `json:"inverters,omitempty"`
	Data      map[string]map[string]interface{} `json:"data,omitempty"`
	PageNum   *int                              `json:"pageNum,omitempty"`
}

// GetInverterAlertList
type AlarmItem struct {
	AlarmCode    *int    `json:"alarm_code,omitempty"`
	Status       *int    `json:"status,omitempty"`
	EndTime      *string `json:"end_time,omitempty"`
	StartTime    *string `json:"start_time,omitempty"`
	AlarmMessage *string `json:"alarm_message,omitempty"`
}

func (a *AlarmItem) GetAlarmCode(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(a.AlarmCode, value)
}

func (a *AlarmItem) GetStatus(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(a.Status, value)
}

func (a *AlarmItem) GetEndTime(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(a.EndTime, value)
}

func (a *AlarmItem) GetStartTime(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(a.StartTime, value)
}

func (a *AlarmItem) GetAlarmMessage(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(a.AlarmMessage, value)
}

type InverterAlertData struct {
	Count  *int         `json:"count,omitempty"`
	SN     *string      `json:"sn,omitempty"`
	Alarms []*AlarmItem `json:"alarms,omitempty"`
}

func (i *InverterAlertData) GetCount(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(i.Count, value)
}

func (i *InverterAlertData) GetSN(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(i.SN, value)
}

type GetInverterAlertListResponse struct {
	DefaultResponse
	Data *InverterAlertData `json:"data,omitempty"`
}

// GetEnergyStorageMachineAlertList
type EnergyStorageMachineAlertData struct {
	Count     *int         `json:"count,omitempty"`
	StorageSN *string      `json:"storage_sn,omitempty"`
	Alarms    []*AlarmItem `json:"alarms,omitempty"`
}

func (e *EnergyStorageMachineAlertData) GetCount(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(e.Count, value)
}

func (e *EnergyStorageMachineAlertData) GetStorageSN(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(e.StorageSN, value)
}

type GetEnergyStorageMachineAlertListResponse struct {
	DefaultResponse
	Data *EnergyStorageMachineAlertData `json:"data,omitempty"`
}

// GetMaxAlertList
type MaxAlertData struct {
	Count  *int         `json:"count,omitempty"`
	MaxSN  *string      `json:"max_sn,omitempty"`
	Alarms []*AlarmItem `json:"alarms,omitempty"`
}

func (m *MaxAlertData) GetCount(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(m.Count, value)
}

func (m *MaxAlertData) GetMaxSN(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(m.MaxSN, value)
}

type GetMaxAlertListResponse struct {
	DefaultResponse
	Data *MaxAlertData `json:"data,omitempty"`
}

// GetMixAlertList
type MixAlertData struct {
	Count  *int         `json:"count,omitempty"`
	MixSN  *string      `json:"mix_sn,omitempty"`
	Alarms []*AlarmItem `json:"alarms,omitempty"`
}

func (m *MixAlertData) GetCount(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(m.Count, value)
}

func (m *MixAlertData) GetMixSN(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(m.MixSN, value)
}

type GetMixAlertListResponse struct {
	DefaultResponse
	Data *MixAlertData `json:"data,omitempty"`
}

// GetMinAlertList
type MinAlertData struct {
	Count  *int         `json:"count,omitempty"`
	MinSN  *string      `json:"min_sn,omitempty"`
	Alarms []*AlarmItem `json:"alarms,omitempty"`
}

func (m *MinAlertData) GetCount(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(m.Count, value)
}

func (m *MinAlertData) GetMinSN(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(m.MinSN, value)
}

type GetMinAlertListResponse struct {
	DefaultResponse
	Data *MinAlertData `json:"data,omitempty"`
}

// GetSpaAlertList
type SpaAlertData struct {
	Count  *int         `json:"count,omitempty"`
	SpaSN  *string      `json:"spa_sn,omitempty"`
	Alarms []*AlarmItem `json:"alarms,omitempty"`
}

func (s *SpaAlertData) GetCount(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(s.Count, value)
}

func (s *SpaAlertData) GetSpaSN(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(s.SpaSN, value)
}

type GetSpaAlertListResponse struct {
	DefaultResponse
	Data *SpaAlertData `json:"data,omitempty"`
}

// GetPcsAlertList
type PcsAlertData struct {
	Count  *int         `json:"count,omitempty"`
	PcsSN  *string      `json:"pcs_sn,omitempty"`
	Alarms []*AlarmItem `json:"alarms,omitempty"`
}

func (p *PcsAlertData) GetCount(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(p.Count, value)
}

func (p *PcsAlertData) GetPcsSN(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.PcsSN, value)
}

type GetPcsAlertListResponse struct {
	DefaultResponse
	Data *PcsAlertData `json:"data,omitempty"`
}

// GetHpsAlertList
type HpsAlertData struct {
	Count  *int         `json:"count,omitempty"`
	HpsSN  *string      `json:"hps_sn,omitempty"`
	Alarms []*AlarmItem `json:"alarms,omitempty"`
}

func (h *HpsAlertData) GetCount(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(h.Count, value)
}

func (h *HpsAlertData) GetHpsSN(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(h.HpsSN, value)
}

type GetHpsAlertListResponse struct {
	DefaultResponse
	Data *HpsAlertData `json:"data,omitempty"`
}

// GetPbdAlertList
type PbdAlertData struct {
	Count  *int         `json:"count,omitempty"`
	PbdSN  *string      `json:"pbd_sn,omitempty"`
	Alarms []*AlarmItem `json:"alarms,omitempty"`
}

func (p *PbdAlertData) GetCount(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(p.Count, value)
}

func (p *PbdAlertData) GetPbdSN(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.PbdSN, value)
}

type GetPbdAlertListResponse struct {
	DefaultResponse
	Data *PbdAlertData `json:"data,omitempty"`
}

// |=> GetHistoricalPlantPowerGeneration
type GetHistoricalPlantPowerGenerationResponse struct {
	DefaultResponse
	Data *GetHistoricalPlantPowerGenerationData `json:"data,omitempty"`
}

type GetHistoricalPlantPowerGenerationData struct {
	Count    *int                                   `json:"count,omitempty"`
	TimeUnit *string                                `json:"time_unit,omitempty"`
	Energys  []HistoricalPlantPowerGenerationEnergy `json:"energys,omitempty"`
}

func (h GetHistoricalPlantPowerGenerationData) GetCount() int {
	return pointy.IntValue(h.Count, 0)
}

func (h GetHistoricalPlantPowerGenerationData) GetTimeUnit(fallback ...string) string {
	val := "day"
	if len(fallback) > 0 {
		val = fallback[0]
	}
	return pointy.StringValue(h.TimeUnit, val)
}

type HistoricalPlantPowerGenerationEnergy struct {
	Date   any     `json:"date"`
	Energy *string `json:"energy"`
}

func (h HistoricalPlantPowerGenerationEnergy) GetDate() *time.Time {
	if h.Date == nil {
		return nil
	}

	switch v := h.Date.(type) {
	case *string:
		date, err := time.Parse(*v, "2006-01-02")
		if err != nil {
			return nil
		}
		return &date
	case int:
		date := time.Date(v, 1, 1, 0, 0, 0, 0, time.Local)
		return &date
	default:
		return nil
	}
}

func (h HistoricalPlantPowerGenerationEnergy) GetEnergy() float64 {
	if h.Energy == nil {
		return 0
	}

	energy, err := strconv.ParseFloat(*h.Energy, 64)
	if err != nil {
		return 0
	}

	return energy
}

// |=> GetPlantBasicInfo
type GetPlantBasicInfoResponse struct {
	DefaultResponse
	Data *GetPlantBasicInfoData `json:"data,omitempty"`
}

type GetPlantBasicInfoData struct {
	Address1            *string  `json:"address1,omitempty"`
	InstalledDcCapacity *string  `json:"installed_dc_capacity,omitempty"`
	City                *string  `json:"city,omitempty"`
	Longitude           *string  `json:"longitude,omitempty"`
	Country             *string  `json:"country,omitempty"`
	Latitude            *string  `json:"latitude,omitempty"`
	Locale              *string  `json:"locale,omitempty"`
	Currency            *string  `json:"currency,omitempty"`
	Name                *string  `json:"name,omitempty"`
	PeakPower           *float64 `json:"peak_power,omitempty"`
}

func (s GetPlantBasicInfoData) GetName(fallback ...string) string {
	val := ""
	if len(fallback) > 0 {
		val = fallback[0]
	}

	return pointy.StringValue(s.Name, val)
}

func (s GetPlantBasicInfoData) GetLongitude(fallback ...string) string {
	val := ""
	if len(fallback) > 0 {
		val = fallback[0]
	}

	return pointy.StringValue(s.Longitude, val)
}

func (s GetPlantBasicInfoData) GetLatitude(fallback ...string) string {
	val := ""
	if len(fallback) > 0 {
		val = fallback[0]
	}

	return pointy.StringValue(s.Latitude, val)
}

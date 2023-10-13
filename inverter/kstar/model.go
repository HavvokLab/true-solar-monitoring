package kstar

import "go.openly.dev/pointy"

type Meta struct {
	Success *bool   `json:"success,omitempty"`
	Code    *string `json:"code,omitempty"`
}

func (m *Meta) GetSuccess(defaultValue ...bool) bool {
	val := false
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.BoolValue(m.Success, val)
}

func (m *Meta) GetCode(defaultValue ...string) string {
	val := ""
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.StringValue(m.Code, val)
}

type MetaResponse struct {
	Meta *Meta `json:"meta,omitempty"`
}

type PlantInfoItem struct {
	ID                *string  `json:"powerId,omitempty"`
	Name              *string  `json:"powerName,omitempty"`
	InstalledCapacity *float64 `json:"powerCap,omitempty"`
	Longitude         *float64 `json:"longitude,omitempty"`
	Latitude          *float64 `json:"latitude,omitempty"`
	CityCode          *string  `json:"cityCode,omitempty"`
	CityName          *string  `json:"cityName,omitempty"`
	DistrictCode      *string  `json:"districtCode,omitempty"`
	Address           *string  `json:"powerArea,omitempty"`
	CreatedTime       *string  `json:"createTime,omitempty"`
	DealerCode        *string  `json:"dealerCode,omitempty"`
	ElectricPrice     *float64 `json:"elecPrice,omitempty"`
	ElectricUnit      *string  `json:"elecUnit,omitempty"`
	TimeZone          *string  `json:"timeZone,omitempty"`
}

func (p *PlantInfoItem) GetID(defaultValue ...string) string {
	val := ""
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.StringValue(p.ID, val)
}

func (p *PlantInfoItem) GetName(defaultValue ...string) string {
	val := ""
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.StringValue(p.Name, val)
}

func (p *PlantInfoItem) GetInstalledCapacity(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(p.InstalledCapacity, val)
}

func (p *PlantInfoItem) GetLongitude(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(p.Longitude, val)
}

func (p *PlantInfoItem) GetLatitude(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(p.Latitude, val)
}

func (p *PlantInfoItem) GetCityCode(defaultValue ...string) string {
	val := ""
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.StringValue(p.CityCode, val)
}

func (p *PlantInfoItem) GetCityName(defaultValue ...string) string {
	val := ""
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.StringValue(p.CityName, val)
}

func (p *PlantInfoItem) GetDistrictCode(defaultValue ...string) string {
	val := ""
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.StringValue(p.DistrictCode, val)
}

func (p *PlantInfoItem) GetAddress(defaultValue ...string) string {
	val := ""
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.StringValue(p.Address, val)
}

func (p *PlantInfoItem) GetCreatedTime(defaultValue ...string) string {
	val := ""
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.StringValue(p.CreatedTime, val)
}

func (p *PlantInfoItem) GetDealerCode(defaultValue ...string) string {
	val := ""
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.StringValue(p.DealerCode, val)
}

func (p *PlantInfoItem) GetElectricPrice(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(p.ElectricPrice, val)
}

func (p *PlantInfoItem) GetElectricUnit(defaultValue ...string) string {
	val := ""
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.StringValue(p.ElectricUnit, val)
}

func (p *PlantInfoItem) GetTimeZone(defaultValue ...string) string {
	val := ""
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.StringValue(p.TimeZone, val)
}

type GetPlantListResponse struct {
	Meta *Meta            `json:"meta,omitempty"`
	Data []*PlantInfoItem `json:"data,omitempty"`
}

type DeviceInfoItem struct {
	ID                 *string  `json:"device_id,omitempty"`
	SN                 *string  `json:"inverter_id,omitempty"`
	Status             *int     `json:"status,omitempty"`
	Version            *int     `json:"version,omitempty"`
	SaveTime           *string  `json:"save_time,omitempty"`
	VoltagePV1         *float64 `json:"voltage_pv1,omitempty"`
	VoltagePV2         *float64 `json:"voltage_pv2,omitempty"`
	VoltagePV3         *float64 `json:"voltage_pv3,omitempty"`
	CurrentPV1         *float64 `json:"current_pv1,omitempty"`
	CurrentPV2         *float64 `json:"current_pv2,omitempty"`
	CurrentPV3         *float64 `json:"current_pv3,omitempty"`
	PowerPV1           *float64 `json:"power_pv1,omitempty"`
	PowerPV2           *float64 `json:"power_pv2,omitempty"`
	PowerPV3           *float64 `json:"power_pv3,omitempty"`
	VoltagePBUS        *float64 `json:"voltage_pbus,omitempty"`
	VoltageNBUS        *float64 `json:"voltage_nbus,omitempty"`
	VoltageRS          *float64 `json:"voltage_rs,omitempty"`
	VoltageST          *float64 `json:"voltage_st,omitempty"`
	VoltageTR          *float64 `json:"voltage_tr,omitempty"`
	FrequencyRS        *float64 `json:"frequency_rs,omitempty"`
	FrequencyST        *float64 `json:"frequency_st,omitempty"`
	FrequencyTR        *float64 `json:"frequency_tr,omitempty"`
	CurrentR           *float64 `json:"current_r,omitempty"`
	CurrentS           *float64 `json:"current_s,omitempty"`
	CurrentT           *float64 `json:"current_t,omitempty"`
	PowerInter         *float64 `json:"power_inter,omitempty"`
	RadiatorTemp       *float64 `json:"radiator_temp,omitempty"`
	ModuleTemp         *float64 `json:"module_temp,omitempty"`
	DSPAlarmCode       *string  `json:"dsp_alarm_code,omitempty"`
	DSPErrorCode       *string  `json:"dsp_error_code,omitempty"`
	InverterType       *int     `json:"inverter_type,omitempty"`
	DeviceModel        *int     `json:"device_model,omitempty"`
	FanA               *float64 `json:"fan_A,omitempty"`
	FanB               *float64 `json:"fan_B,omitempty"`
	FanC               *float64 `json:"fan_C,omitempty"`
	ARMAlarmCode       *string  `json:"arm_alarm_code,omitempty"`
	ARMErrorCode       *string  `json:"arm_error_code,omitempty"`
	InputType          *int     `json:"input_type,omitempty"`
	GenerationStandard *int     `json:"generation_standard,omitempty"`
	TotalGeneration    *float64 `json:"total_generation,omitempty"`
	YearGeneration     *float64 `json:"year_generation,omitempty"`
	DayGeneration      *float64 `json:"day_generation,omitempty"`
	MonthGeneration    *float64 `json:"month_generation,omitempty"`
	VoltageOpen        *float64 `json:"voltage_open,omitempty"`
	DelayTime          *float64 `json:"delay_time,omitempty"`
	VoltageLower       *float64 `json:"voltage_lower,omitempty"`
	VoltageCeiling     *float64 `json:"voltage_ceiling,omitempty"`
	PowerLower         *float64 `json:"power_lower,omitempty"`
	PowerCeiling       *float64 `json:"power_ceiling,omitempty"`
	ReactSet           *float64 `json:"react_set,omitempty"`
	ActiveSet          *float64 `json:"active_set,omitempty"`
	ReactiveSet        *float64 `json:"reactive_set,omitempty"`
	ReactiveType       *int     `json:"reactive_type,omitempty"`
	PowerApparent      *float64 `json:"power_apparent,omitempty"`
	PowerReactive      *float64 `json:"power_reactive,omitempty"`
	PowerFactors       *float64 `json:"power_factors,omitempty"`
	CurrentImpedance   *float64 `json:"current_impedance,omitempty"`
	OverPower          *int     `json:"over_power,omitempty"`
	OverPowerSet       *float64 `json:"over_power_set,omitempty"`
	QVHighSet          *float64 `json:"qv_high_set,omitempty"`
	QVHighPercentage   *float64 `json:"qv_high_percentage,omitempty"`
	QVLowerSet         *float64 `json:"qv_lower_set,omitempty"`
	QVLowerPercentage  *float64 `json:"qv_lower_percentage,omitempty"`
	WeakGeneration     *float64 `json:"weak_generation,omitempty"`
	InverterPower      *float64 `json:"inverter_power,omitempty"`
	InverterMake       *int     `json:"inverter_make,omitempty"`
	DREDMake           *int     `json:"dred_make,omitempty"`
}

func (d *DeviceInfoItem) GetID(defaultValue ...string) string {
	val := ""
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.StringValue(d.ID, val)
}

func (d *DeviceInfoItem) GetSN(defaultValue ...string) string {
	val := ""
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.StringValue(d.SN, val)
}

func (d *DeviceInfoItem) GetStatus(defaultValue ...int) int {
	val := 0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.IntValue(d.Status, val)
}

func (d *DeviceInfoItem) GetVersion(defaultValue ...int) int {
	val := 0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.IntValue(d.Version, val)
}

func (d *DeviceInfoItem) GetSaveTime(defaultValue ...string) string {
	val := ""
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.StringValue(d.SaveTime, val)
}

func (d *DeviceInfoItem) GetVoltagePV1(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.VoltagePV1, val)
}

func (d *DeviceInfoItem) GetVoltagePV2(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.VoltagePV2, val)
}

func (d *DeviceInfoItem) GetVoltagePV3(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.VoltagePV3, val)
}

func (d *DeviceInfoItem) GetCurrentPV1(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.CurrentPV1, val)
}

func (d *DeviceInfoItem) GetCurrentPV2(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.CurrentPV2, val)
}

func (d *DeviceInfoItem) GetCurrentPV3(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.CurrentPV3, val)
}

func (d *DeviceInfoItem) GetPowerPV1(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.PowerPV1, val)
}

func (d *DeviceInfoItem) GetPowerPV2(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.PowerPV2, val)
}

func (d *DeviceInfoItem) GetPowerPV3(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.PowerPV3, val)
}

func (d *DeviceInfoItem) GetVoltagePBUS(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.VoltagePBUS, val)
}

func (d *DeviceInfoItem) GetVoltageNBUS(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.VoltageNBUS, val)
}

func (d *DeviceInfoItem) GetVoltageRS(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.VoltageRS, val)
}

func (d *DeviceInfoItem) GetVoltageST(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.VoltageST, val)
}

func (d *DeviceInfoItem) GetVoltageTR(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.VoltageTR, val)
}

func (d *DeviceInfoItem) GetFrequencyRS(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.FrequencyRS, val)
}

func (d *DeviceInfoItem) GetFrequencyST(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.FrequencyST, val)
}

func (d *DeviceInfoItem) GetFrequencyTR(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.FrequencyTR, val)
}

func (d *DeviceInfoItem) GetCurrentR(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.CurrentR, val)
}

func (d *DeviceInfoItem) GetCurrentS(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.CurrentS, val)
}

func (d *DeviceInfoItem) GetCurrentT(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.CurrentT, val)
}

func (d *DeviceInfoItem) GetPowerInter(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.PowerInter, val)
}

func (d *DeviceInfoItem) GetRadiatorTemp(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.RadiatorTemp, val)
}

func (d *DeviceInfoItem) GetModuleTemp(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.ModuleTemp, val)
}

func (d *DeviceInfoItem) GetDSPAlarmCode(defaultValue ...string) string {
	val := ""
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.StringValue(d.DSPAlarmCode, val)
}

func (d *DeviceInfoItem) GetDSPErrorCode(defaultValue ...string) string {
	val := ""
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.StringValue(d.DSPErrorCode, val)
}

func (d *DeviceInfoItem) GetInverterType(defaultValue ...int) int {
	val := 0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.IntValue(d.InverterType, val)
}

func (d *DeviceInfoItem) GetDeviceModel(defaultValue ...int) int {
	val := 0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.IntValue(d.DeviceModel, val)
}

func (d *DeviceInfoItem) GetFanA(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.FanA, val)
}

func (d *DeviceInfoItem) GetFanB(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.FanB, val)
}

func (d *DeviceInfoItem) GetFanC(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.FanC, val)
}

func (d *DeviceInfoItem) GetARMAlarmCode(defaultValue ...string) string {
	val := ""
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.StringValue(d.ARMAlarmCode, val)
}

func (d *DeviceInfoItem) GetARMErrorCode(defaultValue ...string) string {
	val := ""
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.StringValue(d.ARMErrorCode, val)
}

func (d *DeviceInfoItem) GetInputType(defaultValue ...int) int {
	val := 0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.IntValue(d.InputType, val)
}

func (d *DeviceInfoItem) GetGenerationStandard(defaultValue ...int) int {
	val := 0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.IntValue(d.GenerationStandard, val)
}

func (d *DeviceInfoItem) GetTotalGeneration(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.TotalGeneration, val)
}

func (d *DeviceInfoItem) GetYearGeneration(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.YearGeneration, val)
}

func (d *DeviceInfoItem) GetDayGeneration(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.DayGeneration, val)
}

func (d *DeviceInfoItem) GetMonthGeneration(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.MonthGeneration, val)
}

func (d *DeviceInfoItem) GetVoltageOpen(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.VoltageOpen, val)
}

func (d *DeviceInfoItem) GetDelayTime(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.DelayTime, val)
}

func (d *DeviceInfoItem) GetVoltageLower(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.VoltageLower, val)
}

func (d *DeviceInfoItem) GetVoltageCeiling(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.VoltageCeiling, val)
}

func (d *DeviceInfoItem) GetPowerLower(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.PowerLower, val)
}

func (d *DeviceInfoItem) GetPowerCeiling(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.PowerCeiling, val)
}

func (d *DeviceInfoItem) GetReactSet(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.ReactSet, val)
}

func (d *DeviceInfoItem) GetActiveSet(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.ActiveSet, val)
}

func (d *DeviceInfoItem) GetReactiveSet(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.ReactiveSet, val)
}

func (d *DeviceInfoItem) GetReactiveType(defaultValue ...int) int {
	val := 0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.IntValue(d.ReactiveType, val)
}

func (d *DeviceInfoItem) GetPowerApparent(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.PowerApparent, val)
}

func (d *DeviceInfoItem) GetPowerReactive(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.PowerReactive, val)
}

func (d *DeviceInfoItem) GetPowerFactors(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.PowerFactors, val)
}

func (d *DeviceInfoItem) GetCurrentImpedance(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.CurrentImpedance, val)
}

func (d *DeviceInfoItem) GetOverPower(defaultValue ...int) int {
	val := 0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.IntValue(d.OverPower, val)
}

func (d *DeviceInfoItem) GetOverPowerSet(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.OverPowerSet, val)
}

func (d *DeviceInfoItem) GetQVHighSet(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.QVHighSet, val)
}

func (d *DeviceInfoItem) GetQVHighPercentage(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.QVHighPercentage, val)
}

func (d *DeviceInfoItem) GetQVLowerSet(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.QVLowerSet, val)
}

func (d *DeviceInfoItem) GetQVLowerPercentage(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.QVLowerPercentage, val)
}

func (d *DeviceInfoItem) GetWeakGeneration(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.WeakGeneration, val)
}

func (d *DeviceInfoItem) GetInverterPower(defaultValue ...float64) float64 {
	val := 0.0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.Float64Value(d.InverterPower, val)
}

func (d *DeviceInfoItem) GetInverterMake(defaultValue ...int) int {
	val := 0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.IntValue(d.InverterMake, val)
}

func (d *DeviceInfoItem) GetDREDMake(defaultValue ...int) int {
	val := 0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.IntValue(d.DREDMake, val)
}

type GetRealtimeDeviceDataResponse struct {
	Meta *Meta           `json:"meta,omitempty"`
	Data *DeviceInfoItem `json:"data,omitempty"`
}

type DeviceItem struct {
	ID        *string `json:"deviceId,omitempty"`
	SN        *string `json:"inverterId,omitempty"`
	Name      *string `json:"deviceName,omitempty"`
	Status    *int    `json:"deviceStatus,omitempty"`
	PlantID   *string `json:"powerId,omitempty"`
	PlantName *string `json:"powerName,omitempty"`
	SaveTime  *string `json:"saveTime,omitempty"`
}

func (d *DeviceItem) GetID(defaultValue ...string) string {
	val := ""
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.StringValue(d.ID, val)
}

func (d *DeviceItem) GetSN(defaultValue ...string) string {
	val := ""
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.StringValue(d.SN, val)
}

func (d *DeviceItem) GetName(defaultValue ...string) string {
	val := ""
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.StringValue(d.Name, val)
}

func (d *DeviceItem) GetStatus(defaultValue ...int) int {
	val := 0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.IntValue(d.Status, val)
}

func (d *DeviceItem) GetPlantID(defaultValue ...string) string {
	val := ""
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.StringValue(d.PlantID, val)
}

func (d *DeviceItem) GetPlantName(defaultValue ...string) string {
	val := ""
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.StringValue(d.PlantName, val)
}

func (d *DeviceItem) GetSaveTime(defaultValue ...string) string {
	val := ""
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.StringValue(d.SaveTime, val)
}

type DeviceAlarmInfoItem struct {
	DeviceID   *string `json:"deviceId,omitempty"`
	DeviceName *string `json:"deviceName,omitempty"`
	SaveTime   *string `json:"saveTime,omitempty"`
	Message    *string `json:"message,omitempty"`
	ErrorLevel *int    `json:"errorLevel,omitempty"`
	RemoveTime *string `json:"removeTime,omitempty"`
	PlantID    *string `json:"powerId,omitempty"`
	PlantName  *string `json:"powerName,omitempty"`
}

func (d *DeviceAlarmInfoItem) GetDeviceID(defaultValue ...string) string {
	val := ""
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.StringValue(d.DeviceID, val)
}

func (d *DeviceAlarmInfoItem) GetDeviceName(defaultValue ...string) string {
	val := ""
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.StringValue(d.DeviceName, val)
}

func (d *DeviceAlarmInfoItem) GetSaveTime(defaultValue ...string) string {
	val := ""
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.StringValue(d.SaveTime, val)
}

func (d *DeviceAlarmInfoItem) GetMessage(defaultValue ...string) string {
	val := ""
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.StringValue(d.Message, val)
}

func (d *DeviceAlarmInfoItem) GetErrorLevel(defaultValue ...int) int {
	val := 0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.IntValue(d.ErrorLevel, val)
}

func (d *DeviceAlarmInfoItem) GetRemoveTime(defaultValue ...string) string {
	val := ""
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.StringValue(d.RemoveTime, val)
}

func (d *DeviceAlarmInfoItem) GetPlantID(defaultValue ...string) string {
	val := ""
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.StringValue(d.PlantID, val)
}

func (d *DeviceAlarmInfoItem) GetPlantName(defaultValue ...string) string {
	val := ""
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.StringValue(d.PlantName, val)
}

type GetRealtimeAlarmListOfDevice struct {
	Meta *Meta                  `json:"meta,omitempty"`
	Data []*DeviceAlarmInfoItem `json:"data,omitempty"`
}

type GetDeviceListResponse struct {
	Meta *Meta              `json:"meta,omitempty"`
	Data *GetDeviceListData `json:"data,omitempty"`
}

type GetDeviceListData struct {
	Code  *int          `json:"code,omitempty"`
	Count *int          `json:"count,omitempty"`
	List  []*DeviceItem `json:"list,omitempty"`
}

func (g *GetDeviceListData) GetCode(defaultValue ...int) int {
	val := 0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.IntValue(g.Code, val)
}

func (g *GetDeviceListData) GetCount(defaultValue ...int) int {
	val := 0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.IntValue(g.Count, val)
}

type GetAlarmListResponse struct {
	Meta *Meta             `json:"meta,omitempty"`
	Data *GetAlarmListData `json:"data,omitempty"`
}

type GetAlarmListData struct {
	Code  *int                   `json:"code,omitempty"`
	Count *int                   `json:"count,omitempty"`
	Data  []*DeviceAlarmInfoItem `json:"data,omitempty"`
}

func (g *GetAlarmListData) GetCode(defaultValue ...int) int {
	val := 0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.IntValue(g.Code, val)
}

func (g *GetAlarmListData) GetCount(defaultValue ...int) int {
	val := 0
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	return pointy.IntValue(g.Count, val)
}

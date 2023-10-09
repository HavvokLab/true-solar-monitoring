package huawei

import "go.openly.dev/pointy"

type ResponseData struct {
	Success    *bool        `json:"success,omitempty"`
	FailCode   *int         `json:"failCode,omitempty"`
	Parameters *interface{} `json:"params,omitempty"`
	Message    *string      `json:"message,omitempty"`
}

func (d *ResponseData) GetSuccess(defaultValue ...bool) bool {
	value := false
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.BoolValue(d.Success, value)
}

func (d *ResponseData) GetFailCode(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(d.FailCode, value)
}

func (d *ResponseData) GetParameters(defaultValue ...interface{}) interface{} {
	var value interface{} = nil
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.PointerValue(d.Parameters, value)
}

func (d *ResponseData) GetMessage(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.Message, value)
}

// API 2.1
type GetTokenResponse struct {
	ResponseData
	Data *string `json:"data,omitempty"`
}

func (d *GetTokenResponse) GetData(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.Data, value)
}

// API 2.2
type GetPlantListResponse struct {
	ResponseData
	Data []*PlantItem `json:"data,omitempty"`
}

type PlantItem struct {
	Code           *string  `json:"stationCode,omitempty"`
	Name           *string  `json:"stationName,omitempty"`
	Address        *string  `json:"stationAddr,omitempty"`
	Capacity       *float64 `json:"capacity,omitempty"`
	BuildState     *string  `json:"buildState,omitempty"`
	CombineType    *string  `json:"combineType,omitempty"`
	AIDType        *int     `json:"aidType,omitempty"`
	StationLinkMan *string  `json:"stationLinkman,omitempty"`
	LinkManPhone   *string  `json:"linkmanPho,omitempty"`
}

func (p *PlantItem) GetCode(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.Code, value)
}

func (p *PlantItem) GetName(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.Name, value)
}

func (p *PlantItem) GetAddress(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.Address, value)
}

func (p *PlantItem) GetCapacity(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.Capacity, value)
}

func (p *PlantItem) GetBuildState(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.BuildState, value)
}

func (p *PlantItem) GetCombineType(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.CombineType, value)
}

func (p *PlantItem) GetAIDType(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(p.AIDType, value)
}

func (p *PlantItem) GetStationLinkMan(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.StationLinkMan, value)
}

func (p *PlantItem) GetLinkManPhone(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.LinkManPhone, value)
}

// API 2.3
type GetRealtimePlantDataResponse struct {
	ResponseData
	Data []*RealtimePlantData `json:"data,omitempty"`
}

type RealtimePlantData struct {
	Code        *string            `json:"stationCode,omitempty"`
	DataItemMap *RealtimePlantItem `json:"dataItemMap,omitempty"`
}

func (d *RealtimePlantData) GetCode(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.Code, value)
}

type RealtimePlantItem struct {
	TotalIncome     *float64 `json:"total_income,omitempty"`
	TotalPower      *float64 `json:"total_power,omitempty"`
	DayPower        *float64 `json:"day_power,omitempty"`
	DayIncome       *float64 `json:"day_income,omitempty"`
	RealHealthState *int     `json:"real_health_state,omitempty"`
	MonthPower      *float64 `json:"month_power,omitempty"`
}

func (r *RealtimePlantItem) GetTotalIncome(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.TotalIncome, value)
}

func (r *RealtimePlantItem) GetTotalPower(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.TotalPower, value)
}

func (r *RealtimePlantItem) GetDayPower(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.DayPower, value)
}

func (r *RealtimePlantItem) GetDayIncome(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.DayIncome, value)
}

func (r *RealtimePlantItem) GetRealHealthState(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(r.RealHealthState, value)
}

func (r *RealtimePlantItem) GetMonthPower(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.MonthPower, value)
}

// API 2.5, 2.6, 2.7
type GetHistoricalPlantDataResponse struct {
	ResponseData
	Data []*HistoricalPlantData `json:"data,omitempty"`
}

type HistoricalPlantData struct {
	Code        *string              `json:"stationCode,omitempty"`
	CollectTime *int64               `json:"collectTime,omitempty"`
	DataItemMap *HistoricalPlantItem `json:"dataItemMap,omitempty"`
}

func (d *HistoricalPlantData) GetCode(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.Code, value)
}

func (d *HistoricalPlantData) GetCollectTime(defaultValue ...int64) int64 {
	value := int64(0)
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Int64Value(d.CollectTime, value)
}

type HistoricalPlantItem struct {
	RadiationIntensity *float64 `json:"radiation_intensity,omitempty"`
	InstalledCapacity  *float64 `json:"installed_capacity,omitempty"`
	UsePower           *float64 `json:"use_power,omitempty"`
	InverterPower      *float64 `json:"inverter_power,omitempty"`
	PowerProfit        *float64 `json:"power_profit,omitempty"`
	TheoryPower        *float64 `json:"theory_power,omitempty"`
	PerPowerRatio      *float64 `json:"perpower_ratio,omitempty"`
	OnGridPower        *float64 `json:"ongrid_power,omitempty"`
	PerformanceRatio   *float64 `json:"performance_ratio,omitempty"`
	ReductionTotalCO2  *float64 `json:"reduction_total_co2,omitempty"`
	ReductionTotalCoal *float64 `json:"reduction_total_coal,omitempty"`
	ReductionTotalTree *float64 `json:"reduction_total_tree,omitempty"`
}

func (h *HistoricalPlantItem) GetRadiationIntensity(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(h.RadiationIntensity, value)
}

func (h *HistoricalPlantItem) GetInstalledCapacity(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(h.InstalledCapacity, value)
}

func (h *HistoricalPlantItem) GetUsePower(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(h.UsePower, value)
}

func (h *HistoricalPlantItem) GetInverterPower(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(h.InverterPower, value)
}

func (h *HistoricalPlantItem) GetPowerProfit(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(h.PowerProfit, value)
}

func (h *HistoricalPlantItem) GetTheoryPower(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(h.TheoryPower, value)
}

func (h *HistoricalPlantItem) GetPerPowerRatio(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(h.PerPowerRatio, value)
}

func (h *HistoricalPlantItem) GetOnGridPower(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(h.OnGridPower, value)
}

func (h *HistoricalPlantItem) GetPerformanceRatio(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(h.PerformanceRatio, value)
}

func (h *HistoricalPlantItem) GetReductionTotalCO2(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(h.ReductionTotalCO2, value)
}

func (h *HistoricalPlantItem) GetReductionTotalCoal(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(h.ReductionTotalCoal, value)
}

func (h *HistoricalPlantItem) GetReductionTotalTree(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(h.ReductionTotalTree, value)
}

// API 2.8
type GetDeviceListResponse struct {
	ResponseData
	Data []*DeviceItem `json:"data,omitempty"`
}

type DeviceItem struct {
	ID              *int     `json:"id,omitempty"`
	SN              *string  `json:"esnCode,omitempty"`
	Name            *string  `json:"devName,omitempty"`
	TypeID          *int     `json:"devTypeId,omitempty"`
	InverterModel   *string  `json:"invType,omitempty"`
	Latitude        *float64 `json:"latitude,omitempty"`
	Longitude       *float64 `json:"longitude,omitempty"`
	SoftwareVersion *string  `json:"softwareVersion,omitempty"`
	PlantCode       *string  `json:"stationCode,omitempty"`
}

func (d *DeviceItem) GetID(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(d.ID, value)
}

func (d *DeviceItem) GetSN(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.SN, value)
}

func (d *DeviceItem) GetName(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.Name, value)
}

func (d *DeviceItem) GetTypeID(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(d.TypeID, value)
}

func (d *DeviceItem) GetInverterModel(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.InverterModel, value)
}

func (d *DeviceItem) GetLatitude(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(d.Latitude, value)
}

func (d *DeviceItem) GetLongitude(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(d.Longitude, value)
}

func (d *DeviceItem) GetSoftwareVersion(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.SoftwareVersion, value)
}

func (d *DeviceItem) GetPlantCode(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.PlantCode, value)
}

// API 2.9
type GetRealtimeDeviceDataResponse struct {
	ResponseData
	Data []*RealtimeDeviceData `json:"data,omitempty"`
}

type RealtimeDeviceData struct {
	ID          *int                `json:"devId,omitempty"`
	DataItemMap *RealtimeDeviceItem `json:"dataItemMap,omitempty"`
}

func (d *RealtimeDeviceData) GetID(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(d.ID, value)
}

type RealtimeDeviceItem struct {
	InverterState      *float64     `json:"inverter_state,omitempty"`
	GridABVoltage      *float64     `json:"ab_u,omitempty"`
	GridBCVoltage      *float64     `json:"bc_u,omitempty"`
	GridCAVoltage      *float64     `json:"ca_u,omitempty"`
	PhaseAVoltage      *float64     `json:"a_u,omitempty"`
	PhaseBVoltage      *float64     `json:"b_u,omitempty"`
	PhaseCVoltage      *float64     `json:"c_u,omitempty"`
	GridPhaseACurrent  *float64     `json:"a_i,omitempty"`
	GridPhaseBCurrent  *float64     `json:"b_i,omitempty"`
	GridPhaseCCurrent  *float64     `json:"c_i,omitempty"`
	Efficiency         *float64     `json:"efficiency,omitempty"`
	Temperature        *float64     `json:"temperature,omitempty"`
	PowerFactor        *float64     `json:"power_factor,omitempty"`
	GridFrequency      *float64     `json:"elec_freq,omitempty"`
	ActivePower        *float64     `json:"active_power,omitempty"`
	ReactivePower      *float64     `json:"reactive_power,omitempty"`
	DailyEnergy        *float64     `json:"day_cap,omitempty"`
	MPPTTotalPower     *float64     `json:"mppt_power,omitempty"`
	PV1InputVoltage    *float64     `json:"pv1_u,omitempty"`
	PV2InputVoltage    *float64     `json:"pv2_u,omitempty"`
	PV3InputVoltage    *float64     `json:"pv3_u,omitempty"`
	PV4InputVoltage    *float64     `json:"pv4_u,omitempty"`
	PV5InputVoltage    *float64     `json:"pv5_u,omitempty"`
	PV6InputVoltage    *float64     `json:"pv6_u,omitempty"`
	PV7InputVoltage    *float64     `json:"pv7_u,omitempty"`
	PV8InputVoltage    *float64     `json:"pv8_u,omitempty"`
	PV9InputVoltage    *float64     `json:"pv9_u,omitempty"`
	PV10InputVoltage   *float64     `json:"pv10_u,omitempty"`
	PV11InputVoltage   *float64     `json:"pv11_u,omitempty"`
	PV12InputVoltage   *float64     `json:"pv12_u,omitempty"`
	PV13InputVoltage   *float64     `json:"pv13_u,omitempty"`
	PV14InputVoltage   *float64     `json:"pv14_u,omitempty"`
	PV15InputVoltage   *float64     `json:"pv15_u,omitempty"`
	PV16InputVoltage   *float64     `json:"pv16_u,omitempty"`
	PV17InputVoltage   *float64     `json:"pv17_u,omitempty"`
	PV18InputVoltage   *float64     `json:"pv18_u,omitempty"`
	PV19InputVoltage   *float64     `json:"pv19_u,omitempty"`
	PV20InputVoltage   *float64     `json:"pv20_u,omitempty"`
	PV21InputVoltage   *float64     `json:"pv21_u,omitempty"`
	PV22InputVoltage   *float64     `json:"pv22_u,omitempty"`
	PV23InputVoltage   *float64     `json:"pv23_u,omitempty"`
	PV24InputVoltage   *float64     `json:"pv24_u,omitempty"`
	PV1InputCurrent    *float64     `json:"pv1_i,omitempty"`
	PV2InputCurrent    *float64     `json:"pv2_i,omitempty"`
	PV3InputCurrent    *float64     `json:"pv3_i,omitempty"`
	PV4InputCurrent    *float64     `json:"pv4_i,omitempty"`
	PV5InputCurrent    *float64     `json:"pv5_i,omitempty"`
	PV6InputCurrent    *float64     `json:"pv6_i,omitempty"`
	PV7InputCurrent    *float64     `json:"pv7_i,omitempty"`
	PV8InputCurrent    *float64     `json:"pv8_i,omitempty"`
	PV9InputCurrent    *float64     `json:"pv9_i,omitempty"`
	PV10InputCurrent   *float64     `json:"pv10_i,omitempty"`
	PV11InputCurrent   *float64     `json:"pv11_i,omitempty"`
	PV12InputCurrent   *float64     `json:"pv12_i,omitempty"`
	PV13InputCurrent   *float64     `json:"pv13_i,omitempty"`
	PV14InputCurrent   *float64     `json:"pv14_i,omitempty"`
	PV15InputCurrent   *float64     `json:"pv15_i,omitempty"`
	PV16InputCurrent   *float64     `json:"pv16_i,omitempty"`
	PV17InputCurrent   *float64     `json:"pv17_i,omitempty"`
	PV18InputCurrent   *float64     `json:"pv18_i,omitempty"`
	PV19InputCurrent   *float64     `json:"pv19_i,omitempty"`
	PV20InputCurrent   *float64     `json:"pv20_i,omitempty"`
	PV21InputCurrent   *float64     `json:"pv21_i,omitempty"`
	PV22InputCurrent   *float64     `json:"pv22_i,omitempty"`
	PV23InputCurrent   *float64     `json:"pv23_i,omitempty"`
	PV24InputCurrent   *float64     `json:"pv24_i,omitempty"`
	TotalEnergy        *float64     `json:"total_cap,omitempty"`
	InverterStartup    *interface{} `json:"open_time,omitempty"`
	InverterShutdown   *interface{} `json:"close_time,omitempty"`
	TotalDCInputEnergy *float64     `json:"mppt_total_cap,omitempty"`
	MPPT1DCEnergy      *float64     `json:"mppt_1_cap,omitempty"`
	MPPT2DCEnergy      *float64     `json:"mppt_2_cap,omitempty"`
	MPPT3DCEnergy      *float64     `json:"mppt_3_cap,omitempty"`
	MPPT4DCEnergy      *float64     `json:"mppt_4_cap,omitempty"`
	MPPT5DCEnergy      *float64     `json:"mppt_5_cap,omitempty"`
	MPPT6DCEnergy      *float64     `json:"mppt_6_cap,omitempty"`
	MPPT7DCEnergy      *float64     `json:"mppt_7_cap,omitempty"`
	MPPT8DCEnergy      *float64     `json:"mppt_8_cap,omitempty"`
	MPPT9DCEnergy      *float64     `json:"mppt_9_cap,omitempty"`
	MPPT10DCEnergy     *float64     `json:"mppt_10_cap,omitempty"`
	Status             *int         `json:"run_state,omitempty"`
}

func (r *RealtimeDeviceItem) GetInverterState(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.InverterState, value)
}

func (r *RealtimeDeviceItem) GetGridABVoltage(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.GridABVoltage, value)
}

func (r *RealtimeDeviceItem) GetGridBCVoltage(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.GridBCVoltage, value)
}

func (r *RealtimeDeviceItem) GetGridCAVoltage(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.GridCAVoltage, value)
}

func (r *RealtimeDeviceItem) GetPhaseAVoltage(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PhaseAVoltage, value)
}

func (r *RealtimeDeviceItem) GetPhaseBVoltage(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PhaseBVoltage, value)
}

func (r *RealtimeDeviceItem) GetPhaseCVoltage(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PhaseCVoltage, value)
}

func (r *RealtimeDeviceItem) GetGridPhaseACurrent(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.GridPhaseACurrent, value)
}

func (r *RealtimeDeviceItem) GetGridPhaseBCurrent(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.GridPhaseBCurrent, value)
}

func (r *RealtimeDeviceItem) GetGridPhaseCCurrent(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.GridPhaseCCurrent, value)
}

func (r *RealtimeDeviceItem) GetEfficiency(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Efficiency, value)
}

func (r *RealtimeDeviceItem) GetTemperature(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.Temperature, value)
}

func (r *RealtimeDeviceItem) GetPowerFactor(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PowerFactor, value)
}

func (r *RealtimeDeviceItem) GetGridFrequency(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.GridFrequency, value)
}

func (r *RealtimeDeviceItem) GetActivePower(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.ActivePower, value)
}

func (r *RealtimeDeviceItem) GetReactivePower(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.ReactivePower, value)
}

func (r *RealtimeDeviceItem) GetDailyEnergy(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.DailyEnergy, value)
}

func (r *RealtimeDeviceItem) GetMPPTTotalPower(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.MPPTTotalPower, value)
}

func (r *RealtimeDeviceItem) GetPV1InputVoltage(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV1InputVoltage, value)
}

func (r *RealtimeDeviceItem) GetPV2InputVoltage(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV2InputVoltage, value)
}
func (r *RealtimeDeviceItem) GetPV3InputVoltage(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV3InputVoltage, value)
}
func (r *RealtimeDeviceItem) GetPV4InputVoltage(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV4InputVoltage, value)
}
func (r *RealtimeDeviceItem) GetPV5InputVoltage(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV5InputVoltage, value)
}
func (r *RealtimeDeviceItem) GetPV6InputVoltage(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV6InputVoltage, value)
}
func (r *RealtimeDeviceItem) GetPV7InputVoltage(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV7InputVoltage, value)
}
func (r *RealtimeDeviceItem) GetPV8InputVoltage(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV8InputVoltage, value)
}
func (r *RealtimeDeviceItem) GetPV9InputVoltage(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV9InputVoltage, value)
}
func (r *RealtimeDeviceItem) GetPV10InputVoltage(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV10InputVoltage, value)
}
func (r *RealtimeDeviceItem) GetPV11InputVoltage(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV11InputVoltage, value)
}
func (r *RealtimeDeviceItem) GetPV12InputVoltage(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV12InputVoltage, value)
}
func (r *RealtimeDeviceItem) GetPV13InputVoltage(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV13InputVoltage, value)
}
func (r *RealtimeDeviceItem) GetPV14InputVoltage(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV14InputVoltage, value)
}
func (r *RealtimeDeviceItem) GetPV15InputVoltage(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV15InputVoltage, value)
}
func (r *RealtimeDeviceItem) GetPV16InputVoltage(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV16InputVoltage, value)
}
func (r *RealtimeDeviceItem) GetPV17InputVoltage(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV17InputVoltage, value)
}
func (r *RealtimeDeviceItem) GetPV18InputVoltage(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV18InputVoltage, value)
}
func (r *RealtimeDeviceItem) GetPV19InputVoltage(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV19InputVoltage, value)
}
func (r *RealtimeDeviceItem) GetPV20InputVoltage(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV1InputVoltage, value)
}
func (r *RealtimeDeviceItem) GetPV21InputVoltage(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV1InputVoltage, value)
}
func (r *RealtimeDeviceItem) GetPV22InputVoltage(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV1InputVoltage, value)
}
func (r *RealtimeDeviceItem) GetPV23InputVoltage(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV1InputVoltage, value)
}
func (r *RealtimeDeviceItem) GetPV24InputVoltage(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV1InputVoltage, value)
}

func (r *RealtimeDeviceItem) GetPV1InputCurrent(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV1InputCurrent, value)
}

func (r *RealtimeDeviceItem) GetPV2InputCurrent(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV2InputCurrent, value)
}
func (r *RealtimeDeviceItem) GetPV3InputCurrent(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV3InputCurrent, value)
}
func (r *RealtimeDeviceItem) GetPV4InputCurrent(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV4InputCurrent, value)
}
func (r *RealtimeDeviceItem) GetPV5InputCurrent(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV5InputCurrent, value)
}
func (r *RealtimeDeviceItem) GetPV6InputCurrent(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV6InputCurrent, value)
}
func (r *RealtimeDeviceItem) GetPV7InputCurrent(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV7InputCurrent, value)
}
func (r *RealtimeDeviceItem) GetPV8InputCurrent(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV8InputCurrent, value)
}
func (r *RealtimeDeviceItem) GetPV9InputCurrent(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV9InputCurrent, value)
}
func (r *RealtimeDeviceItem) GetPV10InputCurrent(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV10InputCurrent, value)
}
func (r *RealtimeDeviceItem) GetPV11InputCurrent(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV11InputCurrent, value)
}
func (r *RealtimeDeviceItem) GetPV12InputCurrent(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV12InputCurrent, value)
}
func (r *RealtimeDeviceItem) GetPV13InputCurrent(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV13InputCurrent, value)
}
func (r *RealtimeDeviceItem) GetPV14InputCurrent(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV14InputCurrent, value)
}
func (r *RealtimeDeviceItem) GetPV15InputCurrent(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV15InputCurrent, value)
}
func (r *RealtimeDeviceItem) GetPV16InputCurrent(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV16InputCurrent, value)
}
func (r *RealtimeDeviceItem) GetPV17InputCurrent(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV17InputCurrent, value)
}
func (r *RealtimeDeviceItem) GetPV18InputCurrent(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV18InputCurrent, value)
}
func (r *RealtimeDeviceItem) GetPV19InputCurrent(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV19InputCurrent, value)
}
func (r *RealtimeDeviceItem) GetPV20InputCurrent(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV1InputCurrent, value)
}
func (r *RealtimeDeviceItem) GetPV21InputCurrent(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV1InputCurrent, value)
}
func (r *RealtimeDeviceItem) GetPV22InputCurrent(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV1InputCurrent, value)
}
func (r *RealtimeDeviceItem) GetPV23InputCurrent(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV1InputCurrent, value)
}
func (r *RealtimeDeviceItem) GetPV24InputCurrent(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PV1InputCurrent, value)
}

func (h *RealtimeDeviceItem) GetTotalEnergy(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(h.TotalEnergy, value)
}

func (h *RealtimeDeviceItem) GetInverterStartup(defaultValue ...interface{}) interface{} {
	value := interface{}(nil)
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.PointerValue(h.InverterStartup, value)
}

func (h *RealtimeDeviceItem) GetInverterShutdown(defaultValue ...interface{}) interface{} {
	value := interface{}(nil)
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.PointerValue(h.InverterShutdown, value)
}

func (h *RealtimeDeviceItem) GetTotalDCInputEnergy(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(h.TotalDCInputEnergy, value)
}

func (h *RealtimeDeviceItem) GetMPPT1DCEnergy(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(h.MPPT1DCEnergy, value)
}
func (h *RealtimeDeviceItem) GetMPPT2DCEnergy(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(h.MPPT2DCEnergy, value)
}
func (h *RealtimeDeviceItem) GetMPPT3DCEnergy(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(h.MPPT3DCEnergy, value)
}
func (h *RealtimeDeviceItem) GetMPPT4DCEnergy(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(h.MPPT4DCEnergy, value)
}
func (h *RealtimeDeviceItem) GetMPPT5DCEnergy(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(h.MPPT5DCEnergy, value)
}
func (h *RealtimeDeviceItem) GetMPPT6DCEnergy(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(h.MPPT6DCEnergy, value)
}
func (h *RealtimeDeviceItem) GetMPPT7DCEnergy(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(h.MPPT7DCEnergy, value)
}
func (h *RealtimeDeviceItem) GetMPPT8DCEnergy(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(h.MPPT8DCEnergy, value)
}
func (h *RealtimeDeviceItem) GetMPPT9DCEnergy(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(h.MPPT9DCEnergy, value)
}
func (h *RealtimeDeviceItem) GetMPPT10DCEnergy(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(h.MPPT10DCEnergy, value)
}

func (h *RealtimeDeviceItem) GetStatus(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(h.Status, value)
}

// API 2.11, 2.12, 2.13
type GetHistoricalDeviceDataResponse struct {
	ResponseData
	Data []*HistoricalDeviceData `json:"data,omitempty"`
}

type HistoricalDeviceData struct {
	ID          *interface{}          `json:"devId,omitempty"`
	CollectTime *int64                `json:"collectTime,omitempty"`
	DataItemMap *HistoricalDeviceItem `json:"dataItemMap,omitempty"`
}

func (d *HistoricalDeviceData) GetCollectTime(defaultValue ...int64) int64 {
	value := int64(0)
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Int64Value(d.CollectTime, value)
}

func (d *HistoricalDeviceData) GetID(defaultValue ...interface{}) interface{} {
	value := interface{}(nil)
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.PointerValue(d.ID, value)
}

type HistoricalDeviceItem struct {
	InstalledCapacity *float64 `json:"installed_capacity,omitempty"`
	ProductPower      *float64 `json:"product_power,omitempty"`
	PerPowerRatio     *float64 `json:"perpower_ratio,omitempty"`
}

func (h *HistoricalDeviceItem) GetInstalledCapacity(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(h.InstalledCapacity, value)
}

func (h *HistoricalDeviceItem) GetProductPower(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(h.ProductPower, value)
}

func (h *HistoricalDeviceItem) GetPerPowerRatio(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(h.PerPowerRatio, value)
}

// API 2.17
type GetDeviceAlarmListResponse struct {
	ResponseData
	Data []*DeviceAlarmItem `json:"data,omitempty"`
}

type DeviceAlarmItem struct {
	PlantCode        *string `json:"stationCode,omitempty"`
	PlantName        *string `json:"stationName,omitempty"`
	DeviceSN         *string `json:"esnCode,omitempty"`
	DeviceName       *string `json:"devName,omitempty"`
	DeviceTypeID     *int    `json:"devTypeId,omitempty"`
	AlarmID          *int    `json:"alarmId,omitempty"`
	AlarmName        *string `json:"alarmName,omitempty"`
	AlarmCause       *string `json:"alarmCause,omitempty"`
	AlarmType        *int    `json:"alarmType,omitempty"`
	RepairSuggestion *string `json:"repairSuggestion,omitempty"`
	CauseID          *int    `json:"causeId,omitempty"`
	RaiseTime        *int64  `json:"raiseTime,omitempty"`
	Level            *int    `json:"lev,omitempty"`
	Status           *int    `json:"status,omitempty"`
}

func (d *DeviceAlarmItem) GetPlantCode(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.PlantCode, value)
}

func (d *DeviceAlarmItem) GetPlantName(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.PlantName, value)
}

func (d *DeviceAlarmItem) GetDeviceSN(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.DeviceSN, value)
}

func (d *DeviceAlarmItem) GetDeviceName(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.DeviceName, value)
}

func (d *DeviceAlarmItem) GetDeviceTypeID(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(d.DeviceTypeID, value)
}

func (d *DeviceAlarmItem) GetAlarmID(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(d.AlarmID, value)
}

func (d *DeviceAlarmItem) GetAlarmName(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.AlarmName, value)
}

func (d *DeviceAlarmItem) GetAlarmCause(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.AlarmCause, value)
}

func (d *DeviceAlarmItem) GetAlarmType(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(d.AlarmType, value)
}

func (d *DeviceAlarmItem) GetRepairSuggestion(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.RepairSuggestion, value)
}

func (d *DeviceAlarmItem) GetCauseID(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(d.CauseID, value)
}

func (d *DeviceAlarmItem) GetRaiseTime(defaultValue ...int64) int64 {
	value := int64(0)
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Int64Value(d.RaiseTime, value)
}

func (d *DeviceAlarmItem) GetLevel(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(d.Level, value)
}

func (d *DeviceAlarmItem) GetStatus(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(d.Status, value)
}

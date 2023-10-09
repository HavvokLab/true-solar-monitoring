package solarman

import "go.openly.dev/pointy"

type PaginationRequestBody struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

type StationRequestBody struct {
	StationID int `json:"stationId"`
}

type DeviceRequestBody struct {
	DeviceSN string `json:"deviceSn"`
}

type GetTokenRequestBody struct {
	Username    string `json:"username,omitempty"`
	Password    string `json:"password"`
	AppSecret   string `json:"appSecret"`
	CountryCode string `json:"countryCode,omitempty"`
	Email       string `json:"email,omitempty"`
	Mobile      string `json:"mobile,omitempty"`
	OrgID       int    `json:"orgId,omitempty"`
}

type GetTokenResponse struct {
	Code         *string `json:"code,omitempty"`
	Message      *string `json:"msg,omitempty"`
	Success      *bool   `json:"success,omitempty"`
	RequestID    *string `json:"requestId,omitempty"`
	Uid          *int    `json:"uid,omitempty"`
	TokenType    *string `json:"token_type,omitempty"`
	Scope        *string `json:"scope,omitempty"`
	AccessToken  *string `json:"access_token,omitempty"`
	RefreshToken *string `json:"refresh_token,omitempty"`
	ExpiresIn    *string `json:"expires_in,omitempty"`
}

func (t *GetTokenResponse) GetCode(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(t.Code, value)
}

func (t *GetTokenResponse) GetMessage(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(t.Message, value)
}

func (t *GetTokenResponse) GetSuccess(defaultValue ...bool) bool {
	value := false
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.BoolValue(t.Success, value)
}

func (t *GetTokenResponse) GetRequestID(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(t.RequestID, value)
}

func (t *GetTokenResponse) GetUid(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(t.Uid, value)
}

func (t *GetTokenResponse) GetTokenType(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(t.TokenType, value)
}

func (t *GetTokenResponse) GetScope(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(t.Scope, value)
}

func (t *GetTokenResponse) GetAccessToken(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(t.AccessToken, value)
}

func (t *GetTokenResponse) GetRefreshToken(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(t.RefreshToken, value)
}

func (t *GetTokenResponse) GetExpiresIn(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(t.ExpiresIn, value)
}

type GetUserInfoResponse struct {
	Code        *string                 `json:"code,omitempty"`
	Message     *string                 `json:"msg,omitempty"`
	Success     *bool                   `json:"success,omitempty"`
	RequestID   *string                 `json:"requestId,omitempty"`
	OrgInfoList []*OrganizationInfoItem `json:"orgInfoList,omitempty"`
}

func (u *GetUserInfoResponse) GetCode(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(u.Code, value)
}

func (u *GetUserInfoResponse) GetMessage(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(u.Message, value)
}

func (u *GetUserInfoResponse) GetSuccess(defaultValue ...bool) bool {
	value := false
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.BoolValue(u.Success, value)
}

func (u *GetUserInfoResponse) GetRequestID(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(u.RequestID, value)
}

type OrganizationInfoItem struct {
	CompanyID   *int    `json:"companyId,omitempty"`
	CompanyName *string `json:"companyName,omitempty"`
	RoleName    *string `json:"roleName,omitempty"`
}

func (o *OrganizationInfoItem) GetCompanyID(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(o.CompanyID, value)
}

func (o *OrganizationInfoItem) GetCompanyName(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(o.CompanyName, value)
}

func (o *OrganizationInfoItem) GetRoleName(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(o.RoleName, value)
}

type GetPlantListResponse struct {
	Code        *string      `json:"code,omitempty"`
	Message     *string      `json:"msg,omitempty"`
	Success     *bool        `json:"success,omitempty"`
	RequestID   *string      `json:"requestId,omitempty"`
	Total       *int         `json:"total,omitempty"`
	StationList []*PlantItem `json:"stationList,omitempty"`
}

func (p *GetPlantListResponse) GetCode(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.Code, value)
}

func (p *GetPlantListResponse) GetMessage(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.Message, value)
}

func (p *GetPlantListResponse) GetSuccess(defaultValue ...bool) bool {
	value := false
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.BoolValue(p.Success, value)
}

func (p *GetPlantListResponse) GetRequestID(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.RequestID, value)
}

func (p *GetPlantListResponse) GetTotal(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(p.Total, value)
}

type PlantItem struct {
	ID                      *int     `json:"id,omitempty"`
	Name                    *string  `json:"name,omitempty"`
	LocationLat             *float64 `json:"locationLat,omitempty"`
	LocationLng             *float64 `json:"locationLng,omitempty"`
	LocationAddress         *string  `json:"locationAddress,omitempty"`
	RegionNationID          *int     `json:"regionNationId,omitempty"`
	RegionLevel1            *int     `json:"regionLevel1,omitempty"`
	RegionLevel2            *int     `json:"regionLevel2,omitempty"`
	RegionLevel3            *int     `json:"regionLevel3,omitempty"`
	RegionLevel4            *int     `json:"regionLevel4,omitempty"`
	RegionLevel5            *int     `json:"regionLevel5,omitempty"`
	RegionTimezone          *string  `json:"regionTimezone,omitempty"`
	Type                    *string  `json:"type,omitempty"`
	GridInterconnectionType *string  `json:"gridInterconnectionType,omitempty"`
	InstalledCapacity       *float64 `json:"installedCapacity,omitempty"`
	StartOperatingTime      *float64 `json:"startOperatingTime,omitempty"`
	StationImage            *string  `json:"stationImage,omitempty"`
	CreatedDate             *float64 `json:"createdDate,omitempty"`
	BatterySoc              *float64 `json:"batterySoc,omitempty"`
	NetworkStatus           *string  `json:"networkStatus,omitempty"`
	GenerationPower         *float64 `json:"generationPower,omitempty"`
	LastUpdateTime          *float64 `json:"lastUpdateTime,omitempty"`
}

func (p *PlantItem) GetID(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(p.ID, value)
}

func (p *PlantItem) GetName(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.Name, value)
}

func (p *PlantItem) GetLocationLat(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.LocationLat, value)
}

func (p *PlantItem) GetLocationLng(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.LocationLng, value)
}

func (p *PlantItem) GetLocationAddress(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.LocationAddress, value)
}

func (p *PlantItem) GetRegionNationID(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(p.RegionNationID, value)
}

func (p *PlantItem) GetRegionLevel1(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(p.RegionLevel1, value)
}

func (p *PlantItem) GetRegionLevel2(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(p.RegionLevel2, value)
}

func (p *PlantItem) GetRegionLevel3(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(p.RegionLevel3, value)
}

func (p *PlantItem) GetRegionLevel4(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(p.RegionLevel4, value)
}

func (p *PlantItem) GetRegionLevel5(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(p.RegionLevel5, value)
}

func (p *PlantItem) GetRegionTimezone(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.RegionTimezone, value)
}

func (p *PlantItem) GetType(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.Type, value)
}

func (p *PlantItem) GetGridInterconnectionType(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.GridInterconnectionType, value)
}

func (p *PlantItem) GetInstalledCapacity(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.InstalledCapacity, value)
}

func (p *PlantItem) GetStartOperatingTime(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.StartOperatingTime, value)
}

func (p *PlantItem) GetStationImage(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.StationImage, value)
}

func (p *PlantItem) GetCreatedDate(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.CreatedDate, value)
}

func (p *PlantItem) GetBatterySoc(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.BatterySoc, value)
}

func (p *PlantItem) GetNetworkStatus(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.NetworkStatus, value)
}

func (p *PlantItem) GetGenerationPower(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.GenerationPower, value)
}

func (p *PlantItem) GetLastUpdateTime(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.LastUpdateTime, value)
}

type GetPlantDeviceListRequestBody struct {
	StationID  int    `json:"stationId"`
	DeviceType string `json:"deviceType,omitempty"`
	Page       int    `json:"page,omitempty"`
	Size       int    `json:"size,omitempty"`
}

type GetPlantDeviceListResponse struct {
	Code            *string            `json:"code,omitempty"`
	Message         *string            `json:"msg,omitempty"`
	Success         *bool              `json:"success,omitempty"`
	RequestID       *string            `json:"requestId,omitempty"`
	Total           *int               `json:"total,omitempty"`
	DeviceListItems []*PlantDeviceItem `json:"deviceListItems,omitempty"`
}

func (p *GetPlantDeviceListResponse) GetCode(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.Code, value)
}

func (p *GetPlantDeviceListResponse) GetMessage(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.Message, value)
}

func (p *GetPlantDeviceListResponse) GetSuccess(defaultValue ...bool) bool {
	value := false
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.BoolValue(p.Success, value)
}

func (p *GetPlantDeviceListResponse) GetRequestID(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.RequestID, value)
}

func (p *GetPlantDeviceListResponse) GetTotal(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(p.Total, value)
}

type PlantDeviceItem struct {
	DeviceSN       *string `json:"deviceSn,omitempty"`
	DeviceID       *int    `json:"deviceId,omitempty"`
	DeviceType     *string `json:"deviceType,omitempty"`
	ConnectStatus  *int    `json:"connectStatus,omitempty"`
	CollectionTime *int64  `json:"collectionTime,omitempty"`
}

func (p *PlantDeviceItem) GetDeviceSN(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.DeviceSN, value)
}

func (p *PlantDeviceItem) GetDeviceID(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(p.DeviceID, value)
}

func (p *PlantDeviceItem) GetDeviceType(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.DeviceType, value)
}

func (p *PlantDeviceItem) GetConnectStatus(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(p.ConnectStatus, value)
}

func (p *PlantDeviceItem) GetCollectionTime(defaultValue ...int64) int64 {
	value := int64(0)
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Int64Value(p.CollectionTime, value)
}

type GetDeviceAlertListRequestBody struct {
	DeviceSN       string `json:"deviceSn"`
	DeviceID       int    `json:"deviceId,omitempty"`
	StartTimestamp int64  `json:"startTimestamp"`
	EndTimestamp   int64  `json:"endTimestamp"`
	Page           int    `json:"page,omitempty"`
	Size           int    `json:"size,omitempty"`
}

type GetDeviceAlertListResponse struct {
	Code       *string            `json:"code,omitempty"`
	Message    *string            `json:"msg,omitempty"`
	Success    *bool              `json:"success,omitempty"`
	RequestID  *string            `json:"requestId,omitempty"`
	DeviceSN   *string            `json:"deviceSn,omitempty"`
	DeviceID   *int               `json:"deviceId,omitempty"`
	DeviceType *string            `json:"deviceType,omitempty"`
	Total      *int               `json:"total,omitempty"`
	AlertList  []*DeviceAlertItem `json:"alertList,omitempty"`
}

func (d *GetDeviceAlertListResponse) GetCode(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.Code, value)
}

func (d *GetDeviceAlertListResponse) GetMessage(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.Message, value)
}

func (d *GetDeviceAlertListResponse) GetSuccess(defaultValue ...bool) bool {
	value := false
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.BoolValue(d.Success, value)
}

func (d *GetDeviceAlertListResponse) GetRequestID(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.RequestID, value)
}

func (d *GetDeviceAlertListResponse) GetDeviceSN(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.DeviceSN, value)
}

func (d *GetDeviceAlertListResponse) GetDeviceID(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(d.DeviceID, value)
}

func (d *GetDeviceAlertListResponse) GetDeviceType(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.DeviceType, value)
}

func (d *GetDeviceAlertListResponse) GetTotal(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(d.Total, value)
}

type DeviceAlertItem struct {
	AlertID         *int    `json:"alertId,omitempty"`
	AlertName       *string `json:"addr,omitempty"`
	AlertNameInPAAS *string `json:"alertName,omitempty"`
	Code            *string `json:"code,omitempty"`
	Level           *int    `json:"level,omitempty"`
	Influence       *int    `json:"influence,omitempty"`
	AlertTime       *int64  `json:"alertTime,omitempty"`
}

func (d *DeviceAlertItem) GetAlertID(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(d.AlertID, value)
}

func (d *DeviceAlertItem) GetAlertName(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.AlertName, value)
}

func (d *DeviceAlertItem) GetAlertNameInPAAS(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.AlertNameInPAAS, value)
}

func (d *DeviceAlertItem) GetCode(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.Code, value)
}

func (d *DeviceAlertItem) GetLevel(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(d.Level, value)
}

func (d *DeviceAlertItem) GetInfluence(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(d.Influence, value)
}

func (d *DeviceAlertItem) GetAlertTime(defaultValue ...int64) int64 {
	value := int64(0)
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Int64Value(d.AlertTime, value)
}

type GetPlantBaseInfoResponse struct {
	Code                     *string             `json:"code,omitempty"`
	Message                  *string             `json:"msg,omitempty"`
	Success                  *bool               `json:"success,omitempty"`
	RequestID                *string             `json:"requestId,omitempty"`
	ID                       *int                `json:"id,omitempty"`
	Name                     *string             `json:"name,omitempty"`
	LocationLat              *float64            `json:"locationLat,omitempty"`
	LocationLng              *float64            `json:"locationLng,omitempty"`
	LocationAddress          *string             `json:"locationAddress,omitempty"`
	Region                   Region              `json:"region,omitempty"`
	Type                     *string             `json:"type,omitempty"`
	GridInterconnectionType  *string             `json:"gridInterconnectionType,omitempty"`
	InstalledCapacity        *float64            `json:"installedCapacity,omitempty"`
	InstallationAzimuthAngle *float64            `json:"installationAzimuthAngle,omitempty"`
	InstallationTiltAngle    *float64            `json:"installationTiltAngle,omitempty"`
	StartOperatingTime       *float64            `json:"startOperatingTime,omitempty"`
	Currency                 *string             `json:"currency,omitempty"`
	OwnerName                *string             `json:"ownerName,omitempty"`
	OwnerCompany             *string             `json:"ownerCompany,omitempty"`
	ContactPhone             *string             `json:"contactPhone,omitempty"`
	MergeElectricPrice       *float64            `json:"mergeElectricPrice,omitempty"`
	ConstructionCost         *float64            `json:"constructionCost,omitempty"`
	StationImage             *string             `json:"stationImage,omitempty"`
	StationImages            []*StationImageItem `json:"stationImages,omitempty"`
	CreatedDate              *float64            `json:"createdDate,omitempty"`
}

func (p *GetPlantBaseInfoResponse) GetCode(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.Code, value)
}

func (p *GetPlantBaseInfoResponse) GetMessage(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.Message, value)
}

func (p *GetPlantBaseInfoResponse) GetSuccess(defaultValue ...bool) bool {
	value := false
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.BoolValue(p.Success, value)
}

func (p *GetPlantBaseInfoResponse) GetRequestID(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.RequestID, value)
}

func (p *GetPlantBaseInfoResponse) GetID(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(p.ID, value)
}

func (p *GetPlantBaseInfoResponse) GetName(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.Name, value)
}

func (p *GetPlantBaseInfoResponse) GetLocationLat(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.LocationLat, value)
}

func (p *GetPlantBaseInfoResponse) GetLocationLng(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.LocationLng, value)
}

func (p *GetPlantBaseInfoResponse) GetLocationAddress(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.LocationAddress, value)
}

func (p *GetPlantBaseInfoResponse) GetType(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.Type, value)
}

func (p *GetPlantBaseInfoResponse) GetGridInterconnectionType(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.GridInterconnectionType, value)
}

func (p *GetPlantBaseInfoResponse) GetInstalledCapacity(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.InstalledCapacity, value)
}

func (p *GetPlantBaseInfoResponse) GetInstallationAzimuthAngle(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.InstallationAzimuthAngle, value)
}

func (p *GetPlantBaseInfoResponse) GetInstallationTiltAngle(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.InstallationTiltAngle, value)
}

func (p *GetPlantBaseInfoResponse) GetStartOperatingTime(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.StartOperatingTime, value)
}

func (p *GetPlantBaseInfoResponse) GetCurrency(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.Currency, value)
}

func (p *GetPlantBaseInfoResponse) GetOwnerName(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.OwnerName, value)
}

func (p *GetPlantBaseInfoResponse) GetOwnerCompany(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.OwnerCompany, value)
}

func (p *GetPlantBaseInfoResponse) GetContactPhone(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.ContactPhone, value)
}

func (p *GetPlantBaseInfoResponse) GetMergeElectricPrice(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.MergeElectricPrice, value)
}

func (p *GetPlantBaseInfoResponse) GetConstructionCost(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.ConstructionCost, value)
}

func (p *GetPlantBaseInfoResponse) GetStationImage(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(p.StationImage, value)
}

func (p *GetPlantBaseInfoResponse) GetCreatedDate(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(p.CreatedDate, value)
}

type Region struct {
	NationID *int    `json:"nationId,omitempty"`
	Level1   *int    `json:"level1,omitempty"`
	Level2   *int    `json:"level2,omitempty"`
	Level3   *int    `json:"level3,omitempty"`
	Level4   *int    `json:"level4,omitempty"`
	Level5   *int    `json:"level5,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
}

func (r *Region) GetNationID(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(r.NationID, value)
}

func (r *Region) GetLevel1(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(r.Level1, value)
}

func (r *Region) GetLevel2(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(r.Level2, value)
}

func (r *Region) GetLevel3(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(r.Level3, value)
}

func (r *Region) GetLevel4(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(r.Level4, value)
}

func (r *Region) GetLevel5(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(r.Level5, value)
}

func (r *Region) GetTimezone(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(r.Timezone, value)
}

type StationImageItem struct {
	ID          *int    `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Url         *string `json:"url,omitempty"`
}

func (i *StationImageItem) GetID(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(i.ID, value)
}

func (i *StationImageItem) GetName(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(i.Name, value)
}

func (i *StationImageItem) GetDescription(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(i.Description, value)
}

func (i *StationImageItem) GetURL(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(i.Url, value)
}

type GetRealtimePlantDataResponse struct {
	Code               *string  `json:"code,omitempty"`
	Message            *string  `json:"msg,omitempty"`
	Success            *bool    `json:"success,omitempty"`
	RequestID          *string  `json:"requestId,omitempty"`
	GenerationPower    *float64 `json:"generationPower,omitempty"`
	UsePower           *float64 `json:"usePower,omitempty"`
	GridPower          *float64 `json:"gridPower,omitempty"`
	PurchasePower      *float64 `json:"purchasePower,omitempty"`
	WirePower          *float64 `json:"wirePower,omitempty"`
	ChargePower        *float64 `json:"chargePower,omitempty"`
	DischargePower     *float64 `json:"dischargePower,omitempty"`
	BatteryPower       *float64 `json:"batteryPower,omitempty"`
	BatterySoc         *float64 `json:"batterySoc,omitempty"`
	IrradiateIntensity *float64 `json:"irradiateIntensity,omitempty"`
	LastUpdateTime     *float64 `json:"lastUpdateTime,omitempty"`
}

func (r *GetRealtimePlantDataResponse) GetCode(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(r.Code, value)
}

func (r *GetRealtimePlantDataResponse) GetMessage(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(r.Message, value)
}

func (r *GetRealtimePlantDataResponse) GetSuccess(defaultValue ...bool) bool {
	value := false
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.BoolValue(r.Success, value)
}

func (r *GetRealtimePlantDataResponse) GetRequestID(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(r.RequestID, value)
}

func (r *GetRealtimePlantDataResponse) GetGenerationPower(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.GenerationPower, value)
}

func (r *GetRealtimePlantDataResponse) GetUsePower(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.UsePower, value)
}

func (r *GetRealtimePlantDataResponse) GetGridPower(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.GridPower, value)
}

func (r *GetRealtimePlantDataResponse) GetPurchasePower(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.PurchasePower, value)
}

func (r *GetRealtimePlantDataResponse) GetWirePower(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.WirePower, value)
}

func (r *GetRealtimePlantDataResponse) GetChargePower(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.ChargePower, value)
}

func (r *GetRealtimePlantDataResponse) GetDischargePower(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.DischargePower, value)
}

func (r *GetRealtimePlantDataResponse) GetBatteryPower(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.BatteryPower, value)
}

func (r *GetRealtimePlantDataResponse) GetBatterySoc(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.BatterySoc, value)
}

func (r *GetRealtimePlantDataResponse) GetIrradiateIntensity(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.IrradiateIntensity, value)
}

func (r *GetRealtimePlantDataResponse) GetLastUpdateTime(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(r.LastUpdateTime, value)
}

type GetHistoricalPlantDataRequestBody struct {
	StationID int    `json:"stationId"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	TimeType  int    `json:"timeType"`
}

type GetHistoricalPlantDataResponse struct {
	Code             *string                      `json:"code,omitempty"`
	Message          *string                      `json:"msg,omitempty"`
	Success          *bool                        `json:"success,omitempty"`
	RequestID        *string                      `json:"requestId,omitempty"`
	Total            *int                         `json:"total,omitempty"`
	StationDataItems []*HistoricalStationDataItem `json:"stationDataItems,omitempty"`
}

func (h *GetHistoricalPlantDataResponse) GetCode(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(h.Code, value)
}

func (h *GetHistoricalPlantDataResponse) GetMessage(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(h.Message, value)
}

func (h *GetHistoricalPlantDataResponse) GetSuccess(defaultValue ...bool) bool {
	value := false
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.BoolValue(h.Success, value)
}

func (h *GetHistoricalPlantDataResponse) GetRequestID(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(h.RequestID, value)
}

func (h *GetHistoricalPlantDataResponse) GetTotal(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(h.Total, value)
}

type HistoricalStationDataItem struct {
	GenerationPower       *float64 `json:"generationPower,omitempty"`
	UsePower              *float64 `json:"usePower,omitempty"`
	GridPower             *float64 `json:"gridPower,omitempty"`
	PurchasePower         *float64 `json:"purchasePower,omitempty"`
	WirePower             *float64 `json:"wirePower,omitempty"`
	ChargePower           *float64 `json:"chargePower,omitempty"`
	DischargePower        *float64 `json:"dischargePower,omitempty"`
	BatterPower           *float64 `json:"batteryPower,omitempty"`
	BatterSoc             *float64 `json:"batterySoc,omitempty"`
	IrradiateIntensity    *float64 `json:"irradiateIntensity,omitempty"`
	GenerationValue       *float64 `json:"generationValue,omitempty"`
	GenerationRation      *float64 `json:"generationRatio,omitempty"`
	GridRatio             *float64 `json:"gridRatio,omitempty"`
	ChargeRatio           *float64 `json:"chargeRatio,omitempty"`
	UseValue              *float64 `json:"useValue,omitempty"`
	UseRatio              *float64 `json:"useRatio,omitempty"`
	BuyRatio              *float64 `json:"buyRatio,omitempty"`
	UseDischargeRatio     *float64 `json:"useDischargeRatio,omitempty"`
	GridValue             *float64 `json:"gridValue,omitempty"`
	BuyValue              *float64 `json:"buyValue,omitempty"`
	ChargeValue           *float64 `json:"chargeValue,omitempty"`
	DischargeValue        *float64 `json:"dischargeValue,omitempty"`
	FullPowerHours        *float64 `json:"fullPowerHours,omitempty"`
	Irradiate             *float64 `json:"irradiate,omitempty"`
	TheoreticalGeneration *float64 `json:"theoreticalGeneration,omitempty"`
	PR                    *float64 `json:"pr,omitempty"`
	CPR                   *float64 `json:"cpr,omitempty"`
	DateTime              *float64 `json:"dateTime,omitempty"`
	Year                  *int     `json:"year,omitempty"`
	Month                 *int     `json:"month,omitempty"`
	Day                   *int     `json:"day,omitempty"`
}

func (s *HistoricalStationDataItem) GetGenerationPower(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(s.GenerationPower, value)
}

func (s *HistoricalStationDataItem) GetUsePower(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(s.UsePower, value)
}

func (s *HistoricalStationDataItem) GetGridPower(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(s.GridPower, value)
}

func (s *HistoricalStationDataItem) GetPurchasePower(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(s.PurchasePower, value)
}

func (s *HistoricalStationDataItem) GetWirePower(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(s.WirePower, value)
}

func (s *HistoricalStationDataItem) GetChargePower(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(s.ChargePower, value)
}

func (s *HistoricalStationDataItem) GetDischargePower(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(s.DischargePower, value)
}

func (s *HistoricalStationDataItem) GetBatterPower(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(s.BatterPower, value)
}

func (s *HistoricalStationDataItem) GetBatterSoc(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(s.BatterSoc, value)
}

func (s *HistoricalStationDataItem) GetIrradiateIntensity(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(s.IrradiateIntensity, value)
}

func (s *HistoricalStationDataItem) GetGenerationValue(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(s.GenerationValue, value)
}

func (s *HistoricalStationDataItem) GetGenerationRation(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(s.GenerationRation, value)
}

func (s *HistoricalStationDataItem) GetGridRatio(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(s.GridRatio, value)
}

func (s *HistoricalStationDataItem) GetChargeRatio(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(s.ChargeRatio, value)
}

func (s *HistoricalStationDataItem) GetUseValue(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(s.UseValue, value)
}

func (s *HistoricalStationDataItem) GetUseRatio(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(s.UseRatio, value)
}

func (s *HistoricalStationDataItem) GetBuyRatio(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(s.BuyRatio, value)
}

func (s *HistoricalStationDataItem) GetUseDischargeRatio(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(s.UseDischargeRatio, value)
}

func (s *HistoricalStationDataItem) GetGridValue(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(s.GridValue, value)
}

func (s *HistoricalStationDataItem) GetBuyValue(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(s.BuyValue, value)
}

func (s *HistoricalStationDataItem) GetChargeValue(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(s.ChargeValue, value)
}

func (s *HistoricalStationDataItem) GetDischargeValue(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(s.DischargeValue, value)
}

func (s *HistoricalStationDataItem) GetFullPowerHours(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(s.FullPowerHours, value)
}

func (s *HistoricalStationDataItem) GetIrradiate(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(s.Irradiate, value)
}

func (s *HistoricalStationDataItem) GetTheoreticalGeneration(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(s.TheoreticalGeneration, value)
}

func (s *HistoricalStationDataItem) GetPR(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(s.PR, value)
}

func (s *HistoricalStationDataItem) GetCPR(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(s.CPR, value)
}

func (s *HistoricalStationDataItem) GetDateTime(defaultValue ...float64) float64 {
	value := 0.0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.Float64Value(s.DateTime, value)
}

func (s *HistoricalStationDataItem) GetYear(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(s.Year, value)
}

func (s *HistoricalStationDataItem) GetMonth(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(s.Month, value)
}

func (s *HistoricalStationDataItem) GetDay(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(s.Day, value)
}

type GetRealtimeDeviceDataResponse struct {
	Code        *string     `json:"code,omitempty"`
	Message     *string     `json:"msg,omitempty"`
	Success     *bool       `json:"success,omitempty"`
	RequestID   *string     `json:"requestId,omitempty"`
	DeviceSN    *string     `json:"deviceSn,omitempty"`
	DeviceID    *int        `json:"deviceId,omitempty"`
	DeviceType  *string     `json:"deviceType,omitempty"`
	DeviceState *int        `json:"deviceState,omitempty"`
	DataList    []*DataItem `json:"dataList,omitempty"`
}

func (r *GetRealtimeDeviceDataResponse) GetCode(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(r.Code, value)
}

func (r *GetRealtimeDeviceDataResponse) GetMessage(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(r.Message, value)
}

func (r *GetRealtimeDeviceDataResponse) GetSuccess(defaultValue ...bool) bool {
	value := false
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.BoolValue(r.Success, value)
}

func (r *GetRealtimeDeviceDataResponse) GetRequestID(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(r.RequestID, value)
}

func (r *GetRealtimeDeviceDataResponse) GetDeviceSN(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(r.DeviceSN, value)
}

func (r *GetRealtimeDeviceDataResponse) GetDeviceID(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(r.DeviceID, value)
}

func (r *GetRealtimeDeviceDataResponse) GetDeviceType(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(r.DeviceType, value)
}

func (r *GetRealtimeDeviceDataResponse) GetDeviceState(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(r.DeviceState, value)
}

type GetHistoricalDeviceDataRequestBody struct {
	DeviceSN  string `json:"deviceSn"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	TimeType  int    `json:"timeType"`
}

type GetHistoricalDeviceDataResponse struct {
	Code          *string          `json:"code,omitempty"`
	Message       *string          `json:"msg,omitempty"`
	Success       *bool            `json:"success,omitempty"`
	RequestID     *string          `json:"requestId,omitempty"`
	DeviceSN      *string          `json:"deviceSn,omitempty"`
	DeviceID      *int             `json:"deviceId,omitempty"`
	DeviceType    *string          `json:"deviceType,omitempty"`
	TimeType      *int             `json:"timeType,omitempty"`
	ParamDataList []*ParamDataItem `json:"paramDataList,omitempty"`
}

func (h *GetHistoricalDeviceDataResponse) GetCode(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(h.Code, value)
}

func (h *GetHistoricalDeviceDataResponse) GetMessage(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(h.Message, value)
}

func (h *GetHistoricalDeviceDataResponse) GetSuccess(defaultValue ...bool) bool {
	value := false
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.BoolValue(h.Success, value)
}

func (h *GetHistoricalDeviceDataResponse) GetRequestID(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(h.RequestID, value)
}

func (h *GetHistoricalDeviceDataResponse) GetDeviceSN(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(h.DeviceSN, value)
}

func (h *GetHistoricalDeviceDataResponse) GetDeviceID(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(h.DeviceID, value)
}

func (h *GetHistoricalDeviceDataResponse) GetDeviceType(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(h.DeviceType, value)
}

func (h *GetHistoricalDeviceDataResponse) GetTimeType(defaultValue ...int) int {
	value := 0
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.IntValue(h.TimeType, value)
}

type ParamDataItem struct {
	CollectTime *string     `json:"collectTime,omitempty"`
	DataList    []*DataItem `json:"dataList,omitempty"`
}

func (d *ParamDataItem) GetCollectTime(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.CollectTime, value)
}

type DataItem struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
	Unit  *string `json:"unit,omitempty"`
	Name  *string `json:"name,omitempty"`
}

func (d *DataItem) GetKey(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.Key, value)
}

func (d *DataItem) GetValue(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.Value, value)
}

func (d *DataItem) GetUnit(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.Unit, value)
}

func (d *DataItem) GetName(defaultValue ...string) string {
	value := ""
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return pointy.StringValue(d.Name, value)
}

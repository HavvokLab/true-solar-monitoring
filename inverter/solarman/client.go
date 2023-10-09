package solarman

import (
	"errors"
	"net/http"
)

type SolarmanClient interface {
	GetBasicToken() (*GetTokenResponse, error)
	GetBusinessToken(orgID int) (*GetTokenResponse, error)
	GetUserInfo() (*GetUserInfoResponse, error)
	GetPlantListWithPagination(businessToken string, page, size int) (*GetPlantListResponse, error)
	GetPlantList(businessToken string) ([]*PlantItem, error)
	GetPlantBaseInfo(businessToken string, stationID int) (*GetPlantBaseInfoResponse, error)
	GetPlantRealtimeData(businessToken string, stationID int) (*GetRealtimePlantDataResponse, error)
	GetHistoricalPlantData(businessToken string, stationID int, timeType int, from, to int64) (*GetHistoricalPlantDataResponse, error)
	GetPlantDeviceListWithPagination(businessToken string, stationID, page, size int) (*GetPlantDeviceListResponse, error)
	GetPlantDeviceList(businessToken string, stationID int) ([]*PlantDeviceItem, error)
	GetDeviceRealtimeData(businessToken, deviceSN string) (*GetRealtimeDeviceDataResponse, error)
	GetHistoricalDeviceData(businessToken, deviceSN string, timeType int, from, to int64) (*GetHistoricalDeviceDataResponse, error)
	GetDeviceAlertListWithPagination(businessToken, deviceSN string, from, to int64, page, size int) (*GetDeviceAlertListResponse, error)
	GetDeviceAlertList(businessToken, deviceSN string, from, to int64) ([]*DeviceAlertItem, error)
}

type SolarmanCredential struct {
	Username  string
	Password  string
	AppID     string
	AppSecret string
}

type solarmanClient struct {
	URL        string
	credential *SolarmanCredential
	headers    map[string]string
}

func NewSolarmanClient(credential *SolarmanCredential) (SolarmanClient, error) {
	if credential == nil {
		return nil, errors.New("credential must not be empty")
	}
	credential.Password = DecodePassword(credential.Password)

	client := &solarmanClient{
		URL:        URL_VERSION1,
		credential: credential,
		headers:    make(map[string]string),
	}

	resp, err := client.GetBasicToken()
	if err != nil {
		return nil, err
	}

	if resp.AccessToken == nil {
		return nil, errors.New("access token must not be empty")
	}

	client.headers[AUTHORIZATION_HEADER] = buildAuthorizationToken(resp.GetAccessToken())
	return client, nil
}

func (r *solarmanClient) GetBasicToken() (*GetTokenResponse, error) {
	body := GetTokenRequestBody{
		Username:  r.credential.Username,
		Password:  r.credential.Password,
		AppSecret: r.credential.AppSecret,
	}

	url := r.URL + "/account/v1.0/token?appId=" + r.credential.AppID
	req, err := prepareHttpRequest(http.MethodPost, url, nil, body)
	if err != nil {
		return nil, err
	}

	data, status, err := prepareHttpResponse[GetTokenResponse](req)
	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		return nil, errors.New(data.GetMessage())
	}

	return data, nil
}

func (r *solarmanClient) GetBusinessToken(orgID int) (*GetTokenResponse, error) {
	body := GetTokenRequestBody{
		Username:  r.credential.Username,
		Password:  r.credential.Password,
		AppSecret: r.credential.AppSecret,
		OrgID:     orgID,
	}

	url := r.URL + "/account/v1.0/token?appId=" + r.credential.AppID
	req, err := prepareHttpRequest(http.MethodPost, url, nil, body)
	if err != nil {
		return nil, err
	}

	data, status, err := prepareHttpResponse[GetTokenResponse](req)
	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		return nil, errors.New(data.GetMessage())
	}

	return data, nil
}

func (r *solarmanClient) GetUserInfo() (*GetUserInfoResponse, error) {
	url := r.URL + "/account/v1.0/info?language=en"
	req, err := prepareHttpRequest(http.MethodPost, url, r.headers, nil)
	if err != nil {
		return nil, err
	}

	data, status, err := prepareHttpResponse[GetUserInfoResponse](req)
	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		return nil, errors.New(data.GetMessage())
	}

	return data, nil
}

func (r *solarmanClient) GetPlantListWithPagination(businessToken string, page, size int) (*GetPlantListResponse, error) {
	body := map[string]interface{}{
		"page": page,
		"size": size,
	}

	headers := map[string]string{AUTHORIZATION_HEADER: buildAuthorizationToken(businessToken)}
	url := r.URL + "/station/v1.0/list?language=en"
	req, err := prepareHttpRequest(http.MethodPost, url, headers, body)
	if err != nil {
		return nil, err
	}

	data, status, err := prepareHttpResponse[GetPlantListResponse](req)
	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		return nil, errors.New(data.GetMessage())
	}

	return data, nil
}

func (r *solarmanClient) GetPlantList(businessToken string) ([]*PlantItem, error) {
	data := make([]*PlantItem, 0)
	page := 1

	for {
		res, err := r.GetPlantListWithPagination(businessToken, page, MAX_PAGE_SIZE)
		if err != nil {
			return nil, err
		}

		data = append(data, res.StationList...)
		page += 1

		if len(data) >= res.GetTotal() {
			break
		}
	}

	return data, nil
}

func (r *solarmanClient) GetPlantBaseInfo(businessToken string, stationID int) (*GetPlantBaseInfoResponse, error) {
	body := StationRequestBody{StationID: stationID}
	headers := map[string]string{AUTHORIZATION_HEADER: buildAuthorizationToken(businessToken)}

	url := r.URL + "/station/v1.0/base?language=en"
	req, err := prepareHttpRequest(http.MethodPost, url, headers, body)
	if err != nil {
		return nil, err
	}

	data, status, err := prepareHttpResponse[GetPlantBaseInfoResponse](req)
	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		return nil, errors.New(data.GetMessage())
	}

	return data, nil
}

func (r *solarmanClient) GetPlantRealtimeData(businessToken string, stationID int) (*GetRealtimePlantDataResponse, error) {
	url := r.URL + "/station/v1.0/realTime?language=en"
	body := StationRequestBody{StationID: stationID}
	headers := map[string]string{AUTHORIZATION_HEADER: buildAuthorizationToken(businessToken)}

	req, err := prepareHttpRequest(http.MethodPost, url, headers, body)
	if err != nil {
		return nil, err
	}

	data, status, err := prepareHttpResponse[GetRealtimePlantDataResponse](req)
	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		return nil, errors.New(data.GetMessage())
	}

	return data, nil
}

func (r *solarmanClient) GetHistoricalPlantData(businessToken string, stationID int, timeType int, from, to int64) (*GetHistoricalPlantDataResponse, error) {
	url := r.URL + "/station/v1.0/history?language=en"
	startTime := buildDateFromTimestamp(from, timeType)
	endTime := buildDateFromTimestamp(to, timeType)
	headers := map[string]string{AUTHORIZATION_HEADER: buildAuthorizationToken(businessToken)}
	body := GetHistoricalPlantDataRequestBody{
		StationID: stationID,
		StartTime: startTime,
		EndTime:   endTime,
		TimeType:  timeType,
	}

	req, err := prepareHttpRequest(http.MethodPost, url, headers, body)
	if err != nil {
		return nil, err
	}

	data, status, err := prepareHttpResponse[GetHistoricalPlantDataResponse](req)
	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		return nil, errors.New(data.GetMessage())
	}

	return data, nil
}

func (r *solarmanClient) GetPlantDeviceListWithPagination(businessToken string, stationID, page, size int) (*GetPlantDeviceListResponse, error) {
	url := r.URL + "/station/v1.0/device?language=en"
	headers := map[string]string{AUTHORIZATION_HEADER: buildAuthorizationToken(businessToken)}
	body := GetPlantDeviceListRequestBody{
		StationID: stationID,
		Page:      page,
		Size:      size,
	}

	req, err := prepareHttpRequest(http.MethodPost, url, headers, body)
	if err != nil {
		return nil, err
	}

	data, status, err := prepareHttpResponse[GetPlantDeviceListResponse](req)
	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		return nil, errors.New(data.GetMessage())
	}

	return data, nil
}

func (r *solarmanClient) GetPlantDeviceList(businessToken string, stationID int) ([]*PlantDeviceItem, error) {
	data := make([]*PlantDeviceItem, 0)
	page := 1

	for {
		res, err := r.GetPlantDeviceListWithPagination(businessToken, stationID, page, MAX_PAGE_SIZE)
		if err != nil {
			return nil, err
		}

		data = append(data, res.DeviceListItems...)
		page += 1

		if len(data) >= res.GetTotal() {
			break
		}
	}

	return data, nil
}

func (r *solarmanClient) GetDeviceRealtimeData(businessToken, deviceSN string) (*GetRealtimeDeviceDataResponse, error) {
	url := r.URL + "/device/v1.0/currentData?language=en"
	body := DeviceRequestBody{DeviceSN: deviceSN}
	headers := map[string]string{AUTHORIZATION_HEADER: buildAuthorizationToken(businessToken)}

	req, err := prepareHttpRequest(http.MethodPost, url, headers, body)
	if err != nil {
		return nil, err
	}

	data, status, err := prepareHttpResponse[GetRealtimeDeviceDataResponse](req)
	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		return nil, errors.New(data.GetMessage())
	}

	return data, nil
}

func (r *solarmanClient) GetHistoricalDeviceData(businessToken, deviceSN string, timeType int, from, to int64) (*GetHistoricalDeviceDataResponse, error) {
	url := r.URL + "/device/v1.0/historical?language=en"
	startTime := buildDateFromTimestamp(from, timeType)
	endTime := buildDateFromTimestamp(to, timeType)
	headers := map[string]string{AUTHORIZATION_HEADER: buildAuthorizationToken(businessToken)}
	body := GetHistoricalDeviceDataRequestBody{
		DeviceSN:  deviceSN,
		StartTime: startTime,
		EndTime:   endTime,
		TimeType:  timeType,
	}

	req, err := prepareHttpRequest(http.MethodPost, url, headers, body)
	if err != nil {
		return nil, err
	}

	data, status, err := prepareHttpResponse[GetHistoricalDeviceDataResponse](req)
	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		return nil, errors.New(data.GetMessage())
	}

	return data, nil
}

func (r *solarmanClient) GetDeviceAlertListWithPagination(businessToken, deviceSN string, from, to int64, page, size int) (*GetDeviceAlertListResponse, error) {
	url := r.URL + "/device/v1.0/alertList?language=en"
	headers := map[string]string{AUTHORIZATION_HEADER: buildAuthorizationToken(businessToken)}
	body := GetDeviceAlertListRequestBody{
		DeviceSN:       deviceSN,
		StartTimestamp: from,
		EndTimestamp:   to,
		Page:           page,
		Size:           size,
	}

	req, err := prepareHttpRequest(http.MethodPost, url, headers, body)
	if err != nil {
		return nil, err
	}

	data, status, err := prepareHttpResponse[GetDeviceAlertListResponse](req)
	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		return nil, errors.New(data.GetMessage())
	}

	return data, nil
}

func (r *solarmanClient) GetDeviceAlertList(businessToken, deviceSN string, from, to int64) ([]*DeviceAlertItem, error) {
	data := make([]*DeviceAlertItem, 0)
	page := 1

	for {
		res, err := r.GetDeviceAlertListWithPagination(businessToken, deviceSN, from, to, page, MAX_PAGE_SIZE)
		if err != nil {
			return nil, err
		}

		data = append(data, res.AlertList...)
		page += 1

		if len(data) >= res.GetTotal() {
			break
		}
	}

	return data, nil
}

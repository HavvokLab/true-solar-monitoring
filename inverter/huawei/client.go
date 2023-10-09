package huawei

import (
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/HavvokLab/true-solar-monitoring/util"
	"github.com/avast/retry-go"
)

type HuaweiClient interface {
	GetToken(username, password string) (string, error)
	GetPlantList() (*GetPlantListResponse, error)
	GetRealtimePlantData(plantCode string) (*GetRealtimePlantDataResponse, error)
	GetDailyPlantData(plantCode string, timestamp int64) (*GetHistoricalPlantDataResponse, error)
	GetMonthlyPlantData(plantCode string, timestamp int64) (*GetHistoricalPlantDataResponse, error)
	GetYearlyPlantData(plantCode string, timestamp int64) (*GetHistoricalPlantDataResponse, error)
	GetDeviceList(plantCode string) (*GetDeviceListResponse, error)
	GetRealtimeDeviceData(deviceID, deviceTypeID string) (*GetRealtimeDeviceDataResponse, error)
	GetDailyDeviceData(deviceID, deviceTypeID string, timestamp int64) (*GetHistoricalDeviceDataResponse, error)
	GetMonthlyDeviceData(deviceID, deviceTypeID string, timestamp int64) (*GetHistoricalDeviceDataResponse, error)
	GetYearlyDeviceData(deviceID, deviceTypeID string, timestamp int64) (*GetHistoricalDeviceDataResponse, error)
	GetDeviceAlarmList(plantCode string, from, to int64) (*GetDeviceAlarmListResponse, error)
}

type HuaweiCredential struct {
	Username string
	Password string
}

type huaweiClient struct {
	URL     string
	headers map[string]string
}

func NewHuaweiClient(credential *HuaweiCredential) (HuaweiClient, error) {
	client := &huaweiClient{
		URL:     URL_VERSION1,
		headers: make(map[string]string),
	}

	token, err := client.GetToken(credential.Username, credential.Password)
	if err != nil {
		return nil, err
	}

	client.headers[AUTH_HEADER] = token
	return client, nil
}

func (h *huaweiClient) GetToken(username, password string) (string, error) {
	url := h.URL + "/thirdData/login"
	data := map[string]interface{}{
		"userName":   username,
		"systemCode": password,
	}

	req, err := prepareHttpRequest(http.MethodPost, url, nil, data)
	if err != nil {
		return "", err
	}

	retryOptions := []retry.Option{
		retry.Delay(RETRY_WAIT_TIME),
		retry.Attempts(RETRY_ATTEMPT),
		retry.DelayType(retry.FixedDelay),
	}

	var token string
	retry.Do(func() error {
		client := &http.Client{Timeout: 300 * time.Second}
		res, err := client.Do(req)
		if err != nil {
			return err
		}

		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		var result GetTokenResponse
		if err := util.Recast(body, &result); err != nil {
			return err
		}

		if result.Success != nil {
			if !result.GetSuccess() {
				if result.GetFailCode() == 407 {
					return err
				} else {
					return errors.New(HuaweiMapErrorMessage[result.GetFailCode()])
				}
			} else {
				mapCookies := make(map[string]string)
				for _, c := range res.Cookies() {
					mapCookies[c.Name] = c.Value
				}
				token = mapCookies["XSRF-TOKEN"]
			}
		}

		return nil
	}, retryOptions...)

	return token, nil
}

func (h *huaweiClient) GetPlantList() (*GetPlantListResponse, error) {
	url := h.URL + "/thirdData/getStationList"
	req, err := prepareHttpRequest(http.MethodPost, url, h.headers, nil)
	if err != nil {
		return nil, err
	}

	res, _, err := prepareHttpResponse[GetPlantListResponse](req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *huaweiClient) GetRealtimePlantData(plantCode string) (*GetRealtimePlantDataResponse, error) {
	url := h.URL + "/thirdData/getStationRealKpi"
	data := map[string]interface{}{
		"stationCodes": plantCode,
	}
	req, err := prepareHttpRequest(http.MethodPost, url, h.headers, data)
	if err != nil {
		return nil, err
	}

	res, _, err := prepareHttpResponse[GetRealtimePlantDataResponse](req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *huaweiClient) GetDailyPlantData(plantCode string, timestamp int64) (*GetHistoricalPlantDataResponse, error) {
	url := h.URL + "/thirdData/getKpiStationDay"
	data := map[string]interface{}{
		"stationCodes": plantCode,
		"collectTime":  timestamp,
	}

	req, err := prepareHttpRequest(http.MethodPost, url, h.headers, data)
	if err != nil {
		return nil, err
	}

	res, _, err := prepareHttpResponse[GetHistoricalPlantDataResponse](req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *huaweiClient) GetMonthlyPlantData(plantCode string, timestamp int64) (*GetHistoricalPlantDataResponse, error) {
	url := h.URL + "/thirdData/getKpiStationMonth"
	data := map[string]interface{}{
		"stationCodes": plantCode,
		"collectTime":  timestamp,
	}

	req, err := prepareHttpRequest(http.MethodPost, url, h.headers, data)
	if err != nil {
		return nil, err
	}

	res, _, err := prepareHttpResponse[GetHistoricalPlantDataResponse](req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *huaweiClient) GetYearlyPlantData(plantCode string, timestamp int64) (*GetHistoricalPlantDataResponse, error) {
	url := h.URL + "/thirdData/getKpiStationYear"
	data := map[string]interface{}{
		"stationCodes": plantCode,
		"collectTime":  timestamp,
	}

	req, err := prepareHttpRequest(http.MethodPost, url, h.headers, data)
	if err != nil {
		return nil, err
	}

	res, _, err := prepareHttpResponse[GetHistoricalPlantDataResponse](req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *huaweiClient) GetDeviceList(plantCode string) (*GetDeviceListResponse, error) {
	url := h.URL + "/thirdData/getDevList"
	data := map[string]interface{}{
		"stationCodes": plantCode,
	}

	req, err := prepareHttpRequest(http.MethodPost, url, h.headers, data)
	if err != nil {
		return nil, err
	}

	res, _, err := prepareHttpResponse[GetDeviceListResponse](req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *huaweiClient) GetRealtimeDeviceData(deviceID, deviceTypeID string) (*GetRealtimeDeviceDataResponse, error) {
	url := h.URL + "/thirdData/getDevRealKpi"
	data := map[string]interface{}{
		"devIds":    deviceID,
		"devTypeId": deviceTypeID,
	}

	req, err := prepareHttpRequest(http.MethodPost, url, h.headers, data)
	if err != nil {
		return nil, err
	}

	res, _, err := prepareHttpResponse[GetRealtimeDeviceDataResponse](req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *huaweiClient) GetDailyDeviceData(deviceID, deviceTypeID string, timestamp int64) (*GetHistoricalDeviceDataResponse, error) {
	url := h.URL + "/thirdData/getDevKpiDay"
	data := map[string]interface{}{
		"devIds":      deviceID,
		"devTypeId":   deviceTypeID,
		"collectTime": timestamp,
	}

	req, err := prepareHttpRequest(http.MethodPost, url, h.headers, data)
	if err != nil {
		return nil, err
	}

	res, _, err := prepareHttpResponse[GetHistoricalDeviceDataResponse](req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *huaweiClient) GetMonthlyDeviceData(deviceID, deviceTypeID string, timestamp int64) (*GetHistoricalDeviceDataResponse, error) {
	url := h.URL + "/thirdData/getDevKpiMonth"
	data := map[string]interface{}{
		"devIds":      deviceID,
		"devTypeId":   deviceTypeID,
		"collectTime": timestamp,
	}

	req, err := prepareHttpRequest(http.MethodPost, url, h.headers, data)
	if err != nil {
		return nil, err
	}

	res, _, err := prepareHttpResponse[GetHistoricalDeviceDataResponse](req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *huaweiClient) GetYearlyDeviceData(deviceID, deviceTypeID string, timestamp int64) (*GetHistoricalDeviceDataResponse, error) {
	url := h.URL + "/thirdData/getDevKpiYear"
	data := map[string]interface{}{
		"devIds":      deviceID,
		"devTypeId":   deviceTypeID,
		"collectTime": timestamp,
	}

	req, err := prepareHttpRequest(http.MethodPost, url, h.headers, data)
	if err != nil {
		return nil, err
	}

	res, _, err := prepareHttpResponse[GetHistoricalDeviceDataResponse](req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *huaweiClient) GetDeviceAlarmList(plantCode string, from, to int64) (*GetDeviceAlarmListResponse, error) {
	url := h.URL + "/thirdData/getAlarmList"
	data := map[string]interface{}{
		"stationCodes": plantCode,
		"beginTime":    from,
		"endTime":      to,
		"language":     HUAWEI_LANG_ENGLISH,
	}

	req, err := prepareHttpRequest(http.MethodPost, url, h.headers, data)
	if err != nil {
		return nil, err
	}

	res, _, err := prepareHttpResponse[GetDeviceAlarmListResponse](req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

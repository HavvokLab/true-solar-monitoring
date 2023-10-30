package growatt

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"dario.cat/mergo"
	"go.openly.dev/pointy"
)

type GrowattClient interface {
	GetPlantList() ([]*PlantItem, error)
	GetPlantListWithPagination(page, size int) (*GetPlantListResponse, error)
	GetPlantOverviewInfo(plantID int) (*GetPlantOverviewInfoResponse, error)
	GetPlantDataLoggerInfo(plantID int) (*GetPlantDataLoggerInfoResponse, error)
	GetPlantDeviceList(plantID int) ([]*DeviceItem, error)
	GetPlantDeviceListWithPagination(plantID, page, size int) (*GetPlantDeviceListResponse, error)
	GetRealtimeDeviceData(deviceSN string) (*GetRealtimeDeviceDataResponse, error)
	GetRealtimeDeviceBatchesData(deviceSNs []string) (*GetRealtimeDeviceBatchesDataResponse, error)
	GetRealtimeDeviceBatchesDataWithPagination(deviceSNs []string, page int) (*GetRealtimeDeviceBatchesDataResponse, error)
	GetInverterAlertList(deviceSN string) ([]*AlarmItem, error)
	GetInverterAlertListWithPagination(deviceSN string, page, size int) (*GetInverterAlertListResponse, error)
	GetEnergyStorageMachineAlertList(deviceSN string, timestamp int64) (*GetEnergyStorageMachineAlertListResponse, error)
	GetMaxAlertList(deviceSN string, timestamp int64) ([]*AlarmItem, error)
	GetMaxAlertListWithPagination(deviceSN string, timestamp int64, page, size int) (*GetMaxAlertListResponse, error)
	GetMixAlertList(deviceSN string, timestamp int64) ([]*AlarmItem, error)
	GetMixAlertListWithPagination(deviceSN string, timestamp int64, page, size int) (*GetMixAlertListResponse, error)
	GetSpaAlertList(deviceSN string, timestamp int64) ([]*AlarmItem, error)
	GetSpaAlertListWithPagination(deviceSN string, timestamp int64, page, size int) (*GetSpaAlertListResponse, error)
	GetMinAlertList(deviceSN string, timestamp int64) ([]*AlarmItem, error)
	GetMinAlertListWithPagination(deviceSN string, timestamp int64, page, size int) (*GetMinAlertListResponse, error)
	GetPcsAlertList(deviceSN string, timestamp int64) ([]*AlarmItem, error)
	GetPcsAlertListWithPagination(deviceSN string, timestamp int64, page, size int) (*GetPcsAlertListResponse, error)
	GetHpsAlertList(deviceSN string, timestamp int64) ([]*AlarmItem, error)
	GetHpsAlertListWithPagination(deviceSN string, timestamp int64, page, size int) (*GetHpsAlertListResponse, error)
	GetPbdAlertList(deviceSN string, timestamp int64) ([]*AlarmItem, error)
	GetPbdAlertListWithPagination(deviceSN string, timestamp int64, page, size int) (*GetPbdAlertListResponse, error)
}

type growattClient struct {
	URL      string
	username string
	token    string
	headers  map[string]string
}

type GrowattCredential struct {
	Username string
	Token    string
}

func NewGrowattClient(credential *GrowattCredential) (GrowattClient, error) {
	if credential == nil {
		return nil, errors.New("credential must not be nil")
	}

	client := &growattClient{
		URL:      URL_VERSION1,
		username: credential.Username,
		token:    credential.Token,
		headers:  map[string]string{AUTH_HEADER: credential.Token},
	}

	return client, nil
}

func (r *growattClient) GetPlantListWithPagination(page, size int) (*GetPlantListResponse, error) {
	queryMap := map[string]interface{}{
		"user_name": r.username,
		"page":      page,
		"perpage":   size,
	}

	query := BuildQueryParams(queryMap)
	url := r.URL + "/plant/user_plant_list" + query

	req, err := prepareHttpRequest(http.MethodPost, url, r.headers, nil)
	if err != nil {
		return nil, err
	}

	res, _, err := prepareHttpResponse[GetPlantListResponse](req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *growattClient) GetPlantList() ([]*PlantItem, error) {
	plants := make([]*PlantItem, 0)
	page := 1
	for {
		res, err := r.GetPlantListWithPagination(page, MAX_PAGE_SIZE)
		if err != nil {
			return nil, err
		}

		plants = append(plants, res.Data.Plants...)

		if len(plants) >= res.Data.GetCount() {
			break
		}

		page += 1
	}

	return plants, nil
}

func (r *growattClient) GetPlantOverviewInfo(plantID int) (*GetPlantOverviewInfoResponse, error) {
	queryMap := map[string]interface{}{
		"plant_id": plantID,
	}

	query := BuildQueryParams(queryMap)
	url := r.URL + "/plant/data" + query

	req, err := prepareHttpRequest(http.MethodGet, url, r.headers, nil)
	if err != nil {
		return nil, err
	}

	res, _, err := prepareHttpResponse[GetPlantOverviewInfoResponse](req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *growattClient) GetPlantDataLoggerInfo(plantID int) (*GetPlantDataLoggerInfoResponse, error) {
	queryMap := map[string]interface{}{
		"plant_id": plantID,
	}

	query := BuildQueryParams(queryMap)
	url := r.URL + "/device/datalogger/list" + query

	req, err := prepareHttpRequest(http.MethodGet, url, r.headers, nil)
	if err != nil {
		return nil, err
	}

	res, _, err := prepareHttpResponse[GetPlantDataLoggerInfoResponse](req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *growattClient) GetPlantDeviceListWithPagination(plantID, page, size int) (*GetPlantDeviceListResponse, error) {
	queryMap := map[string]interface{}{
		"plant_id": plantID,
		"page":     page,
		"perpage":  size,
	}

	query := BuildQueryParams(queryMap)
	url := r.URL + "/device/list" + query
	req, err := prepareHttpRequest(http.MethodGet, url, r.headers, nil)
	if err != nil {
		return nil, err
	}

	res, _, err := prepareHttpResponse[GetPlantDeviceListResponse](req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *growattClient) GetPlantDeviceList(plantID int) ([]*DeviceItem, error) {
	devices := make([]*DeviceItem, 0)
	page := 1
	for {
		res, err := r.GetPlantDeviceListWithPagination(plantID, page, MAX_PAGE_SIZE)
		if err != nil {
			return nil, err
		}

		devices = append(devices, res.Data.Devices...)

		if len(devices) >= res.Data.GetCount() {
			break
		}

		page += 1
	}

	return devices, nil
}

func (r *growattClient) GetRealtimeDeviceData(deviceSN string) (*GetRealtimeDeviceDataResponse, error) {
	queryMap := map[string]interface{}{
		"device_sn": deviceSN,
	}

	query := BuildQueryParams(queryMap)
	url := r.URL + "/device/inverter/last_new_data" + query

	req, err := prepareHttpRequest(http.MethodGet, url, r.headers, nil)
	if err != nil {
		return nil, err
	}

	res, _, err := prepareHttpResponse[GetRealtimeDeviceDataResponse](req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *growattClient) GetRealtimeDeviceBatchesDataWithPagination(deviceSNs []string, page int) (*GetRealtimeDeviceBatchesDataResponse, error) {
	queryMap := map[string]interface{}{
		"inverters": strings.Join(deviceSNs, ","),
		"pageNum":   page,
	}

	query := BuildQueryParams(queryMap)
	url := r.URL + "/device/inverter/invs_data" + query

	req, err := prepareHttpRequest(http.MethodPost, url, r.headers, nil)
	if err != nil {
		return nil, err
	}

	res, _, err := prepareHttpResponse[GetRealtimeDeviceBatchesDataResponse](req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *growattClient) GetRealtimeDeviceBatchesData(deviceSNs []string) (*GetRealtimeDeviceBatchesDataResponse, error) {
	batches := make([][]string, 0)
	var j int
	for i := 0; i < len(deviceSNs); i += HALF_A_BATCH_SIZE {
		j += HALF_A_BATCH_SIZE
		if j > len(deviceSNs) {
			j = len(deviceSNs)
		}

		batches = append(batches, deviceSNs[i:j])
	}

	result := GetRealtimeDeviceBatchesDataResponse{
		Inverters: make([]*string, 0),
		Data:      make(map[string]map[string]interface{}, 0),
		PageNum:   pointy.Int(1),
	}

	for _, batch := range batches {
		resp, err := r.GetRealtimeDeviceBatchesDataWithPagination(batch, 1)
		if err != nil {
			return nil, err
		}

		err = mergo.Merge(&result.Data, resp.Data)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func (r *growattClient) GetInverterAlertListWithPagination(deviceSN string, page, size int) (*GetInverterAlertListResponse, error) {
	queryMap := map[string]interface{}{
		"device_sn": deviceSN,
		"page":      page,
		"perpage":   size,
	}

	query := BuildQueryParams(queryMap)
	url := r.URL + "/device/inverter/alarm" + query
	req, err := prepareHttpRequest(http.MethodGet, url, r.headers, nil)
	if err != nil {
		return nil, err
	}

	res, _, err := prepareHttpResponse[GetInverterAlertListResponse](req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *growattClient) GetInverterAlertList(deviceSN string) ([]*AlarmItem, error) {
	alarms := make([]*AlarmItem, 0)
	page := 1
	for {
		res, err := r.GetInverterAlertListWithPagination(deviceSN, page, MAX_PAGE_SIZE)
		if err != nil {
			return nil, err
		}

		alarms = append(alarms, res.Data.Alarms...)

		if len(alarms) >= res.Data.GetCount() {
			break
		}

		page += 1
	}

	return alarms, nil
}

func (r *growattClient) GetEnergyStorageMachineAlertList(deviceSN string, timestamp int64) (*GetEnergyStorageMachineAlertListResponse, error) {
	queryMap := map[string]interface{}{
		"storage_sn": deviceSN,
		"date":       time.Unix(timestamp, 0).Format("2006-01-02"),
	}

	query := BuildQueryParams(queryMap)
	url := r.URL + "/device/storage/alarm_date" + query
	req, err := prepareHttpRequest(http.MethodPost, url, r.headers, nil)
	if err != nil {
		return nil, err
	}

	res, _, err := prepareHttpResponse[GetEnergyStorageMachineAlertListResponse](req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *growattClient) GetMaxAlertListWithPagination(deviceSN string, timestamp int64, page, size int) (*GetMaxAlertListResponse, error) {
	queryMap := map[string]interface{}{
		"max_sn":  deviceSN,
		"page":    page,
		"perpage": size,
		"date":    time.Unix(timestamp, 0).Format("2006-01-02"),
	}

	query := BuildQueryParams(queryMap)
	url := r.URL + "/device/max/alarm_data" + query
	req, err := prepareHttpRequest(http.MethodGet, url, r.headers, nil)
	if err != nil {
		return nil, err
	}

	res, _, err := prepareHttpResponse[GetMaxAlertListResponse](req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *growattClient) GetMaxAlertList(deviceSN string, timestamp int64) ([]*AlarmItem, error) {
	alarms := make([]*AlarmItem, 0)
	page := 1
	for {
		res, err := r.GetMaxAlertListWithPagination(deviceSN, timestamp, page, MAX_PAGE_SIZE)
		if err != nil {
			return nil, err
		}

		alarms = append(alarms, res.Data.Alarms...)

		if len(alarms) >= res.Data.GetCount() {
			break
		}

		page += 1
	}

	return alarms, nil
}

func (r *growattClient) GetMixAlertListWithPagination(deviceSN string, timestamp int64, page, size int) (*GetMixAlertListResponse, error) {
	queryMap := map[string]interface{}{
		"mix_sn":  deviceSN,
		"page":    page,
		"perpage": size,
		"date":    time.Unix(timestamp, 0).Format("2006-01-02"),
	}

	query := BuildQueryParams(queryMap)
	url := r.URL + "/device/max/alarm_data" + query
	req, err := prepareHttpRequest(http.MethodGet, url, r.headers, nil)
	if err != nil {
		return nil, err
	}

	res, _, err := prepareHttpResponse[GetMixAlertListResponse](req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *growattClient) GetMixAlertList(deviceSN string, timestamp int64) ([]*AlarmItem, error) {
	alarms := make([]*AlarmItem, 0)
	page := 1
	for {
		res, err := r.GetMixAlertListWithPagination(deviceSN, timestamp, page, MAX_PAGE_SIZE)
		if err != nil {
			return nil, err
		}

		alarms = append(alarms, res.Data.Alarms...)

		if len(alarms) >= res.Data.GetCount() {
			break
		}

		page += 1
	}

	return alarms, nil
}

func (r *growattClient) GetSpaAlertListWithPagination(deviceSN string, timestamp int64, page, size int) (*GetSpaAlertListResponse, error) {
	queryMap := map[string]interface{}{
		"spa_sn":  deviceSN,
		"page":    page,
		"perpage": size,
		"date":    time.Unix(timestamp, 0).Format("2006-01-02"),
	}

	query := BuildQueryParams(queryMap)
	url := r.URL + "/device/spa/alarm_data" + query
	req, err := prepareHttpRequest(http.MethodGet, url, r.headers, nil)
	if err != nil {
		return nil, err
	}

	res, _, err := prepareHttpResponse[GetSpaAlertListResponse](req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *growattClient) GetSpaAlertList(deviceSN string, timestamp int64) ([]*AlarmItem, error) {
	alarms := make([]*AlarmItem, 0)
	page := 1
	for {
		res, err := r.GetSpaAlertListWithPagination(deviceSN, timestamp, page, MAX_PAGE_SIZE)
		if err != nil {
			return nil, err
		}

		alarms = append(alarms, res.Data.Alarms...)

		if len(alarms) >= res.Data.GetCount() {
			break
		}

		page += 1
	}

	return alarms, nil
}

func (r *growattClient) GetMinAlertListWithPagination(deviceSN string, timestamp int64, page, size int) (*GetMinAlertListResponse, error) {
	queryMap := map[string]interface{}{
		"tlx_sn":  deviceSN,
		"page":    page,
		"perpage": size,
		"date":    time.Unix(timestamp, 0).Format("2006-01-02"),
	}

	query := BuildQueryParams(queryMap)
	url := r.URL + "/device/tlx/alarm_data" + query
	req, err := prepareHttpRequest(http.MethodGet, url, r.headers, nil)
	if err != nil {
		return nil, err
	}

	res, _, err := prepareHttpResponse[GetMinAlertListResponse](req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *growattClient) GetMinAlertList(deviceSN string, timestamp int64) ([]*AlarmItem, error) {
	alarms := make([]*AlarmItem, 0)
	page := 1
	for {
		res, err := r.GetMinAlertListWithPagination(deviceSN, timestamp, page, MAX_PAGE_SIZE)
		if err != nil {
			return nil, err
		}

		alarms = append(alarms, res.Data.Alarms...)

		if len(alarms) >= res.Data.GetCount() {
			break
		}

		page += 1
	}

	return alarms, nil
}

func (r *growattClient) GetPcsAlertListWithPagination(deviceSN string, timestamp int64, page, size int) (*GetPcsAlertListResponse, error) {
	queryMap := map[string]interface{}{
		"pcs_sn":  deviceSN,
		"page":    page,
		"perpage": size,
		"date":    time.Unix(timestamp, 0).Format("2006-01-02"),
	}

	query := BuildQueryParams(queryMap)
	url := r.URL + "/device/pcs/alarm_data" + query
	req, err := prepareHttpRequest(http.MethodGet, url, r.headers, nil)
	if err != nil {
		return nil, err
	}

	res, _, err := prepareHttpResponse[GetPcsAlertListResponse](req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *growattClient) GetPcsAlertList(deviceSN string, timestamp int64) ([]*AlarmItem, error) {
	alarms := make([]*AlarmItem, 0)
	page := 1
	for {
		res, err := r.GetPcsAlertListWithPagination(deviceSN, timestamp, page, MAX_PAGE_SIZE)
		if err != nil {
			return nil, err
		}

		alarms = append(alarms, res.Data.Alarms...)

		if len(alarms) >= res.Data.GetCount() {
			break
		}

		page += 1
	}

	return alarms, nil
}

func (r *growattClient) GetHpsAlertListWithPagination(deviceSN string, timestamp int64, page, size int) (*GetHpsAlertListResponse, error) {
	queryMap := map[string]interface{}{
		"hps_sn":  deviceSN,
		"page":    page,
		"perpage": size,
		"date":    time.Unix(timestamp, 0).Format("2006-01-02"),
	}

	query := BuildQueryParams(queryMap)
	url := r.URL + "/device/hps/alarm_data" + query
	req, err := prepareHttpRequest(http.MethodGet, url, r.headers, nil)
	if err != nil {
		return nil, err
	}

	res, _, err := prepareHttpResponse[GetHpsAlertListResponse](req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *growattClient) GetHpsAlertList(deviceSN string, timestamp int64) ([]*AlarmItem, error) {
	alarms := make([]*AlarmItem, 0)
	page := 1
	for {
		res, err := r.GetHpsAlertListWithPagination(deviceSN, timestamp, page, MAX_PAGE_SIZE)
		if err != nil {
			return nil, err
		}

		alarms = append(alarms, res.Data.Alarms...)

		if len(alarms) >= res.Data.GetCount() {
			break
		}

		page += 1
	}

	return alarms, nil
}

func (r *growattClient) GetPbdAlertListWithPagination(deviceSN string, timestamp int64, page, size int) (*GetPbdAlertListResponse, error) {
	queryMap := map[string]interface{}{
		"pbd_sn":  deviceSN,
		"page":    page,
		"perpage": size,
		"date":    time.Unix(timestamp, 0).Format("2006-01-02"),
	}

	query := BuildQueryParams(queryMap)
	url := r.URL + "/device/pbd/alarm_data" + query
	req, err := prepareHttpRequest(http.MethodGet, url, r.headers, nil)
	if err != nil {
		return nil, err
	}

	res, _, err := prepareHttpResponse[GetPbdAlertListResponse](req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *growattClient) GetPbdAlertList(deviceSN string, timestamp int64) ([]*AlarmItem, error) {
	alarms := make([]*AlarmItem, 0)
	page := 1
	for {
		res, err := r.GetPbdAlertListWithPagination(deviceSN, timestamp, page, MAX_PAGE_SIZE)
		if err != nil {
			return nil, err
		}

		alarms = append(alarms, res.Data.Alarms...)

		if len(alarms) >= res.Data.GetCount() {
			break
		}

		page += 1
	}

	return alarms, nil
}

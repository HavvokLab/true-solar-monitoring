package kstar

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

type KStarClient interface {
	GetPlantList() (*GetPlantListResponse, error)
	GetDeviceListWithPagination(page, size int) (*GetDeviceListResponse, error)
	GetDeviceList() ([]*DeviceItem, error)
	GetRealtimeDeviceData(deviceID string) (*GetRealtimeDeviceDataResponse, error)
	GetRealtimeAlarmListOfDevice(deviceID string) (*GetRealtimeAlarmListOfDevice, error)
	GetRealtimeAlarmListOfPlantWithPagination(from, to int64, page, size int) (*GetAlarmListResponse, error)
	GetRealtimeAlarmListOfPlant(from, to int64) ([]*DeviceAlarmInfoItem, error)
}

type kstarClient struct {
	URL      string
	username string
	password string
}

func NewKStarClient(
	username, password string,
) KStarClient {
	client := kstarClient{
		URL:      URL_VERSION1,
		username: username,
		password: password,
	}

	client.password = client.DecodePassword(password)
	return &client
}

func (c *kstarClient) DecodePassword(password string) string {
	hashPassword := md5.Sum([]byte(password))
	return strings.ToUpper(fmt.Sprintf("%x", hashPassword[:]))
}

func (obj *kstarClient) EncodeParameter(data map[string]string) string {
	data["userCode"] = obj.username
	data["password"] = obj.password

	var keys []string
	for k := range data {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	var params []string
	for _, k := range keys {
		params = append(params, fmt.Sprintf("%s=%s", k, data[k]))
	}

	paramStr := strings.Join(params, "&")
	hash := sha1.New()
	hash.Write([]byte(paramStr))
	result := fmt.Sprintf("%x", hash.Sum(nil))
	return result
}

func (k *kstarClient) GetPlantList() (*GetPlantListResponse, error) {
	sign := k.EncodeParameter(make(map[string]string))
	data := url.Values{}
	data.Set("userCode", k.username)
	data.Set("password", k.password)
	data.Set("sign", sign)

	url := k.URL + "/power/info"
	req, err := prepareHttpRequest(http.MethodPost, url, data)
	if err != nil {
		return nil, err
	}

	result, _, err := prepareHttpResponse[GetPlantListResponse](req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (k *kstarClient) GetDeviceListWithPagination(page, size int) (*GetDeviceListResponse, error) {
	pageStr := strconv.Itoa(page)
	sizeStr := strconv.Itoa(size)
	params := map[string]string{
		"PageNum":  pageStr,
		"PageSize": sizeStr,
	}

	sign := k.EncodeParameter(params)
	data := url.Values{}
	data.Set("userCode", k.username)
	data.Set("password", k.password)
	data.Set("PageNum", pageStr)
	data.Set("PageSize", sizeStr)
	data.Set("sign", sign)

	url := k.URL + "/inverter/list"
	req, err := prepareHttpRequest(http.MethodPost, url, data)
	if err != nil {
		return nil, err
	}

	result, _, err := prepareHttpResponse[GetDeviceListResponse](req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (k *kstarClient) GetDeviceList() ([]*DeviceItem, error) {
	result := []*DeviceItem{}
	var page int = 1
	for {
		resp, err := k.GetDeviceListWithPagination(page, MAX_PAGE_SIZE)
		if err != nil {
			return nil, err
		}

		if len(resp.Data.List) == 0 {
			break
		}

		result = append(result, resp.Data.List...)
		if len(result) >= resp.Data.GetCount() {
			break
		}

		page += 1
	}
	return result, nil
}

func (k *kstarClient) GetRealtimeDeviceData(deviceID string) (*GetRealtimeDeviceDataResponse, error) {
	params := map[string]string{"deviceId": deviceID}
	sign := k.EncodeParameter(params)

	data := url.Values{}
	data.Set("userCode", k.username)
	data.Set("password", k.password)
	data.Set("deviceId", deviceID)
	data.Set("sign", sign)

	url := k.URL + "/device/real"
	req, err := prepareHttpRequest(http.MethodPost, url, data)
	if err != nil {
		return nil, err
	}

	result, _, err := prepareHttpResponse[GetRealtimeDeviceDataResponse](req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (k *kstarClient) GetRealtimeAlarmListOfDevice(deviceID string) (*GetRealtimeAlarmListOfDevice, error) {
	params := map[string]string{"deviceId": deviceID}
	sign := k.EncodeParameter(params)

	data := url.Values{}
	data.Set("userCode", k.username)
	data.Set("password", k.password)
	data.Set("deviceId", deviceID)
	data.Set("sign", sign)

	url := k.URL + "/alarm/device/list"
	req, err := prepareHttpRequest(http.MethodPost, url, data)
	if err != nil {
		return nil, err
	}

	result, _, err := prepareHttpResponse[GetRealtimeAlarmListOfDevice](req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (k *kstarClient) GetRealtimeAlarmListOfPlantWithPagination(from, to int64, page, size int) (*GetAlarmListResponse, error) {
	fromTime := time.Unix(from, 0)
	fromDate := fromTime.Format("2006-01-02")
	toTime := time.Unix(to, 0)
	toDate := toTime.Format("2006-01-02")

	pageStr := strconv.Itoa(page)
	sizeStr := strconv.Itoa(size)
	params := map[string]string{
		"stime":    fromDate,
		"etime":    toDate,
		"PageNum":  pageStr,
		"PageSize": sizeStr,
	}

	sign := k.EncodeParameter(params)
	data := url.Values{}
	data.Set("userCode", k.username)
	data.Set("password", k.password)
	data.Set("PageNum", pageStr)
	data.Set("PageSize", sizeStr)
	data.Set("stime", fromDate)
	data.Set("etime", toDate)
	data.Set("sign", sign)

	url := k.URL + "/real/alarm/list"
	req, err := prepareHttpRequest(http.MethodPost, url, data)
	if err != nil {
		return nil, err
	}

	result, _, err := prepareHttpResponse[GetAlarmListResponse](req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (k *kstarClient) GetRealtimeAlarmListOfPlant(from, to int64) ([]*DeviceAlarmInfoItem, error) {
	result := make([]*DeviceAlarmInfoItem, 0)
	var page int = 1
	for {
		resp, err := k.GetRealtimeAlarmListOfPlantWithPagination(from, to, page, MAX_PAGE_SIZE)
		if err != nil {
			return nil, err
		}

		if len(resp.Data.Data) == 0 {
			break
		}

		result = append(result, resp.Data.Data...)
		if len(result) >= resp.Data.GetCount() {
			break
		}
	}

	return result, nil
}

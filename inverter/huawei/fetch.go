package huawei

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/HavvokLab/true-solar-monitoring/util"
	"github.com/avast/retry-go"
)

func prepareHttpRequest(method, url string, headers map[string]string, data interface{}) (*http.Request, error) {
	var req *http.Request
	var err error

	if data != nil {
		encoded_data, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}

		buffered_data := bytes.NewBuffer(encoded_data)
		req, err = http.NewRequest(method, url, buffered_data)
		if err != nil {
			return nil, err
		}
	} else {
		req, err = http.NewRequest(method, url, nil)
		if err != nil {
			return nil, err
		}
	}

	req.Header.Set("Content-Type", "application/json")
	for key, val := range headers {
		req.Header.Set(key, val)
	}

	return req, nil
}

func prepareHttpResponse[R interface{}](req *http.Request) (*R, int, error) {
	// request to endpoint
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		retryOptions := []retry.Option{
			retry.Delay(RETRY_WAIT_TIME),
			retry.Attempts(RETRY_ATTEMPT),
			retry.DelayType(retry.FixedDelay),
		}

		err := retry.Do(func() error {
			res, err = client.Do(req)
			if err != nil {
				return err
			}

			return nil
		}, retryOptions...)

		if err != nil {
			return nil, http.StatusInternalServerError, err
		}
	}
	defer res.Body.Close()

	// read a bytes of data
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	// check empty response body
	if len(resBody) == 0 {
		return nil, res.StatusCode, nil
	}

	// decode response body
	var result R
	if err := util.Recast(resBody, &result); err != nil {
		huaweiErr := ResponseData{}
		util.Recast(resBody, &huaweiErr)
		return nil, http.StatusInternalServerError, huaweiErr
	}

	return &result, res.StatusCode, nil
}

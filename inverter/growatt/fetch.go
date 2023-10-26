package growatt

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/HavvokLab/true-solar-monitoring/util"
	"github.com/avast/retry-go"
	"github.com/bytedance/sonic"
)

func prepareHttpRequest(method, url string, headers map[string]string, data interface{}) (*http.Request, error) {
	var req *http.Request
	var err error

	if data != nil {
		encoded_data, err := sonic.Marshal(data)
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
	retryOptions := []retry.Option{
		retry.Delay(RETRY_WAIT_TIME),
		retry.Attempts(RETRY_ATTEMPT),
		retry.DelayType(retry.FixedDelay),
	}

	var resp *http.Response
	var result R
	err := retry.Do(func() error {
		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		if res.StatusCode == http.StatusTooManyRequests {
			fmt.Println("Too many request")
			return errors.New("too many request")
		}

		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		if len(resBody) == 0 {
			fmt.Println("empty body")
			return errors.New("empty response")
		}

		if err := checkHTMLResponse(resBody); err != nil {
			fmt.Println("HTML response")
			return err
		}

		if err := util.Recast(resBody, &result); err != nil {
			errResp := ErrorResponse{}
			if err := util.Recast(resBody, &errResp); err != nil {
				return err
			}

			fmt.Printf("[%v] %v\n", errResp.GetErrorCode(), errResp.GetErrorMsg())
			return fmt.Errorf("[%v] %v", errResp.GetErrorCode(), errResp.GetErrorMsg())
		}

		return nil
	}, retryOptions...)

	if err != nil {
		if resp != nil {
			return nil, resp.StatusCode, err
		}

		return nil, http.StatusInternalServerError, err
	}

	return &result, http.StatusOK, nil
}

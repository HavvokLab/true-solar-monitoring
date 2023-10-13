package kstar

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/HavvokLab/true-solar-monitoring/util"
	"github.com/avast/retry-go"
)

func prepareHttpRequest(method string, url string, values url.Values) (*http.Request, error) {
	req, err := http.NewRequest(method, url, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}

	contentType := "application/x-www-form-urlencoded"
	contentLength := strconv.Itoa(len(values.Encode()))
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Content-Length", contentLength)

	return req, nil
}

func prepareHttpResponse[R interface{}](req *http.Request) (*R, int, error) {
	// request to endpoint
	client := &http.Client{}
	retryOptions := []retry.Option{
		retry.Delay(RETRY_WAIT_TIME),
		retry.Attempts(RETRY_ATTEMPT),
		retry.DelayType(retry.FixedDelay),
	}

	var res *http.Response
	var err error
	err = retry.Do(func() error {
		res, err = client.Do(req)
		if err != nil {
			return err
		}

		if res != nil {
			if res.StatusCode == http.StatusTooManyRequests {
				return ErrorTooManyRequest
			}
		}

		return nil
	}, retryOptions...)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	defer res.Body.Close()

	// read a bytes of data
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	// check empty response body
	if len(resBody) == 0 {
		return nil, res.StatusCode, errors.New("empty body")
	}

	var result R
	var response MetaResponse
	if err := util.Recast(resBody, &response); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if response.Meta == nil {
		return nil, http.StatusInternalServerError, errors.New("meta is nil")
	}

	if !response.Meta.GetSuccess() {
		return nil, http.StatusInternalServerError, errors.New("response return success: false")
	}

	if err := util.Recast(resBody, &result); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &result, res.StatusCode, nil

}

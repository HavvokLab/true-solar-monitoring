package growatt

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mitchellh/mapstructure"
)

type GrowattInverterProduction struct {
	Total *float64
	Today *float64
}

func CalculateInverterProductions(credential *GrowattCredential, inverterSNs []string) (map[string]GrowattInverterProduction, error) {
	client, err := NewGrowattClient(credential)
	if err != nil {
		return nil, err
	}

	resp, err := client.GetRealtimeDeviceBatchesData(inverterSNs)
	if err != nil {
		return nil, err
	}

	if resp.Data == nil {
		return nil, errors.New("empty response data")
	}

	deviceMap := make(map[string]GrowattInverterProduction)
	for sn, data := range resp.Data {
		if mappedData, ok := data[sn].(map[string]interface{}); ok {
			var decoded RealtimeDeviceData
			if err := mapstructure.Decode(&mappedData, &decoded); err == nil {
				deviceMap[sn] = GrowattInverterProduction{
					Total: decoded.PowerTotal,
					Today: decoded.PowerToday,
				}
			}
		}
	}

	return deviceMap, nil
}

func checkHTMLResponse(body []byte) error {
	if len(body) > 0 && htmlTagsRegExp.Find(body) != nil {
		return ErrResponseMustNotBeHTML
	}

	return nil
}

func BuildQueryParams(params map[string]interface{}) string {
	var queries []string
	for key, val := range params {
		query := fmt.Sprintf("%v=%v", key, val)
		queries = append(queries, query)
	}

	if len(queries) > 0 {
		return "?" + strings.Join(queries, "&")
	}
	return ""
}

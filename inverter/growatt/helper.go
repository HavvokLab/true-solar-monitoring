package growatt

import (
	"fmt"
	"strings"
)

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

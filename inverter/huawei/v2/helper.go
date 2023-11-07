package huawei

import "fmt"

func checkHTMLResponse(body []byte) error {
	if len(body) > 0 && htmlTagsRegExp.Find(body) != nil {
		return fmt.Errorf("response must not be HTML: %s", string(body))
	}

	return nil
}

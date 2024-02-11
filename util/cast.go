package util

import "encoding/json"

func Recast(from, to interface{}) error {
	switch v := from.(type) {
	case []byte:
		return json.Unmarshal(v, to)
	default:
		buf, err := json.Marshal(from)
		if err != nil {
			return err
		}

		return json.Unmarshal(buf, to)
	}
}

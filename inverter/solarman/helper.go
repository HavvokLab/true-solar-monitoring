package solarman

import (
	"crypto/sha256"
	"fmt"
	"time"
)

func checkHTMLResponse(body []byte) error {
	if len(body) > 0 && htmlTagsRegExp.Find(body) != nil {
		return ErrResponseMustNotBeHTML
	}

	return nil
}

func buildAuthorizationToken(token string) string {
	return "Bearer " + token
}

func buildDateFromTimestamp(timestamp int64, timeType int) string {
	switch timeType {
	case TIME_TYPE_YEAR:
		return time.Unix(timestamp, 0).Format("2006")
	case TIME_TYPE_MONTH:
		return time.Unix(timestamp, 0).Format("2006-01")
	default:
		return time.Unix(timestamp, 0).Format("2006-01-02")
	}
}

func DecodePassword(password string) string {
	hashPassword := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", hashPassword[:])
}

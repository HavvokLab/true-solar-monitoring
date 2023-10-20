package util

import (
	"encoding/json"
	"net/http"

	"github.com/HavvokLab/true-solar-monitoring/errors"
)

var sqliteErrorCodes = map[int]error{
	1555: errors.NewServerError(http.StatusBadRequest, "Duplicate entry"),
	2067: errors.NewServerError(http.StatusBadRequest, "Duplicate entry"),
	12:   errors.NewServerError(http.StatusNotFound, "Record not found"),
}

type SqliteError struct {
	Code         int `json:"Code"`
	ExtendedCode int `json:"ExtendedCode"`
	SystemErrno  int `json:"SystemErrno"`
}

func TranslateSqliteError(err error) error {
	parsedErr, marshalErr := json.Marshal(err)
	if marshalErr != nil {
		return err
	}

	var errMsg SqliteError
	unmarshalErr := json.Unmarshal(parsedErr, &errMsg)
	if unmarshalErr != nil {
		return err
	}

	if translatedErr, found := sqliteErrorCodes[errMsg.ExtendedCode]; found {
		return translatedErr
	}
	return err
}

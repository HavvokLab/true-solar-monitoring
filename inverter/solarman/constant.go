package solarman

import (
	"regexp"
	"time"
)

var htmlTagsRegExp = regexp.MustCompile(`<\/?[a-z][\s\S]*>`)

const BRAND = "solarman"

const (
	URL_VERSION1                  = "https://globalapi.solarmanpv.com"
	AUTHORIZATION_HEADER          = "Authorization"
	MAX_PAGE_SIZE                 = 200
	SOLARMAN_DEVICE_TYPE_INVERTER = "INVERTER"
	RETRY_WAIT_TIME               = 30 * time.Second
	RETRY_ATTEMPT                 = 3
)

const (
	TIME_TYPE_TIMEFRAME = iota + 1
	TIME_TYPE_DAY
	TIME_TYPE_MONTH
	TIME_TYPE_YEAR
)

const (
	DATA_LIST_KEY_CUMULATIVE_PRODUCTION = "Et_ge0"
	DATA_LIST_KEY_GENERATION            = "generation"
)

const (
	SOLARMAN_PLANT_STATUS_ON    = "ONLINE"
	SOLARMAN_PLANT_STATUS_OFF   = "OFFLINE"
	SOLARMAN_PLANT_STATUS_ALARM = "ALARM"
)

const (
	DEVICE_STATUS_ON      = "ONLINE"
	DEVICE_STATUS_OFF     = "OFFLINE"
	DEVICE_STATUS_FAILURE = "FAILURE"
)

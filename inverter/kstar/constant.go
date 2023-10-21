package kstar

import "time"

const BRAND = "kstar"

const (
	URL_VERSION1    = "http://solar.kstar.com:9000/public"
	MAX_PAGE_SIZE   = 200
	RETRY_WAIT_TIME = 30 * time.Second
	RETRY_ATTEMPT   = 3
)

const (
	KSTAR_DEVICE_TYPE_INVERTER = "INVERTER"
)

const (
	KSTAR_API_CALL_TYPE_PLANT  = "plant"
	KSTAR_API_CALL_TYPE_DEVICE = "device"
	KSTAR_API_CALL_TYPE_ALARM  = "alarm"
)

const (
	KSTAR_DEVICE_STATUS_ON    = "ONLINE"
	KSTAR_DEVICE_STATUS_OFF   = "OFFLINE"
	KSTAR_DEVICE_STATUS_ALARM = "ALARM"
)

const (
	KSTAR_COLLECTOR_TASK_TYPE_ALL_USER_DATA = iota
	KSTAR_COLLECTOR_TASK_TYPE_USER_DATA
)
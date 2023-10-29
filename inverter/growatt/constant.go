package growatt

import (
	"fmt"
	"regexp"
	"time"
)

var htmlTagsRegExp = regexp.MustCompile(`<\/?[a-z][\s\S]*>`)

const BRAND = "growatt"

const (
	URL_VERSION1      = "https://openapi.growatt.com/v1"
	AUTH_HEADER       = "Token"
	MAX_PAGE_SIZE     = 100
	HALF_A_BATCH_SIZE = 50
	WAIT_TIME         = 15 * time.Second
	ERROR_WAIT_TIME   = 5 * time.Minute
	RETRY_WAIT_TIME   = 5 * time.Minute
	RETRY_ATTEMPT     = 3
)

const (
	GROWATT_TYPE_PLANT  = "plant"
	GROWATT_TYPE_DEVICE = "device"
	GROWATT_TYPE_ALARM  = "alarm"
)

const (
	GROWATT_DEVICE_STATUS_ONLINE      = "ONLINE"
	GROWATT_DEVICE_STATUS_OFFLINE     = "OFFLINE"
	GROWATT_DEVICE_STATUS_DISCONNECT  = "DISCONNECT"
	GROWATT_DEVICE_STATUS_STAND_BY    = "STAND BY"
	GROWATT_DEVICE_STATUS_FAILURE     = "FAILURE"
	GROWATT_DEVICE_STATUS_CHARGING    = "CHARGING"
	GROWATT_DEVICE_STATUS_DISCHARGING = "DISCHARGING"
	GROWATT_DEVICE_STATUS_BURNING     = "BURNING"
	GROWATT_DEVICE_STATUS_WAITING     = "WAITING"
	GROWATT_DEVICE_STATUS_SELF_CHECK  = "SELF CHECK"
	GROWATT_DEVICE_STATUS_UPGRADING   = "UPGRADING"
)

const (
	GROWATT_PLANT_STATUS_ON    = "ONLINE"
	GROWATT_PLANT_STATUS_OFF   = "OFFLINE"
	GROWATT_PLANT_STATUS_ALARM = "ALARM"
)

const (
	GROWATT_COLLECTOR_TASK_TYPE_ALL_USER_DATA = iota
	GROWATT_COLLECTOR_TASK_TYPE_USER_DATA
)

const (
	GROWATT_DEVICE_TYPE_INVERTER = iota + 1
	GROWATT_DEVICE_TYPE_ENERGY_STORAGE_MACHINE
	GROWATT_DEVICE_TYPE_OTHER_EQUIPMENT
	GROWATT_DEVICE_TYPE_MAX
	GROWATT_DEVICE_TYPE_MIX
	GROWATT_DEVICE_TYPE_SPA
	GROWATT_DEVICE_TYPE_MIN
	GROWATT_DEVICE_TYPE_PCS
	GROWATT_DEVICE_TYPE_HPS
	GROWATT_DEVICE_TYPE_PBD
)

func ParseGrowattDeviceType(deviceType int) string {
	switch deviceType {
	case GROWATT_DEVICE_TYPE_INVERTER:
		return "INVERTER"
	case GROWATT_DEVICE_TYPE_ENERGY_STORAGE_MACHINE:
		return "ENERGY STORAGE MACHINE"
	case GROWATT_DEVICE_TYPE_OTHER_EQUIPMENT:
		return "OTHER EQUIPMENT"
	case GROWATT_DEVICE_TYPE_MAX:
		return "MAX"
	case GROWATT_DEVICE_TYPE_MIX:
		return "MIX"
	case GROWATT_DEVICE_TYPE_SPA:
		return "SPA"
	case GROWATT_DEVICE_TYPE_MIN:
		return "MIN"
	case GROWATT_DEVICE_TYPE_PCS:
		return "PCS"
	case GROWATT_DEVICE_TYPE_HPS:
		return "HPS"
	case GROWATT_DEVICE_TYPE_PBD:
		return "PBD"
	default:
		return ""
	}
}

const (
	ERROR_NORMAL           = 0
	ERROR_NO_ACCESS        = 10011
	ERROR_FREQUENCY_ACCESS = 10012
)

func ParseGrowattErrorMessage(errorCode int) string {
	switch errorCode {
	case ERROR_NORMAL:
		return "error: normal(general)"
	case ERROR_NO_ACCESS:
		return "error: no privilege access (generic)"
	case ERROR_FREQUENCY_ACCESS:
		return "error: frequency access"
	default:
		return fmt.Sprintf("error code: %q", errorCode)
	}
}

var (
	GROWATT_EQUIP_TYPE_MAPPER      = []string{"", "Inverter", "Energy storage machine", "Other equipment", "MAX", "MIX", "SPA", "MIN", "PCS", "HPS", "PBD"}
	GROWATT_INVERTER_STATUS_MAPPER = []string{"Disconnect", "Online", "Stanby", "Fault"}
)

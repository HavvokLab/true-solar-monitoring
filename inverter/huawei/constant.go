package huawei

import "time"

const BRAND = "huawei"

const (
	URL_VERSION1        = "https://sg5.fusionsolar.huawei.com"
	AUTH_HEADER         = "XSRF-TOKEN"
	HUAWEI_LANG_ENGLISH = "en_UK"
	HUAWEI_CURRENCY_USD = "USD"
	RETRY_WAIT_TIME     = 60 * time.Second
	RETRY_ATTEMPT       = 3
)

const (
	HUAWEI_TYPE_PLANT  = "plant"
	HUAWEI_TYPE_DEVICE = "device"
	HUAWEI_TYPE_ALARM  = "alarm"
)

const (
	HUAWEI_STATUS_ON    = "ONLINE"
	HUAWEI_STATUS_OFF   = "OFFLINE"
	HUAWEI_STATUS_ALARM = "ALARM"
)

var HuaweiMapPlantStatus = map[int]string{
	1: HUAWEI_STATUS_OFF,
	2: HUAWEI_STATUS_ALARM,
	3: HUAWEI_STATUS_ON,
}

var HuaweiMapDeviceStatus = map[int]string{
	0: HUAWEI_STATUS_OFF,
	1: HUAWEI_STATUS_ON,
}

const (
	HUAWEI_DEVICE_TYPE_INVERTER                                      = "INVERTER"
	HUAWEI_DEVICE_TYPE_SMART_LOGGER                                  = "SMART LOGGER"
	HUAWEI_DEVICE_TYPE_STRING                                        = "STRING"
	HUAWEI_DEVICE_TYPE_BAY                                           = "BAY"
	HUAWEI_DEVICE_TYPE_BUSBAR                                        = "BUSBAR"
	HUAWEI_DEVICE_TYPE_TRANSFORMER                                   = "TRANSFORMER"
	HUAWEI_DEVICE_TYPE_TRANSFORMER_METER                             = "TRANSFORMER METER"
	HUAWEI_DEVICE_TYPE_EMI                                           = "EMI"
	HUAWEI_DEVICE_TYPE_AC_COMBINER_BOX                               = "AC COMBINER BOX"
	HUAWEI_DEVICE_TYPE_DPU                                           = "DPU"
	HUAWEI_DEVICE_TYPE_CENTRAL_INVERTER                              = "CENTRAL INVERTER"
	HUAWEI_DEVICE_TYPE_DC_COMBINER_BOX                               = "DC COMBINER BOX"
	HUAWEI_DEVICE_TYPE_GENERAL_DEVICE                                = "GENERAL DEVICE"
	HUAWEI_DEVICE_TYPE_GATEWAY_POWER_METER                           = "GATEWAY POWER METER"
	HUAWEI_DEVICE_TYPE_STEP_UP_STATION                               = "STEP-UP STATION"
	HUAWEI_DEVICE_TYPE_FACTORY_USED_ENERGY_GENERATION_AREA_METER     = "FACTORY-USED ENERGY GENERATION AREA METER"
	HUAWEI_DEVICE_TYPE_SOLAR_POWER_FORECASTING_SYSTEM                = "SOLAR POWER FORECASTING SYSTEM"
	HUAWEI_DEVICE_TYPE_FACTORY_USED_ENERGY_NON_GENERATION_AREA_METER = "FACTORY-USED ENERGY NON-GENERATION AREA METER"
	HUAWEI_DEVICE_TYPE_PID                                           = "PID"
	HUAWEI_DEVICE_TYPE_VIRTUAL_DEVICE_OF_PLANT_MONITORING_SYSTEM     = "VIRTUAL DEVICE OF PLANT MONITORING SYSTEM"
	HUAWEI_DEVICE_TYPE_POWER_QUALITY_DEVICE                          = "POWER QUALITY DEVICE"
	HUAWEI_DEVICE_TYPE_STEP_UP_TRANSFORMER                           = "STEP-UP TRANSFORMER"
	HUAWEI_DEVICE_TYPE_PHOTOVOLTAIC_GRID_CONNECTION_CABINET          = "PHOTOVOLTAIC GRID CONNECTION CABINET"
	HUAWEI_DEVICE_TYPE_PHOTOVOLTAIC_GRID_CONNECTION_PANEL            = "PHOTOVOLTAIC GRID CONNECTION PANEL"
	HUAWEI_DEVICE_TYPE_PINNET_SMART_LOGGER                           = "PINNET SMART LOGGER"
	HUAWEI_DEVICE_TYPE_SMART_ENERGY_CENTER                           = "SMART ENERGY CENTER"
	HUAWEI_DEVICE_TYPE_BATTERY                                       = "BATTERY"
	HUAWEI_DEVICE_TYPE_SMART_BACKUP_BOX                              = "SMART BACKUP BOX"
	HUAWEI_DEVICE_TYPE_PLC                                           = "PLC"
	HUAWEI_DEVICE_TYPE_OPTIMIZER                                     = "OPTIMIZER"
	HUAWEI_DEVICE_TYPE_POWER_SENSOR                                  = "POWER SENSOR"
	HUAWEI_DEVICE_TYPE_SANSHO_SMART_LOGGER                           = "SANSHO SMART LOGGER"
	HUAWEI_DEVICE_TYPE_STEP_UP_TRANSFORMER_HIGH_VOLTAGE_BAY          = "STEP-UP TRANSFORMER HIGH VOLTAGE BAY"
	HUAWEI_DEVICE_TYPE_STEP_UP_TRANSFORMER_LOW_VOLTAGE_BAY           = "STEP-UP TRANSFORMER LOW VOLTAGE BAY"
	HUAWEI_DEVICE_TYPE_LINE_BAY                                      = "LINE BAY"
	HUAWEI_DEVICE_TYPE_TRANSFORMER_BAY                               = "TRANSFORMER BAY"
	HUAWEI_DEVICE_TYPE_SVC_SVG_BAY                                   = "SVC/SVG BAY"
	HUAWEI_DEVICE_TYPE_BUSBAR_SEGMENT_BAY                            = "BUSBAR/SEGMENT BAY"
	HUAWEI_DEVICE_TYPE_PLANT_POWER_SUPPLY                            = "PLANT POWER SUPPLY"
	HUAWEI_DEVICE_TYPE_SMART_DONGLE                                  = "SMART DONGLE"
	HUAWEI_DEVICE_TYPE_SMART_LOGGER_1000A                            = "SMART LOGGER 1000A"
	HUAWEI_DEVICE_TYPE_SAFETY_BOX                                    = "SAFETY BOX"
)

var HuaweiMapDeviceType = map[int]string{
	1:  HUAWEI_DEVICE_TYPE_INVERTER,
	2:  HUAWEI_DEVICE_TYPE_SMART_LOGGER,
	3:  HUAWEI_DEVICE_TYPE_STRING,
	6:  HUAWEI_DEVICE_TYPE_BAY,
	7:  HUAWEI_DEVICE_TYPE_BUSBAR,
	8:  HUAWEI_DEVICE_TYPE_TRANSFORMER,
	9:  HUAWEI_DEVICE_TYPE_TRANSFORMER_METER,
	10: HUAWEI_DEVICE_TYPE_EMI,
	11: HUAWEI_DEVICE_TYPE_AC_COMBINER_BOX,
	13: HUAWEI_DEVICE_TYPE_DPU,
	14: HUAWEI_DEVICE_TYPE_CENTRAL_INVERTER,
	15: HUAWEI_DEVICE_TYPE_DC_COMBINER_BOX,
	16: HUAWEI_DEVICE_TYPE_GENERAL_DEVICE,
	17: HUAWEI_DEVICE_TYPE_GATEWAY_POWER_METER,
	18: HUAWEI_DEVICE_TYPE_STEP_UP_STATION,
	19: HUAWEI_DEVICE_TYPE_FACTORY_USED_ENERGY_GENERATION_AREA_METER,
	20: HUAWEI_DEVICE_TYPE_SOLAR_POWER_FORECASTING_SYSTEM,
	21: HUAWEI_DEVICE_TYPE_FACTORY_USED_ENERGY_NON_GENERATION_AREA_METER,
	22: HUAWEI_DEVICE_TYPE_PID,
	23: HUAWEI_DEVICE_TYPE_VIRTUAL_DEVICE_OF_PLANT_MONITORING_SYSTEM,
	24: HUAWEI_DEVICE_TYPE_POWER_QUALITY_DEVICE,
	25: HUAWEI_DEVICE_TYPE_STEP_UP_TRANSFORMER,
	26: HUAWEI_DEVICE_TYPE_PHOTOVOLTAIC_GRID_CONNECTION_CABINET,
	27: HUAWEI_DEVICE_TYPE_PHOTOVOLTAIC_GRID_CONNECTION_PANEL,
	37: HUAWEI_DEVICE_TYPE_PINNET_SMART_LOGGER,
	38: HUAWEI_DEVICE_TYPE_SMART_ENERGY_CENTER,
	39: HUAWEI_DEVICE_TYPE_BATTERY,
	40: HUAWEI_DEVICE_TYPE_SMART_BACKUP_BOX,
	45: HUAWEI_DEVICE_TYPE_PLC,
	46: HUAWEI_DEVICE_TYPE_OPTIMIZER,
	47: HUAWEI_DEVICE_TYPE_POWER_SENSOR,
	52: HUAWEI_DEVICE_TYPE_SANSHO_SMART_LOGGER,
	53: HUAWEI_DEVICE_TYPE_STEP_UP_TRANSFORMER_HIGH_VOLTAGE_BAY,
	54: HUAWEI_DEVICE_TYPE_STEP_UP_TRANSFORMER,
	55: HUAWEI_DEVICE_TYPE_STEP_UP_TRANSFORMER_LOW_VOLTAGE_BAY,
	56: HUAWEI_DEVICE_TYPE_BUSBAR,
	57: HUAWEI_DEVICE_TYPE_LINE_BAY,
	58: HUAWEI_DEVICE_TYPE_TRANSFORMER_BAY,
	59: HUAWEI_DEVICE_TYPE_SVC_SVG_BAY,
	60: HUAWEI_DEVICE_TYPE_BUSBAR_SEGMENT_BAY,
	61: HUAWEI_DEVICE_TYPE_PLANT_POWER_SUPPLY,
	62: HUAWEI_DEVICE_TYPE_SMART_DONGLE,
	63: HUAWEI_DEVICE_TYPE_SMART_LOGGER_1000A,
	70: HUAWEI_DEVICE_TYPE_SAFETY_BOX,
}

var HuaweiMapErrorMessage = map[int]string{
	20001: "The third-party system ID does not exist.",
	20002: "The third-party system is forbidden.",
	20003: "The third-party system expires.",
	20004: "The server is abnormal.",
	20005: "The device ID cannot be empty.",
	20006: "Some devices do not match the device type.",
	20007: "The system does not have the desired power plant resources.",
	20008: "The system does not have the desired device resources.",
	20009: "Queried KPIs are not configured in the system.",
	20010: "The plant list cannot be empty.",
	20011: "The device list cannot be empty.",
	20012: "The query time cannot be empty.",
	20013: "The device type is incorrect. The interface does not support operations on some devices.",
	20014: "A maximum of 100 plants can be queried at a time.",
	20015: "A maximum of 100 plants can be queried at a time.",
	20016: "A maximum of 100 devices can be queried at a time.",
	20017: "A maximum of 100 devices can be queried at a time.",
	20018: "A maximum of 10 devices can be manipulated at a time.",
	20019: "The switch type is incorrect. 1 and 2 indicate switch-on and switch-off respectively.",
	20020: "The upgrade package specific to the device version cannot be found.",
	20021: "The upgrade file does not exist.",
	20022: "The upgrade records of the devices in the system are not found.",
	305:   "You are not in the login state. You need to log in again.",
	401:   "You do not have the related data interface permission.",
	407:   "The interface access frequency is too high.",
	0:     "",
}

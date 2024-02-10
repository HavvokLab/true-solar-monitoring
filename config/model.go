package config

type Config struct {
	API                   APIConfig            `mapstructure:"api"`
	Authentication        AuthenticationConfig `mapstructure:"authentication"`
	Elastic               ElasticsearchConfig  `mapstructure:"elasticsearch"`
	Redis                 RedisConfig          `mapstructure:"redis"`
	Snmp                  SnmpConfig           `mapstructure:"snmp"`
	CalculateProduction   AlarmConfig          `mapstructure:"calculate_production"`
	LowPerformanceAlarm   AlarmConfig          `mapstructure:"low_performance_alarm"`
	SumPerformanceAlarm   AlarmConfig          `mapstructure:"sum_performance_alarm"`
	DailyPerformanceAlarm AlarmConfig          `mapstructure:"daily_performance_alarm"`
	Solarman              InverterConfig       `mapstructure:"solarman"`
	Huawei                InverterConfig       `mapstructure:"huawei"`
	Growatt               InverterConfig       `mapstructure:"growatt"`
	KStar                 InverterConfig       `mapstructure:"kstar"`
	PlantAggregate        PlantAggregateConfig `mapstructure:"plant_aggr"`
}

type ElasticsearchConfig struct {
	Host                   string `mapstructure:"host"`
	Username               string `mapstructure:"username"`
	Password               string `mapstructure:"password"`
	SolarIndex             string `mapstructure:"solar_index"`
	SiteStationIndex       string `mapstructure:"site_station_index"`
	DailyProductionIndex   string `mapstructure:"daily_production_index"`
	MonthlyProductionIndex string `mapstructure:"monthly_production_index"`
	AlarmIndex             string `mapstructure:"alarm_index"`
	PerformanceAlarmIndex  string `mapstructure:"performance_alarm_index"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type SnmpConfig struct {
	AgentHost  string `mapstructure:"agent_host"`
	TargetHost string `mapstructure:"target_host"`
	TargetPort uint16 `mapstructure:"target_port"`
}

type InverterConfig struct {
	CollectorCrontab      string `mapstructure:"collector_crontab"`
	NightCollectorCrontab string `mapstructure:"night_collector_crontab"`
	AlarmCrontab          string `mapstructure:"alarm_crontab"`
}

type AlarmConfig struct {
	Crontab string `mapstructure:"crontab"`
}

type PlantAggregateConfig struct {
	Crontab string `mapstructure:"crontab"`
}

type AuthenticationConfig struct {
	Secret  string `mapstructure:"secret"`
	Expired int64  `mapstructure:"expired"`
}

type APIConfig struct {
	Port string `mapstructure:"port"`
}

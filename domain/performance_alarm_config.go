package domain

type UpdatePerformanceAlarmConfigRequest struct {
	Interval   int     `json:"interval" validate:"required"`
	HitDay     *int    `json:"hit_day"`
	Percentage float64 `json:"percentage" validate:"required"`
	Duration   *int    `json:"duration"`
}

package domain

type UpdateInstalledCapacityRequest struct {
	EfficiencyFactor float64 `json:"efficiency_factor" validate:"required"`
	FocusHour        int     `json:"focus_hour" validate:"required"`
}

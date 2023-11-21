package domain

import "github.com/HavvokLab/true-solar-monitoring/model"

type FindAllSiteRegionMappingsResponse struct {
	Count              int64                     `json:"count"`
	SiteRegionMappings []model.SiteRegionMapping `json:"site_region_mappings"`
}

type UpdateRegionRequest struct {
	Area   string   `json:"area"  validate:"required"`
	Cities []string `json:"cities" validate:"required"`
}

type UpdateCityRequest struct {
	Code string  `json:"code" validate:"required"`
	Name string  `json:"city"`
	Area *string `json:"area,omitempty"`
}

type CreateCityRequest struct {
	Code string  `json:"code" validate:"required"`
	Name string  `json:"city"  validate:"required"`
	Area *string `json:"area,omitempty"`
}

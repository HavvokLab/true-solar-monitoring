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

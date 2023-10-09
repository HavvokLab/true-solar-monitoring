package repo

import (
	"github.com/HavvokLab/true-solar-monitoring/model"
	"gorm.io/gorm"
)

type SiteRegionMappingRepo interface {
	GetSiteRegionMappings() ([]model.SiteRegionMapping, error)
}

type siteRegionMappingRepo struct {
	db *gorm.DB
}

func NewSiteRegionMappingRepo(db *gorm.DB) SiteRegionMappingRepo {
	return &siteRegionMappingRepo{db: db}
}

func (r *siteRegionMappingRepo) GetSiteRegionMappings() ([]model.SiteRegionMapping, error) {
	tx := r.db.Session(&gorm.Session{})
	var siteRegionMappings []model.SiteRegionMapping
	err := tx.Find(&siteRegionMappings).Error
	if err != nil {
		return nil, err
	}

	return siteRegionMappings, nil
}

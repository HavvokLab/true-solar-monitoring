package repo

import (
	"os"

	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/gocarina/gocsv"
)

type MasterSiteRepo interface {
	ExportToMap() map[string]model.MasterSite
}

type masterSiteRepo struct {
	masterSites []model.MasterSite
}

func NewMasterSiteRepo() (MasterSiteRepo, error) {
	obj := masterSiteRepo{masterSites: make([]model.MasterSite, 0)}
	in, err := os.Open("master_site.csv")
	if err != nil {
		return nil, err
	}

	sites := make([]model.MasterSite, 0)
	if err := gocsv.UnmarshalFile(in, &sites); err != nil {
		return nil, err
	}

	obj.masterSites = sites
	return &obj, nil
}

func (r *masterSiteRepo) ExportToMap() map[string]model.MasterSite {
	m := make(map[string]model.MasterSite)
	for _, site := range r.masterSites {
		m[site.GetKey()] = site
	}
	return m
}

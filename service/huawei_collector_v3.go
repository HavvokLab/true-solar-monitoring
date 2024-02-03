package service

import (
	"strings"

	"github.com/HavvokLab/true-solar-monitoring/config"
	"github.com/HavvokLab/true-solar-monitoring/constant"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/HavvokLab/true-solar-monitoring/repo"
)

type HuaweiCollectorV3Service interface {
	Run(*model.HuaweiCredential) error
}

type huaweiCollectorV3Service struct {
	vendorType         string
	huaweiAltervimRepo repo.HuaweiAltervimRepo
	siteRegionRepo     repo.SiteRegionMappingRepo
	siteRegions        []model.SiteRegionMapping
	solarRepo          repo.SolarRepo
	elasticConfig      config.ElasticsearchConfig
	logger             logger.Logger
}

func NewHuaweiCollectorV3Service(
	huaweiAltervimRepo repo.HuaweiAltervimRepo,
	solarRepo repo.SolarRepo,
	siteRegionRepo repo.SiteRegionMappingRepo,
	logger logger.Logger,
) *huaweiCollectorV3Service {
	return &huaweiCollectorV3Service{
		vendorType:         strings.ToUpper(constant.VENDOR_TYPE_HUAWEI),
		huaweiAltervimRepo: huaweiAltervimRepo,
		siteRegionRepo:     siteRegionRepo,
		solarRepo:          solarRepo,
		logger:             logger,
		siteRegions:        make([]model.SiteRegionMapping, 0),
		elasticConfig:      config.GetConfig().Elastic,
	}
}

func (s *huaweiCollectorV3Service) run(
	credential *model.HuaweiCredential,
	documentCh chan interface{},
	doneCh chan bool,
	errorCh chan error,
) {
	doneCh <- true
}

func (s *huaweiCollectorV3Service) preparePlantAndDevice() error {
	return nil
}

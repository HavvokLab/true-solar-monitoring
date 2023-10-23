package api

import (
	"github.com/HavvokLab/true-solar-monitoring/handler"
	"github.com/HavvokLab/true-solar-monitoring/infra"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/middleware"
	"github.com/HavvokLab/true-solar-monitoring/repo"
	"github.com/HavvokLab/true-solar-monitoring/service"
	"github.com/gofiber/fiber/v2"
)

func bindPrivateSiteMappingRegionAPI(router fiber.Router) {
	siteRegionMappingRepo := repo.NewSiteRegionMappingRepo(infra.GormDB)
	siteRegionMappingServ := service.NewSiteRegionMappingService(siteRegionMappingRepo, logger.GetLogger())
	hdl := handler.NewSiteRegionMappingHandler(siteRegionMappingServ)

	authMiddleware, authWrapper := middleware.NewAuthMiddleware()
	sub := router.Group("/site-region")
	sub.Use(authMiddleware())
	sub.Get("/city", authWrapper(hdl.GetCities))
	sub.Post("/city", authWrapper(hdl.CreateCity))
	sub.Put("/city/:id", authWrapper(hdl.UpdateCity))
	sub.Delete("/city/:id", authWrapper(hdl.DeleteCity))
	sub.Get("/region", authWrapper(hdl.GetRegions))
	sub.Put("/region", authWrapper(hdl.UpdateRegion))
	sub.Delete("/area/:area", authWrapper(hdl.DeleteArea))
}

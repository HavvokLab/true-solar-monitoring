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

/*
GET | /install-capacity
PUT | /install-capacity
*/

func bindPrivateInstalledCapacityAPI(router fiber.Router) {
	installedCapacityRepo := repo.NewInstalledCapacityRepo(infra.GormDB)
	installedCapacityServ := service.NewInstalledCapacityService(installedCapacityRepo, logger.GetLogger())
	hdl := handler.NewInstalledCapacityHandler(installedCapacityServ)

	authMiddleware, authWrapper := middleware.NewAuthMiddleware()
	sub := router.Group("/installed-capacity")
	sub.Use(authMiddleware())
	sub.Get("", authWrapper(hdl.FindOne))
	sub.Put("", authWrapper(hdl.Update))

}

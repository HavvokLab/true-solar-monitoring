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

func bindPrivatePerformanceAlarmConfigAPI(router fiber.Router) {
	performanceAlarmConfigRepo := repo.NewPerformanceAlarmConfigRepo(infra.GormDB)
	performanceAlarmConfigServ := service.NewPerformanceAlarmConfigService(performanceAlarmConfigRepo, logger.GetLogger())
	hdl := handler.NewPerformanceAlarmConfigHandler(performanceAlarmConfigServ)

	authMiddleware, authWrapper := middleware.NewAuthMiddleware()
	sub := router.Group("/performance-alarm")
	sub.Use(authMiddleware())
	sub.Get("/low", authWrapper(hdl.GetLowPerformanceAlarmConfig))
	sub.Put("/low", authWrapper(hdl.UpdateLowPerformanceAlarmConfig))
}

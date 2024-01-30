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

func bindPrivatePlantAPI(router fiber.Router) {
	plantRepo := repo.NewPlantRepo(infra.GormDB)
	plantServ := service.NewPlantService(plantRepo, logger.GetLogger())
	hdl := handler.NewPlantHandler(plantServ)

	authMiddleware, authWrapper := middleware.NewAuthMiddleware()
	sub := router.Group("/plant")
	sub.Use(authMiddleware())
	sub.Get("", authWrapper(hdl.FindAll))
	sub.Get("all", authWrapper(hdl.All))
	sub.Delete("/:id", authWrapper(hdl.Delete))
}

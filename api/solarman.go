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
GET    | /solarman
POST   | /solarman
PUT    | /solarman
DELETE | /solarman
*/

func bindPrivateSolarmanAPI(router fiber.Router) {
	credentialRepo := repo.NewSolarmanCredentialRepo(infra.GormDB)
	credentialServ := service.NewSolarmanCredentialService(credentialRepo, logger.GetLogger())
	credentialHdl := handler.NewSolarmanCredentialHandler(credentialServ)

	authMiddleware, authWrapper := middleware.NewAuthMiddleware()
	sub := router.Group("/solarman")
	sub.Use(authMiddleware())
	sub.Get("/credential", authWrapper(credentialHdl.FindAll))
	sub.Post("/credential", authWrapper(credentialHdl.Create))
	sub.Put("/credential/:id", authWrapper(credentialHdl.Update))
	sub.Delete("/credential/:id", authWrapper(credentialHdl.Delete))
}

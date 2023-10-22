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

func bindPrivateKStarAPI(router fiber.Router) {
	credentialRepo := repo.NewKStarCredentialRepo(infra.GormDB)
	credentialServ := service.NewKStarCredentialService(credentialRepo, logger.GetLogger())
	credentialHdl := handler.NewKStarCredentialHandler(credentialServ)

	authMiddleware, authWrapper := middleware.NewAuthMiddleware()
	sub := router.Group("/kstar")
	sub.Use(authMiddleware())
	sub.Get("/credential", authWrapper(credentialHdl.FindAll))
	sub.Post("/credential", authWrapper(credentialHdl.Create))
	sub.Put("/credential/:id", authWrapper(credentialHdl.Update))
	sub.Delete("/credential/:id", authWrapper(credentialHdl.Delete))
}

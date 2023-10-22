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
GET | /kibana/credentials
*/
func bindPrivateKibanaAPI(router fiber.Router) {
	credentialRepo := repo.NewKibanaCredentialRepo(infra.GormDB)
	credentialServ := service.NewKibanaCredentialService(credentialRepo, logger.GetLogger())
	credentialHdl := handler.NewKibanaCredentialHandler(credentialServ)

	authMiddleware, authWrapper := middleware.NewAuthMiddleware()
	sub := router.Group("/kibana")
	sub.Use(authMiddleware())
	sub.Get("/credential", authWrapper(credentialHdl.FindOne))
}

package api

import (
	"github.com/HavvokLab/true-solar-monitoring/handler"
	"github.com/HavvokLab/true-solar-monitoring/infra"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/repo"
	"github.com/HavvokLab/true-solar-monitoring/service"
	"github.com/gofiber/fiber/v2"
)

func bindPublicAuthRouter(router fiber.Router) {
	userRepo := repo.NewUserRepo(infra.GormDB)
	authServ := service.NewAuthService(userRepo, logger.GetLogger())
	hdl := handler.NewAuthHandler(authServ)

	sub := router.Group("/auth")
	sub.Post("/login", hdl.Login)
}

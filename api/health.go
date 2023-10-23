package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

/*
GET | /health
*/

func bindPublicHealthAPI(router fiber.Router) {
	router.Get("/health", monitor.New())
}

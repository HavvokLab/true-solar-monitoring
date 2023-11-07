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
	router.Get("/metrics", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(map[string]interface{}{
			"code":    200,
			"message": "OK",
		})
	})
}

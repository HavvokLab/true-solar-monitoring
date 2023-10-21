package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func InitAPI(app *fiber.App) {
	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(recover.New())

	// |=> API GROUP <=|
	api := app.Group("/api")

	// |=> Authentication
	bindPublicAuthRouter(api)
}

package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkaXNwbGF5X25hbWUiOiJzb2xhcmFkbSIsImV4cCI6MTY5Nzk2NDQxMSwianRpIjoiMlg0dWhNRjNHYVlGMTBqMDNBcHhhek42NTR4In0.bmLugqSMXZcma1D13CBOoI8I9ha6n8aRXpN4O2-_SHk

func InitAPI(app *fiber.App) {
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(recover.New())

	// |=> API GROUP <=|
	api := app.Group("/api")

	// |=> Authentication
	bindPublicAuthRouter(api)

	// |=> Huawei API
	bindPrivateHuaweiAPI(api)

	// |=> KStar API
	bindPrivateKStarAPI(api)

	// |=> Solarman API
	bindPrivateSolarmanAPI(api)

	// |=> Kibana API
	bindPrivateKibanaAPI(api)

	// |=> Installed Capacity API
	bindPrivateInstalledCapacityAPI(api)
}

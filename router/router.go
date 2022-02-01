package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	v1 "github.com/vinicarvalhosantos/fawkes-api/router/v1"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api", logger.New())

	//Performance Monitor Route
	app.Get("/dashboard", monitor.New())
	v1.SetupV1Routes(api)

}

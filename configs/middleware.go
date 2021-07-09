package configs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Middleware(app *fiber.App) {
	app.Use(
		cors.New(),
		logger.New(),
	)
}

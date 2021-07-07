package main

import (
	"log"

	"github.com/brenobaptista/library-go-fiber/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupRoutes(app *fiber.App) {
	route := app.Group("/api/todos")

	routes.PublicRoutes(route)
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

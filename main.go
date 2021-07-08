package main

import (
	"log"

	"github.com/brenobaptista/library-go-fiber/config"
	"github.com/brenobaptista/library-go-fiber/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	route := app.Group("/api/todos")

	routes.PublicRoutes(route)
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

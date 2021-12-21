package main

import (
	"fmt"
	"log"
	"os"

	"github.com/brenobaptista/go-todo-fiber/configs"
	_ "github.com/brenobaptista/go-todo-fiber/docs"
	"github.com/brenobaptista/go-todo-fiber/pkg/db"
	"github.com/brenobaptista/go-todo-fiber/pkg/routes"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

// @title To-do Go Fiber
// @version 1.0
// @description To-do API made using Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name Contact me
// @contact.email brenobaptista@protonmail.com
// @license.name MIT License
// @license.url https://mit-license.org/
// @host localhost:8080
// @BasePath /
func main() {
	app := fiber.New()

	configs.Middleware(app)

	db.ConnectDB()

	routes.SwaggerRoute(app)
	routes.TodoPublicRoutes(app)

	port := os.Getenv("PORT")
	err := app.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
	}
}

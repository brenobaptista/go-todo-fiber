package main

import (
	"log"

	"github.com/brenobaptista/library-go-fiber/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/todos", controllers.GetAllTodos)
	app.Get("/api/todos/:id", controllers.GetTodoByID)
	app.Post("/api/todos", controllers.CreateTodo)
	app.Patch("/api/todos/:id", controllers.ToggleTodoStatus)
	app.Delete("/api/todos/:id", controllers.DeleteTodo)
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

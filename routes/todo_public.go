package routes

import (
	"github.com/brenobaptista/go-todo-fiber/controllers"
	"github.com/gofiber/fiber/v2"
)

func TodoPublicRoutes(app *fiber.App) {
	route := app.Group("/api/todos")

	route.Get("", controllers.GetTodos)
	route.Get("/:id", controllers.GetTodo)
	route.Post("", controllers.CreateTodo)
	route.Put("/:id", controllers.UpdateTodo)
	route.Delete("/:id", controllers.DeleteTodo)
}

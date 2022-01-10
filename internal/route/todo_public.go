package route

import (
	"github.com/brenobaptista/go-todo-fiber/internal/controller"
	"github.com/gofiber/fiber/v2"
)

func TodoPublicRoutes(app *fiber.App) {
	route := app.Group("/api/todos")

	route.Get("", controller.GetTodos)
	route.Get("/:id", controller.GetTodo)
	route.Post("", controller.CreateTodo)
	route.Put("/:id", controller.UpdateTodo)
	route.Delete("/:id", controller.DeleteTodo)
}

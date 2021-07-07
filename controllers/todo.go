package controllers

import (
	"github.com/brenobaptista/library-go-fiber/models"
	"github.com/gofiber/fiber/v2"
)

var todos = []*models.Todo{
	{
		Id:        1,
		Title:     "Walk the dog",
		Completed: false,
	},
}

func GetTodos(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todos": todos,
		},
	})
}

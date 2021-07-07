package controllers

import "github.com/gofiber/fiber/v2"

func GetAllTodos(c *fiber.Ctx) error {
	return c.SendString("All Todos")
}

func GetTodoByID(c *fiber.Ctx) error {
	return c.SendString("Single Todo")
}

func CreateTodo(c *fiber.Ctx) error {
	return c.SendString("New Todo")
}

func ToggleTodoStatus(c *fiber.Ctx) error {
	return c.SendString("Updated Todo")
}

func DeleteTodo(c *fiber.Ctx) error {
	return c.SendString("Delete Todo")
}

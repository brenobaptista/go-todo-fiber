package controllers

import "github.com/gofiber/fiber/v2"

func GetAllTodos(c *fiber.Ctx) error {
	return c.SendString("Got all to-dos")
}

func GetTodoByID(c *fiber.Ctx) error {
	return c.SendString("Got a single to-do")
}

func CreateTodo(c *fiber.Ctx) error {
	return c.SendString("Created to-do")
}

func ToggleTodoStatus(c *fiber.Ctx) error {
	return c.SendString("Updated to-do")
}

func DeleteTodo(c *fiber.Ctx) error {
	return c.SendString("Deleted to-do")
}

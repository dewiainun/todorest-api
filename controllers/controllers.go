package controllers

import (
	"github.com/dewi/todo-app/database"
	"github.com/dewi/todo-app/models"
	"github.com/gofiber/fiber/v2"
)

// GEt all todos
func GetTodos(c *fiber.Ctx) error {
	var todos []models.ToDo
	database.DB.Find(&todos)
	return c.JSON(todos)
}

// get  spesific todo by ID
func GetTodosById(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo models.ToDo
	result := database.DB.First(&todo, id)

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Todo not found"})
	}
	return c.JSON(todo)
}

// create a new todo
func CreateTodo(c *fiber.Ctx) error {
	todo := new(models.ToDo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	database.DB.Create(todo)
	return c.JSON(todo)
}

// update an existing todo by ID
func UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo models.ToDo

	result := database.DB.First(&todo, id)

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Todo not found"})
	}

	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	database.DB.Save(todo)
	return c.JSON(todo)
}

// delete an existing todo by ID
func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo models.ToDo

	result := database.DB.First(&todo, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Todo not found"})
	}
	database.DB.Delete(&todo)
	return c.SendStatus(fiber.StatusNoContent)
}

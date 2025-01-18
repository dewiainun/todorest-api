package routes

import (
	"github.com/dewi/todo-app/controllers"
	"github.com/gofiber/fiber/v2"
)

func TodoRoutes(app *fiber.App) {
	app.Get("/todos", controllers.GetTodos)
	app.Get("/todos/:id", controllers.GetTodosById)
	app.Post("/todos", controllers.CreateTodo)
	app.Put("/todos/:id", controllers.UpdateTodo)
	app.Delete("/todos/:id", controllers.DeleteTodo)
}

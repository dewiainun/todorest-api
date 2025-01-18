package main

import (
  "log"
  "github.com/dewi/todo-app/database"
  "github.com/gofiber/fiber/v2"
  "github.com/dewi/todo-app/routes"
)

func main() {
    database.Connect()

    app := fiber.New()

    routes.TodoRoutes(app)

    log.Fatal(app.Listen(":8000"))
}
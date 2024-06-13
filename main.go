package main

import (
	"log"

	"github.com/Jahn16/habitshare/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/habits", handlers.HabitList)
	app.Post("/habits", handlers.HabitCreate)

	log.Fatal(app.Listen(":8000"))
}

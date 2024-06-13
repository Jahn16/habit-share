package main

import (
	"log"

	"github.com/Jahn16/habitshare/database"
	"github.com/Jahn16/habitshare/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db, _ := database.Setup()

	app.Get("/habits", handlers.HabitList(db))
	app.Post("/habits", handlers.HabitCreate(db))

	log.Fatal(app.Listen(":8000"))
}

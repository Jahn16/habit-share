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
	app.Get("/habits/:id", handlers.HabitGet(db))
	app.Post("/habits", handlers.HabitCreate(db))
	app.Post("/habits/:id/record", handlers.RecordHabit(db))

	app.Get("/users", handlers.UserList(db))
	app.Post("/users", handlers.UserCreate(db))
	app.Get("/users/:id", handlers.UserGet(db))

	log.Fatal(app.Listen(":8000"))
}

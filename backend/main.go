package main

import (
	"log"

	"github.com/Jahn16/habitshare/database"
	"github.com/Jahn16/habitshare/handlers"

	"os"

	"github.com/Jahn16/habitshare/middlewares/googleauth"
	"github.com/gofiber/fiber/v2"

	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	db, _ := database.Setup()
	godotenv.Load()
	authMiddleware := googleauth.New(googleauth.Config{GoogleID: os.Getenv("AUTH_GOOGLE_ID")})

	app.Get("/habits", handlers.HabitList(db))
	app.Get("/habits/:id", handlers.HabitGet(db))
	app.Post("/habits", authMiddleware, handlers.HabitCreate(db))
	app.Post("/habits/:id/record", authMiddleware, handlers.RecordHabit(db))

	app.Get("/users", handlers.UserList(db))
	app.Post("/users", authMiddleware, handlers.UserCreate(db))
	app.Get("/users/:id", handlers.UserGet(db))

	log.Fatal(app.Listen(":8000"))
}

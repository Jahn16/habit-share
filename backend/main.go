package main

import (
	"log"

	"github.com/Jahn16/socialhabits/database"
	"github.com/Jahn16/socialhabits/handlers"

	"os"

	"github.com/Jahn16/socialhabits/middlewares/auth0"
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	godotenv.Load()
	db, _ := database.New(os.Getenv("DB_DSN"))
	authMiddleware := auth0.New(auth0.Config{Issuer: os.Getenv("AUTH_AUTH0_ISSUER"), Audience: []string{os.Getenv("AUTH_AUTH0_ID")}})

	app.Get("/habits", handlers.HabitList(db))
	app.Get("/habits/:id", handlers.HabitGet(db))
	app.Post("/habits", authMiddleware, handlers.HabitCreate(db))
	app.Delete("/habits/:id", authMiddleware, handlers.DeleteHabit(db))
	app.Patch("/habits/:id", authMiddleware, handlers.UpdateHabit(db))

	app.Post("/habits/:id/record", authMiddleware, handlers.RecordHabit(db))
	app.Delete("/records/:id", authMiddleware, handlers.DeleteRecord(db))

	app.Get("/users", handlers.UserList(db))
	app.Post("/users", authMiddleware, handlers.UserCreate(db))

	app.Get("/users/me", authMiddleware, handlers.GetAuthenticatedUser(db))
	app.Patch("/users/me", authMiddleware, handlers.UserUpdate(db))
	app.Post("/users/me/friends", authMiddleware, handlers.AddFriend(db))
	app.Delete("/users/me/friends", authMiddleware, handlers.RemoveFriend(db))
	app.Get("/users/:slug", handlers.UserGet(db))

	log.Fatal(app.Listen(":8000"))
}

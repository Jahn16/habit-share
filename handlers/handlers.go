package handlers

import (
	"github.com/Jahn16/habitshare/database"
	"github.com/Jahn16/habitshare/models"
	"github.com/gofiber/fiber/v2"
)

func HabitList(c *fiber.Ctx) error {
	habits := database.List()
	return c.JSON(fiber.Map{
		"success": true,
		"value":   habits,
	})
}

func HabitCreate(c *fiber.Ctx) error {
	habit := new(models.Habit)
	if err := c.BodyParser(habit); err != nil {
		return err
	}
	database.Insert(habit)
	return c.JSON(fiber.Map{
		"sucess": true,
		"value":  habit,
	})
}

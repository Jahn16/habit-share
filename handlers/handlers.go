package handlers

import (
	"github.com/Jahn16/habitshare/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func HabitList(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var habits []models.Habit
		db.Find(&habits)
		return c.JSON(fiber.Map{
			"success": true,
			"value":   habits,
		})

	}
}

func HabitCreate(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		habit := new(models.Habit)
		if err := c.BodyParser(habit); err != nil {
			return err
		}
		db.Create(&habit)
		return c.JSON(fiber.Map{
			"sucess": true,
			"value":  habit,
		})
	}
}

func HabitGet(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var habit models.Habit
		db.First(&habit, c.Params("id"))
		return c.JSON(fiber.Map{
			"success": true,
			"value":   habit,
		})
	}
}

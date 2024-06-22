package handlers

import (
	"github.com/Jahn16/habitshare/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserCreate(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		username := c.Locals("username").(string)
		email := c.Locals("email").(string)
		picture := c.Locals("picture").(string)
		user := models.User{
			Name:    username,
			Email:   email,
			Picture: picture,
		}
		db.Create(&user)
		return c.JSON(fiber.Map{
			"sucess": true,
			"value":  user,
		})

	}
}

func UserGet(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user models.User
		db.Preload("Habits").First(&user, c.Params("id"))
		return c.JSON(fiber.Map{
			"success": true,
			"value":   user,
		})
	}
}

func UserList(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var users []models.User
		db.Find(&users)
		return c.JSON(fiber.Map{
			"success": true,
			"value":   users,
		})

	}
}

func HabitList(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var habits []models.Habit
		db.Preload("Records").Find(&habits)
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
		db.Preload("Records").First(&habit, c.Params("id"))
		return c.JSON(fiber.Map{
			"success": true,
			"value":   habit,
		})
	}
}

func RecordHabit(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		habitRecord := new(models.HabitRecord)
		if err := c.BodyParser(habitRecord); err != nil {
			return err
		}
		habitID, err := c.ParamsInt("id")
		if err != nil {
			return err
		}
		habitRecord.HabitID = uint(habitID)
		db.Create(&habitRecord)
		return c.JSON(fiber.Map{
			"sucess": true,
			"value":  habitRecord,
		})
	}
}

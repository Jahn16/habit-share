package handlers

import (
	"github.com/Jahn16/socialhabits/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAuthenticatedUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		email := c.Locals("email").(string)
		var user models.User
		result := db.Where(&models.User{Email: email}).Preload("Habits.Records").First(&user)
		if result.Error != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"sucess": "false", "message": "User not found"})
		}
		return c.JSON(fiber.Map{
			"success": true,
			"value":   user,
		})
	}
}

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
		result := db.Preload("Habits.Records").First(&user, c.Params("id"))
		if result.Error != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"sucess": "false", "message": "User not found"})
		}
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

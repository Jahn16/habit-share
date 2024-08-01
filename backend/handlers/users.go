package handlers

import (
	"github.com/Jahn16/socialhabits/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAuthenticatedUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("id").(string)
		var user models.User
		result := db.Preload("Habits.Records").First(&user, "id = ?", userID)
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
		userID := c.Locals("id").(string)
		username := c.Locals("username").(string)
		picture := c.Locals("picture").(string)
		user := models.User{
			ID:      userID,
			Name:    username,
			Picture: picture,
		}
		db.Create(&user)
		return c.JSON(fiber.Map{
			"sucess": true,
			"value":  user,
		})

	}
}

func UserUpdate(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("id").(string)
		var user models.User
		result := db.Preload("Habits.Records").First(&user, "id = ?", userID)
		if result.Error != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"sucess": "false", "message": "User not found"})
		}
		updatedUser := new(models.User)
		if err := c.BodyParser(updatedUser); err != nil {
			return err
		}
		user.Name = updatedUser.Name
		user.ColorPalette = updatedUser.ColorPalette
		updateResult := db.Save(&user)
		if updateResult.Error != nil {
			return updateResult.Error
		}
		return c.JSON(fiber.Map{
			"success": true,
			"value":   user,
		})
	}
}

func UserGet(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Params("id")
		var user models.User
		result := db.Preload("Habits.Records").First(&user, "id = ?", userID)
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

func AddFriend(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("id").(string)
		var user models.User
		result := db.First(&user, "id = ?", userID)
		if result.Error != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"sucess": "false", "message": "User not found"})
		}

		data := new(
			struct {
				FriendID string `json:"friendID"`
			})
		if err := c.BodyParser(data); err != nil {
			return err
		}

		var friend models.User
		friendQuery := db.First(&friend, "id = ?", data.FriendID)
		if friendQuery.Error != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"sucess": "false", "message": "Friend not found"})
		}

		user.Friends = append(user.Friends, friend)
		updateResult := db.Save(&user)
		if updateResult.Error != nil {
			return updateResult.Error
		}
		return c.JSON(fiber.Map{
			"success": true,
			"value":   user,
		})
	}
}

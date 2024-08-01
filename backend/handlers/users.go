package handlers

import (
	"github.com/Jahn16/socialhabits/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

func GetAuthenticatedUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("id").(string)
		var user models.User
		result := db.Preload("Friends").Preload("Habits.Records").First(&user, "id = ?", userID)
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
		name := c.Locals("name").(string)
		userSlug := slug.Make(name)
		user := models.User{
			ID:      userID,
			Name:    username,
			Slug:    userSlug,
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
		userSlug := c.Params("slug")
		var user models.User
		result := db.Where("slug = ?", userSlug).Preload("Friends").Preload("Habits.Records").First(&user)
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
		if result := db.Preload("Friends").First(&user, "id = ?", userID); result.Error != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"sucess": "false", "message": "User not found"})
		}

		data := new(
			struct {
				FriendID string `json:"friendID"`
			},
		)
		if err := c.BodyParser(data); err != nil {
			return err
		}
		if user.ID == data.FriendID {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"sucess": "false", "message": "Cant add yourself as a friend"})
		}

		var friend models.User
		if result := db.First(&friend, "id = ?", data.FriendID); result.Error != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"sucess": "false", "message": "Friend not found"})
		}

		if err := db.Model(&user).Association("Friends").Append(&friend); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": "false", "message": "error"})
		}
		return c.JSON(fiber.Map{
			"success": true,
			"value":   user,
		})
	}
}

package googleauth

import (
	"context"
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/idtoken"
)

func extractToken(header string) (string, error) {
	if !strings.HasPrefix(header, "Bearer ") {
		return "", errors.New("missing or malformed token")
	}
	return header[7:], nil
}

func New(config Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorizationHeader := c.Get(fiber.HeaderAuthorization)
		token, err := extractToken(authorizationHeader)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
		}
		payload, err := idtoken.Validate(context.Background(), token, config.GoogleID)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString("invalid token")
		}
		c.Locals("username", payload.Claims["name"])
		c.Locals("email", payload.Claims["email"])
		c.Locals("picture", payload.Claims["picture"])
		return c.Next()
	}

}

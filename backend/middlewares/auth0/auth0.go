package auth0

import (
	"context"
	"errors"
	"net/url"
	"strings"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gofiber/fiber/v2"
)

func extractToken(header string) (string, error) {
	if !strings.HasPrefix(header, "Bearer ") {
		return "", errors.New("missing or malformed token")
	}
	return header[7:], nil
}

type CustomClaims struct {
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Picture  string `json:"picture"`
}

func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

func New(config Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorizationHeader := c.Get(fiber.HeaderAuthorization)
		token, err := extractToken(authorizationHeader)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
		}
		issuerURL, _ := url.Parse("https://" + config.Issuer + "/")
		provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

		jwtValidator, err := validator.New(
			provider.KeyFunc,
			validator.RS256,
			issuerURL.String(),
			config.Audience,
			validator.WithCustomClaims(
				func() validator.CustomClaims {
					return &CustomClaims{}
				},
			),
		)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
		}
		claims, err := jwtValidator.ValidateToken(c.Context(), token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
		}
		validatedClaims, ok := claims.(*validator.ValidatedClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).SendString("Invalid claims")
		}
		customClaims, ok := validatedClaims.CustomClaims.(*CustomClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).SendString("Invalid custom token")
		}

		c.Locals("email", customClaims.Email)
		c.Locals("username", customClaims.Nickname)
		c.Locals("picture", customClaims.Picture)

		return c.Next()
	}

}

package middleware

import (
	"fmt"
	"strings"

	"github.com/HavvokLab/true-solar-monitoring/config"
	"github.com/HavvokLab/true-solar-monitoring/domain"
	"github.com/HavvokLab/true-solar-monitoring/errors"
	"github.com/HavvokLab/true-solar-monitoring/util"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthWrapper func(func(*domain.UserContext, *fiber.Ctx) error) func(*fiber.Ctx) error
type AuthMiddleware func() func(*fiber.Ctx) error

func NewAuthMiddleware() (AuthMiddleware, AuthWrapper) {
	return authMiddleware, authWrapper
}

// |=> Function AuthMiddleware
func authMiddleware() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		conf := config.GetConfig().Authentication
		header := c.Get("Authorization")
		fmt.Println(header)
		if util.EmptyString(header) {
			fmt.Println("empty header")
			return util.ResponseUnauthorized(c)
		}

		raw := extractBearerToken(header)
		fmt.Println(raw)
		if util.EmptyString(raw) {
			fmt.Println("empty raw")
			return util.ResponseUnauthorized(c)
		}

		token, err := jwt.ParseWithClaims(raw, &domain.AccessToken{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(conf.Secret), nil
		})

		if err != nil {
			fmt.Println("error parse token", err)
			return util.ResponseUnauthorized(c)
		}

		claims, ok := token.Claims.(*domain.AccessToken)
		if !ok || !token.Valid {
			fmt.Println("invalid token")
			return util.ResponseUnauthorized(c)
		}

		c.Locals("user_context", domain.UserContext{ID: claims.ID, DisplayName: claims.DisplayName})
		return c.Next()
	}
}

func extractBearerToken(raw string) string {
	return strings.TrimSpace(strings.Replace(raw, "Bearer", "", 1))
}

// |=> Function AuthWrapper
func authWrapper(handler func(*domain.UserContext, *fiber.Ctx) error) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		utx, ok := c.Locals("user_context").(domain.UserContext)
		if !ok {
			err := errors.NewServerError(fiber.StatusUnauthorized, "empty user context")
			return util.ResponseError(c, err)
		}

		return handler(&utx, c)
	}
}

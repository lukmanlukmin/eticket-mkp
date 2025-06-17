// Package middleware ...
package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/lukmanlukmin/go-lib/util"
)

// JWT ...
func (m *Midleware) JWT() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		authHeaderLower := strings.ToLower(authHeader)
		if !strings.HasPrefix(authHeaderLower, "bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing or invalid Authorization header",
			})
		}
		tokenStr := strings.TrimSpace(authHeader[len("Bearer "):])

		claimData, err := util.ValidateJWT(m.Conf.Security.JWTSecret, tokenStr)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		for k, v := range claimData {
			if k != "exp" {
				c.Locals(k, v)
			}
		}
		return c.Next()
	}
}

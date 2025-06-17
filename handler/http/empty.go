// Package http ...
package http

import "github.com/gofiber/fiber/v2"

// EmptyHandler ...
func (h *Handler) EmptyHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "ok",
	})
}

// Package masterdata ...
package masterdata

import (
	"eticket/model/payload"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// CreateTerminal ...
// @Summary Create Terminal
// @Description Create Terminal
// @Tags Masterdata
// @Accept json
// @Produce json
// @Router /v1/masterdata/terminals [post]
// @Param payload body payload.CreateTerminalPayload true "Create Payload"
// @Success 200 {object} payload.BaseResponse{success=bool}
// @Response 400
// @Response 500
// @Security BearerAuth
func (h *Handler) CreateTerminal(c *fiber.Ctx) error {
	req := &payload.CreateTerminalPayload{}
	if err := c.BodyParser(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(payload.BaseResponse{
			Success: false,
			Error:   err.Error(),
		})
	}
	if err := h.Validate.Struct(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(payload.BaseResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	err := h.Service.MasterData.CreateTerminal(c.Context(), req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(payload.BaseResponse{
			Success: false,
			Error:   err.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(payload.BaseResponse{
		Success: true,
	})
}

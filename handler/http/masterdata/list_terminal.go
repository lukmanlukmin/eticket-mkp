// Package masterdata ...
package masterdata

import (
	"eticket/model/payload"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// ListTerminals ...
// @Summary Get list of Terminals
// @Description Retrieve list of Terminals with optional filters
// @Tags Masterdata
// @Accept json
// @Produce json
// @Router /v1/masterdata/terminals [get]
// @Param params query payload.BaseQueryFilter true "Pagination and Filter Request"
// @Success 200 {object} payload.BaseResponse{data=[]payload.TerminalResponse{},meta=payload.MetaResponse{}}
// @Failure 500
// @Security BearerAuth
func (h *Handler) ListTerminals(c *fiber.Ctx) error {
	order := c.Query("order", "created_at")
	direction := c.Query("direction", "desc")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("per_page", "10"))

	queryParam := payload.BaseQueryFilter{
		PaginationFilter: payload.PaginationFilter{
			Page:    uint64(page),
			PerPage: uint64(perPage),
		},
		OrderingFilter: payload.OrderingFilter{
			Order:     order,
			Direction: direction,
		},
	}

	res, count, err := h.Service.MasterData.ListTerminals(c.Context(), queryParam)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(payload.BaseResponse{
			Success: false,
			Error:   err,
		})
	}
	return c.Status(http.StatusOK).JSON(payload.BaseResponse{
		Success: true,
		Data:    res,
		Meta: &payload.MetaResponse{
			CurrentPage: queryParam.PaginationFilter.Page,
			Limit:       queryParam.PaginationFilter.PerPage,
			TotalData:   count,
		},
	})
}

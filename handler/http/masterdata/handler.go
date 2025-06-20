// Package masterdata ...
package masterdata

import (
	"eticket/bootstrap"

	"github.com/go-playground/validator/v10"
)

// Handler ...
type Handler struct {
	*bootstrap.Bootstrap
	Validate *validator.Validate
}

// NewHandler ...
func NewHandler(bs *bootstrap.Bootstrap) *Handler {
	return &Handler{
		Bootstrap: bs,
		Validate:  validator.New(),
	}
}

// Package terminals ...
package terminals

import (
	"context"
	model "eticket/model/db"

	"github.com/jmoiron/sqlx"
)

// Repository ...
type Repository struct {
	DB *sqlx.DB
}

// NewRepository ...
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

// Iterminal ...
type Iterminal interface {
	CreateTerminal(ctx context.Context, user *model.Terminal) error
	FindTerminalList(ctx context.Context, order, direction string, page, perPage uint64) ([]model.Terminal, uint64, error)
}

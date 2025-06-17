// Package users ...
package users

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

// Iuser ...
type Iuser interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByID(ctx context.Context, ID int64) (*model.User, error)
	GetUserByUsername(ctx context.Context, uname string) (*model.User, error)
}

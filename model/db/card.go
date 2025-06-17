// Package db ...
package db

import (
	"time"

	"github.com/google/uuid"
)

// Card ...
type Card struct {
	ID         uuid.UUID `db:"id"`
	CardNumber string    `db:"card_number"`
	Balance    int       `db:"balance"`
	LastSynced time.Time `db:"last_synced"`
	CreatedAt  time.Time `db:"created_at"`
}

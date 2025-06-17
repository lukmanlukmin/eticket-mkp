// Package db ...
package db

import (
	"time"

	"github.com/google/uuid"
)

// Terminal ...
type Terminal struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	Location  string    `db:"location"`
	CreatedAt time.Time `db:"created_at"`
}

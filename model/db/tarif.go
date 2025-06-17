// Package db ...
package db

import (
	"time"

	"github.com/google/uuid"
)

// Tariff ...
type Tariff struct {
	ID            uuid.UUID `db:"id"`
	EntryTerminal uuid.UUID `db:"entry_terminal_id"`
	ExitTerminal  uuid.UUID `db:"exit_terminal_id"`
	Fare          int       `db:"fare"`
	EffectiveDate time.Time `db:"effective_date"`
	CreatedAt     time.Time `db:"created_at"`
}

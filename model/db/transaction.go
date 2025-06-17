// Package db ...
package db

import (
	"time"

	"github.com/google/uuid"
)

// Transaction ...
type Transaction struct {
	ID            uuid.UUID `db:"id"`
	CardID        string    `db:"card_id"`
	EntryTerminal uuid.UUID `db:"entry_terminal_id"`
	ExitTerminal  uuid.UUID `db:"exit_terminal_id"`
	Fare          int       `db:"fare"`
	EntryTime     time.Time `db:"entry_time"`
	ExitTime      time.Time `db:"exit_time"`
	SourceDevice  string    `db:"source_device_id"`
	SyncedAt      time.Time `db:"synced_at"`
	CreatedAt     time.Time `db:"created_at"`
}

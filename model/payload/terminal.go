// Package payload ...
package payload

import (
	"time"

	"github.com/google/uuid"
)

// CreateTerminalPayload ...
type CreateTerminalPayload struct {
	Name     string `json:"name"`
	Location string `json:"alias"`
}

// TerminalResponse ...
type TerminalResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Location  string    `json:"alias"`
	CreatedAt time.Time `json:"created_at"`
}

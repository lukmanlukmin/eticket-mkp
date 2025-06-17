// Package masterdata ...
package masterdata

import (
	"context"
	"eticket/model/db"
	"eticket/model/payload"
)

// CreateTerminal ...
func (s *Service) CreateTerminal(ctx context.Context, terminal *payload.CreateTerminalPayload) error {
	return s.Repository.DB.Terminal.CreateTerminal(ctx, &db.Terminal{
		Name:     terminal.Name,
		Location: terminal.Location,
	})
}

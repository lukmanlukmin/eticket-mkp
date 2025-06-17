// Package masterdata ...
package masterdata

import (
	"context"
	"eticket/bootstrap/repository"
	"eticket/model/payload"
)

// Service ...
type Service struct {
	*repository.Repository
}

// NewService ...
func NewService(bs *repository.Repository) *Service {
	return &Service{
		Repository: bs,
	}
}

// IMasterData ...
type IMasterData interface {
	CreateTerminal(ctx context.Context, terminal *payload.CreateTerminalPayload) error
	ListTerminals(ctx context.Context, filter payload.BaseQueryFilter) ([]payload.TerminalResponse, uint64, error)
}

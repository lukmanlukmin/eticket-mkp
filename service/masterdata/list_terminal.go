// Package masterdata ...
package masterdata

import (
	"context"
	"eticket/model/payload"
)

// ListTerminals ...
func (s *Service) ListTerminals(ctx context.Context, filter payload.BaseQueryFilter) ([]payload.TerminalResponse, uint64, error) {
	res := []payload.TerminalResponse{}

	data, count, err := s.DB.Terminal.FindTerminalList(ctx, filter.Order, filter.Direction, filter.Page, filter.PerPage)
	if err != nil {
		return nil, 0, err
	}

	for _, t := range data {
		res = append(res, payload.TerminalResponse{
			ID:        t.ID,
			Name:      t.Name,
			Location:  t.Location,
			CreatedAt: t.CreatedAt,
		})
	}
	return res, count, nil
}

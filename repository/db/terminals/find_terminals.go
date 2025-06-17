// Package terminals ...
package terminals

import (
	"context"
	model "eticket/model/db"

	sq "github.com/Masterminds/squirrel"
	"github.com/lukmanlukmin/go-lib/database"
)

// FindTerminalList ...
func (r *Repository) FindTerminalList(ctx context.Context, order, direction string, page, perPage uint64) ([]model.Terminal, uint64, error) {
	var db database.SQLQueryExec = r.DB
	if tx := database.GetTxFromContext(ctx); tx != nil {
		db = tx
	}

	countSQ := sq.Select("COUNT(*)").
		From("terminals").
		PlaceholderFormat(sq.Dollar)

	countQuery, countArgs, err := countSQ.ToSql()
	if err != nil {
		return nil, 0, err
	}

	var total uint64
	if err := db.GetContext(ctx, &total, countQuery, countArgs...); err != nil {
		return nil, 0, err
	}

	dataSQ := sq.
		Select("id", "name", "location", "created_at").
		From("terminals")

	if order != "" && direction != "" {
		dataSQ = dataSQ.OrderBy(order + " " + direction)
	}
	if page > 0 && perPage > 0 {
		dataSQ = dataSQ.Offset((page - 1) * perPage).Limit(perPage)
	}

	dataQuery, dataArgs, err := dataSQ.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, 0, err
	}

	var data []model.Terminal
	if err = db.SelectContext(ctx, &data, dataQuery, dataArgs...); err != nil {
		return nil, 0, err
	}

	return data, total, nil
}

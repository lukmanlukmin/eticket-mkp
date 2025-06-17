// Package terminals ...
package terminals

import (
	model "eticket/model/db"

	"github.com/google/uuid"
	"github.com/lukmanlukmin/go-lib/database"

	"context"

	sq "github.com/Masterminds/squirrel"
)

// CreateTerminal ...
func (r *Repository) CreateTerminal(ctx context.Context, terminal *model.Terminal) error {

	var db database.SQLQueryExec = r.DB
	if tx := database.GetTxFromContext(ctx); tx != nil {
		db = tx
	}

	terminal.ID = uuid.New()
	query, args, err := sq.
		Insert("terminals").
		SetMap(sq.Eq{
			"id":         terminal.ID,
			"name":       terminal.Name,
			"location":   terminal.Location,
			"created_at": sq.Expr("NOW()"),
		}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}
	if _, err = db.ExecContext(ctx, query, args...); err != nil {
		return err
	}
	return nil
}

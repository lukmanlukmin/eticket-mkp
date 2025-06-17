// Package users ...
package users

import (
	model "eticket/model/db"

	"github.com/google/uuid"
	"github.com/lukmanlukmin/go-lib/database"

	"context"

	sq "github.com/Masterminds/squirrel"
)

// CreateUser ...
func (r *Repository) CreateUser(ctx context.Context, user *model.User) error {

	var db database.SQLQueryExec = r.DB
	if tx := database.GetTxFromContext(ctx); tx != nil {
		db = tx
	}

	user.ID = uuid.New()
	query, args, err := sq.
		Insert("users").
		SetMap(sq.Eq{
			"id":            user.ID,
			"username":      user.Username,
			"password_hash": user.Password,
			"created_at":    sq.Expr("NOW()"),
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

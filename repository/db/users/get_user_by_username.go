// Package users ...
package users

import (
	model "eticket/model/db"

	"github.com/lukmanlukmin/go-lib/database"

	"context"

	sq "github.com/Masterminds/squirrel"
)

// GetUserByUsername ...
func (r *Repository) GetUserByUsername(ctx context.Context, uname string) (*model.User, error) {

	var db database.SQLQueryExec = r.DB
	if tx := database.GetTxFromContext(ctx); tx != nil {
		db = tx
	}

	data := &model.User{}
	query, args, err := sq.
		Select("id", "username", "password_hash", "created_at").
		From("users").
		Where("username = ?", uname).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return data, err
	}
	if err = db.GetContext(ctx, data, query, args...); err != nil {
		return data, err
	}
	return data, nil
}

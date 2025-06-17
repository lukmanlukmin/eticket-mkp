// Package bootstrap ...
package bootstrap

import (
	"context"
	"database/sql"
	"errors"
	"eticket/model/db"
	"fmt"

	"github.com/lukmanlukmin/go-lib/log"
	"github.com/lukmanlukmin/go-lib/util"
)

// LoadDefaultData ...
func LoadDefaultData(bs *Bootstrap) {
	log.Info("make sure has default data")
	LoadDefaultUser(context.TODO(), bs)
	log.Info("done making sure has default data")
}

// DefaultUser ...
type DefaultUser struct {
	Username string
	Password string
}

// LoadDefaultUser ...
func LoadDefaultUser(ctx context.Context, bs *Bootstrap) {

	users := []DefaultUser{}
	for i := 0; i < 10; i++ {
		hashPasswd, err := util.HashPassword(fmt.Sprintf("password%d", i))
		if err != nil {
			log.WithError(err).Fatal("failed to prepare default data - 1", i)
		}
		users = append(users, DefaultUser{
			Username: fmt.Sprintf("user%d", i),
			Password: hashPasswd,
		})
	}

	for _, user := range users {
		dataUser, err := bs.Repository.DB.Users.GetUserByUsername(ctx, user.Username)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			log.WithError(err).Fatal("failed to prepare default data - 2", user.Username)
		}
		if dataUser == nil || errors.Is(err, sql.ErrNoRows) {
			err := bs.Repository.DB.Users.CreateUser(ctx, &db.User{
				Username: user.Username,
				Password: user.Password,
			})
			if err != nil {
				log.WithError(err).Fatal("failed to prepare default data - 3", user.Username)
			}
		}
	}
}

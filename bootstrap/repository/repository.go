// Package repository ...
package repository

import (
	"eticket/repository/db"
	terminalRepoDB "eticket/repository/db/terminals"
	userRepoDB "eticket/repository/db/users"

	connDB "github.com/lukmanlukmin/go-lib/database/connection"
)

// Repository ...
type Repository struct {
	DB db.Repository
}

// LoadRepository ...
func LoadRepository(connectionDB *connDB.Store) *Repository {
	return &Repository{
		DB: db.Repository{
			Users:    userRepoDB.NewRepository(connectionDB.GetMaster()),
			Terminal: terminalRepoDB.NewRepository(connectionDB.GetMaster()),
		},
	}
}

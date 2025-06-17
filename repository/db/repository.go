// Package db ...
package db

import (
	"eticket/repository/db/terminals"
	"eticket/repository/db/users"
)

// Repository ...
type Repository struct {
	Users    users.Iuser
	Terminal terminals.Iterminal
}

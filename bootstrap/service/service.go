// Package service ...
package service

import (
	"eticket/bootstrap/repository"
	"eticket/config"
	"eticket/service/auth"
	"eticket/service/masterdata"
)

// Service ...
type Service struct {
	Auth       auth.IAuth
	MasterData masterdata.IMasterData
}

// LoadServices ...
func LoadServices(repo *repository.Repository, conf *config.Config) *Service {
	return &Service{
		Auth:       auth.NewService(repo, conf),
		MasterData: masterdata.NewService(repo),
	}
}

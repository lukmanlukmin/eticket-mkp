// Package auth ...
package auth

import (
	"context"
	"eticket/bootstrap/repository"
	"eticket/config"
	"eticket/model/payload"
)

// Service ...
type Service struct {
	*repository.Repository
	Conf *config.Config
}

// NewService ...
func NewService(bs *repository.Repository, conf *config.Config) *Service {
	return &Service{
		Repository: bs,
		Conf:       conf,
	}
}

// IAuth ...
type IAuth interface {
	ValidateUserByCredential(ctx context.Context, data payload.LoginCredential) (*payload.Token, error)
}

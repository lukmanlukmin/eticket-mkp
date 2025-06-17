package middleware

import (
	"eticket/bootstrap"
	"eticket/config"
)

// Midleware ...
type Midleware struct {
	Bs   *bootstrap.Bootstrap
	Conf *config.Config
}

// New ...
func New(bs *bootstrap.Bootstrap, conf *config.Config) *Midleware {
	return &Midleware{
		Bs:   bs,
		Conf: conf,
	}
}

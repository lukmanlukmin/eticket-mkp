// Package bootstrap ...
package bootstrap

import (
	"eticket/bootstrap/repository"
	"eticket/bootstrap/service"
	"eticket/config"

	_ "github.com/lukmanlukmin/go-lib/cache/redis"
	connDB "github.com/lukmanlukmin/go-lib/database/connection"
)

// Bootstrap ...
type Bootstrap struct {
	Repository *repository.Repository
	Service    *service.Service
}

// NewBootstrap ...
func NewBootstrap(conf *config.Config) *Bootstrap {
	connectionDB := connDB.New(connDB.DBConfig{
		MasterDSN:     conf.PostgreSQLConfig.DSN,
		EnableSlave:   false,
		RetryInterval: conf.PostgreSQLConfig.RetryInterval,
		MaxIdleConn:   conf.PostgreSQLConfig.MaxIdleConn,
		MaxConn:       conf.PostgreSQLConfig.MaxConn,
	}, connDB.DriverPostgres)

	repo := repository.LoadRepository(connectionDB)
	svc := service.LoadServices(repo, conf)
	bs := &Bootstrap{
		Repository: repo,
		Service:    svc,
	}
	LoadDefaultData(bs)
	return bs
}

package cmd

import (
	"context"
	"eticket/config"
	"eticket/server"
)

func StartHttp(ctx context.Context, cfg *config.Config) {
	api := server.NewHTTPApi(cfg)
	api.Run()
}

// Package cmd ...
package cmd

import (
	"context"
	"eticket/config"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

// Start ...
func Start() {
	ctx, cancel := context.WithCancel(context.Background())
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-quit
		cancel()
	}()

	rootCmd := &cobra.Command{}
	serveHTTPCmd := &cobra.Command{
		Use:   "serve-http",
		Short: "Run HTTP Server",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := &config.Config{}
			err := config.ReadModuleConfig(cfg, "") /// set default config // working on it later
			if err == nil {
				StartHttp(ctx, cfg)
			}
		},
	}
	rootCmd.AddCommand(serveHTTPCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

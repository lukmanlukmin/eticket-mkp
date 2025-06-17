// Package config ...
package config

import (
	"eticket/constant"
	"log"

	fileConfig "github.com/lukmanlukmin/go-lib/file"
)

// Config ...
type Config struct {
	// ServiceName      string         `yaml:"ServiceName"`
	Server           ServerConfig   `yaml:"Server"`
	Security         Security       `yaml:"Security"`
	PostgreSQLConfig PostgresConfig `yaml:"PostgresDBConfig"`
}

type (
	// ServerConfig Configuration
	ServerConfig struct {
		HTTPPort string `yaml:"HttpPort"`
	}

	// Security Configuration
	Security struct {
		JWTSecret   string `yaml:"JWTSecret"`
		JWTDuration string `yaml:"JWTDuration"`
	}

	// PostgresConfig Configuration
	PostgresConfig struct {
		DSN           string `yaml:"DSN"`
		RetryInterval int    `yaml:"RetryInterval"`
		MaxIdleConn   int    `yaml:"MaxIdleConn"`
		MaxConn       int    `yaml:"MaxConn"`
	}
)

func ReadModuleConfig(cfg interface{}, filePath string) error {
	if filePath != "" {
		err := fileConfig.ReadConfig(cfg, filePath)
		if err != nil {
			log.Fatalf("failed to read config. %v", err)
		}
		return nil
	}
	err := fileConfig.ReadConfig(cfg, constant.DEFAULT_CONFIG_FILE)
	if err != nil {
		log.Fatalf("failed to read config. %v", err)
	}
	return nil
}

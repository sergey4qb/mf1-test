package config

import (
	"log"
	"os"
	"sync"
)

type Config struct {
	GRPCProtocol string
	GRPCAddress  string
	GRPCPort     string
}

var (
	cfg      *Config
	loadOnce sync.Once
)

func LoadConfig() *Config {
	loadOnce.Do(func() {
		cfg = &Config{
			GRPCProtocol: os.Getenv("GRPC_PROTOCOL"),
			GRPCAddress:  os.Getenv("GRPC_ADDRESS"),
			GRPCPort:     os.Getenv("GRPC_PORT"),
		}
		if cfg.GRPCProtocol == "" {
			log.Fatalf("ERROR: GRPC_PROTOCOL not passed")
		}
		if cfg.GRPCAddress == "" {
			log.Fatalf("ERROR: GRPC_ADDRESS not passed")
		}
		if cfg.GRPCPort == "" {
			log.Fatalf("ERROR: GRPC_PORT not passed")
		}
	})

	return cfg
}

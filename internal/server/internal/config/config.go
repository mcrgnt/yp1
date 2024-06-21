package config

import (
	"flag"
	"fmt"
	"time"

	"github.com/caarlos0/env/v11"
)

type ConfigPrased struct {
	StoreInterval time.Duration
}

type Config struct {
	Address         string `env:"ADDRESS"`
	StorageType     string `env:"MEMORY"`
	StoreInterval   string `env:"STORE_INTERVAL"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
	Restore         bool   `env:"RESTORE"`
	Parsed          *ConfigPrased
}

func (t *Config) paramsParseEnv() error {
	err := env.Parse(t)
	if err != nil {
		return fmt.Errorf("parse env: %w", err)
	}
	return nil
}

func (t *Config) paramsParseFlag() {
	flag.StringVar(&t.Address, "a", "localhost:8080", "")
	flag.StringVar(&t.StoreInterval, "i", "300", "")
	flag.StringVar(&t.FileStoragePath, "f", "/tmp/metrics-db.json", "")
	flag.BoolVar(&t.Restore, "r", true, "")
	flag.Parse()
}

func (t *Config) validate() error {
	if storeInterval, err := time.ParseDuration(t.StoreInterval + "s"); err != nil {
		return fmt.Errorf("parse duration failed: %w", err)
	} else {
		t.Parsed.StoreInterval = storeInterval
	}
	return nil
}

func NewConfig() (*Config, error) {
	config := &Config{
		Parsed: &ConfigPrased{},
	}
	config.paramsParseFlag()
	if err := config.paramsParseEnv(); err != nil {
		return nil, fmt.Errorf("parse env failed: %w", err)
	}
	if err := config.validate(); err != nil {
		return nil, fmt.Errorf("validate failed: %w", err)
	}
	return config, nil
}

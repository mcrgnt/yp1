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
	Parsed          *ConfigPrased
	Address         string `env:"ADDRESS"`
	StorageType     string `env:"MEMORY"`
	StoreInterval   string `env:"STORE_INTERVAL"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
	Restore         bool   `env:"RESTORE"`
	DataBaseDSN     string `env:"DATABASE_DSN"`
}

func (t *Config) paramsParseEnv() error {
	err := env.Parse(t)
	if err != nil {
		return fmt.Errorf("parse env: %w", err)
	}
	return nil
}

func (t *Config) paramsParseFlag() {
	flag.StringVar(&t.Address, "a", "localhost:8080", "The interface and port that the server listens on.")
	flag.StringVar(&t.StoreInterval, "i", "300", "The interval for saving metrics to disk.")
	flag.StringVar(&t.FileStoragePath, "f", "/tmp/metrics-db.json", "Path to the file to save the metrics.")
	flag.BoolVar(&t.Restore, "r", true, "Whether to read metrics from a file when the server starts.")
	flag.StringVar(&t.DataBaseDSN, "d", "", "DSN line for connecting to the database.")
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

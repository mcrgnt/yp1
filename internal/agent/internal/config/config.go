package config

import (
	"flag"
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Address        string `env:"ADDRESS"`
	StorageType    string `env:"MEMORY"`
	PollInterval   string `env:"POLL_INTERVAL"`
	ReportInterval string `env:"REPORT_INTERVAL"`
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
	flag.StringVar(&t.PollInterval, "p", "2", "")
	flag.StringVar(&t.ReportInterval, "r", "10", "")
	flag.Parse()
}

func NewConfig() (config *Config, err error) {
	config = &Config{}
	config.paramsParseFlag()
	err = config.paramsParseEnv()
	return
}

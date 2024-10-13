package config

import (
	"flag"
	"os"
	"path/filepath"
)

const (
	RunAddress           = "RUN_ADDRESS"
	DatabaseURI          = "DATABASE_URI"
	AccrualSystemAddress = "ACCRUAL_SYSTEM_ADDRESS"
	MigrationsDir        = "migrations"
)

var Env = map[string]string{}

type Config struct {
	RunAddress           string
	DatabaseURI          string
	AccrualSystemAddress string
	MigrationsDir        string
}

func (config *Config) Init() {
	flag.StringVar(&config.RunAddress, "a", "", "host and port of server")

	flag.StringVar(&config.DatabaseURI, "d", "", "connection address of database")

	flag.StringVar(&config.AccrualSystemAddress, "r", "", "accrual system address")

	flag.Parse()

	runAddress := os.Getenv(RunAddress)
	databaseURI := os.Getenv(DatabaseURI)
	accrualSystemAddress := os.Getenv(AccrualSystemAddress)

	if config.RunAddress == "" {
		config.RunAddress = runAddress
	}

	if config.DatabaseURI == "" {
		config.DatabaseURI = databaseURI
	}

	if config.AccrualSystemAddress == "" {
		config.AccrualSystemAddress = accrualSystemAddress
	}
}

func (config *Config) GetDatabaseURI() string {
	return config.DatabaseURI
}

func (config *Config) GetAccrualSystemAddress() string {
	return config.AccrualSystemAddress
}

func (config *Config) GetRunAddress() string { return config.RunAddress }

func (config *Config) GetMigrationsDir() string { return config.MigrationsDir }

func NewConfig() *Config {
	path, _ := filepath.Abs(MigrationsDir)

	return &Config{
		MigrationsDir: path,
	}
}

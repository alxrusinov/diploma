package config

import (
	"flag"
	"fmt"
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
}

func (config *Config) Parse() {
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

func CreateConfig() *Config {
	path, _ := filepath.Abs(MigrationsDir)
	fmt.Printf("%#v\n", path)
	return &Config{
		MigrationsDir: path,
	}
}

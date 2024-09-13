package config

import (
	"flag"
	"os"
)

const (
	RunAddress           = "RUN_ADDRESS"
	DatabaseURI          = "DATABASE_URI"
	AccrualSystemAddress = "ACCRUAL_SYSTEM_ADDRESS"
)

var Env = map[string]string{}

type Config struct {
	RunAddress           string
	DatabaseURI          string
	AccrualSystemAddress string
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
	return &Config{}
}

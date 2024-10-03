package config

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/alxrusinov/diploma/internal/authenticate"
	"github.com/alxrusinov/diploma/internal/handler"
	"github.com/alxrusinov/diploma/internal/migrator"
	"github.com/alxrusinov/diploma/internal/server"
	"github.com/alxrusinov/diploma/internal/store"
	"github.com/alxrusinov/diploma/internal/usecase"
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

func (config *Config) Run() {
	migratorInst := migrator.NewMigrator()
	store := store.NewStore(config.DatabaseURI, migratorInst)

	authClient := authenticate.NewAuth()
	uc := usecase.NewUsecase(store, config.AccrualSystemAddress)
	router := handler.NewHandler(uc, config.AccrualSystemAddress, authClient)
	server := server.NewServer(router, config.RunAddress)

	err := store.RunMigration()

	if err != nil {
		log.Fatal(err)
	}

	server.Run()
}

func NewConfig() *Config {
	path, _ := filepath.Abs(MigrationsDir)

	return &Config{
		MigrationsDir: path,
	}
}

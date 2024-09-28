package app

import (
	"log"

	"github.com/alxrusinov/diploma/internal/authenticate"
	"github.com/alxrusinov/diploma/internal/config"
	"github.com/alxrusinov/diploma/internal/handler"
	"github.com/alxrusinov/diploma/internal/migrator"
	"github.com/alxrusinov/diploma/internal/server"
	"github.com/alxrusinov/diploma/internal/store"
	"github.com/alxrusinov/diploma/internal/use"
)

type App struct {
	Config *config.Config
}

func (app *App) Init() {
	app.Config.Init()
}

func (app *App) Run() {
	app.Config.Parse()

	migratorInst := migrator.CreateMigrator()
	store := store.CreateDBStore(app.Config.DatabaseURI, migratorInst)
	authClient := authenticate.CreateAuth()
	uc := use.CreateUsecase(store)
	router := handler.CreateHandler(uc, app.Config.AccrualSystemAddress, authClient)
	server := server.CreateServer(router, app.Config.RunAddress)

	err := store.RunMigration()

	if err != nil {
		log.Fatal(err)
	}

	server.Run()
}

func CreateApp() *App {
	var appConfig = config.CreateConfig()
	return &App{Config: appConfig}
}

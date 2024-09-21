package app

import (
	"github.com/alxrusinov/diploma/internal/auth"
	"github.com/alxrusinov/diploma/internal/config"
	"github.com/alxrusinov/diploma/internal/handler"
	"github.com/alxrusinov/diploma/internal/server"
	"github.com/alxrusinov/diploma/internal/store"
)

type App struct {
	Config *config.Config
}

func (app *App) Init() {
	app.Config.Init()
}

func (app *App) Run() {
	app.Config.Parse()

	store := store.CreateDBStore(app.Config.DatabaseURI)
	authClient := auth.CreateAuth()
	router := handler.CreateHandler(store, app.Config.AccrualSystemAddress, authClient)
	server := server.CreateServer(router, app.Config.RunAddress)

	server.Run()
}

func CreateApp() *App {
	var appConfig *config.Config = config.CreateConfig()
	return &App{Config: appConfig}
}

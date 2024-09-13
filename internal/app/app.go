package app

import "github.com/alxrusinov/diploma/internal/config"

type App struct {
	Config *config.Config
}

func (app *App) Init() {
	app.Config.Init()
}

func (app *App) Run() {
	app.Config.Parse()
}

func CreateApp() *App {
	var appConfig *config.Config = config.CreateConfig()
	return &App{Config: appConfig}
}

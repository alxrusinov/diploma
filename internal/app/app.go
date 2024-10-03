package app

import (
	"github.com/alxrusinov/diploma/internal/config"
)

type App struct {
	Config Config
}

type Config interface {
	Init()
	Run()
}

func (app *App) Run() {
	app.Config.Init()
	app.Config.Run()
}

func NewApp() *App {
	var appConfig = config.NewConfig()
	return &App{Config: appConfig}
}

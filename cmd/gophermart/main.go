package main

import "github.com/alxrusinov/diploma/internal/app"

var App = app.CreateApp()

func init() {
	App.Init()
}

func main() {
	App.Run()
}

package main

import "github/gyu-young-park/go-archive/app"

func main() {
	app := app.NewApp()
	app.Ready()
	app.Start()
}

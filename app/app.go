package app

import "github/gyu-young-park/go-archive/server"

type app struct {
	httpServer *server.WebEngine
}

func NewApp() *app {
	return &app{
		httpServer: server.NewWebEngine(),
	}
}

func (a *app) Ready() {
	a.httpServer.Register(server.NewbasicRouter())
}

func (a *app) Start() {
	a.httpServer.Run()
}

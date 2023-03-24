package app

import (
	"github/gyu-young-park/go-archive/server"
	"github/gyu-young-park/go-archive/worker"
)

type app struct {
	httpServer *server.WebEngine
	worker     *worker.Service
}

func NewApp() *app {
	return &app{
		httpServer: server.NewWebEngine(),
		worker:     worker.NewWorker(),
	}
}

func (a *app) Ready() {
	a.worker.Execute()
	a.httpServer.Register(server.NewbasicRouter())
}

func (a *app) Start() {
	a.httpServer.Run()
}

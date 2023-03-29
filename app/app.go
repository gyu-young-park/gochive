package app

import (
	api "github/gyu-young-park/go-archive/api/post"
	"github/gyu-young-park/go-archive/repository"
	"github/gyu-young-park/go-archive/server"
	"github/gyu-young-park/go-archive/worker"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type app struct {
	httpServer *server.WebEngine
	worker     *worker.Service
	Store      *repository.Storer
}

func NewApp() *app {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	dsn := os.Getenv("DB_DSN")

	store := repository.NewStorer(dsn)
	return &app{
		httpServer: server.NewWebEngine(),
		worker:     worker.NewWorker(store),
		Store:      store,
	}
}

func (a *app) Ready() {
	a.worker.Execute()
	a.httpServer.Register(server.NewbasicRouter(a.Store))
	a.httpServer.Register(api.NewPostRouter(a.Store))
}

func (a *app) Start() {
	defer a.Store.Close()
	a.httpServer.Run()
}

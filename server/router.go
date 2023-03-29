package server

import (
	"github/gyu-young-park/go-archive/repository"

	"github.com/gin-gonic/gin"
)

type Router interface {
	SetupRoutes(*gin.Engine)
}

type basicRouter struct {
	basicHandler *basicHandler
}

func NewbasicRouter(storer *repository.Storer) *basicRouter {
	r := &basicRouter{}
	r.basicHandler = newBasicHandler(storer)
	return r
}

func (b *basicRouter) SetupRoutes(core *gin.Engine) {
	core.GET("", b.basicHandler.greeting)
	core.GET("healthcheck", b.basicHandler.healthcheck)
}

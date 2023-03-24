package server

import "github.com/gin-gonic/gin"

type Router interface {
	SetupRoutes(*gin.Engine)
}

type basicRouter struct {
	basicHandler *basicHandler
}

func NewbasicRouter() *basicRouter {
	r := &basicRouter{}
	r.basicHandler = &basicHandler{}
	return r
}

func (b *basicRouter) SetupRoutes(core *gin.Engine) {
	core.GET("", b.basicHandler.greeting)
	core.GET("healthcheck", b.basicHandler.healthcheck)
}

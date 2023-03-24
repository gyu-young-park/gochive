package medium

import "github.com/gin-gonic/gin"

type aPIRouter struct {
}

func NewAPIRouter() *aPIRouter {
	return &aPIRouter{}
}

func (a *aPIRouter) SetupRoutes(core *gin.Engine) {

}

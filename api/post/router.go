package api

import (
	"github/gyu-young-park/go-archive/repository"

	"github.com/gin-gonic/gin"
)

type postRouter struct {
	postHandler *postHandler
}

func NewPostRouter(storer *repository.Storer) *postRouter {
	r := &postRouter{}
	r.postHandler = newPostHandler(storer)
	return r
}

func (b *postRouter) SetupRoutes(core *gin.Engine) {
	post := core.Group(PREFIX_POST_ROUTE)
	post.GET("", b.postHandler.getLastestPostsHandler)
	post.GET("/latest", b.postHandler.getLastestPostHandler)
}

package api

import (
	"fmt"
	"github/gyu-young-park/go-archive/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type postHandler struct {
	store *repository.Storer
}

func newPostHandler(storer *repository.Storer) *postHandler {
	return &postHandler{store: storer}
}

func (p *postHandler) getLastestPostHandler(c *gin.Context) {
	var requestPostQueryParam RequestPostQueryParam
	if err := c.ShouldBind(&requestPostQueryParam); err != nil {
		fmt.Println(err)
		c.String(http.StatusBadRequest, "Please check your 'origin' route path")
		return
	}

	post, err := p.store.RetriveLatestPostInDB(requestPostQueryParam.Origin)

	if err != nil {
		fmt.Println(err)
		c.String(http.StatusBadRequest, "Please check your query params")
		return
	}
	c.JSON(http.StatusOK, post)
}

func (p *postHandler) getLastestPostsHandler(c *gin.Context) {
	var requestPostQueryParam RequestPostQueryParam
	if err := c.ShouldBind(&requestPostQueryParam); err != nil {
		fmt.Println(err)
		c.String(http.StatusBadRequest, "Please check your 'origin' route path")
		return
	}
	id, err := strconv.ParseInt(requestPostQueryParam.Id, 0, 64)
	if err != nil {
		post, _ := p.store.RetriveLatestPostInDB(requestPostQueryParam.Origin)
		id = int64(post.Id)
	}
	limit, err := strconv.ParseInt(requestPostQueryParam.Limit, 0, 64)
	if err != nil || limit < 0 {
		limit = 10
	}

	posts, err := p.store.RetriveLatestPostsInDB(requestPostQueryParam.Origin, int(id), int(limit))

	if err != nil {
		fmt.Println(err)
		c.String(http.StatusBadRequest, "Please check your query params")
		return
	}

	c.JSON(http.StatusOK, posts)
}

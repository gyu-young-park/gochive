package server

import (
	"github/gyu-young-park/go-archive/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type basicHandler struct {
	store *repository.Storer
}

func newBasicHandler(storer *repository.Storer) *basicHandler {
	return &basicHandler{store: storer}
}

func (b *basicHandler) greeting(c *gin.Context) {
	c.String(http.StatusOK, "Hello! World")
}

func (b *basicHandler) healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

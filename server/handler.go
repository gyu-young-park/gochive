package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type basicHandler struct {
}

func NewBasicHandler() *basicHandler {
	return &basicHandler{}
}

func (b *basicHandler) healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type WebEngine struct {
	core    *gin.Engine
	routers []Router
}

func NewWebEngine() *WebEngine {
	core := gin.Default()

	return &WebEngine{
		core: core,
	}
}

func (s *WebEngine) Register(router Router) {
	s.routers = append(s.routers, router)
}

func (s *WebEngine) Run() {
	for _, router := range s.routers {
		router.SetupRoutes(s.core)
	}
	server := &http.Server{
		Addr:           port,
		Handler:        s.core,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf("Start server: %v\n", port)
	server.ListenAndServe()
}

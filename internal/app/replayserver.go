package app

import (
	"cqrs-practise/internal/app/adapter/ctrl/healthz"
	"cqrs-practise/internal/cfg"

	"github.com/gin-gonic/gin"
)

// Server is the server
type Server struct {
	*gin.Engine
	Cfg *cfg.Config
}

// NewReplayServer creates a go-gin server
func NewReplayServer(cfg *cfg.Config) *Server {
	server := &Server{
		Engine: gin.Default(),
		Cfg:    cfg,
	}
	setup(server)
	return server
}

func setup(server *Server) {
	indexGroup := server.Group("/")
	{
		healthzCtrl := healthz.Controller{}
		indexGroup.Any("/healthz", healthzCtrl.Healthz)
	}
}

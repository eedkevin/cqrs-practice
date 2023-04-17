package app

import (
	"cqrs-practise/internal/app/adapter/ctrl/cqrs/command"
	"cqrs-practise/internal/app/adapter/ctrl/healthz"
	"cqrs-practise/internal/app/adapter/repository"
	serviceImpl "cqrs-practise/internal/app/adapter/service"
	"cqrs-practise/internal/app/application/service"
	"cqrs-practise/internal/app/domain/event/repo"
	"cqrs-practise/internal/cfg"

	"github.com/gin-gonic/gin"
)

type Server struct {
	*gin.Engine
	Cfg        *cfg.Config
	EventBus   service.EventBus
	EventStore repo.EventRepo
}

func NewReplayServer(cfg *cfg.Config) *Server {
	server := &Server{
		Engine: gin.Default(),
		Cfg:    cfg,
	}
	setup(server)
	return server
}

func setup(server *Server) {
	server.EventBus = serviceImpl.NewNatsEventBus(server.Cfg)
	server.EventStore = repository.NewSQLiteEventRepo(server.Cfg)

	indexGroup := server.Group("/")
	{
		healthzCtrl := healthz.Controller{}
		indexGroup.Any("/healthz", healthzCtrl.Healthz)
	}

	cqrsGroup := server.Group("/cqrs")
	{
		cqrsCmdCtrl := command.NewController(server.EventBus, server.EventStore)
		cqrsGroup.POST("/commands", command.CommandReqValidator, cqrsCmdCtrl.Handle)
	}
}

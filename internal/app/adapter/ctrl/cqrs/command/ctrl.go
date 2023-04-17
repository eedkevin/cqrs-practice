package command

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"cqrs-practise/internal/app/application/service"
	"cqrs-practise/internal/app/application/usecase"
	"cqrs-practise/internal/app/domain/event/repo"
	"cqrs-practise/internal/cfg"
)

type Controller struct {
	EventBus   service.EventBus
	EventStore repo.EventRepo
}

func NewController(eventbus service.EventBus, eventstore repo.EventRepo) *Controller {
	return &Controller{
		EventBus:   eventbus,
		EventStore: eventstore,
	}
}

func (ctrl Controller) Handle(c *gin.Context) {
	var in CommandReq
	c.ShouldBindBodyWith(&in, binding.JSON)
	log.Printf("Receiving cqrs command(payload has been omitted): %v\n", in)

	switch in.Command {
	case "REPLAY":
		var replayCmd ReplayCommandReq
		c.ShouldBindBodyWith(&replayCmd, binding.JSON)
		go usecase.ReplayEvents(ctrl.EventBus, ctrl.EventStore, usecase.ReplayParams{
			StartTime:       replayCmd.Payload.StartTime,
			EndTime:         replayCmd.Payload.EndTime,
			ENV:             replayCmd.Payload.ENV,
			EventBusSubject: cfg.Cfg.EventWorker.Odds.Event,
		})
	default:
		log.Printf("Unsupported command: %s", in.Command)
	}
	c.String(http.StatusOK, "ok")
}

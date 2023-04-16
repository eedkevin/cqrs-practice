package replayservice

import (
	"cqrs-practise/internal/app"
	"cqrs-practise/internal/cfg"
	"fmt"
)

func Bootstrap() {
	server := app.NewReplayServer(&cfg.Cfg)
	server.Run(fmt.Sprintf(":%v", server.Cfg.App.ServerPort))
}

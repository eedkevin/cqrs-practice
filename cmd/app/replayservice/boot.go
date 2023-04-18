package replayservice

import (
	"cqrs-practise/internal/app"
	"cqrs-practise/internal/cfg"
	"fmt"
	"log"
)

func Bootstrap() {
	cfg.Init()
	log.Printf("Boost using cfg: %v\n", cfg.Cfg)
	server := app.NewReplayServer(&cfg.Cfg)
	server.Run(fmt.Sprintf(":%v", server.Cfg.App.ServerPort))
}

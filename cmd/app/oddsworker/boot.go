package oddsworker

import (
	"cqrs-practise/internal/app"
	"cqrs-practise/internal/cfg"
	"log"
)

func Bootstrap() {
	cfg.Init()
	log.Printf("Boost using cfg: %v\n", cfg.Cfg)
	oddsworker := app.NewOddsWorker(&cfg.Cfg)
	oddsworker.Loop()
}

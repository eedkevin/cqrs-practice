package oddsworker

import (
	"cqrs-practise/internal/app"
	"cqrs-practise/internal/cfg"
)

func Bootstrap() {
	cfg.Init()
	oddsworker := app.NewOddsWorker(&cfg.Cfg)
	oddsworker.Loop()
}

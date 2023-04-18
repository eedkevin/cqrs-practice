package app

import (
	"cqrs-practise/internal/app/adapter/repository"
	serviceImpl "cqrs-practise/internal/app/adapter/service"
	"cqrs-practise/internal/app/application/service"
	"cqrs-practise/internal/app/application/usecase"
	"cqrs-practise/internal/app/domain/odds/repo"
	"cqrs-practise/internal/cfg"
	"log"
	"sync"
)

type OddsWorker struct {
	Cfg      *cfg.Config
	EventBus service.EventBus
	OddsRepo repo.OddsRepo
}

func NewOddsWorker(cfg *cfg.Config) *OddsWorker {
	return &OddsWorker{
		Cfg:      cfg,
		EventBus: serviceImpl.NewNatsEventBus(cfg),
		OddsRepo: repository.NewSQLiteOddsRepo(cfg),
	}
}

func (ow *OddsWorker) Loop() {
	var wg sync.WaitGroup
	wg.Add(ow.Cfg.EventWorker.Odds.Concurrency)

	for i := 0; i < ow.Cfg.EventWorker.Odds.Concurrency; i++ {
		log.Printf("Starting worker thread: #%v\n", i)
		go func() {
			defer wg.Done()
			for {
				usecase.HandleOddsEvent(ow.EventBus, ow.OddsRepo, usecase.HandleOddsEventParams{
					EventBusSubject: ow.Cfg.EventWorker.Odds.Event,
					EventBusQueue:   ow.Cfg.EventWorker.Odds.Name,
				})
			}
		}()
	}

	wg.Wait()
}

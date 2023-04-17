package usecase

import (
	"log"
	"time"

	"github.com/pkg/errors"

	"cqrs-practise/internal/app/application/service"
	"cqrs-practise/internal/app/domain/event/repo"
)

var PAGE_SIZE int = 20

type ReplayParams struct {
	StartTime       time.Time
	EndTime         time.Time
	ENV             string
	EventBusSubject string
}

func ReplayEvents(eventbus service.EventBus, eventStore repo.EventRepo, params ReplayParams) error {
	var USECASE_NAME = "ReplayEvents"
	log.Printf("%s has started\n", USECASE_NAME)

	var pageNo int = 1
	for {
		events, err := eventStore.List(repo.QueryParams{
			StartTime: params.StartTime,
			EndTime:   params.EndTime,
			PageSize:  PAGE_SIZE,
			PageNo:    pageNo,
		})
		if err != nil {
			return errors.Wrap(err, USECASE_NAME)
		}

		for _, event := range events {
			err := eventbus.Send(params.EventBusSubject, *event)
			if err != nil {
				// TODO: send it to failure queue
				log.Printf("%s: error on sending to event bus. detail: %v\n", USECASE_NAME, event)
			}
		}

		if len(events) == 0 {
			break
		}

		pageNo++
	}

	log.Printf("%s finished\n", USECASE_NAME)
	return nil
}

package usecase

import (
	"fmt"
	"log"
	"time"

	"github.com/pkg/errors"

	"cqrs-practise/internal/app/application/service"
	"cqrs-practise/internal/app/domain/event/repo"
	"cqrs-practise/internal/cfg"
)

var PAGE_SIZE int32 = 20

type ReplayQueryParams struct {
	StartTime time.Time
	EndTime   time.Time
	ENV       string
}

func ReplayEvents(eventbus service.EventBus, eventStore repo.EventRepo, params ReplayQueryParams) error {
	var USECASE_NAME = "ReplayEvents"

	if params.ENV != cfg.Cfg.App.ENV {
		return errors.New(fmt.Sprintf("%s: will not apply the replay due to incorrect env from request", USECASE_NAME))
	}

	var pageNo int32 = 1
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
			err := eventbus.Send(event)
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

	return nil
}

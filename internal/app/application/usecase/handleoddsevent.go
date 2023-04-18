package usecase

import (
	"cqrs-practise/internal/app/application/service"
	"cqrs-practise/internal/app/domain/cqrs/model"
	"cqrs-practise/internal/app/domain/odds/repo"
	"log"

	"github.com/pkg/errors"
)

type HandleOddsEventParams struct {
	EventBusSubject string
	EventBusQueue   string
}

func HandleOddsEvent(eventbus service.EventBus, oddsrepo repo.OddsRepo, params HandleOddsEventParams) error {
	var USECASE_NAME = "HandleOddsEvent"
	log.Printf("%s has started\n", USECASE_NAME)

	event, err := eventbus.QueueReceive(params.EventBusSubject, params.EventBusQueue)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return errors.Wrap(err, "Error on consuming eventbus")
	}

	if event.Domain == "odds" {
		oddscmd := event.Payload.(model.OddsCommand)

		switch oddscmd.Command.Command {
		case "UpdateOdds":
			oddsrepo.Update(oddscmd.Payload.UUID, oddscmd.Payload)
		case "NewOdds":
			oddsrepo.Save(oddscmd.Payload)
		case "DeleteOdds":
			oddsrepo.Delete(oddscmd.Payload.UUID)
		default:
			log.Printf("%s: unsupported command: %s\n", USECASE_NAME, oddscmd.Command.Command)
		}
	}

	return nil
}

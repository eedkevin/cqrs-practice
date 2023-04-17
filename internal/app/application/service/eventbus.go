package service

import (
	"cqrs-practise/internal/app/domain/event/model"
)

type EventBus interface {
	Send(subject string, event model.Event) error
	Receive(subject string) (*model.Event, error)
	QueueReceive(subject string, queue string) (*model.Event, error)
}

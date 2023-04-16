package model

import (
	"cqrs-practise/internal/app/domain/event/model"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	model.Event
}

func (e *Event) ToModel() *model.Event {
	return &model.Event{
		UUID:    e.UUID,
		Domain:  e.Domain,
		Payload: e.Payload,
	}
}

func (e *Event) FromModel(event model.Event) {
	e.UUID = event.UUID
	e.Domain = event.Domain
	e.Payload = event.Payload
}

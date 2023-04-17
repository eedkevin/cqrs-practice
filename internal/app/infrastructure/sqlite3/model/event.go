package model

import (
	"cqrs-practise/internal/app/domain/event/model"
	"cqrs-practise/internal/util"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	UUID    string
	Domain  string
	Payload datatypes.JSON
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
	b, _ := util.Encode(event.Payload)
	e.Payload = datatypes.JSON(b)
}

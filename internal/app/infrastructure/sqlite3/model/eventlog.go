package model

import (
	"cqrs-practise/internal/app/domain/event/model"

	"gorm.io/gorm"
)

type EventLog struct {
	gorm.Model
	model.EventLog
}

func (l *EventLog) ToModel() *model.EventLog {
	return &model.EventLog{
		UUID:             l.UUID,
		Domain:           l.Domain,
		DomainObjectUUID: l.DomainObjectUUID,
	}
}

func (l *EventLog) FromModel(eventlog model.EventLog) {
	l.UUID = eventlog.UUID
	l.EventUUID = eventlog.EventUUID
	l.Domain = eventlog.Domain
	l.DomainObjectUUID = eventlog.DomainObjectUUID
}

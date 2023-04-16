package repo

import (
	"cqrs-practise/internal/app/domain/event/model"
	"time"
)

type QueryParams struct {
	StartTime time.Time
	EndTime   time.Time
	PageSize  int
	PageNo    int
}

type EventRepo interface {
	Save(event model.Event) (*model.Event, error)
	Get(UUID string) (*model.Event, error)
	List(params QueryParams) ([]*model.Event, error)
}

type EventLogRepo interface {
	Save(event model.EventLog) (*model.Event, error)
	Get(UUID string) (*model.EventLog, error)
	List(params QueryParams) ([]*model.EventLog, error)
}

package repository

import (
	domainModel "cqrs-practise/internal/app/domain/event/model"
	"cqrs-practise/internal/app/domain/event/repo"
	"cqrs-practise/internal/app/infrastructure/sqlite3"
	"cqrs-practise/internal/app/infrastructure/sqlite3/model"
	"cqrs-practise/internal/cfg"

	"github.com/pkg/errors"
)

type SQLiteEventRepo struct {
	DB *sqlite3.Client
}

type SQLiteEventLogRepo struct {
	DB *sqlite3.Client
}

func NewSQLiteEventRepo(cfg *cfg.Config) *SQLiteEventRepo {
	return &SQLiteEventRepo{
		DB: sqlite3.Connection(cfg),
	}
}

func NewSQLiteEventLogRepo(cfg *cfg.Config) *SQLiteEventLogRepo {
	return &SQLiteEventLogRepo{
		DB: sqlite3.Connection(cfg),
	}
}

func (ser *SQLiteEventRepo) Save(event domainModel.Event) (*domainModel.Event, error) {
	evt := model.Event{}
	evt.FromModel(event)

	err := ser.DB.Create(&evt).Error
	if err != nil {
		return nil, errors.Wrap(err, "Error saving event")
	}
	return evt.ToModel(), nil
}

func (ser *SQLiteEventRepo) Get(UUID string) (*domainModel.Event, error) {
	var evt model.Event

	err := ser.DB.Where("uuid = ?", UUID).First(&evt).Error
	if err != nil {
		return nil, errors.Wrap(err, "Error getting event")
	}
	return evt.ToModel(), nil
}

func (ser *SQLiteEventRepo) List(p repo.QueryParams) ([]*domainModel.Event, error) {
	evts := make([]*model.Event, 0)

	err := ser.DB.Limit(p.PageSize).Offset((p.PageNo - 1) * p.PageSize).Find(&evts).Error
	if err != nil {
		return nil, errors.Wrap(err, "Error listing events")
	}

	evtsModel := make([]*domainModel.Event, 0)
	for _, e := range evts {
		m := e.ToModel()
		evtsModel = append(evtsModel, m)
	}
	return evtsModel, nil
}

func (selr *SQLiteEventLogRepo) Save(eventlog domainModel.EventLog) error {
	evtlog := model.EventLog{}
	evtlog.FromModel(eventlog)

	err := selr.DB.Create(&evtlog).Error
	if err != nil {
		return errors.Wrap(err, "Error saving eventlog")
	}
	return nil
}

func (selr *SQLiteEventLogRepo) Get(UUID string) (*domainModel.EventLog, error) {
	var evtlog model.EventLog

	err := selr.DB.Where("uuid = ?", UUID).First(&evtlog).Error
	if err != nil {
		return nil, errors.Wrap(err, "Error getting eventlog")
	}
	return evtlog.ToModel(), nil
}

func (selr *SQLiteEventLogRepo) List(p repo.QueryParams) ([]*domainModel.EventLog, error) {
	evtlogs := make([]*model.EventLog, 0)

	err := selr.DB.Limit(p.PageSize).Offset((p.PageNo - 1) * p.PageSize).Find(&evtlogs).Error
	if err != nil {
		return nil, errors.Wrap(err, "Error listing eventlogs")
	}

	evtlogsModel := make([]*domainModel.EventLog, 0)
	for _, e := range evtlogs {
		m := e.ToModel()
		evtlogsModel = append(evtlogsModel, m)
	}
	return evtlogsModel, nil
}

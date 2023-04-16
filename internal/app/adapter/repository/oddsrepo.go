package repository

import (
	domainModel "cqrs-practise/internal/app/domain/odds/model"
	"cqrs-practise/internal/app/domain/odds/repo"
	"cqrs-practise/internal/app/infrastructure/sqlite3"
	"cqrs-practise/internal/app/infrastructure/sqlite3/model"

	"github.com/pkg/errors"
)

type SQLiteOddsRepo struct {
	DB *sqlite3.Client
}

func (sor *SQLiteOddsRepo) Save(odds domainModel.Odds) error {
	o := model.Odds{}
	o.FromModel(odds)

	err := ser.DB.Create(&evt).Error
	if err != nil {
		return errors.Wrap(err, "Error saving event")
	}
	return nil
}

func (sor *SQLiteOddsRepo) Get(UUID string) (*domainModel.Event, error) {
	var evt model.Event

	err := ser.DB.First(&evt, UUID).Error
	if err != nil {
		return nil, errors.Wrap(err, "Error getting event")
	}
	return evt.ToModel(), nil
}

func (sor *SQLiteEventRepo) List(p repo.QueryParams) ([]*domainModel.Event, error) {
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

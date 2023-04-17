package repository

import (
	domainModel "cqrs-practise/internal/app/domain/odds/model"
	"cqrs-practise/internal/app/infrastructure/sqlite3"
	"cqrs-practise/internal/app/infrastructure/sqlite3/model"
	"cqrs-practise/internal/cfg"

	"github.com/pkg/errors"
)

type SQLiteOddsRepo struct {
	DB *sqlite3.Client
}

func NewSQLiteOddsRepo(cfg *cfg.Config) *SQLiteOddsRepo {
	return &SQLiteOddsRepo{
		DB: sqlite3.Connection(cfg),
	}
}

func (sor *SQLiteOddsRepo) Save(odds domainModel.Odds) (*domainModel.Odds, error) {
	o := model.Odds{}
	o.FromModel(odds)

	err := sor.DB.Create(&o).Error
	if err != nil {
		return nil, errors.Wrap(err, "Error saving odds")
	}
	return o.ToModel(), nil
}

func (sor *SQLiteOddsRepo) Get(UUID string) (*domainModel.Odds, error) {
	var o model.Odds

	err := sor.DB.Where("uuid = ? AND deleted = 0", UUID).First(&o).Error
	if err != nil {
		return nil, errors.Wrap(err, "Error getting odds")
	}
	return o.ToModel(), nil
}

func (sor *SQLiteOddsRepo) Update(UUID string, odds domainModel.Odds) error {
	err := sor.DB.Where("uuid = ? AND deleted = 0", UUID).Updates(odds).Error
	if err != nil {
		return errors.Wrap(err, "Error updating odds")
	}
	return nil
}

func (sor *SQLiteOddsRepo) Delete(UUID string) error {
	err := sor.DB.Where("uuid = ? AND deleted = 0", UUID).Updates(&model.Odds{Deleted: true}).Error
	if err != nil {
		return errors.Wrap(err, "Error deleting odds")
	}
	return nil
}

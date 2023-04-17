package repo

import (
	"cqrs-practise/internal/app/domain/odds/model"
)

type OddsRepo interface {
	Save(odds model.Odds) (*model.Odds, error)
	Get(UUID string) (*model.Odds, error)
	Update(UUID string, odds model.Odds) error
	Delete(UUID string) error
}

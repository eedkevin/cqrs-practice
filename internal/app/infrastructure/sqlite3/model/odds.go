package model

import (
	"cqrs-practise/internal/app/domain/odds/model"

	"gorm.io/gorm"
)

type Odds struct {
	gorm.Model
	UUID      string
	GameUUID  string
	MoneyLine struct {
		Home float64
		Away float64
		Draw float64
	}
	Deleted bool
}

func (o *Odds) ToModel() *model.Odds {
	return &model.Odds{
		UUID:      o.UUID,
		GameUUID:  o.GameUUID,
		MoneyLine: o.MoneyLine,
	}
}

func (o *Odds) FromModel(odds model.Odds) {
	o.UUID = odds.UUID
	o.GameUUID = odds.GameUUID
	o.MoneyLine = odds.MoneyLine
}

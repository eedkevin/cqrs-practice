package command

import (
	"cqrs-practise/internal/app/domain/cqrs/model"
)

type OddsCommandReq struct {
	CommandReq
	model.OddsCommand
}

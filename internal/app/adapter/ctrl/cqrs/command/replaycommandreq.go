package command

import (
	"cqrs-practise/internal/app/domain/cqrs/model"
)

type ReplayCommandReq struct {
	CommandReq
	model.ReplayCommand
}

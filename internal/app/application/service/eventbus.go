package service

type EventBus interface {
	Send(event interface{}) error
}

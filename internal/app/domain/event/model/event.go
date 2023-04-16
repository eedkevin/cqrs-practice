package model

type Event struct {
	UUID    string
	Domain  string
	Payload interface{}
}

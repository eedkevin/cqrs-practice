package nats

import (
	"log"

	"cqrs-practise/internal/cfg"

	"github.com/nats-io/nats.go"
)

type Client = nats.Conn
type Message = nats.Msg

func Connection(cfg *cfg.Config) *Client {
	nc, err := nats.Connect(cfg.Nats.URL)
	if err != nil {
		log.Fatal("Error on connecting to nats", err)
	}
	return nc
}

package service

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"

	"github.com/pkg/errors"

	"cqrs-practise/internal/app/domain/event/model"
	"cqrs-practise/internal/app/infrastructure/nats"
	"cqrs-practise/internal/cfg"
)

var TIMEOUT = 15 * time.Second

type NatsEventBus struct {
	Client *nats.Client
}

func NewNatsEventBus(cfg *cfg.Config) *NatsEventBus {
	return &NatsEventBus{
		Client: nats.Connection(cfg),
	}
}

func (neb *NatsEventBus) Send(subject string, event model.Event) error {
	b, err := encode(event)
	if err != nil {
		return errors.Wrap(err, "Error on encoding event")
	}
	err = neb.Client.Publish(subject, b)
	if err != nil {
		return errors.Wrap(err, "Error on sending to eventbus")
	}
	return nil
}

func (neb *NatsEventBus) Receive(subject string) (*model.Event, error) {
	sub, err := neb.Client.SubscribeSync(subject)
	if err != nil {
		return nil, errors.Wrap(err, "Timeout on receiving from eventbus")
	}
	m, err := sub.NextMsg(TIMEOUT)
	if err != nil {
		return nil, errors.Wrap(err, "Error on receiving from eventbus")
	}
	fmt.Printf("Received a message: %s\n", string(m.Data))
	event, err := decode(m.Data)
	if err != nil {
		return nil, errors.Wrap(err, "Error on decoding event")
	}
	return event, nil
}

func (neb *NatsEventBus) QueueReceive(subject string, queue string) (*model.Event, error) {
	sub, err := neb.Client.QueueSubscribeSync(subject, queue)
	if err != nil {
		return nil, errors.Wrap(err, "Timeout on (queue)receiving from eventbus")
	}
	m, err := sub.NextMsg(TIMEOUT)
	if err != nil {
		return nil, errors.Wrap(err, "Error on (queue)receiving from eventbus")
	}
	fmt.Printf("Received a message: %s\n", string(m.Data))
	event, err := decode(m.Data)
	if err != nil {
		return nil, errors.Wrap(err, "Error on decoding event")
	}
	return event, nil
}

func encode(data model.Event) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func decode(b []byte) (*model.Event, error) {
	var data model.Event
	var buf bytes.Buffer
	dec := gob.NewDecoder(&buf)
	err := dec.Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

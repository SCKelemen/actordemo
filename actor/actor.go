package actor

import (
	"fmt"

	e "github.com/SCKelemen/actordemo/event"
)

type Actor interface {
	Listen(address string) error
	Dispatch(e.Event) error
	Handle(e.Event) error
}

type AdderActor struct {
	State int

	id     uint64
	events chan e.Event
}

func (adder *AdderActor) Listen(address string) error {

	adder.events = make(chan e.Event, 1024)

	for event := range adder.events {
		if err := adder.Handle(event); err != nil {
			return err
		}
	}

	return nil
}

func (adder *AdderActor) Dipatch(event e.Event) error {
	return nil
}

func (adder *AdderActor) Handle(event e.Event) error {
	switch event.Type {
	case e.ADD:
		return adder.onAddEvent(event)
	default:
		return fmt.Errorf("no handler defined for event %s", event.Type)
	}
}

func (adder *AdderActor) Dispatch(event e.Event) error {
	adder.events <- event
	return nil
}

func (adder *AdderActor) onAddEvent(event e.Event) error {
	adder.State++
	return nil
}

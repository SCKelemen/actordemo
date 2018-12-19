package actor

import "github.com/sckelemen/actordemo/event"

type Actor interface {
	Listen(address string) error
	Dispatch(event.Event) error
	Handle(event.Event) error
}

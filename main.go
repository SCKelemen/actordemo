package main

import (
	"fmt"

	ev "github.com/SCKelemen/actordemo/event"
	"github.com/sckelemen/actordemo/actor"
)

func main() {
	fmt.Println("Starting Actor System")
	adder := actor.AdderActor{State: 0}
	adder.Listen("1111")
	adder.Handle(ev.Event{Type: ev.ADD})
	fmt.Println("Shutting down...")
}

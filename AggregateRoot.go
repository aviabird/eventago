package eventago

import (
	"fmt"
	// EventStore "dummy_bank/dummycqrs/eventago/EventStore"
)

// AggregateMain struct used for all aggregate function.
type AggregateMain struct {
	id    string
	event []EventStream
}

// AggregateRoot interface used to declare all functions.
type AggregateRoot interface {
	id() string
	event()
	SetID(UUID string)
	GetID() string
	LoadFromEventStream()
	pullDomainEvents()
}

// SetID stores the EventMessage in the changes collection.
func (a *AggregateMain) SetID(UUID string) {
	a.id = UUID
}

// GetID method return id
func (a *AggregateMain) GetID() string {
	return a.id
}

// LoadFromEventStream implements the Events method of the Aggregate interface.
func (a AggregateMain) LoadFromEventStream(eventstream EventStream) {
	if a.event == nil {
		fmt.Println("AggregateRoot was already created from event stream and cannot be hydrated again.")
	}
	a.SetID(eventstream.getUUID())
	for _, ev := range a.event {
		fmt.Println(ev)
	}
}

// pullDomainEvents is used for get event and return events
func (a *AggregateMain) pullDomainEvents() []EventStream {
	event := a.event
	a.event = nil
	return event
}

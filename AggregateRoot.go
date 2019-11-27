package eventago

import (
	"fmt"
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
	IncrementVersion()
	LoadFromEventStream()
	pullDomainEvents()
	executeEvent()
}

// SetID stores the EventMessage in the changes collection.
func (a *AggregateMain) SetID(UUID string) {
	a.id = UUID
}

// GetID method return id
func (a *AggregateMain) GetID() string {
	return a.id
}

// Apply method return for event present or not in this programm.
// func (a *AggregateMain) Apply(ev DomainEvent) {
// 	// go executeEvent(ev)
// 	// var event AggregateMain
// 	a.event = append(a.event, ev)
// }

// func executeEvent(event DomainEvent) {
// 	var evnetname EventNameRoot
// 	evnetname = evnetname(event)
// 	method := fmt.Sprintf("apply%s", eventname)

// 	if event != method {
// 		fmt.Println("There is no event named %s that can be applied", method)
// 	}
// 	method(event)
// }

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

// IncrementVersion increments the v of the aggregate and should be called
// after an event has been applied successfully in ApplyEvent.
// func (a *AggregateMain) IncrementVersion() {
// 	a.v++
// }

//ClearAllEvents removes all unpersisted events from the aggregate.
// func (a *AggregateMain) ClearAllEvents() {
// 	a.events = []string
// }

// TrackAllEvents stores the EventMessage in the changes collection.
// func (a *AggregateMain) TrackAllEvents(event EventMessage) {
// 	a.events = append(a.events, event)
// }

// pullDomainEvents is used for get event and return events
func (a *AggregateMain) pullDomainEvents() []EventStream {
	event := a.event
	a.event = nil
	return event
}

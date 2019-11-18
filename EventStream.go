package eventago


import (
	uuid "dummy_bank/dummycqrs/eventago/internal/uuid"
)
// EventStream define all member veriable for event stram
type EventStream struct {
	eid       string
	event     []string
	newEvent  []string
	className string
	version   int
}

func init() {
	eid := uuid.NewUUID()
	event := make([]string,10)
	newEvent := make([]string,10)
	className := ""
	version := 0
}

// GetClassName return class name
func (es *EventStream) GetClassName() string {
	return es.className
}

// getUUID return event uid
func (es *EventStream) getUUID() string {
	return es.eid
}

// GetVersion return class name
func (es *EventStream) GetVersion() int {
	return es.version
}

// NewEvents return new event
// func (es EventStream) NewEvents() []EventMessage {
// 	return es.newEvent
// }

// AddEvents call each events in array and send to AddEvent function.
func (es *EventStream) AddEvents() {
	for _, events := range es.newEvent {
				go AddEvent(events)
			}
}

// AddEvent is used for add event in array and update new events.
func AddEvent(ev *DomainEvent) {
	events = []ev
	newEvent = []ev
}

// MarkNewEventsProcessed is used for create first event and make new event array.
func (es *EventStream) MarkNewEventsProcessed() {
	es.version = 0
	es.newEvent = make([]string)
}


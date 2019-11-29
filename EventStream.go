package eventago

// EventStream define all member veriable for event stram
type EventStream struct {
	eid          string
	event        []string
	eventDetails string
	version      int
}

// EventStreamRoot EventStreamRoot
type EventStreamRoot interface {
	getUUID() string
	GetClassName() string
	GetVersion() int
}

func (es EventStream) new(eid string, event []string, eventDetails string, version int) EventStream {
	es = EventStream{eid, event, eventDetails, version}
	return es
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
func (es *EventStream) NewEvents() string {
	return es.eventDetails
}

// AddEvents call each events in array and send to AddEvent function.
func (es *EventStream) AddEvents() {
	for _, events := range es.eventDetails {
		var events string = string(events)
		es.AddEvent(events)
	}
}

// AddEvent is used for add event in array and update new events.
func (es *EventStream) AddEvent(events string) {
	es.event = append(es.event, events)
	es.eventDetails = es.eventDetails + events
}

// MarkNewEventsProcessed is used for create first event and make new event array.
func (es *EventStream) MarkNewEventsProcessed() {
	es.version = 0
	es.eventDetails = ""
}

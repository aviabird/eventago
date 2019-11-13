package cqrs

import "time"

// EventMessage events details
type EventMessage interface {
	AggregateID() string

	// Returns the actual event which is the payload of the event message.
	EventData() interface{}

	// EventType returns a string descriptor of the command name
	EventType() string

	// Version returns the version of the event
	Version() *int

	// Timestamp of when the event was created.
	Timestamp() time.Time
}

// EventInformation is an implementation of the event message interface.
type EventInformation struct {
	id        string
	eventdata interface{}
	version   *int
	timestamp time.Time
}

// NewEvent returns a new event descriptor
func NewEvent(aggregateID string, eventdata interface{}, timestamp time.Time, version *int) *EventInformation {
	return &EventInformation{
		id:        aggregateID,
		eventdata: eventdata,
		timestamp: timestamp,
		version:   version,
	}
}

// EventType returns the name of the event type as a string.
func (c *EventInformation) EventType() string {
	return typeOf(c.eventdata)
}

// AggregateID returns the ID of the Aggregate that the event relates to.
func (c *EventInformation) AggregateID() string {
	return c.id
}

// EventData the event payload of the event message
func (c *EventInformation) EventData() interface{} {
	return c.eventdata
}

// Version returns the version of the event
func (c *EventInformation) Version() *int {
	return c.version
}

// Timestamp returns the Timestamp of the event
func (c *EventInformation) Timestamp() time.Time {
	return c.timestamp
}

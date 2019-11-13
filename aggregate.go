package cqrs

// AggregateRoot is used to implementation
type AggregateRoot struct {
	id      string
	version int
	events  []EventMessage
}

// AggregateMain is used to implementation
type AggregateMain interface {
	Id() string
	Version() int
	IncrementVersion()
	Apply(events EventMessage)
	TrackAllEvents(EventMessage)
	GetAllEvents() []EventMessage
	ClearAllEvents()
}

// NewAggregateFields is for create field.
func NewAggregateFields(id string) *AggregateRoot {
	return &AggregateRoot{
		id:      id,
		version: -1,
		events:  []EventMessage{},
	}
}

// GetAggregateID returns the ID
func (a *AggregateRoot) GetAggregateID() string {
	return a.id
}

// IncrementVersion increments the aggregate version number by one.
func (a *AggregateRoot) IncrementVersion() {
	a.version++
}

// TrackAllEvents stores the EventMessage in the changes collection.
func (a *AggregateRoot) TrackAllEvents(event EventMessage) {
	a.events = append(a.events, event)
}

// GetAllEvents returns the collection of new unpersisted events that have
// been applied to the aggregate.
func (a *AggregateRoot) GetAllEvents() []EventMessage {
	return a.events
}

//ClearAllEvents removes all unpersisted events from the aggregate.
func (a *AggregateRoot) ClearAllEvents() {
	a.events = []EventMessage{}
}

// func (a *AggregateRoot) apply(events EventMessage) {

// }

package eventago

// EventNameRoot decleare all event name root
type EventNameRoot struct {
	event string
	name  string
}

// EventName interface used to declare all functions.
type EventName interface {
	GetEventName()
	SetEventName()
}

// GetEventName return event name
// func (e DomainEvent) GetEventName() string {
// 	return a.events
// }

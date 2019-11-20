package eventago

// DomainEvent DomainEvent
type DomainEvent struct {
	aggregateID string
	event       string
}

// DomainEventRoot interface for get and set aggregate id
type DomainEventRoot interface {
	setAggregateId()
	getAggregateId()
}

// type DomainEvent struct {
// 	aggregateId string
// }

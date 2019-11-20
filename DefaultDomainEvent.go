package eventago

// DefaultDomainEventBase declaration of fields
type DefaultDomainEventBase struct {
	aggregateID string
	data        []string
}

// DefaultDomainEvent DefaultDomainEvent
type DefaultDomainEvent interface {
	assertPropertyExists(name string)
	setAggregateId(aggregateID string)
	getAggregateId()
}

func (dde *DefaultDomainEventBase) init() {
	// for key, value := range dde.data {

	// }
}

func (dde *DefaultDomainEventBase) assertPropertyExists(name string) {
	// if dde == name {

	// }
}

func (dde *DefaultDomainEventBase) setAggregateID(aggregateID string) {
	dde.aggregateID = aggregateID
}

func (dde *DefaultDomainEventBase) getAggregateID() string {
	return dde.aggregateID
}

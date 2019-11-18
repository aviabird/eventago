package eventago

type DefaultDomainEventBase struct {
	aggregateId string
	data        []string
}

type DefaultDomainEvent interface {
	assertPropertyExists(name string)
	setAggregateId(aggregateId string)
	getAggregateId()
}

func (dde *DefaultDomainEventBase) init() {
	for key, value := range dde.data {

	}
}

func (dde *DefaultDomainEventBase) assertPropertyExists(name string) {
	// if dde == name {

	// }
}

func (dde *DefaultDomainEventBase) setAggregateId(aggregateId string) {
	dde.aggregateId = aggregateId
}

func (dde *DefaultDomainEventBase) getAggregateId() string {
	return dde.aggregateId
}

package eventago

// DomainEvent interface for get and set aggregate id
type DomainEvent interface {
	setAggregateId()
	getAggregateId()
}

// type DomainEvent struct {
// 	aggregateId string
// }

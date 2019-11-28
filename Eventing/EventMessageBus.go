package eventago

// EventMessageBus EventMessageBus
type EventMessageBus interface {
	publish(event DomainEvent)
}

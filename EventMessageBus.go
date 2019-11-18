package eventago

type EventMessageBus interface {
	publish(event DomainEvent)
}

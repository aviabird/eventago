package eventago

import eventago "github.com/aviabird/eventago"

// EventMessageBus EventMessageBus
type EventMessageBus interface {
	publish(event eventago.DomainEvent)
}

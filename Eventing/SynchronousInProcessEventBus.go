package eventago

import eventago "github.com/aviabird/eventago"

//SynchronousInProcessEventBus SynchronousInProcessEventBus
type SynchronousInProcessEventBus struct {
	locator string
}

// SynchronousInProcessEventBusRoot SynchronousInProcessEventBusRoot
type SynchronousInProcessEventBusRoot interface {
	publish(ev eventago.DomainEvent, mehl MemoryEventHandlerLocatorRoot)
	invokeEventHandler(service string, eventname string, ev eventago.DomainEvent)
}

func (s *SynchronousInProcessEventBus) init(event EventHandlerLocator) {
	s.locator = event.locator
}

func publish(ev eventago.DomainEvent, mehl MemoryEventHandlerLocatorRoot) {
	eventdata := new(eventago.EventNameRoot)
	eventdata.aggregateID = eventago.aggregateID
	eventdata.event = eventago.event

	// services := mehl.getHandlersForEvent()

	// for _, service := range services {
	// 	// SynchronousInProcessEventBusRoot.invokeEventHandler(service, eventdata, ev)
	// }
}

func invokeEventHandler(service string, eventname string, ev eventago.DomainEvent) {
	// methodName := eventname
	// err := service.methodName(ev)
	// if err != nil {
	// 	SynchronousInProcessEventBusRoot.publish(EventExecutionFailed(service, eventname), ev)
	// }
}

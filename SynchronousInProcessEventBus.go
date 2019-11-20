package eventago

//SynchronousInProcessEventBus SynchronousInProcessEventBus
type SynchronousInProcessEventBus struct {
	locator string
}

// SynchronousInProcessEventBusRoot SynchronousInProcessEventBusRoot
type SynchronousInProcessEventBusRoot interface {
	publish(ev DomainEvent, mehl MemoryEventHandlerLocatorRoot)
	invokeEventHandler(service string, eventname string, ev DomainEvent)
}

func (s *SynchronousInProcessEventBus) init(event EventHandlerLocator) {
	s.locator = event.locator
}

func publish(ev DomainEvent, mehl MemoryEventHandlerLocatorRoot) {
	eventdata := new(EventNameRoot)
	eventdata.aggregateID = ev.aggregateID
	eventdata.event = ev.event

	// services := mehl.getHandlersForEvent()

	// for _, service := range services {
	// 	// SynchronousInProcessEventBusRoot.invokeEventHandler(service, eventdata, ev)
	// }
}

func invokeEventHandler(service string, eventname string, ev DomainEvent) {
	// methodName := eventname
	// err := service.methodName(ev)
	// if err != nil {
	// 	SynchronousInProcessEventBusRoot.publish(EventExecutionFailed(service, eventname), ev)
	// }
}

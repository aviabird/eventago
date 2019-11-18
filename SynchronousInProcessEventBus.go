package eventago

type SynchronousInProcessEventBus struct {
	locator string
}

func (s *SynchronousInProcessEventBus) init(event EventHandlerLocator) {
	s.locator = event.locator
}

func publish(ev DomainEvent, es EventHandlerLocator) {
	eventname = EventName(ev)
	services = es.getHandlersFor(eventname)

	for _, service := range services {
		go invokeEventHandler(service, eventname, ev)
	}
}

func invokeEventHandler(service string, eventname string, ev DomainEvent) {
	methodName := "on".eventname
	err := service.methodName(ev)
	if err != nil {
		go publish(EventExecutionFailed(service, eventname, ev))
	}
}

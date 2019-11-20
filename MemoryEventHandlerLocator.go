package eventago

// MemoryEventHandlerLocator MemoryEventHandlerLocator
type MemoryEventHandlerLocator struct {
	handlers []string
}

// MemoryEventHandlerLocatorRoot MemoryEventHandlerLocatorRoot
type MemoryEventHandlerLocatorRoot interface {
	getHandlersForEvent(eventname *EventNameRoot)
}

func (mehl *MemoryEventHandlerLocator) getHandlersForEvent(eventname EventNameRoot) []string {
	if eventname.event == "" {
		return nil
	}
	mehl.handlers = append(mehl.handlers, eventname.event)
	return mehl.handlers
}

func (mehl *MemoryEventHandlerLocator) register(eventname EventNameRoot) []string {
	if eventname.name == "" {
		mehl.handlers = append(mehl.handlers, eventname.name)
	}
	return mehl.handlers
}

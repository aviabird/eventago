package eventago

// EventHandlerLocator EventHandlerLocator
type EventHandlerLocator struct {
	locator string
}

// EventHandlerLocatorRoot EventHandlerLocatorRoot
type EventHandlerLocatorRoot interface {
	getHandlersForEvent(eventname *EventNameRoot)
}

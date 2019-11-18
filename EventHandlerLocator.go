package eventago

type EventHandlerLocator interface {
	getHandlersFor(eventname EventName)
}

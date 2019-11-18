package eventago

type MemoryEventHandlerLocator struct {
	handlers  []string 
}

func getHandlersFor(eventname EventName) []string {
	if eventname == nil {
		return []nil
	}
	handlers = append(eventname)
	return handlers
}

func (m *MemoryEventHandlerLocator)register(eventname EventName) []string {
	if (m.handlers[eventname] == nil) {
		m.handlers = append(eventname)
	}
	return handlers
}
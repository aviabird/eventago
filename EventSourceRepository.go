package eventago

// EventSourceRepository EventSourceRepository
type EventSourceRepository struct {
	eventStore string
	eventBus   string
	streams    []string
}

func (esr EventSourceRepository) init(eventstore EventStore, eventbus EventMessageBus) {
	// esr.eventStore = eventstore
	// esr.eventBus = eventbus
}

func (esr EventSourceRepository) find(classname string, expectedVersion int, UUID string) {
	// err, eventStream := EventStore.find(UUId)
	// if err != nil {
	// 	return
	// }
	// esr.streams[UUId] = eventStream

	// aggregateRootClass := EventStream.GetClassName(eventStream)

	// if aggregateRootClass != classname {
	// 	return fmt.Println("AggregateRootNotFoundException")
	// }

	// if expectedVersion && EventStream.GetVersion != expectedVersion {
	// 	return fmt.Println("AggregateRootNotFoundException")
	// }

	// AggregateRoot = AggregateRoot.LoadFromEventStream(eventStream)
}

func (esr EventSourceRepository) save(aggregateroot AggregateRoot) {
	// id := AggregateRoot.getId()

	// if esr.streams[id] != nil {
	// 	esr.streams[id] = EventStream(AggregateRoot.GetClassName(), AggregateRoot.getId())
	// }

	// EventStream = est.streams[id]
	// EventStream.AddEvents(AggregateRoot.pullDomainEvents())
	// transaction := EventStore.commit(EventStream)

	// for _, events := range transaction {
	// 	// AggregateRoot
	// }
}

package eventago

// EventSourceRepository EventSourceRepository
type EventSourceRepository struct {
	// eventStore EventStore
	// eventBus   EventMessageBus
	streams []string
}

// func (esr EventSourceRepository) init() {
// 	esr.eventStore = EventStore
// 	esr.eventBus = EventMessageBus
// }

// func (esr EventSourceRepository) find(classname string, expectedVersion int, UUID string) (string, error) {

// 	err, eventStream := EventStore.find(UUID)
// 	if err != "" {
// 		return "", err
// 	}
// 	esr.streams = eventStream

// 	aggregateRootClass := EventStreamRoot.GetClassName(eventStream)

// 	if aggregateRootClass != classname {
// 		// return fmt.Println("AggregateRootNotFoundException")
// 	}
// 	var versionInfo int
// 	versionInfo = EventStreamRoot.GetVersion()
// 	if expectedVersion != 0 && versionInfo != expectedVersion {
// 		return fmt.Println("AggregateRootNotFoundException")
// 	}

// 	AggregateRoot = AggregateRoot.LoadFromEventStream(eventStream)
// }

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

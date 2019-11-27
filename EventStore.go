package eventago

// EventStoreBase event store base
type EventStoreBase struct {
}

// EventStore   EventStore
type EventStore interface {
	find(UUID string) (string, error)
	commit(eventstram EventStream) (string, error)
}

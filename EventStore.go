package eventago

// EventStore   EventStore
type EventStore interface {
	find(UUID string)
	commit(eventstram EventStream)
}

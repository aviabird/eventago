package eventago

type EventStore interface {
	find(UUId string)
	commit(eventstram EventStream)
}

package eventago

type Transaction interface {
	getEventStream()
	getCommittedEvents()
}

type TransactionBase struct {
	eventStream     string
	committedEvents []string
}

func (transaction *TransactionBase) init(eventStream EventStream, committedEvents string) {
	transaction.eventStream = EventStream.eventstream
	transaction.committedEvents = committedEvents
}

func (transaction *TransactionBase) getEventStream() {
	return transaction.eventStream
}

func (transaction *TransactionBase) getCommittedEvents() {
	return transaction.committedEvents
}

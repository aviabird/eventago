package eventago

// Transaction Transaction
type Transaction interface {
	getEventStream()
	getCommittedEvents()
}

// TransactionBase TransactionBase
type TransactionBase struct {
	eventStream     EventStream
	committedEvents []string
}

func (transaction *TransactionBase) init(eventStream EventStream, committedEvents string) {
	transaction.eventStream = eventStream
	transaction.committedEvents = append(transaction.committedEvents, committedEvents)
}

func (transaction *TransactionBase) getEventStream() EventStream {
	return transaction.eventStream
}

func (transaction *TransactionBase) getCommittedEvents() []string {
	return transaction.committedEvents
}

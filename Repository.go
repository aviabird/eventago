package eventago

// Repository used for find and save method declaration.
type Repository interface {
	find(classname string, uuid string, expectedVersion int)
	save(AggregateRoot)
}

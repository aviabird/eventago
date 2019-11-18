package eventago

type Repository interface {
	find(classname string, uuid string, expectedVersion int)
	save(AggregateRoot)
}

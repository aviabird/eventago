package eventago

// Load Load
func Load(eid string) *EventStream {
	if eid != "" {
		return FindEntryQueryModel(eid)
	}
	return nil
}

func store(eid string, event string, newEvent string, version int) {
	err := createOrderQueryModel(eid, event, newEvent, version)
	if err != false {
		panic(err)
	}
}

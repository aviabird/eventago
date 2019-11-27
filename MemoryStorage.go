package eventago

// Load Load
func Load(eid string) *EventStream {
	if eid != "" {
		return FindEntryQueryModel(eid)
	}
	return nil
}

func store(eid string, event []string, newEvent []string, className string, version int) {
	eventdetails1 := createOrderQueryModel(eid, event, newEvent, className, version)
}

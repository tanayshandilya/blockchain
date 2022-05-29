package event

func GenesisEvent() *Event {
	e := new(Event)
	e.create(event_GenesisData, "origin")
	e.TimeStamp = event_GenesisTimestamp
	e.updateHash()
	return e
}

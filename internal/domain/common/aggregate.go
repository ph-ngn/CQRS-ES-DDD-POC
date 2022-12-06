package common

type Entity interface {
	GetID() string
}

type AggregateRoot interface {
	Entity
	Apply([]Event)
	AddEvent(Event)
	GetEvents() []Event
	ClearEvents()
}

type AggregateBase struct {
	id      string
	version int
	events  []Event
}

func (a *AggregateBase) GetID() string {
	return a.id
}

func (a *AggregateBase) AddEvent(event Event) {
	a.events = append(a.events, event)
}

func (a *AggregateBase) GetEvents() []Event {
	return a.events
}

func (a *AggregateBase) ClearEvents() {
	a.events = []Event{}
}

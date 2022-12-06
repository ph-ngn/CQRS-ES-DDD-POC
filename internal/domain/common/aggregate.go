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
	ID     string
	Events []Event
}

func (a *AggregateBase) GetID() string {
	return a.ID
}

func (a *AggregateBase) AddEvent(event Event) {
	a.Events = append(a.Events, event)
}

func (a *AggregateBase) GetEvents() []Event {
	return a.Events
}

func (a *AggregateBase) ClearEvents() {
	a.Events = []Event{}
}

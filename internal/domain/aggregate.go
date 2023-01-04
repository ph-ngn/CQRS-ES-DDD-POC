package domain

type Entity interface {
	GetID() string
}

type AggregateRoot interface {
	Entity
	When(Event, bool) error
	TrackChange(Event)
	GetChanges() []Event
	ClearChanges()
}

type AggregateBase struct {
	ID      string
	changes []Event
}

func (a *AggregateBase) GetID() string {
	return a.ID
}

func (a *AggregateBase) TrackChange(event Event) {
	a.changes = append(a.changes, event)
}

func (a *AggregateBase) GetChanges() []Event {
	return a.changes
}

func (a *AggregateBase) ClearChanges() {
	a.changes = []Event{}
}

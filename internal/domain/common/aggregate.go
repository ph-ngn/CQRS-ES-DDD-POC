package common

type Entity interface {
	GetID() string
}

type AggregateRoot interface {
	Entity
	Apply(Event)
	TrackChange(Event)
	GetChanges() []Event
	ClearChanges()
}

type AggregateBase struct {
	ID      string
	changes []Event
}

func NewAggregateBase(id string) *AggregateBase {
	return &AggregateBase{ID: id}
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

package common

type Entity interface {
	ID() string
}

type AggregateRoot interface {
	Entity
	Apply(Event, bool)
	TrackChange(Event)
	GetChanges() []Event
	ClearChanges()
	OriginalVersion() int
	CurrentVersion() int
	IncrementVersion()
}

type AggregateBase struct {
	id      string
	version int
	changes []Event
}

func (a *AggregateBase) ID() string {
	return a.id
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

func (a *AggregateBase) OriginalVersion() int {
	return a.version
}

func (a *AggregateBase) CurrentVersion() int {
	return a.version + len(a.changes)
}

func (a *AggregateBase) IncrementVersion() {
	a.version++
}

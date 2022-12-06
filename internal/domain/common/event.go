package common

import "reflect"

type Event interface {
	AggregateID() string
	EventType() string
}

type EventBase struct {
	aggregateID string
}

func (e *EventBase) AggregateID() string {
	return e.aggregateID
}

func (e *EventBase) EventType() string {
	return reflect.TypeOf(e).Elem().Name()
}

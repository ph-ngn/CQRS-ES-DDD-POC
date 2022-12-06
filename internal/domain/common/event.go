package common

import "reflect"

type Event interface {
	GetAggregateID() string
	GetEventType() string
}

type EventBase struct {
	AggregateID string
}

func (e *EventBase) GetAggregateID() string {
	return e.AggregateID
}

func (e *EventBase) GetEventType() string {
	return reflect.TypeOf(e).Elem().Name()
}

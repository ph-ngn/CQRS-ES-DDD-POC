package command

import "reflect"

type Interface interface {
	GetAggregateID() string
	GetCommandType() string
}

type base struct {
	AggregateID string
}

func (c *base) GetAggregateID() string {
	return c.AggregateID
}

func (c *base) GetCommandType() string {
	return reflect.TypeOf(c).Elem().Name()
}

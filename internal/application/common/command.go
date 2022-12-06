package common

import "reflect"

type Comnmand interface {
	GetAggregateID() string
	GetCommandType() string
}

type CommandBase struct {
	AggregateID string
}

func (c *CommandBase) GetAggregateID() string {
	return c.AggregateID
}

func (c *CommandBase) GetCommandType() string {
	return reflect.TypeOf(c).Elem().Name()
}

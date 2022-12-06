package common

import "reflect"

type Command interface {
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

type CommandHandler interface {
	Handle(Command)
}

type CommandDispatchcer interface {
	Dispatch(Command) error
	RegisterHandler(Command, CommandHandler) error
}

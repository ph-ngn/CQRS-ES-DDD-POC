package common

import (
	"reflect"
	"sync"
)

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
	Handle(Command) error
}

type Dispatchcer interface {
	Dispatch(Command) error
	RegisterHandler(Command, CommandHandler) error
}

type InMemoryDispatcher struct {
	handlers map[string]CommandHandler
}

func NewInMemoryDispatcher() *InMemoryDispatcher {
	return &InMemoryDispatcher{
		handlers: make(map[string]CommandHandler),
	}
}

func (d *InMemoryDispatcher) Dispatch(cmd Command) error {
	cmdMu.RLock()
	defer cmdMu.Unlock()

	if handler, ok := d.handlers[cmd.GetCommandType()]; ok {
		return handler.Handle(cmd)
	}
	return CommandHandlerNotFound
}

func (d *InMemoryDispatcher) RegisterHandler(cmd Command, handler CommandHandler) error {
	cmdMu.Lock()
	defer cmdMu.Unlock()

	if _, ok := d.handlers[cmd.GetCommandType()]; ok {
		return DuplicateCommandHandler
	}
	d.handlers[cmd.GetCommandType()] = handler
	return nil
}

var cmdMu sync.RWMutex
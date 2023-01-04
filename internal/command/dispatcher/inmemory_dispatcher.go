package dispatcher

import (
	"github.com/andyj29/wannabet/internal/command"
	"github.com/andyj29/wannabet/internal/command/handler"
	"sync"
)

type InMemoryDispatcher struct {
	handlers map[string]handler.Interface
	mu       sync.RWMutex
}

func NewInMemoryDispatcher() *InMemoryDispatcher {
	return &InMemoryDispatcher{
		handlers: make(map[string]handler.Interface),
	}
}

func (d *InMemoryDispatcher) Dispatch(cmd command.Interface) error {
	d.mu.RLock()
	defer d.mu.RUnlock()

	if handler, ok := d.handlers[cmd.GetCommandType()]; ok {
		return handler.Handle(cmd)
	}
	return CommandHandlerNotFound
}

func (d *InMemoryDispatcher) RegisterHandler(cmd command.Interface, handler handler.Interface) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if _, ok := d.handlers[cmd.GetCommandType()]; ok {
		return DuplicateCommandHandler
	}

	d.handlers[cmd.GetCommandType()] = handler
	return nil
}

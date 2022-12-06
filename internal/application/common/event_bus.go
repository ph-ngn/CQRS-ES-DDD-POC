package common

import "github.com/andyj29/wannabet/internal/domain/common"

type EventHandler interface {
	Handle(common.Event)
}

type EventBus interface {
	Publish(common.Event) error
	RegisterHandlers(common.Event, ...EventHandler) error
}

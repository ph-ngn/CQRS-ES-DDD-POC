package handler

import "github.com/andyj29/wannabet/internal/domain/common"

type Interface interface {
	Handle(common.Event)
}

package handler

import "github.com/andyj29/wannabet/internal/domain"

type Interface interface {
	Handle(domain.Event)
}

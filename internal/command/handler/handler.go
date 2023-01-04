package handler

import "github.com/andyj29/wannabet/internal/command"

type Interface interface {
	Handle(cmd command.Interface) error
}

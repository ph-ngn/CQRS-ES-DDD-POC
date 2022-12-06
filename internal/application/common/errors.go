package common

import "errors"

var (
	CommandHandlerNotFound  = errors.New("The command dispatcher does not have any registered handlers for this command")
	DuplicateCommandHandler = errors.New("A command handler for this command already exists")
)

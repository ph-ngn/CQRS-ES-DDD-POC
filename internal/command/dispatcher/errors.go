package dispatcher

import "errors"

var (
	CommandHandlerNotFound  = errors.New("the command dispatcher does not have any registered handlers for this command")
	DuplicateCommandHandler = errors.New("a command handler for this command already exists")
)

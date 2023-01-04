package domain

type ErrorType string

const (
	NotFoundError         ErrorType = "NotFound"
	InvalidOperationError           = "InvalidOperation"
	UnauthorizedError               = "Unauthorized"
)

type Error struct {
	err     error
	errType ErrorType
}

func (e Error) Error() string {
	return e.err.Error()
}

func (e Error) Type() ErrorType {
	return e.errType
}

func NewNotFoundError(err error) *Error {
	return &Error{
		err:     err,
		errType: NotFoundError,
	}
}

func NewInvalidOperationError(err error) *Error {
	return &Error{
		err:     err,
		errType: InvalidOperationError,
	}
}

func NewUnauthorizedError(err error) *Error {
	return &Error{
		err:     err,
		errType: UnauthorizedError,
	}
}

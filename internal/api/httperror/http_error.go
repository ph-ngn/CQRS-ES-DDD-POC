package httperror

import "net/http"

type HttpError struct {
	Error   error
	Status  int
	Message string
}

type ErrorResponse struct {
	Title  string `json:"title"`
	Status int    `json:"status"`
	Detail string `json:"detail"`
}

func (e *HttpError) Payload() *ErrorResponse {
	return &ErrorResponse{
		Title:  http.StatusText(e.Status),
		Status: e.Status,
		Detail: e.Message,
	}
}

func NewBadRequestError(err error) *HttpError {
	return &HttpError{
		Error:   err,
		Status:  http.StatusBadRequest,
		Message: err.Error(),
	}
}

func NewInternalError(err error) *HttpError {
	return &HttpError{
		Error:   err,
		Status:  http.StatusInternalServerError,
		Message: err.Error(),
	}
}

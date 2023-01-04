package controller

type response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func newResponse(success bool, message string) *response {
	return &response{
		Success: success,
		Message: message,
	}
}

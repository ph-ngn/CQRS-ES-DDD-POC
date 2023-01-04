package controller

type responseBase struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func newResponseBase(success bool, message string) *responseBase {
	return &responseBase{
		Success: success,
		Message: message,
	}
}

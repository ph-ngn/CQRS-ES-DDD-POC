package account

type accountRegistrationRequest struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type accountRegistrationResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func NewAccountRegistrationResponse(success bool, message string) *accountRegistrationResponse {
	return &accountRegistrationResponse{
		Success: success,
		Message: message,
	}
}

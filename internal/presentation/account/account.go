package account

type registerAccountRequest struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type registerAccountResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func NewRegisterAccountResponse(success bool, message string) *registerAccountResponse {
	return &registerAccountResponse{
		Success: success,
		Message: message,
	}
}

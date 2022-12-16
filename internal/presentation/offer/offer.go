package offer

type createOfferRequest struct {
	OffererID    string `json:"offerer_id"`
	GameID       string `json:"game_id"`
	Favorite     string `json:"favorite"`
	Limit        int64  `json:"limit"`
	CurrencyCode string `json:"currency_code"`
}

type createOfferResponse struct {
	Success      bool   `json:"success"`
	Message      string `json:"message"`
	GameID       string `json:"game_id"`
	Favorite     string `json:"favorite"`
	Limit        int64  `json:"limit"`
	CurrencyCode string `json:"currency_code"`
}

func newCreateOfferResponse(success bool, message, gameID, favorite string, limit int64, currencyCode string) *createOfferResponse {
	return &createOfferResponse{
		Success:      success,
		Message:      message,
		GameID:       gameID,
		Favorite:     favorite,
		Limit:        limit,
		CurrencyCode: currencyCode,
	}
}

type placeBetRequest struct {
	Success      bool   `json:"success"`
	Message      string `json:"message"`
	OfferID      string `json:"offer_id"`
	Stake        int64  `json:"stake"`
	CurrencyCode string `json:"currency_code"`
}

func newPlaceBetResponse(success bool, message, offerID string, stake int64, currencyCode string) *placeBetRequest {
	return &placeBetRequest{
		Success:      success,
		Message:      message,
		OfferID:      offerID,
		Stake:        stake,
		CurrencyCode: currencyCode,
	}
}

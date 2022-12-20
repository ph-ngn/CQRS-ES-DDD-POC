package offer

type createOfferRequest struct {
	BookMakerID  string `json:"book_maker_id"`
	FixtureID    string `json:"fixture_id"`
	HomeOdds     string `json:"home_odds"`
	AwayOdds     string `json:"away_odds"`
	Limit        int64  `json:"limit"`
	CurrencyCode string `json:"currency_code"`
}

type createOfferResponse struct {
	Success      bool   `json:"success"`
	Message      string `json:"message"`
	FixtureID    string `json:"fixture_id"`
	HomeOdds     string `json:"home_odds"`
	AwayOdds     string `json:"away_odds"`
	Limit        int64  `json:"limit"`
	CurrencyCode string `json:"currency_code"`
}

func newCreateOfferResponse(success bool, message, fixtureID, homeOdds, awayOdds string, limit int64, currencyCode string) *createOfferResponse {
	return &createOfferResponse{
		Success:      success,
		Message:      message,
		FixtureID:    fixtureID,
		HomeOdds:     homeOdds,
		AwayOdds:     awayOdds,
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

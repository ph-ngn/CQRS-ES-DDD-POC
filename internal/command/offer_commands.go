package command

type CreateOffer struct {
	*base
	BookMakerID, FixtureID, HomeOdds, AwayOdds string
	Limit                                      int64
	CurrencyCode                               string
}

func NewCreateOfferCommand(aggregateID, bookMakerID, fixtureID, homeOdds, awayOdds string, limit int64, currencyCode string) *CreateOffer {
	return &CreateOffer{
		base:         &base{AggregateID: aggregateID},
		BookMakerID:  bookMakerID,
		FixtureID:    fixtureID,
		HomeOdds:     homeOdds,
		AwayOdds:     awayOdds,
		Limit:        limit,
		CurrencyCode: currencyCode,
	}
}

type PlaceBet struct {
	*base
	BetID        string
	BettorID     string
	Stake        int64
	CurrencyCode string
}

func NewPlaceBetCommand(aggregateID, bettorID string, stake int64, currencyCode string) *PlaceBet {
	return &PlaceBet{
		base:         &base{AggregateID: aggregateID},
		BettorID:     bettorID,
		Stake:        stake,
		CurrencyCode: currencyCode,
	}
}
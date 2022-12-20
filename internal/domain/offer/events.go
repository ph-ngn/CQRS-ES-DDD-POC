package offer

import "github.com/andyj29/wannabet/internal/domain/common"

type offerCreated struct {
	*common.EventBase
	BookMakerID, FixtureID, HomeOdds, AwayOdds string
	Limit                                      common.Money
}

type betPlaced struct {
	*common.EventBase
	Bet *bet
}

func NewOfferCreatedEvent(aggregateID, bookMakerID, fixtureID, homeOdds, awayOdds string, limit common.Money) *offerCreated {
	return &offerCreated{
		EventBase:   &common.EventBase{AggregateID: aggregateID},
		BookMakerID: bookMakerID,
		FixtureID:   fixtureID,
		HomeOdds:    homeOdds,
		AwayOdds:    awayOdds,
		Limit:       limit,
	}
}

func NewBetPlacedEvent(aggregateID string, bet *bet) *betPlaced {
	return &betPlaced{
		EventBase: &common.EventBase{AggregateID: aggregateID},
		Bet:       bet,
	}
}

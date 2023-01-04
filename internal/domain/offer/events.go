package offer

import (
	"github.com/andyj29/wannabet/internal/domain"
)

type offerCreated struct {
	*domain.EventBase
	BookMakerID, FixtureID, HomeOdds, AwayOdds string
	Limit                                      domain.Money
}

type betPlaced struct {
	*domain.EventBase
	Bet *bet
}

func NewOfferCreatedEvent(aggregateID, bookMakerID, fixtureID, homeOdds, awayOdds string, limit domain.Money) *offerCreated {
	return &offerCreated{
		EventBase:   &domain.EventBase{AggregateID: aggregateID},
		BookMakerID: bookMakerID,
		FixtureID:   fixtureID,
		HomeOdds:    homeOdds,
		AwayOdds:    awayOdds,
		Limit:       limit,
	}
}

func NewBetPlacedEvent(aggregateID string, bet *bet) *betPlaced {
	return &betPlaced{
		EventBase: &domain.EventBase{AggregateID: aggregateID},
		Bet:       bet,
	}
}

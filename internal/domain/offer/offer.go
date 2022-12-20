package offer

import (
	"github.com/andyj29/wannabet/internal/domain/common"
)

var _ common.AggregateRoot = (*offer)(nil)
var _ Offer = (*offer)(nil)

type Offer interface {
	common.AggregateRoot
	PlaceBet(*bet) error
}

type offer struct {
	*common.AggregateBase
	BookMakerID, FixtureID, HomeOdds, AwayOdds string
	Limit                                      common.Money
	Bets                                       []*bet
}

func (o *offer) When(event common.Event, isNew bool) (err error) {
	switch e := event.(type) {
	case *offerCreated:
		o.onOfferCreated(e)

	case *betPlaced:
		err = o.onBetPlaced(e)
	}

	if isNew && err == nil {
		o.TrackChange(event)
	}
	return err
}

func NewOffer(id, bookMakerID, fixtureID, homeOdds, awayOdds string, limit common.Money) Offer {
	offerCreatedEvent := NewOfferCreatedEvent(id, bookMakerID, fixtureID, homeOdds, awayOdds, limit)
	newOffer := &offer{}
	newOffer.When(offerCreatedEvent, true)
	return newOffer
}

func (o *offer) PlaceBet(bet *bet) error {
	betPlacedEvent := NewBetPlacedEvent(o.GetID(), bet)
	return o.When(betPlacedEvent, true)
}

func (o *offer) onOfferCreated(event *offerCreated) {
	o.AggregateBase = &common.AggregateBase{ID: event.GetAggregateID()}
	o.BookMakerID = event.BookMakerID
	o.FixtureID = event.FixtureID
	o.HomeOdds = event.HomeOdds
	o.AwayOdds = event.AwayOdds
	o.Limit = event.Limit
}

func (o *offer) onBetPlaced(event *betPlaced) error {
	newLimit, err := o.Limit.Deduct(event.Bet.Stake)
	if err != nil {
		return err
	}

	o.Limit = newLimit
	o.Bets = append(o.Bets, event.Bet)
	return nil
}

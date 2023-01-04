package offer

import (
	"github.com/andyj29/wannabet/internal/domain"
)

var _ domain.AggregateRoot = (*Offer)(nil)

type Offer struct {
	*domain.AggregateBase
	BookMakerID, FixtureID, HomeOdds, AwayOdds string
	Limit                                      domain.Money
	Bets                                       []*bet
}

func (o *Offer) When(event domain.Event, isNew bool) (err error) {
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

func NewOffer(id, bookMakerID, fixtureID, homeOdds, awayOdds string, limit domain.Money) (*Offer, error) {
	offerCreatedEvent := NewOfferCreatedEvent(id, bookMakerID, fixtureID, homeOdds, awayOdds, limit)
	newOffer := &Offer{}
	if err := newOffer.When(offerCreatedEvent, true); err != nil {
		return &Offer{}, err
	}
	return newOffer, nil
}

func (o *Offer) PlaceBet(bet *bet) error {
	betPlacedEvent := NewBetPlacedEvent(o.GetID(), bet)
	return o.When(betPlacedEvent, true)
}

func (o *Offer) onOfferCreated(event *offerCreated) {
	o.AggregateBase = &domain.AggregateBase{ID: event.GetAggregateID()}
	o.BookMakerID = event.BookMakerID
	o.FixtureID = event.FixtureID
	o.HomeOdds = event.HomeOdds
	o.AwayOdds = event.AwayOdds
	o.Limit = event.Limit
}

func (o *Offer) onBetPlaced(event *betPlaced) error {
	newLimit, err := o.Limit.Deduct(event.Bet.Stake)
	if err != nil {
		return err
	}

	o.Limit = newLimit
	o.Bets = append(o.Bets, event.Bet)
	return nil
}

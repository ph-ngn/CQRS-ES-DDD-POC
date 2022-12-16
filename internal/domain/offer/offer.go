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
	OffererID string
	GameID    string
	Favorite  string
	Limit     common.Money
	Bets      []*bet
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

func NewOffer(id, offererID, gameID, favorite string, limit common.Money) Offer {
	offerCreatedEvent := NewOfferCreatedEvent(id, offererID, gameID, favorite, limit)
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
	o.OffererID = event.OffererID
	o.GameID = event.GameID
	o.Favorite = event.Favorite
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

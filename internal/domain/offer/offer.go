package offer

import (
	"github.com/andyj29/wannabet/internal/domain/common"
)

var _ common.AggregateRoot = (*Offer)(nil)

type Offer struct {
	*common.AggregateBase
	OffererID string
	GameID    string
	Favorite  string
	Limit     common.Money
	Bets      []*Bet
}

func NewOffer() *Offer {
	return &Offer{}
}

func (o *Offer) When(event common.Event, isNew bool) (err error) {
	if isNew {
		o.TrackChange(event)
	}

	switch e := event.(type) {
	case *OfferCreated:
		err = o.onOfferCreated(e)

	case *BetPlaced:
		err = o.onBetPlaced(e)
	}

	return err
}

func (o *Offer) onOfferCreated(event *OfferCreated) error {
	o.ID = event.GetAggregateID()
	o.OffererID = event.OffererID
	o.GameID = event.GameID
	o.Favorite = event.Favorite
	o.Limit = event.Limit

	return nil
}

func (o *Offer) onBetPlaced(event *BetPlaced) error {
	// TO BE VALIDATED
	o.Bets = append(o.Bets, event.Bet)
	return nil
}

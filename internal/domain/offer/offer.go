package offer

import (
	"github.com/andyj29/wannabet/internal/domain/common"
)

var _ common.AggregateRoot = (*offer)(nil)

type offer struct {
	*common.AggregateBase
	OffererID string
	GameID    string
	Favorite  string
	Limit     common.Money
	Bets      []*bet
}

func NewOffer() *offer {
	return &offer{}
}

func (o *offer) When(event common.Event, isNew bool) (err error) {
	switch e := event.(type) {
	case *offerCreated:
		err = o.onOfferCreated(e)

	case *betPlaced:
		err = o.onBetPlaced(e)
	}

	if isNew && err == nil {
		o.TrackChange(event)
	}

	return err
}

func (o *offer) onOfferCreated(event *offerCreated) error {
	o.ID = event.GetAggregateID()
	o.OffererID = event.OffererID
	o.GameID = event.GameID
	o.Favorite = event.Favorite
	o.Limit = event.Limit

	return nil
}

func (o *offer) onBetPlaced(event *betPlaced) error {
	// TO CHECK INVARIANT LATER
	o.Bets = append(o.Bets, event.Bet)
	return nil
}

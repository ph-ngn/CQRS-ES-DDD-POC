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

type OfferSettings struct {
	ID        string
	OffererID string
	GameID    string
	Favorite  string
	Limit     common.Money
}

func (o *Offer) Apply(event common.Event) {
	o.TrackChange(event)
	switch e := event.(type) {
	case OfferCreated:
		o.ID = e.GetAggregateID()
		o.OffererID = e.OffererID
		o.GameID = e.GameID
		o.Favorite = e.Favorite
		o.Limit = e.Limit

	case BetPlaced:
		o.Bets = append(o.Bets, e.Bet)
	}
}

func NewOffer(offerSettings OfferSettings) *Offer {
	OfferCreatedEvent := OfferCreated{}
	OfferCreatedEvent.AggregateID = offerSettings.ID
	OfferCreatedEvent.OffererID = offerSettings.OffererID
	OfferCreatedEvent.GameID = offerSettings.GameID
	OfferCreatedEvent.Favorite = offerSettings.Favorite
	OfferCreatedEvent.Limit = offerSettings.Limit

	offer := Offer{}
	offer.Apply(OfferCreatedEvent)

	return &offer
}

func (o *Offer) PlaceBet(bet *Bet) error {
	BetPlacedEvent := BetPlaced{}
	BetPlacedEvent.AggregateID = o.GetID()
	BetPlacedEvent.Bet = bet

	o.Apply(BetPlacedEvent)

	return nil
}

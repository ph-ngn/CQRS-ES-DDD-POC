package offer

import "github.com/andyj29/wannabet/internal/domain/common"

type offerCreated struct {
	*common.EventBase
	OffererID string
	GameID    string
	Favorite  string
	Limit     common.Money
}

type betPlaced struct {
	*common.EventBase
	Bet *bet
}

func NewOfferCreatedEvent(aggregateID, offererID, gameID, favorite string, limit common.Money) *offerCreated {
	return &offerCreated{
		EventBase: &common.EventBase{AggregateID: aggregateID},
		OffererID: offererID,
		GameID:    gameID,
		Favorite:  favorite,
		Limit:     limit,
	}
}

func NewBetPlacedEvent(aggregateID string, bet *bet) *betPlaced {
	return &betPlaced{
		EventBase: &common.EventBase{AggregateID: aggregateID},
		Bet:       bet,
	}
}

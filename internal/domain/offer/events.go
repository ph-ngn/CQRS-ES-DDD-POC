package offer

import "github.com/andyj29/wannabet/internal/domain/common"

type OfferCreated struct {
	*common.EventBase
	OffererID string
	GameID    string
	Favorite  string
	Limit     common.Money
}

type BetPlaced struct {
	*common.EventBase
	Bet *Bet
}

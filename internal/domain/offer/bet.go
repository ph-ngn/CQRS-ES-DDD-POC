package offer

import (
	"github.com/andyj29/wannabet/internal/domain"
)

type bet struct {
	BetID    string
	BettorID string
	Home     bool
	Away     bool
	Stake    domain.Money
}

func NewBet(betID, bettorID string, stake domain.Money) *bet {
	return &bet{
		BetID:    betID,
		BettorID: bettorID,
		Home:     false,
		Away:     false,
		Stake:    stake,
	}
}

func (b *bet) setAway() error {
	if b.Away || b.Home {
		return PickAlreadySet
	}
	b.Away = true
	return nil
}

func (b *bet) setHome() error {
	if b.Home || b.Away {
		return PickAlreadySet
	}
	b.Home = true
	return nil
}

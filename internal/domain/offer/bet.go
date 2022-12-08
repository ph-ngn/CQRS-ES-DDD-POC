package offer

import "github.com/andyj29/wannabet/internal/domain/common"

type bet struct {
	BettorID string
	Stake    common.Money
}

func NewBet(bettorID string, stake common.Money) *bet {
	return &bet{
		BettorID: bettorID,
		Stake:    stake,
	}
}

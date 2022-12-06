package account

import "github.com/andyj29/wannabet/internal/domain/common"

type FundsAdded struct {
	common.EventBase
	Funds Funds
}

type FundsUsed struct {
	common.EventBase
	amount int64
}

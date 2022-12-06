package account

import "github.com/andyj29/wannabet/internal/domain/common"

type AccountCreated struct {
	*common.EventBase
	Email Email
	Name  string
}

type FundsAdded struct {
	*common.EventBase
	Funds Money
}

type FundsUsed struct {
	*common.EventBase
	amount Money
}

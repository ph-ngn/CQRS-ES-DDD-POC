package account

import "github.com/andyj29/wannabet/internal/domain/common"

type AccountCreated struct {
	*common.EventBase
	Email Email
	Name  string
}

type FundsAdded struct {
	*common.EventBase
	Funds common.Money
}

type FundsDeducted struct {
	*common.EventBase
	Amount common.Money
}

func NewAccountCreatedEvent(aggregateID, name string, email Email) *AccountCreated {
	return &AccountCreated{
		EventBase: &common.EventBase{AggregateID: aggregateID},
		Email:     email,
		Name:      name,
	}
}

func NewFundsAddedEvent(aggregateID string, funds common.Money) *FundsAdded {
	return &FundsAdded{
		EventBase: &common.EventBase{AggregateID: aggregateID},
		Funds:     funds,
	}
}

func NewFundsDeductedEvent(aggregateID string, amount common.Money) *FundsDeducted {
	return &FundsDeducted{
		EventBase: &common.EventBase{AggregateID: aggregateID},
		Amount:    amount,
	}
}

package account

import "github.com/andyj29/wannabet/internal/domain/common"

type accountCreated struct {
	*common.EventBase
	Email Email
	Name  string
}

type fundsAdded struct {
	*common.EventBase
	Funds common.Money
}

type fundsDeducted struct {
	*common.EventBase
	Amount common.Money
}

func NewAccountCreatedEvent(aggregateID, name string, email Email) *accountCreated {
	return &accountCreated{
		EventBase: &common.EventBase{AggregateID: aggregateID},
		Email:     email,
		Name:      name,
	}
}

func NewFundsAddedEvent(aggregateID string, funds common.Money) *fundsAdded {
	return &fundsAdded{
		EventBase: &common.EventBase{AggregateID: aggregateID},
		Funds:     funds,
	}
}

func NewFundsDeductedEvent(aggregateID string, amount common.Money) *fundsDeducted {
	return &fundsDeducted{
		EventBase: &common.EventBase{AggregateID: aggregateID},
		Amount:    amount,
	}
}

package account

import (
	"github.com/andyj29/wannabet/internal/domain"
)

type accountCreated struct {
	*domain.EventBase
	Email Email
	Name  string
}

type fundsAdded struct {
	*domain.EventBase
	Funds domain.Money
}

type fundsDeducted struct {
	*domain.EventBase
	Amount domain.Money
}

func NewAccountCreatedEvent(aggregateID string, email Email, name string) *accountCreated {
	return &accountCreated{
		EventBase: &domain.EventBase{AggregateID: aggregateID},
		Email:     email,
		Name:      name,
	}
}

func NewFundsAddedEvent(aggregateID string, funds domain.Money) *fundsAdded {
	return &fundsAdded{
		EventBase: &domain.EventBase{AggregateID: aggregateID},
		Funds:     funds,
	}
}

func NewFundsDeductedEvent(aggregateID string, amount domain.Money) *fundsDeducted {
	return &fundsDeducted{
		EventBase: &domain.EventBase{AggregateID: aggregateID},
		Amount:    amount,
	}
}

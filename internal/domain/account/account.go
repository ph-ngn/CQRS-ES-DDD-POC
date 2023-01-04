package account

import (
	"github.com/andyj29/wannabet/internal/domain"
)

var _ domain.AggregateRoot = (*Account)(nil)

type Email string

type Account struct {
	*domain.AggregateBase
	Email   Email
	Name    string
	Balance domain.Money
}

func (a *Account) When(event domain.Event, isNew bool) (err error) {
	switch e := event.(type) {
	case *accountCreated:
		err = a.onAccountCreated(e)

	case *fundsAdded:
		err = a.onFundsAdded(e)

	case *fundsDeducted:
		err = a.onFundsDeducted(e)
	}

	if isNew && err == nil {
		a.TrackChange(event)
	}
	return err
}

func NewAccount(id string, email Email, name string) (*Account, error) {
	accountCreatedEvent := NewAccountCreatedEvent(id, email, name)
	newAccount := &Account{}
	if err := newAccount.When(accountCreatedEvent, true); err != nil {
		return &Account{}, err
	}
	return newAccount, nil
}

func (a *Account) AddFunds(funds domain.Money) error {
	fundsAddedEvent := NewFundsAddedEvent(a.GetID(), funds)
	return a.When(fundsAddedEvent, true)
}

func (a *Account) DeductFunds(amount domain.Money) error {
	fundsDeductedEvent := NewFundsDeductedEvent(a.GetID(), amount)
	return a.When(fundsDeductedEvent, true)
}

func (a *Account) onAccountCreated(event *accountCreated) error {
	a.AggregateBase = &domain.AggregateBase{ID: event.GetAggregateID()}
	a.Email = event.Email
	a.Name = event.Name
	initialBalance, err := domain.NewMoney(0, "CAD")
	if err != nil {
		return err
	}
	a.Balance = initialBalance
	return nil

}

func (a *Account) onFundsAdded(event *fundsAdded) error {
	newBalance, err := a.Balance.Add(event.Funds)
	if err != nil {
		return err
	}

	a.Balance = newBalance
	return nil
}

func (a *Account) onFundsDeducted(event *fundsDeducted) error {
	newBalance, err := a.Balance.Deduct(event.Amount)
	if err != nil {
		return err
	}

	a.Balance = newBalance
	return nil
}

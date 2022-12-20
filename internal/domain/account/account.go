package account

import (
	"github.com/andyj29/wannabet/internal/domain/common"
)

var _ common.AggregateRoot = (*account)(nil)
var _ Account = (*account)(nil)

type Email string

type Account interface {
	common.AggregateRoot
	AddFunds(common.Money) error
	DeductFunds(common.Money) error
}

type account struct {
	*common.AggregateBase
	Email   Email
	Name    string
	Balance common.Money
}

func (a *account) When(event common.Event, isNew bool) (err error) {
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

func NewAccount(id string, email Email, name string) (*account, error) {
	accountCreatedEvent := NewAccountCreatedEvent(id, email, name)
	newAccount := &account{}
	if err := newAccount.When(accountCreatedEvent, true); err != nil {
		return &account{}, err
	}
	return newAccount, nil
}

func (a *account) AddFunds(funds common.Money) error {
	fundsAddedEvent := NewFundsAddedEvent(a.GetID(), funds)
	return a.When(fundsAddedEvent, true)
}

func (a *account) DeductFunds(amount common.Money) error {
	fundsDeductedEvent := NewFundsDeductedEvent(a.GetID(), amount)
	return a.When(fundsDeductedEvent, true)
}

func (a *account) onAccountCreated(event *accountCreated) error {
	a.AggregateBase = &common.AggregateBase{ID: event.GetAggregateID()}
	a.Email = event.Email
	a.Name = event.Name
	initialBalance, err := common.NewMoney(0, "CAD")
	if err != nil {
		return err
	}
	a.Balance = initialBalance
	return nil

}

func (a *account) onFundsAdded(event *fundsAdded) error {
	newBalance, err := a.Balance.Add(event.Funds)
	if err != nil {
		return err
	}

	a.Balance = newBalance
	return nil
}

func (a *account) onFundsDeducted(event *fundsDeducted) error {
	newBalance, err := a.Balance.Deduct(event.Amount)
	if err != nil {
		return err
	}

	a.Balance = newBalance
	return nil
}

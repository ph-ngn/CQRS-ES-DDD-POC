package account

import "github.com/andyj29/wannabet/internal/domain/common"

var _ common.AggregateRoot = (*Account)(nil)

type Email string

func EmailFromString(str string) Email {
	return Email(str)
}

type Account struct {
	*common.AggregateBase
	Email   Email
	Name    string
	Balance common.Money
}

func (a *Account) When(event common.Event, isNew bool) (err error) {
	if isNew {
		a.TrackChange(event)
	}

	switch e := event.(type) {
	case *AccountCreated:
		a.onAccountCreated(e)

	case *FundsAdded:
		err = a.onFundsAdded(e)

	case *FundsDeducted:
		err = a.onFundsDeducted(e)
	}

	return err
}

func (a *Account) onAccountCreated(event *AccountCreated) {
	a.ID = event.GetAggregateID()
	a.Email = event.Email
	a.Name = event.Name
	a.Balance = common.NewMoney(0, "CAD")
}

func (a *Account) onFundsAdded(event *FundsAdded) error {
	if err := a.Balance.Add(event.Funds); err != nil {
		return err
	}
	return nil
}

func (a *Account) onFundsDeducted(event *FundsDeducted) error {
	if err := a.Balance.Deduct(event.Amount); err != nil {
		return err
	}
	return nil
}

func NewAccount() *Account {
	return &Account{}
}

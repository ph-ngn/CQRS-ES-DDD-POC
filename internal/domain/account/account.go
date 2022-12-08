package account

import "github.com/andyj29/wannabet/internal/domain/common"

var _ common.AggregateRoot = (*account)(nil)

type Email string

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

func (a *account) onAccountCreated(event *accountCreated) error {
	a.ID = event.GetAggregateID()
	a.Email = event.Email
	a.Name = event.Name
	a.Balance = common.NewMoney(0, "CAD")

	return nil
}

func (a *account) onFundsAdded(event *fundsAdded) error {
	if err := a.Balance.Add(event.Funds); err != nil {
		return err
	}
	return nil
}

func (a *account) onFundsDeducted(event *fundsDeducted) error {
	if err := a.Balance.Deduct(event.Amount); err != nil {
		return err
	}
	return nil
}

func NewAccount() *account {
	return &account{}
}

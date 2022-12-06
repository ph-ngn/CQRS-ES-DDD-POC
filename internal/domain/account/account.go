package account

import "github.com/andyj29/wannabet/internal/domain/common"

var _ common.AggregateRoot = (*Account)(nil)

type Account struct {
	common.AggregateBase
	Email   string
	Name    string
	Balance Funds
}

func (a *Account) Apply(events []common.Event) {
	return
}
func (a *Account) AddFunds(funds Funds) error {
	if err := a.Balance.Add(funds); err != nil {
		return err
	}

	fundsAddedEvent := FundsAdded{}
	fundsAddedEvent.AggregateID = a.GetID()
	fundsAddedEvent.Funds = funds

	a.AddEvent(&fundsAddedEvent)
	return nil
}

func (a *Account) UseFunds(amount int64) error {
	if err := a.Balance.Deduct(amount); err != nil {
		return err
	}

	FundsUsedEvent := FundsUsed{}
	FundsUsedEvent.AggregateID = a.GetID()
	FundsUsedEvent.amount = amount

	a.AddEvent(&FundsUsedEvent)
	return nil
}

func New(email, name string) *Account {
	return &Account{
		Email:   email,
		Name:    name,
		Balance: Funds{Currency: Currency{Name: "Canadian Dollar", Code: "CAD"}},
	}
}

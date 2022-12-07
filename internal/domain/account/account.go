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

func (a *Account) Apply(event common.Event) {
	a.TrackChange(event)
	switch e := event.(type) {
	case AccountCreated:
		a.ID = e.GetAggregateID()
		a.Email = e.Email
		a.Name = e.Name
		a.Balance = common.Money{Currency: common.CurrencyFromCode("CAD")}

	case FundsAdded:
		a.Balance.Add(e.Funds.Amount)

	case FundsUsed:
		a.Balance.Deduct(e.amount.Amount)
	}
}

func (a *Account) AddFunds(funds common.Money) error {
	if !a.Balance.CanBeAdded(funds) {
		return FundsNotAddable
	}

	fundsAddedEvent := FundsAdded{}
	fundsAddedEvent.AggregateID = a.GetID()
	fundsAddedEvent.Funds = funds

	a.Apply(fundsAddedEvent)

	return nil
}

func (a *Account) UseFunds(amount common.Money) error {
	if !a.Balance.CanBeDeducted(amount) {
		return FundsNotDeductible
	}

	FundsUsedEvent := FundsUsed{}
	FundsUsedEvent.AggregateID = a.GetID()
	FundsUsedEvent.amount = amount

	a.Apply(FundsUsedEvent)

	return nil
}

func New(id, name string, email Email) *Account {
	AccountCreatedEvent := AccountCreated{}
	AccountCreatedEvent.AggregateID = id
	AccountCreatedEvent.Email = email
	AccountCreatedEvent.Name = name

	account := Account{}
	account.Apply(AccountCreatedEvent)

	return &account
}

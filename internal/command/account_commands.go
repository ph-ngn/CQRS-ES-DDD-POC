package command

type RegisterAccount struct {
	*base
	Email string
	Name  string
}

func NewRegisterAccountCommand(aggregateID, email, name string) *RegisterAccount {
	return &RegisterAccount{
		base:  &base{AggregateID: aggregateID},
		Email: email,
		Name:  name,
	}
}

type AddFunds struct {
	*base
	Amount       int64
	CurrencyCode string
}

func NewAddFundsCommand(aggregateID string, amount int64, currencyCode string) *AddFunds {
	return &AddFunds{
		base:         &base{AggregateID: aggregateID},
		Amount:       amount,
		CurrencyCode: currencyCode,
	}
}

type DeductFunds struct {
	*base
	Amount       int64
	CurrencyCode string
}

func NewDeductFundsCommand(aggregateID string, amount int64, currencyCode string) *DeductFunds {
	return &DeductFunds{
		base:         &base{AggregateID: aggregateID},
		Amount:       amount,
		CurrencyCode: currencyCode,
	}
}

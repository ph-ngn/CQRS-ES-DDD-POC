package account

type Account struct {
	ID      string
	Email   string
	Name    string
	Balance Funds
}

func (a *Account) AddFunds(funds Funds) error {
	return a.Balance.Add(funds)
}

func (a *Account) UseFunds(amount int64) error {
	return a.Balance.Deduct(amount)
}

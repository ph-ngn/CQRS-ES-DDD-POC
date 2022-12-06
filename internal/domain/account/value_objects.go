package account

type Currency struct {
	Name string
	Code string
}

type Funds struct {
	Amount   int64
	Currency Currency
}

func (f *Funds) EqualTo(other Funds) bool {
	return f.Amount == other.Amount && f.Currency == other.Currency
}

func (f *Funds) GreaterThan(other Funds) (bool, error) {
	if f.Currency != other.Currency {
		return false, IncompatibleCurrency
	}
	return f.Amount > other.Amount, nil
}

func (f *Funds) LessThan(other Funds) (bool, error) {
	if f.Currency != other.Currency {
		return false, IncompatibleCurrency
	}
	return f.Amount < other.Amount, nil
}

func (f *Funds) Add(other Funds) error {
	if f.Currency != other.Currency {
		return IncompatibleCurrency
	}
	f.Amount += other.Amount
	return nil
}

func (f *Funds) Deduct(amount int64) error {
	if f.Amount < amount {
		return InsufficientBalance
	}
	f.Amount -= amount
	return nil
}

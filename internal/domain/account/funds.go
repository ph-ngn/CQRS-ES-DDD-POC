package account

type Currency struct {
	Name string
	Code string
}

func (c *Currency) Is(other Currency) bool {
	return c.Name == other.Name && c.Code == other.Code
}

type Funds struct {
	Amount   int64
	Currency Currency
}

func (f *Funds) EqualTo(other Funds) bool {
	return f.Amount == other.Amount && f.Currency.Is(other.Currency)
}

func (f *Funds) GreaterThan(other Funds) bool {
	return f.Amount > other.Amount && f.Currency.Is(other.Currency)
}

func (f *Funds) LessThan(other Funds) bool {
	return !f.GreaterThan(other) && !f.EqualTo(other)
}

func (f *Funds) Add(other Funds) error {
	if !f.Currency.Is(other.Currency) {
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

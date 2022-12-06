package account

type Email string

func EmailFromString(str string) Email {
	return Email(str)
}

type Currency struct {
	Name string
	Code string
}

type Funds struct {
	Amount   int64
	Currency Currency
}

func (f *Funds) IsCompatible(other Funds) bool {
	return f.Currency == other.Currency
}

func (f *Funds) CanBeAdded(other Funds) bool {
	return f.IsCompatible(other) && other.Amount > 0
}

func (f *Funds) CanBeDeducted(other Funds) bool {
	return f.IsCompatible(other) && f.Amount >= other.Amount
}

func (f *Funds) EqualTo(other Funds) bool {
	return f.Amount == other.Amount
}

func (f *Funds) GreaterThan(other Funds) bool {
	return f.Amount > other.Amount
}

func (f *Funds) LessThan(other Funds) bool {
	return f.Amount < other.Amount
}

func (f *Funds) Add(amount int64) {
	f.Amount += amount
}

func (f *Funds) Deduct(amount int64) {
	f.Amount -= amount
}

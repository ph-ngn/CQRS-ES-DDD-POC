package common

type Currency struct {
	Name string
	Code string
}

func CurrencyFromCode(code string) Currency {
	return Currency{Name: "Canadian Dollar", Code: "CAD"}
}

type Money struct {
	Amount   int64
	Currency Currency
}

func NewMoney(amount int64, code string) Money {
	return Money{Amount: amount, Currency: CurrencyFromCode(code)}
}

func (f *Money) IsCompatible(other Money) bool {
	return f.Currency == other.Currency
}

func (f *Money) CanBeAdded(other Money) bool {
	return f.IsCompatible(other) && other.Amount > 0
}

func (f *Money) CanBeDeducted(other Money) bool {
	return f.IsCompatible(other) && f.Amount >= other.Amount
}

func (f *Money) EqualTo(other Money) bool {
	return f.Amount == other.Amount
}

func (f *Money) GreaterThan(other Money) bool {
	return f.Amount > other.Amount
}

func (f *Money) LessThan(other Money) bool {
	return f.Amount < other.Amount
}

func (f *Money) Add(amount int64) {
	f.Amount += amount
}

func (f *Money) Deduct(amount int64) {
	f.Amount -= amount
}

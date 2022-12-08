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

func (f *Money) isCompatible(other Money) bool {
	return f.Currency == other.Currency
}

func (f *Money) canBeAdded(other Money) bool {
	return f.isCompatible(other) && other.Amount > 0
}

func (f *Money) canBeDeducted(other Money) bool {
	return f.isCompatible(other) && f.Amount >= other.Amount && other.Amount > 0
}

func (f *Money) Add(funds Money) error {
	if ok := f.canBeAdded(funds); ok {
		f.Amount += funds.Amount
		return nil
	}
	return FundsNotAddable
}

func (f *Money) Deduct(amount Money) error {
	if ok := f.canBeDeducted(amount); ok {
		f.Amount -= amount.Amount
		return nil
	}
	return FundsNotDeductible
}

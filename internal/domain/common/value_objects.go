package common

type Currency struct {
	Name string
	Code string
}

func CurrencyFromCode(code string) Currency {
	switch code {
	case "CAD":
		return Currency{Name: "Canadian Dollar", Code: "CAD"}
	case "VND":
		return Currency{Name: "Vietnamese Dong", Code: "VND"}
	default:
		return Currency{Name: "Canadian Dollar", Code: "CAD"}
	}
}

type Money struct {
	Amount   int64
	Currency Currency
}

func NewMoney(amount int64, code string) (Money, error) {
	if amount < 0 || code == "" {
		return Money{}, InvalidAmount
	}
	return Money{Amount: amount, Currency: CurrencyFromCode(code)}, nil
}

func (f Money) isCompatible(other Money) bool {
	return f.Currency == other.Currency
}

func (f Money) canBeAdded(other Money) bool {
	return f.isCompatible(other) && other.Amount > 0
}

func (f Money) canBeDeducted(other Money) bool {
	return f.isCompatible(other) && f.Amount >= other.Amount && other.Amount > 0
}

func (f Money) Add(funds Money) (Money, error) {
	if ok := f.canBeAdded(funds); ok {
		f.Amount += funds.Amount
		return f, nil
	}
	return Money{}, FundsNotAddable
}

func (f Money) Deduct(amount Money) (Money, error) {
	if ok := f.canBeDeducted(amount); ok {
		f.Amount -= amount.Amount
		return f, nil
	}
	return Money{}, FundsNotDeductible
}

func (f Money) IsEqual(other Money) bool {
	return f.isCompatible(other) && f.Amount == other.Amount
}

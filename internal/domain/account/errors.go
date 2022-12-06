package account

import "errors"

var (
	IncompatibleCurrency = errors.New("You can not add funds of different type of currency than that of your account")
	UnsupportedCurrency  = errors.New("Oops! We're not supporting this currency at this moment")
	InsufficientBalance  = errors.New("Oops! Your balance is insufficient to make the transaction")
)

package common

import "errors"

var (
	FundsNotAddable    = errors.New("Funds can not be added because it might have incompatible currency or invalid amount")
	FundsNotDeductible = errors.New("Funds can not be deducted because it might have incompatible currency or is insufficient")
)

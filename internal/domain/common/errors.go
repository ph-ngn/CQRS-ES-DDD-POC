package common

import "errors"

var (
	FundsNotAddable    = errors.New("funds can not be added because it might have incompatible currency or invalid amount")
	FundsNotDeductible = errors.New("funds can not be deducted because it might have incompatible currency or is insufficient")
	InvalidAmount      = errors.New(" amount can not be negative")
)

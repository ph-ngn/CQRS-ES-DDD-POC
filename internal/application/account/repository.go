package account

import (
	"github.com/andyj29/wannabet/internal/domain/account"
)

type Repository interface {
	Load(string) account.Account
	Save(account.Account) error
}

package account

import "github.com/andyj29/wannabet/internal/domain/account"

type AccountRepository interface {
	Load(string) *account.Account
	Save(*account.Account) error
}

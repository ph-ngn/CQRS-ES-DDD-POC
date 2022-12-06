package account

import (
	"github.com/andyj29/wannabet/internal/application/common"
	"github.com/andyj29/wannabet/internal/domain/account"
)

type RegisterAccount struct {
	*common.CommandBase
	Email string
	Name  string
}

type RegisterAccountHandler struct {
	repo     AccountRepository
	eventBus common.EventBus
}

func (h *RegisterAccountHandler) Handle(cmd RegisterAccount) error {
	account := account.New(cmd.GetAggregateID(), cmd.Name, account.EmailFromString(cmd.Email))
	return h.repo.Save(account)
}

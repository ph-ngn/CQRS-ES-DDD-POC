package handler

import (
	"fmt"

	"github.com/andyj29/wannabet/internal/application/common"
	"github.com/andyj29/wannabet/internal/command"
	"github.com/andyj29/wannabet/internal/domain/account"
)

type RegisterAccountHandler struct {
	Repo     common.Repository[*account.Account]
	EventBus common.EventBus
}

func (h *RegisterAccountHandler) Handle(cmd common.Command) error {
	c, ok := cmd.(*command.RegisterAccount)
	if !ok {
		panic(fmt.Sprintf("Unexpected command type %T", cmd))
	}

	newAccount, err := account.NewAccount(cmd.GetAggregateID(), account.Email(c.Email), c.Name)
	if err != nil {
		return err
	}

	if err := h.Repo.Save(newAccount); err != nil {
		return err
	}

	for _, event := range newAccount.GetChanges() {
		h.EventBus.Publish(event)
	}
	return nil
}

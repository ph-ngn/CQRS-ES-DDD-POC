package handler

import (
	"fmt"
	"github.com/andyj29/wannabet/internal/command"
	"github.com/andyj29/wannabet/internal/domain/account"
	"github.com/andyj29/wannabet/internal/eventbus"
	"github.com/andyj29/wannabet/internal/repository"
)

type RegisterAccountHandler struct {
	Repo     repository.Interface[*account.Account]
	EventBus eventbus.Interface
}

func (h *RegisterAccountHandler) Handle(cmd command.Interface) error {
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

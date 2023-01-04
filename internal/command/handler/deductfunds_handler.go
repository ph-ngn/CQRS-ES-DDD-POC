package handler

import (
	"fmt"
	"github.com/andyj29/wannabet/internal/domain"
	"github.com/andyj29/wannabet/internal/eventbus"
	"github.com/andyj29/wannabet/internal/repository"

	"github.com/andyj29/wannabet/internal/command"
	"github.com/andyj29/wannabet/internal/domain/account"
)

type DeductFundsHandler struct {
	Repo     repository.Interface[*account.Account]
	EventBus eventbus.Interface
}

func (h *DeductFundsHandler) Handle(cmd command.Interface) error {
	c, ok := cmd.(*command.DeductFunds)
	if !ok {
		panic(fmt.Sprintf("Unexpected command type %T", cmd))
	}

	loadedAccount, err := h.Repo.Load(cmd.GetAggregateID())
	if err != nil {
		return err
	}

	amount, err := domain.NewMoney(c.Amount, c.CurrencyCode)
	if err != nil {
		return err
	}

	if err := loadedAccount.DeductFunds(amount); err != nil {
		return err
	}

	if err := h.Repo.Save(loadedAccount); err != nil {
		return err
	}

	for _, event := range loadedAccount.GetChanges() {
		h.EventBus.Publish(event)
	}
	return nil
}

package handler

import (
	"fmt"

	"github.com/andyj29/wannabet/internal/application/common"
	"github.com/andyj29/wannabet/internal/command"
	"github.com/andyj29/wannabet/internal/domain/account"
	dc "github.com/andyj29/wannabet/internal/domain/common"
)

type DeductFundsHandler struct {
	Repo     common.Repository[*account.Account]
	EventBus common.EventBus
}

func (h *DeductFundsHandler) Handle(cmd common.Command) error {
	c, ok := cmd.(*command.DeductFunds)
	if !ok {
		panic(fmt.Sprintf("Unexpected command type %T", cmd))
	}

	loadedAccount, err := h.Repo.Load(cmd.GetAggregateID())
	if err != nil {
		return err
	}

	amount, err := dc.NewMoney(c.Amount, c.CurrencyCode)
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

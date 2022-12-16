package account

import (
	"github.com/andyj29/wannabet/internal/application/common"
	"github.com/andyj29/wannabet/internal/domain/account"
	domainCommon "github.com/andyj29/wannabet/internal/domain/common"
)

type registerAccount struct {
	*common.CommandBase
	Email string
	Name  string
}

type RegisterAccountHandler struct {
	repo     Repository
	eventBus common.EventBus
}

type addFunds struct {
	*common.CommandBase
	Amount       int64
	CurrencyCode string
}

type AddFundsHandler struct {
	repo     Repository
	eventBus common.EventBus
}

type deductFunds struct {
	*common.CommandBase
	Amount       int64
	CurrencyCode string
}

type DeductFundsHandler struct {
	repo     Repository
	eventBus common.EventBus
}

func NewRegisterAccountCommand(aggregateID, email, name string) *registerAccount {
	return &registerAccount{
		CommandBase: &common.CommandBase{AggregateID: aggregateID},
		Email:       email,
		Name:        name,
	}
}

func NewAddFundsCommand(aggregateID string, amount int64, currencyCode string) *addFunds {
	return &addFunds{
		CommandBase:  &common.CommandBase{AggregateID: aggregateID},
		Amount:       amount,
		CurrencyCode: currencyCode,
	}
}

func NewDeductFundsCommand(aggregateID string, amount int64, currencyCode string) *deductFunds {
	return &deductFunds{
		CommandBase:  &common.CommandBase{AggregateID: aggregateID},
		Amount:       amount,
		CurrencyCode: currencyCode,
	}
}

func (h *RegisterAccountHandler) Handle(cmd registerAccount) error {
	newAccount := account.NewAccount(cmd.GetAggregateID(), account.Email(cmd.Email), cmd.Name)
	if err := h.repo.Save(newAccount); err != nil {
		return err
	}

	for _, event := range newAccount.GetChanges() {
		h.eventBus.Publish(event)
	}
	return nil
}

func (h *AddFundsHandler) Handle(cmd addFunds) error {
	loadedAccount := h.repo.Load(cmd.GetAggregateID())
	if err := loadedAccount.AddFunds(domainCommon.NewMoney(cmd.Amount, cmd.CurrencyCode)); err != nil {
		return err
	}

	if err := h.repo.Save(loadedAccount); err != nil {
		return err
	}

	for _, event := range loadedAccount.GetChanges() {
		h.eventBus.Publish(event)
	}
	return nil
}

func (h *DeductFundsHandler) Handle(cmd deductFunds) error {
	loadedAccount := h.repo.Load(cmd.GetAggregateID())
	if err := loadedAccount.DeductFunds(domainCommon.NewMoney(cmd.Amount, cmd.CurrencyCode)); err != nil {
		return err
	}

	if err := h.repo.Save(loadedAccount); err != nil {
		return err
	}

	for _, event := range loadedAccount.GetChanges() {
		h.eventBus.Publish(event)
	}
	return nil
}

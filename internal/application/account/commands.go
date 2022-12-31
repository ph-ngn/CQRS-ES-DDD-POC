package account

import (
	"github.com/andyj29/wannabet/internal/application/common"
	"github.com/andyj29/wannabet/internal/domain/account"
	domainCommon "github.com/andyj29/wannabet/internal/domain/common"
)

var _ common.Command = &RegisterAccount{}

type RegisterAccount struct {
	*common.CommandBase
	Email string
	Name  string
}

type RegisterAccountHandler struct {
	Repo     common.Repository[*account.Account]
	EventBus common.EventBus
}

type AddFunds struct {
	*common.CommandBase
	Amount       int64
	CurrencyCode string
}

type AddFundsHandler struct {
	Repo     common.Repository[*account.Account]
	EventBus common.EventBus
}

type DeductFunds struct {
	*common.CommandBase
	Amount       int64
	CurrencyCode string
}

type DeductFundsHandler struct {
	Repo     common.Repository[*account.Account]
	EventBus common.EventBus
}

func NewRegisterAccountCommand(aggregateID, email, name string) *RegisterAccount {
	return &RegisterAccount{
		CommandBase: &common.CommandBase{AggregateID: aggregateID},
		Email:       email,
		Name:        name,
	}
}

func NewAddFundsCommand(aggregateID string, amount int64, currencyCode string) *AddFunds {
	return &AddFunds{
		CommandBase:  &common.CommandBase{AggregateID: aggregateID},
		Amount:       amount,
		CurrencyCode: currencyCode,
	}
}

func NewDeductFundsCommand(aggregateID string, amount int64, currencyCode string) *DeductFunds {
	return &DeductFunds{
		CommandBase:  &common.CommandBase{AggregateID: aggregateID},
		Amount:       amount,
		CurrencyCode: currencyCode,
	}
}

func (h *RegisterAccountHandler) Handle(cmd common.Command) error {
	c, ok := cmd.(*RegisterAccount)
	if !ok {
		return nil
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

func (h *AddFundsHandler) Handle(cmd common.Command) error {
	c, ok := cmd.(*AddFunds)
	if !ok {
		return nil
	}
	loadedAccount, err := h.Repo.Load(cmd.GetAggregateID())
	if err != nil {
		return err
	}
	amount, err := domainCommon.NewMoney(c.Amount, c.CurrencyCode)
	if err != nil {
		return err
	}
	if err := loadedAccount.AddFunds(amount); err != nil {
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

func (h *DeductFundsHandler) Handle(cmd common.Command) error {
	c, ok := cmd.(*DeductFunds)
	if !ok {
		return nil
	}
	loadedAccount, err := h.Repo.Load(cmd.GetAggregateID())
	if err != nil {
		return err
	}
	amount, err := domainCommon.NewMoney(c.Amount, c.CurrencyCode)
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

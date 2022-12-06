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

type DepositFunds struct {
	*common.CommandBase
	Amount       int64
	CurrencyCode string
}

type DepositFundsHandler struct {
	repo     AccountRepository
	eventBus common.EventBus
}

func (h *DepositFundsHandler) Handle(cmd DepositFunds) error {
	loadedAccount := h.repo.Load(cmd.GetAggregateID())
	funds := account.Money{
		Amount:   cmd.Amount,
		Currency: account.CurrencyFromCode(cmd.CurrencyCode)}

	err := loadedAccount.AddFunds(funds)
	if err != nil {
		return err
	}

	return h.repo.Save(loadedAccount)
}

type UseFunds struct {
	*common.CommandBase
	Amount       int64
	CurrencyCode string
}

type UseFundsHandler struct {
	repo     AccountRepository
	eventBus common.EventBus
}

func (h *UseFundsHandler) Handle(cmd UseFunds) error {
	loadedAccount := h.repo.Load(cmd.GetAggregateID())
	amount := account.Money{
		Amount:   cmd.Amount,
		Currency: account.CurrencyFromCode(cmd.CurrencyCode)}

	err := loadedAccount.UseFunds(amount)
	if err != nil {
		return err
	}

	return h.repo.Save(loadedAccount)
}

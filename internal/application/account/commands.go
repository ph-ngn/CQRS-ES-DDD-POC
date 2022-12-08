package account

import (
	"github.com/andyj29/wannabet/internal/application/common"
	"github.com/andyj29/wannabet/internal/domain/account"
	domainCommon "github.com/andyj29/wannabet/internal/domain/common"
)

type RegisterAccount struct {
	*common.CommandBase
	Email string
	Name  string
}

type RegisterAccountHandler struct {
	repo     common.Repository
	eventBus common.EventBus
}

type AddFunds struct {
	*common.CommandBase
	Amount       int64
	CurrencyCode string
}

type AddFundsHandler struct {
	repo     common.Repository
	eventBus common.EventBus
}

type UseFunds struct {
	*common.CommandBase
	Amount       int64
	CurrencyCode string
}

type DeductFundsHandler struct {
	repo     common.Repository
	eventBus common.EventBus
}

func (h *RegisterAccountHandler) Handle(cmd RegisterAccount, errChannel chan<- error) {
	newAccount := account.NewAccount()
	accountCreatedEvent := account.NewAccountCreatedEvent(cmd.GetAggregateID(), cmd.Name, account.EmailFromString(cmd.Email))

	if err := newAccount.When(accountCreatedEvent, true); err != nil {
		errChannel <- err
		return
	}
	errChannel <- nil
	h.repo.Save(newAccount)
}

func (h *AddFundsHandler) Handle(cmd AddFunds, errChannel chan<- error) {
	loadedAccount := h.repo.Load(cmd.GetAggregateID())
	fundsAddedEvent := account.NewFundsAddedEvent(cmd.GetAggregateID(), domainCommon.NewMoney(cmd.Amount, cmd.CurrencyCode))

	if err := loadedAccount.When(fundsAddedEvent, true); err != nil {
		errChannel <- err
		return
	}
	errChannel <- nil
	h.repo.Save(loadedAccount)
}

func (h *DeductFundsHandler) Handle(cmd UseFunds, errChannel chan<- error) {
	loadedAccount := h.repo.Load(cmd.GetAggregateID())
	fundsDeductedEvent := account.NewFundsDeductedEvent(cmd.GetAggregateID(), domainCommon.NewMoney(cmd.Amount, cmd.CurrencyCode))

	if err := loadedAccount.When(fundsDeductedEvent, true); err != nil {
		errChannel <- err
		return
	}
	errChannel <- nil
	h.repo.Save(loadedAccount)
}

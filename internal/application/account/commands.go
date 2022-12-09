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

type registerAccountHandler struct {
	repo     common.Repository
	eventBus common.EventBus
}

type addFunds struct {
	*common.CommandBase
	Amount       int64
	CurrencyCode string
}

type addFundsHandler struct {
	repo     common.Repository
	eventBus common.EventBus
}

type deductFunds struct {
	*common.CommandBase
	Amount       int64
	CurrencyCode string
}

type deductFundsHandler struct {
	repo     common.Repository
	eventBus common.EventBus
}

func NewRegisterAccountCommand(aggregateID, email, name string) *registerAccount {
	return &registerAccount{
		CommandBase: common.NewCommandBase(aggregateID),
		Email:       email,
		Name:        name,
	}
}

func NewAddFundsCommand(aggregateID string, amount int64, currencyCode string) *addFunds {
	return &addFunds{
		CommandBase:  common.NewCommandBase(aggregateID),
		Amount:       amount,
		CurrencyCode: currencyCode,
	}
}

func NewDeductFundsCommand(aggregateID string, amount int64, currencyCode string) *deductFunds {
	return &deductFunds{
		CommandBase:  common.NewCommandBase(aggregateID),
		Amount:       amount,
		CurrencyCode: currencyCode,
	}
}

func NewRegisterAccountHandler(repo common.Repository, eventBus common.EventBus) *registerAccountHandler {
	return &registerAccountHandler{
		repo:     repo,
		eventBus: eventBus,
	}
}

func NewAddFundsHandler(repo common.Repository, eventBus common.EventBus) *addFundsHandler {
	return &addFundsHandler{
		repo:     repo,
		eventBus: eventBus,
	}
}

func NewDeductFundsHandler(repo common.Repository, eventBus common.EventBus) *deductFundsHandler {
	return &deductFundsHandler{
		repo:     repo,
		eventBus: eventBus,
	}
}

func (h *registerAccountHandler) Handle(cmd registerAccount) error {
	newAccount := account.NewAccount()
	accountCreatedEvent := account.NewAccountCreatedEvent(cmd.GetAggregateID(), account.Email(cmd.Email), cmd.Name)

	if err := newAccount.When(accountCreatedEvent, true); err != nil {
		return err
	}

	if err := h.eventBus.Publish(accountCreatedEvent); err != nil {
		return err
	}

	return h.repo.Save(newAccount)
}

func (h *addFundsHandler) Handle(cmd addFunds) error {
	loadedAccount := h.repo.Load(cmd.GetAggregateID())
	fundsAddedEvent := account.NewFundsAddedEvent(cmd.GetAggregateID(), domainCommon.NewMoney(cmd.Amount, cmd.CurrencyCode))

	if err := loadedAccount.When(fundsAddedEvent, true); err != nil {
		return err
	}

	if err := h.eventBus.Publish(fundsAddedEvent); err != nil {
		return err
	}

	return h.repo.Save(loadedAccount)
}

func (h *deductFundsHandler) Handle(cmd deductFunds) error {
	loadedAccount := h.repo.Load(cmd.GetAggregateID())
	fundsDeductedEvent := account.NewFundsDeductedEvent(cmd.GetAggregateID(), domainCommon.NewMoney(cmd.Amount, cmd.CurrencyCode))

	if err := loadedAccount.When(fundsDeductedEvent, true); err != nil {
		return err
	}

	if err := h.eventBus.Publish(fundsDeductedEvent); err != nil {
		return err
	}

	return h.repo.Save(loadedAccount)
}

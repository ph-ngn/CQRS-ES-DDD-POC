package handler

import (
	"fmt"
	"github.com/andyj29/wannabet/internal/command"
	"github.com/andyj29/wannabet/internal/domain"

	"github.com/andyj29/wannabet/internal/domain/offer"
	"github.com/andyj29/wannabet/internal/eventbus"
	"github.com/andyj29/wannabet/internal/repository"
)

type PlaceBetHandler struct {
	Repo     repository.Interface[*offer.Offer]
	EventBus eventbus.Interface
}

func (h *PlaceBetHandler) Handle(cmd command.Interface) error {
	c, ok := cmd.(*command.PlaceBet)
	if !ok {
		panic(fmt.Sprintf("Unexpected command type %T", cmd))
	}

	loadedOffer, err := h.Repo.Load(cmd.GetAggregateID())
	if err != nil {
		return err
	}

	stake, err := domain.NewMoney(c.Stake, c.CurrencyCode)
	if err != nil {
		return err
	}

	newBet := offer.NewBet(c.BetID, c.BettorID, stake)
	if err := loadedOffer.PlaceBet(newBet); err != nil {
		return err
	}

	if err := h.Repo.Save(loadedOffer); err != nil {
		return err
	}

	for _, event := range loadedOffer.GetChanges() {
		h.EventBus.Publish(event)
	}
	return nil
}

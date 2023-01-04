package handler

import (
	"fmt"
	"github.com/andyj29/wannabet/internal/command"
	"github.com/andyj29/wannabet/internal/domain"
	"github.com/andyj29/wannabet/internal/domain/offer"
	"github.com/andyj29/wannabet/internal/eventbus"
	"github.com/andyj29/wannabet/internal/repository"
)

type CreateOfferHandler struct {
	Repo     repository.Interface[*offer.Offer]
	EventBus eventbus.Interface
}

func (h *CreateOfferHandler) Handle(cmd command.Interface) error {
	c, ok := cmd.(*command.CreateOffer)
	if !ok {
		panic(fmt.Sprintf("Unexpected command type %T", cmd))
	}

	limit, err := domain.NewMoney(c.Limit, c.CurrencyCode)
	if err != nil {
		return err
	}

	newOffer, err := offer.NewOffer(cmd.GetAggregateID(),
		c.BookMakerID,
		c.FixtureID,
		c.HomeOdds,
		c.AwayOdds,
		limit)
	if err != nil {
		return err
	}

	if err := h.Repo.Save(newOffer); err != nil {
		return err
	}

	for _, event := range newOffer.GetChanges() {
		h.EventBus.Publish(event)
	}
	return nil
}

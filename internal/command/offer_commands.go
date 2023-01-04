package command

import (
	"fmt"
	"github.com/andyj29/wannabet/internal/application/common"
	domainCommon "github.com/andyj29/wannabet/internal/domain/common"
	"github.com/andyj29/wannabet/internal/domain/offer"
)

type CreateOffer struct {
	*base
	BookMakerID, FixtureID, HomeOdds, AwayOdds string
	Limit                                      int64
	CurrencyCode                               string
}

func NewCreateOfferCommand(aggregateID, bookMakerID, fixtureID, homeOdds, awayOdds string, limit int64, currencyCode string) *CreateOffer {
	return &CreateOffer{
		base:         &base{AggregateID: aggregateID},
		BookMakerID:  bookMakerID,
		FixtureID:    fixtureID,
		HomeOdds:     homeOdds,
		AwayOdds:     awayOdds,
		Limit:        limit,
		CurrencyCode: currencyCode,
	}
}

type CreateOfferHandler struct {
	Repo     common.Repository[*offer.Offer]
	EventBus common.EventBus
}

func (h *CreateOfferHandler) Handle(cmd common.Command) error {
	c, ok := cmd.(*CreateOffer)
	if !ok {
		panic(fmt.Sprintf("Unexpected command type %T", cmd))
	}

	limit, err := domainCommon.NewMoney(c.Limit, c.CurrencyCode)
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

type PlaceBet struct {
	*base
	BetID        string
	BettorID     string
	Stake        int64
	CurrencyCode string
}

func NewPlaceBetCommand(aggregateID, bettorID string, stake int64, currencyCode string) *PlaceBet {
	return &PlaceBet{
		base:         &base{AggregateID: aggregateID},
		BettorID:     bettorID,
		Stake:        stake,
		CurrencyCode: currencyCode,
	}
}

type PlaceBetHandler struct {
	Repo     common.Repository[*offer.Offer]
	EventBus common.EventBus
}

func (h *PlaceBetHandler) Handle(cmd common.Command) error {
	c, ok := cmd.(*PlaceBet)
	if !ok {
		panic(fmt.Sprintf("Unexpected command type %T", cmd))
	}

	loadedOffer, err := h.Repo.Load(cmd.GetAggregateID())
	if err != nil {
		return err
	}

	stake, err := domainCommon.NewMoney(c.Stake, c.CurrencyCode)
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

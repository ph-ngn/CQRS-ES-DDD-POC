package offer

import (
	"github.com/andyj29/wannabet/internal/application/common"
	domainCommon "github.com/andyj29/wannabet/internal/domain/common"
	"github.com/andyj29/wannabet/internal/domain/offer"
)

type CreateOffer struct {
	*common.CommandBase
	BookMakerID, FixtureID, HomeOdds, AwayOdds string
	Limit                                      int64
	CurrencyCode                               string
}

type CreateOfferHandler struct {
	Repo     common.Repository[*offer.Offer]
	EventBus common.EventBus
}

type PlaceBet struct {
	*common.CommandBase
	BetID        string
	BettorID     string
	Stake        int64
	CurrencyCode string
}

type PlaceBetHandler struct {
	Repo     common.Repository[*offer.Offer]
	EventBus common.EventBus
}

func NewCreateOfferCommand(aggregateID, bookMakerID, fixtureID, homeOdds, awayOdds string, limit int64, currencyCode string) *CreateOffer {
	return &CreateOffer{
		CommandBase:  &common.CommandBase{AggregateID: aggregateID},
		BookMakerID:  bookMakerID,
		FixtureID:    fixtureID,
		HomeOdds:     homeOdds,
		AwayOdds:     awayOdds,
		Limit:        limit,
		CurrencyCode: currencyCode,
	}
}

func NewPlaceBetCommand(aggregateID, bettorID string, stake int64, currencyCode string) *PlaceBet {
	return &PlaceBet{
		CommandBase:  &common.CommandBase{AggregateID: aggregateID},
		BettorID:     bettorID,
		Stake:        stake,
		CurrencyCode: currencyCode,
	}
}

func (h *CreateOfferHandler) Handle(cmd common.Command) error {
	c, ok := cmd.(*CreateOffer)
	if !ok {
		return nil
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

func (h *PlaceBetHandler) Handle(cmd common.Command) error {
	c, ok := cmd.(*PlaceBet)
	if !ok {
		return nil
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

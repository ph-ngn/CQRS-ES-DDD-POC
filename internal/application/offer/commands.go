package offer

import (
	"github.com/andyj29/wannabet/internal/application/common"
	domainCommon "github.com/andyj29/wannabet/internal/domain/common"
	"github.com/andyj29/wannabet/internal/domain/offer"
)

type createOffer struct {
	*common.CommandBase
	BookMakerID, FixtureID, HomeOdds, AwayOdds string
	Limit                                      int64
	CurrencyCode                               string
}

type CreateOfferHandler struct {
	repo     common.Repository[*offer.Offer]
	eventBus common.EventBus
}

type placeBet struct {
	*common.CommandBase
	BetID        string
	BettorID     string
	Stake        int64
	CurrencyCode string
}

type PlaceBetHandler struct {
	repo     common.Repository[*offer.Offer]
	eventBus common.EventBus
}

func NewCreateOfferCommand(aggregateID, bookMakerID, fixtureID, homeOdds, awayOdds string, limit int64, currencyCode string) *createOffer {
	return &createOffer{
		CommandBase:  &common.CommandBase{AggregateID: aggregateID},
		BookMakerID:  bookMakerID,
		FixtureID:    fixtureID,
		HomeOdds:     homeOdds,
		AwayOdds:     awayOdds,
		Limit:        limit,
		CurrencyCode: currencyCode,
	}
}

func NewPlaceBetCommand(aggregateID, bettorID string, stake int64, currencyCode string) *placeBet {
	return &placeBet{
		CommandBase:  &common.CommandBase{AggregateID: aggregateID},
		BettorID:     bettorID,
		Stake:        stake,
		CurrencyCode: currencyCode,
	}
}

func (h *CreateOfferHandler) Handle(cmd createOffer) error {
	limit, err := domainCommon.NewMoney(cmd.Limit, cmd.CurrencyCode)
	if err != nil {
		return err
	}
	newOffer, err := offer.NewOffer(cmd.GetAggregateID(),
		cmd.BookMakerID,
		cmd.FixtureID,
		cmd.HomeOdds,
		cmd.AwayOdds,
		limit)
	if err != nil {
		return err
	}
	if err := h.repo.Save(newOffer); err != nil {
		return err
	}
	for _, event := range newOffer.GetChanges() {
		h.eventBus.Publish(event)
	}
	return nil
}

func (h *PlaceBetHandler) Handle(cmd placeBet) error {
	loadedOffer := h.repo.Load(cmd.GetAggregateID())
	stake, err := domainCommon.NewMoney(cmd.Stake, cmd.CurrencyCode)
	if err != nil {
		return err
	}
	newBet := offer.NewBet(cmd.BetID, cmd.BettorID, stake)
	if err := loadedOffer.PlaceBet(newBet); err != nil {
		return err
	}
	if err := h.repo.Save(loadedOffer); err != nil {
		return err
	}
	for _, event := range loadedOffer.GetChanges() {
		h.eventBus.Publish(event)
	}
	return nil
}

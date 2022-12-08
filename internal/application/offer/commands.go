package offer

import (
	"github.com/andyj29/wannabet/internal/application/common"
	domainCommon "github.com/andyj29/wannabet/internal/domain/common"
	"github.com/andyj29/wannabet/internal/domain/offer"
)

type CreateOffer struct {
	*common.CommandBase
	OffererID    string
	GameID       string
	Favorite     string
	Limit        int64
	CurrencyCode string
}

type CreateOfferHandler struct {
	repo     common.Repository
	eventBus common.EventBus
}

type PlaceBet struct {
	*common.CommandBase
	BettorID     string
	Stake        int64
	CurrencyCode string
}

type PlaceBetHandler struct {
	repo     common.Repository
	eventBus common.EventBus
}

func (h *CreateOfferHandler) Handle(cmd CreateOffer) error {
	newOffer := offer.NewOffer()
	offerCreatedEvent := offer.NewOfferCreatedEvent(cmd.GetAggregateID(), cmd.OffererID, cmd.GameID, cmd.Favorite, domainCommon.NewMoney(cmd.Limit, cmd.CurrencyCode))

	if err := newOffer.When(offerCreatedEvent, true); err != nil {
		return err
	}
	return h.repo.Save(newOffer)
}

func (h *PlaceBetHandler) Handle(cmd PlaceBet) error {
	loadedOffer := h.repo.Load(cmd.GetAggregateID())
	newBet := offer.NewBet(cmd.BettorID, domainCommon.NewMoney(cmd.Stake, cmd.CurrencyCode))
	betPlacedEvent := offer.NewBetPlacedEvent(cmd.GetAggregateID(), newBet)

	if err := loadedOffer.When(betPlacedEvent, true); err != nil {
		return err
	}
	return h.repo.Save(loadedOffer)
}

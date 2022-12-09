package offer

import (
	"github.com/andyj29/wannabet/internal/application/common"
	domainCommon "github.com/andyj29/wannabet/internal/domain/common"
	"github.com/andyj29/wannabet/internal/domain/offer"
)

type createOffer struct {
	*common.CommandBase
	OffererID    string
	GameID       string
	Favorite     string
	Limit        int64
	CurrencyCode string
}

type createOfferHandler struct {
	repo     common.Repository
	eventBus common.EventBus
}

type placeBet struct {
	*common.CommandBase
	BettorID     string
	Stake        int64
	CurrencyCode string
}

type placeBetHandler struct {
	repo     common.Repository
	eventBus common.EventBus
}

func NewCreateOfferCommand(aggregateID, offererID, gameID, favorite string, limit int64, currencyCode string) *createOffer {
	return &createOffer{
		CommandBase:  common.NewCommandBase(aggregateID),
		OffererID:    offererID,
		GameID:       gameID,
		Favorite:     favorite,
		Limit:        limit,
		CurrencyCode: currencyCode,
	}
}

func NewCreateOfferHandler(repo common.Repository, eventBus common.EventBus) *createOfferHandler {
	return &createOfferHandler{
		repo:     repo,
		eventBus: eventBus,
	}
}

func NewPlaceBetCommand(aggregateID, bettorID string, stake int64, currencyCode string) *placeBet {
	return &placeBet{
		CommandBase:  common.NewCommandBase(aggregateID),
		BettorID:     bettorID,
		Stake:        stake,
		CurrencyCode: currencyCode,
	}
}

func NewPlaceBetHandler(repo common.Repository, eventBus common.EventBus) *placeBetHandler {
	return &placeBetHandler{
		repo:     repo,
		eventBus: eventBus,
	}
}

func (h *createOfferHandler) Handle(cmd createOffer) error {
	newOffer := offer.NewOffer()
	offerCreatedEvent := offer.NewOfferCreatedEvent(cmd.GetAggregateID(), cmd.OffererID, cmd.GameID, cmd.Favorite, domainCommon.NewMoney(cmd.Limit, cmd.CurrencyCode))

	if err := newOffer.When(offerCreatedEvent, true); err != nil {
		return err
	}

	if err := h.eventBus.Publish(offerCreatedEvent); err != nil {
		return err
	}

	return h.repo.Save(newOffer)
}

func (h *placeBetHandler) Handle(cmd placeBet) error {
	loadedOffer := h.repo.Load(cmd.GetAggregateID())
	newBet := offer.NewBet(cmd.BettorID, domainCommon.NewMoney(cmd.Stake, cmd.CurrencyCode))
	betPlacedEvent := offer.NewBetPlacedEvent(cmd.GetAggregateID(), newBet)

	if err := loadedOffer.When(betPlacedEvent, true); err != nil {
		return err
	}

	if err := h.eventBus.Publish(betPlacedEvent); err != nil {
		return err
	}

	return h.repo.Save(loadedOffer)
}

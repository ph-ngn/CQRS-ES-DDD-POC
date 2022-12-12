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
	repo     Repository
	eventBus common.EventBus
}

type placeBet struct {
	*common.CommandBase
	BettorID     string
	Stake        int64
	CurrencyCode string
}

type placeBetHandler struct {
	repo     Repository
	eventBus common.EventBus
}

func NewCreateOfferCommand(aggregateID, offererID, gameID, favorite string, limit int64, currencyCode string) *createOffer {
	return &createOffer{
		CommandBase:  &common.CommandBase{AggregateID: aggregateID},
		OffererID:    offererID,
		GameID:       gameID,
		Favorite:     favorite,
		Limit:        limit,
		CurrencyCode: currencyCode,
	}
}

func NewCreateOfferHandler(repo Repository, eventBus common.EventBus) *createOfferHandler {
	return &createOfferHandler{
		repo:     repo,
		eventBus: eventBus,
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

func NewPlaceBetHandler(repo Repository, eventBus common.EventBus) *placeBetHandler {
	return &placeBetHandler{
		repo:     repo,
		eventBus: eventBus,
	}
}

func (h *createOfferHandler) Handle(cmd createOffer) error {
	newOffer := offer.NewOffer(cmd.GetAggregateID(),
		cmd.OffererID,
		cmd.GameID,
		cmd.Favorite,
		domainCommon.NewMoney(cmd.Limit, cmd.CurrencyCode))
	if err := h.repo.Save(newOffer); err != nil {
		return err
	}

	for _, event := range newOffer.GetChanges() {
		h.eventBus.Publish(event)
	}

	return nil
}

func (h *placeBetHandler) Handle(cmd placeBet) error {
	loadedOffer := h.repo.Load(cmd.GetAggregateID())
	newBet := offer.NewBet(cmd.BettorID, domainCommon.NewMoney(cmd.Stake, cmd.CurrencyCode))

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

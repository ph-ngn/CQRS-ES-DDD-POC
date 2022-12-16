package offer

import (
	"encoding/json"
	"net/http"

	appCommon "github.com/andyj29/wannabet/internal/application/common"
	"github.com/andyj29/wannabet/internal/application/offer"
	"github.com/andyj29/wannabet/internal/presentation/common"
)

type Controller struct {
	appCommon.Dispatcher
	getRequestingAccount func(*http.Request) string
}

func (c *Controller) CreateOffer(w http.ResponseWriter, r *http.Request) {
	var request createOfferRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		common.WriteJSONErrorResponse(w, r, common.NewBadRequestError(err))
		return
	}

	requestingAccount := c.getRequestingAccount(r)
	cmd := offer.NewCreateOfferCommand(requestingAccount,
		request.OffererID,
		request.GameID,
		request.Favorite,
		request.Limit,
		request.CurrencyCode)
	if err := c.Dispatch(cmd); err != nil {
		common.WriteJSONErrorResponse(w, r, common.NewInternalError(err))
		return
	}

	common.WriteJSONResponseWithStatus(w, r, http.StatusCreated, newCreateOfferResponse(true,
		"Offer is succesfully created",
		request.GameID,
		request.Favorite,
		request.Limit,
		request.CurrencyCode))
}

func (c *Controller) PlaceBet(w http.ResponseWriter, r *http.Request) {
	var request placeBetRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		common.WriteJSONErrorResponse(w, r, common.NewBadRequestError(err))
		return
	}

	requestingAccount := c.getRequestingAccount(r)
	cmd := offer.NewPlaceBetCommand(request.OfferID, requestingAccount, request.Stake, request.CurrencyCode)
	if err := c.Dispatch(cmd); err != nil {
		common.WriteJSONErrorResponse(w, r, common.NewInternalError(err))
		return
	}

	common.WriteJSONResponseWithStatus(w, r, http.StatusCreated,
		newPlaceBetResponse(true, "Bet is succesfully placed", request.OfferID, request.Stake, request.CurrencyCode))
}
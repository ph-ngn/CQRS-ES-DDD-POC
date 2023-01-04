package controller

import (
	"encoding/json"
	"net/http"

	"github.com/andyj29/wannabet/internal/api/httperror"
	rw "github.com/andyj29/wannabet/internal/api/responsewriter"
	"github.com/andyj29/wannabet/internal/application/common"
	"github.com/andyj29/wannabet/internal/command"
)

type OfferController struct {
	common.Dispatcher
	GetRequestingAccount func(*http.Request) string
}

func (c *OfferController) CreateOffer(w http.ResponseWriter, r *http.Request) {
	var request createOfferRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		rw.WriteJSONErrorResponse(w, r, httperror.NewBadRequestError(err))
		return
	}

	requestingAccount := c.GetRequestingAccount(r)
	cmd := command.NewCreateOfferCommand("testID",
		requestingAccount,
		request.FixtureID,
		request.HomeOdds,
		request.AwayOdds,
		request.Limit,
		request.CurrencyCode)
	if err := c.Dispatch(cmd); err != nil {
		rw.WriteJSONErrorResponse(w, r, httperror.NewInternalError(err))
		return
	}

	rw.WriteJSONResponseWithStatus(w, r,
		http.StatusCreated,
		createOfferResponse{
			newResponse(true, "Offer successfully created"),
			&request,
		})
}

func (c *OfferController) PlaceBet(w http.ResponseWriter, r *http.Request) {
	var request placeBetRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		rw.WriteJSONErrorResponse(w, r, httperror.NewBadRequestError(err))
		return
	}

	requestingAccount := c.GetRequestingAccount(r)
	cmd := command.NewPlaceBetCommand(request.OfferID, requestingAccount, request.Stake, request.CurrencyCode)
	if err := c.Dispatch(cmd); err != nil {
		rw.WriteJSONErrorResponse(w, r, httperror.NewInternalError(err))
		return
	}

	rw.WriteJSONResponseWithStatus(w, r,
		http.StatusCreated,
		placeBetResponse{
			newResponse(true, "Bet successfully placed"),
			&request,
		})
}

type createOfferRequest struct {
	FixtureID    string `json:"fixture_id"`
	HomeOdds     string `json:"home_odds"`
	AwayOdds     string `json:"away_odds"`
	Limit        int64  `json:"limit"`
	CurrencyCode string `json:"currency_code"`
}

type createOfferResponse struct {
	*response
	*createOfferRequest
}

type placeBetRequest struct {
	OfferID      string `json:"offer_id"`
	Stake        int64  `json:"stake"`
	CurrencyCode string `json:"currency_code"`
}
type placeBetResponse struct {
	*response
	*placeBetRequest
}

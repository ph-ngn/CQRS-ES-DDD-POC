package controller

import (
	"encoding/json"
	"net/http"

	"github.com/andyj29/wannabet/internal/api/httperror"
	rw "github.com/andyj29/wannabet/internal/api/responsewriter"
	"github.com/andyj29/wannabet/internal/application/common"
	"github.com/andyj29/wannabet/internal/command"
)

type AccountController struct {
	common.Dispatcher
}

func (c *AccountController) RegisterAccount(w http.ResponseWriter, r *http.Request) {
	var request registerAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		rw.WriteJSONErrorResponse(w, r, httperror.NewBadRequestError(err))
		return
	}

	cmd := command.NewRegisterAccountCommand(request.ID, request.Email, request.Name)
	if err := c.Dispatch(cmd); err != nil {
		rw.WriteJSONErrorResponse(w, r, httperror.NewInternalError(err))
		return
	}

	rw.WriteJSONResponseWithStatus(w, r, http.StatusCreated, newResponse(true, "Account successfully registered"))
}

type registerAccountRequest struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type registerAccountResponse struct {
	*response
}

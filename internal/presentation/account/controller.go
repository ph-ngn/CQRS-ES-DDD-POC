package account

import (
	"encoding/json"
	"net/http"

	"github.com/andyj29/wannabet/internal/application/account"
	appCommon "github.com/andyj29/wannabet/internal/application/common"
	"github.com/andyj29/wannabet/internal/presentation/common"
)

type Controller struct {
	appCommon.Dispatcher
}

func (c *Controller) RegisterAccount(w http.ResponseWriter, r *http.Request) {
	var request accountRegistrationRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		common.WriteJSONErrorResponse(w, r, common.NewBadRequestError(err))
		return
	}

	cmd := account.NewRegisterAccountCommand(request.ID, request.Email, request.Name)
	if err := c.Dispatcher.Dispatch(cmd); err != nil {
		common.WriteJSONErrorResponse(w, r, common.NewInternalError(err))
		return
	}

	common.WriteJSONResponse(w, r, NewAccountRegistrationResponse(true, "Account is succesfully registered"))
}

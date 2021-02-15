package app

import (
	"encoding/json"
	"net/http"

	"github.com/ThePianoDentist/toast-notification/utils"

	"github.com/ThePianoDentist/toast-notification/app_context"

	"github.com/ThePianoDentist/toast-notification/storage"
)

func PostUser(appCtx *app_context.AppContext, w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var u storage.User
	if err := decoder.Decode(&u); err != nil {
		utils.ErrorResp(appCtx.Lgr, w, http.StatusBadRequest, "Invalid request payload", err)
		return
	}
	// I think reading body is weird/dumb. and defering before reading body leads to panic in some scenarios.
	// (add stack overflow link here if find/know)
	defer r.Body.Close()
	appCtx.Token = u.FirebaseToken
	utils.SuccessResp(appCtx.Lgr, w, 201, map[string]string{"userId": u.UserId.String()})
}

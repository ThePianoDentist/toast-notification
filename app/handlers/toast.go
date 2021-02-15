package app

import (
	"net/http"

	"github.com/ThePianoDentist/toast-notification/app_context"
	"go.uber.org/zap"
)

func ToastNotify(appCtx *app_context.AppContext, w http.ResponseWriter, r *http.Request) {
	err := appCtx.FcmController.SendFCM(appCtx.Token)
	if err != nil {
		appCtx.Lgr.Error("error publishing fcm message", zap.Error(err))
	}
}

package app

import (
	"log"
	"net/http"

	"github.com/ThePianoDentist/toast-notification/app/middleware"

	"github.com/ThePianoDentist/toast-notification/app_context"
	"github.com/ThePianoDentist/toast-notification/fcm_client"

	_ "github.com/lib/pq"

	//_ "github.com/jackc/pgx/v4"

	handlers "github.com/ThePianoDentist/toast-notification/app/handlers"

	"go.uber.org/zap"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	appCtx *app_context.AppContext
}

func NewApp(lgr *zap.Logger) *App {
	fcmClient := fcm_client.NewFCMController(lgr)
	appCtx := &app_context.AppContext{Lgr: lgr, FcmController: fcmClient}

	router := mux.NewRouter()
	app := &App{Router: router, appCtx: appCtx}
	app.setupRouter()
	return app
}

func (a *App) Run(addr string) {
	// prob need smarter way of authing user/kettle.
	//a.Router.HandleFunc("/kettles/{kettleId}/{userId}/offer/", app.PostOffer).Methods(http.MethodPost)
	//a.Router.HandleFunc("/kettles/{kettleId}/{userId}/request/", app.PostDrinkRequest).Methods(http.MethodPost)
	// Need to auth to a kettle. (Is a webserver needed, or can peer-2-peea.Router. that sounds hard.)
	if err := http.ListenAndServe(addr, a.Router); err != nil {
		log.Fatal("error running server: ", zap.Error(err))
	}
}

func (a *App) setupRouter() {
	// handle preflight/CORS requests
	a.Router.Methods(http.MethodOptions).HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			return
		})
	a.Router.Methods(http.MethodPost).Path("/register/").Handler(&app_context.CtxHandler{a.appCtx, handlers.PostUser})
	a.Router.Methods(http.MethodGet).Path("/ready/").Handler(&app_context.CtxHandler{a.appCtx, handlers.ToastNotify})
	a.Router.Use(middleware.AccessControl)
	a.Router.Use(middleware.RequireJsonContentType)
}

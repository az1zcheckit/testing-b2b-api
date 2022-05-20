package router

import (
	"b2b-api/internal/config"
	_ "b2b-api/internal/docs"
	"b2b-api/internal/handlers/authentication"
	"b2b-api/internal/handlers/merchant"
	"b2b-api/internal/middleware"
	"b2b-api/internal/pkg/service"
	"context"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/fx"
	"net/http"
)

var EntryPoint = fx.Options(
	fx.Invoke(
		NewRouter,
	),
)

type dependencies struct {
	fx.In
	Lifecycle fx.Lifecycle
	Config    config.IConfig
	Auth      authentication.IAuthHandler
	MW        middleware.IMiddleware
	SVC       service.IService
	Merch     merchant.IMerchHandler
}

// @title B2B Example API
// @version 1.0
// @description This is a B2B-API server.
// @host localhost:8070
// @BasePath /api/v1
func NewRouter(d dependencies) {
	server := mux.NewRouter()
	mainRoute := server.PathPrefix("/api").Subrouter()
	routeVer := mainRoute.PathPrefix("/v1").Subrouter()

	routeVer.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL(d.Config.GetString("api.swagger.url")), //The url pointing to API definition. // Has been added by Aziz.
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))

	generalMiddleware := []mux.MiddlewareFunc{
		d.MW.Cors,
		d.MW.SetRequestID,
		d.MW.TokenValidator,
	}

	authMiddleware := []mux.MiddlewareFunc{
		d.MW.Cors,
		d.MW.SetRequestID,
	}

	// auth routes
	loginRoute := routeVer.PathPrefix("/auth").Subrouter()
	loginRoute.Use(authMiddleware...)
	loginRoute.HandleFunc("/login", d.Auth.Authentication()).Methods("POST", "OPTIONS")
	loginRoute.HandleFunc("/refresh-token", d.Auth.RefreshToken()).Methods("GET", "OPTIONS")

	// 2fa routes
	TwoFAPath := routeVer.PathPrefix("/2fa").Subrouter()
	TwoFAPath.Use(authMiddleware...)
	TwoFAPath.HandleFunc("/send-sms-otp", d.Auth.SendSmsOTP()).Methods("GET", "OPTIONS")
	TwoFAPath.HandleFunc("/confirm-sms-otp", d.Auth.SmsOtpConfirmation()).Methods("GET", "OPTIONS")
	TwoFAPath.HandleFunc("/confirm-google-totp", d.Auth.GAuthenticator()).Methods("GET", "OPTIONS")
	TwoFAPath.HandleFunc("/generate-google-se", d.Auth.GenerateGoogleSE()).Methods("GET", "OPTIONS")
	
	//merchant routes
	merchantPath := routeVer.PathPrefix("/merchant").Subrouter()
	merchantPath.Use(generalMiddleware...)
	merchantPath.Path("/confirm-transaction").
		Queries("transactionId", "{transactionId}"). 
		HandlerFunc(d.Merch.ConfirmTransaction()).Methods("GET", "OPTIONS")
	merchantPath.Path("/get-all-amount").
		Queries("currency", "{currency}").
		HandlerFunc(d.Merch.GetAllAmount()).Methods("GET", "OPTIONS")
	merchantPath.Path("/get-limit-conversion").
		Queries("currency", "{currency}").
		HandlerFunc(d.Merch.GetLimitConversion()).Methods("GET", "OPTIONS")
	merchantPath.HandleFunc("/get-rates", d.Merch.GetExchangeRates()).Methods("GET", "OPTIONS")
	merchantPath.HandleFunc("/get-accounts", d.Merch.GetAccounts()).Methods("GET", "OPTIONS")
	merchantPath.HandleFunc("/conversion", d.Merch.Conversion()).Methods("POST", "OPTIONS")
	merchantPath.Path("/get-converted-value").
		Queries("value", "{value}").
		Queries("currencyFrom", "{currencyFrom}").
		Queries("currencyTo", "{currencyTo}").
		HandlerFunc(d.Merch.GetConvertedValue()).Methods("GET", "OPTIONS")
	merchantPath.Path("/get-account-details").
		Queries("accountNumber", "{accountNumber}").
		HandlerFunc(d.Merch.GetAccountDetails()).Methods("GET", "OPTIONS")
	merchantPath.HandleFunc("/transactions", d.Merch.GetHistoryOfTransactions()).Methods("POST", "OPTIONS")

	srv := http.Server{
		Addr:    d.SVC.ConfigInstance().GetString("api.server.port"),
		Handler: server,
	}

	d.Lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				d.SVC.LoggerInstance().Info("Application started")
				go srv.ListenAndServe()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				d.SVC.LoggerInstance().Info("Application stopped")
				return srv.Shutdown(ctx)
			},
		},
	)

}

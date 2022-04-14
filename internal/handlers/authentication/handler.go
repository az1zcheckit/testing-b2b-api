package authentication

import (
	"b2b-api/internal/pkg/service"
	"go.uber.org/fx"
	"net/http"
)

var NewAuthHandler = fx.Provide(newAuthHandler)

type IAuthHandler interface {
	Authentication() http.HandlerFunc
	SendSmsOTP() http.HandlerFunc
	SmsOtpConfirmation() http.HandlerFunc
	GenerateGoogleSE() http.HandlerFunc
	GAuthenticator() http.HandlerFunc
}

type dependencies struct {
	fx.In
	SVC service.IService
}

type authHandler struct {
	svc service.IService
}

func newAuthHandler(d dependencies) IAuthHandler {
	return authHandler{svc: d.SVC}
}

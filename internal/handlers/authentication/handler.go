package authentication

import (
	"b2b-api/internal/pkg/logger"
	"b2b-api/internal/pkg/service"
	"go.uber.org/fx"
	"net/http"
)

const (
	requestDesc  = "request"
	responseDesc = "response"
)

var NewAuthHandler = fx.Provide(newAuthHandler)

type IAuthHandler interface {
	Authentication() http.HandlerFunc
	SendSmsOTP() http.HandlerFunc
	SmsOtpConfirmation() http.HandlerFunc
	GenerateGoogleSE() http.HandlerFunc
	GAuthenticator() http.HandlerFunc
	RefreshToken() http.HandlerFunc
}

type dependencies struct {
	fx.In
	SVC    service.IService
	Logger logger.ILogger
}

type authHandler struct {
	svc    service.IService
	Logger logger.ILogger
}

func newAuthHandler(d dependencies) IAuthHandler {
	return authHandler{
		svc:    d.SVC,
		Logger: d.Logger,
	}
}

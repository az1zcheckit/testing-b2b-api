package merchant

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

var NewMerchHandler = fx.Provide(newMerchHandler)

type IMerchHandler interface {
	GetAccounts() http.HandlerFunc
	GetConvertedValue() http.HandlerFunc
	GetExchangeRates() http.HandlerFunc
	GetAllAmount() http.HandlerFunc
	Conversion() http.HandlerFunc
	GetLimitConversion() http.HandlerFunc
	ConfirmTransaction() http.HandlerFunc
	GetAccountDetails() http.HandlerFunc
	GetHistoryOfTransactions() http.HandlerFunc
}

type dependencies struct {
	fx.In
	SVC    service.IService
	Logger logger.ILogger
}

type merchHandler struct {
	svc    service.IService
	Logger logger.ILogger
}

func newMerchHandler(d dependencies) IMerchHandler {
	return merchHandler{
		svc:    d.SVC,
		Logger: d.Logger,
	}
}

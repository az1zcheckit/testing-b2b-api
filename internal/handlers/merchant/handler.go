package merchant

import (
	"b2b-api/internal/pkg/service"
	"go.uber.org/fx"
	"net/http"
)

var NewMerchHandler = fx.Provide(newMerchHandler)

type IMerchHandler interface {
	GetTransactions() http.HandlerFunc
	GetAccounts() http.HandlerFunc
	GetExchangeRates() http.HandlerFunc
	GetAllAmount() http.HandlerFunc
	Conversion() http.HandlerFunc
	GetLimitConversion() http.HandlerFunc
	TwoFAInit() http.HandlerFunc
}

type dependencies struct {
	fx.In
	SVC service.IService
}

type merchHandler struct {
	svc service.IService
}

func newMerchHandler(d dependencies) IMerchHandler {
	return merchHandler{svc: d.SVC}
}

package middleware

import (
	"b2b-api/internal/pkg/service"
	"go.uber.org/fx"
	"net/http"
)

var NewMiddleware = fx.Provide(newMiddleware)

type IMiddleware interface {
	SetRequestID(next http.Handler) http.Handler
	TokenValidator(next http.Handler) http.Handler
	Cors(next http.Handler) http.Handler
	Log(next http.Handler) http.Handler
}

type dependencies struct {
	fx.In
	SVC service.IService
}

type middleware struct {
	Service service.IService
}

func newMiddleware(d dependencies) IMiddleware {
	return &middleware{
		d.SVC,
	}
}

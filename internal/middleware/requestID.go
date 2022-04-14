package middleware

import (
	"b2b-api/internal/pkg/utils"
	"context"
	"github.com/google/uuid"
	"net/http"
)

func (mw *middleware) SetRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		UUID := uuid.New().String()
		w.Header().Set("X-Request-ID", UUID)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), utils.CTXRequestID, UUID)))
	})
}

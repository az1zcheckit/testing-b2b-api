package middleware

import (
	"github.com/gorilla/mux"
	"net/http"
)

type LogRequest struct {
	Headers map[string][]string
	Body    string
}

func (mw *middleware) Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		IPAddress := r.Header.Get("X-Real-Ip")
		if IPAddress == "" {
			IPAddress = r.Header.Get("X-Forwarded-For")
		}
		if IPAddress == "" {
			IPAddress = r.RemoteAddr
		}

		mw.Service.LoggerInstance().Info("", r.Method, r.RequestURI, " IP-Addres: ", IPAddress, "\n Headers: ",
			r.Header, "\n Queries: ", mux.Vars(r), "\n Body: ", r.Body)

		next.ServeHTTP(w, r)
	})
}

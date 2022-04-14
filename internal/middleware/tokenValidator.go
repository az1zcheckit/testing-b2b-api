package middleware

import (
	response "b2b-api/internal/models"
	"b2b-api/internal/pkg/utils"
	"context"
	"github.com/Nerzal/gocloak/v9"
	"net/http"
)

func (mw *middleware) TokenValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			mw.Service.LoggerInstance().Error("%s", response.TokenIsEmpty)
			response.ToJson(w, http.StatusForbidden, response.TokenIsEmpty)
			return
		}

		client := gocloak.NewClient(mw.Service.ConfigInstance().GetString("api.keycloak.server"))

		ret, err := client.RetrospectToken(r.Context(), tokenHeader, mw.Service.ConfigInstance().GetString("api.keycloak.clientID"), mw.Service.ConfigInstance().GetString("api.keycloak.clientSecret"), mw.Service.ConfigInstance().GetString("api.keycloak.realm"))
		if err != nil {
			response.ToJson(w, http.StatusForbidden, response.SetError(err))
			return
		}

		if !*ret.Active && ret.Exp != nil {
			response.ToJson(w, http.StatusForbidden, response.TokenExpired)
			return
		} else if !*ret.Active && ret.Exp == nil {
			response.ToJson(w, http.StatusForbidden, response.MalformedToken)
			return
		}

		userInfo, err := client.GetUserInfo(r.Context(), tokenHeader, mw.Service.ConfigInstance().GetString("api.keycloak.realm"))
		if err != nil {
			response.ToJson(w, http.StatusNotFound, response.SetError(err))
			return
		}

		userRole, err := client.GetRealmRolesByUserID(r.Context(), tokenHeader, mw.Service.ConfigInstance().GetString("api.keycloak.realm"), *userInfo.Sub)
		if err != nil {
			response.ToJson(w, http.StatusNotFound, response.SetError(err))
			return
		}

		if len(userRole) < 1 {
			response.ToJson(w, http.StatusNotFound, response.RoleIsEmpty)
			return
		}

		if len(userRole) > 1 {
			response.ToJson(w, http.StatusNotFound, response.TooManyRoles)
			return
		}

		getRealmRoles, err := client.GetRealmRoleByID(r.Context(), tokenHeader, mw.Service.ConfigInstance().GetString("api.keycloak.realm"), *userRole[0].ID)
		if err != nil {
			response.ToJson(w, http.StatusForbidden, response.SetError(err))
			return
		}

		for key, element := range *getRealmRoles.Attributes {
			if key == r.URL.Path && element[0] == "true" {
				ctx := context.WithValue(r.Context(), utils.JWTkey, utils.GenerateSession(tokenHeader))
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
		}

		response.ToJson(w, http.StatusForbidden, response.AccessDenied)
		return

	})
}

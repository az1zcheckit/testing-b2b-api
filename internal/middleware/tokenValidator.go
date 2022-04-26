package middleware

import (
	response "b2b-api/internal/models"
	"b2b-api/internal/pkg/utils"
	"context"
	"fmt"
	"github.com/Nerzal/gocloak/v9"
	"net/http"
)

func (mw *middleware) TokenValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		tokenHeader := r.Header.Get("Authorization")
		serviceDesc := "Token-Validator"
		requestID := fmt.Sprintf("%v", r.Context().Value(utils.CTXRequestID))
		if tokenHeader == "" {
			go utils.Logger(ctx, mw.Logger.Error, requestDesc, serviceDesc, requestID, response.TokenIsEmpty)
			response.ToJson(w, http.StatusForbidden, response.TokenIsEmpty)
			return
		}
		go utils.Logger(ctx, mw.Logger.Info, responseDesc, serviceDesc, requestID, "")

		client := gocloak.NewClient(mw.Service.ConfigInstance().GetString("api.keycloak.server"))

		ret, err := client.RetrospectToken(r.Context(), tokenHeader, mw.Service.ConfigInstance().GetString("api.keycloak.clientID"), mw.Service.ConfigInstance().GetString("api.keycloak.clientSecret"), mw.Service.ConfigInstance().GetString("api.keycloak.realm"))
		if err != nil {
			go utils.Logger(ctx, mw.Logger.Error, requestDesc, serviceDesc, requestID, response.SetError(err))
			response.ToJson(w, http.StatusForbidden, response.SetError(err))
			return
		}

		if !*ret.Active && ret.Exp != nil {
			go utils.Logger(ctx, mw.Logger.Error, requestDesc, serviceDesc, requestID, response.TokenExpired)
			response.ToJson(w, http.StatusForbidden, response.TokenExpired)
			return
		} else if !*ret.Active && ret.Exp == nil {
			go utils.Logger(ctx, mw.Logger.Error, requestDesc, serviceDesc, requestID, response.MalformedToken)
			response.ToJson(w, http.StatusForbidden, response.MalformedToken)
			return
		}

		userInfo, err := client.GetUserInfo(r.Context(), tokenHeader, mw.Service.ConfigInstance().GetString("api.keycloak.realm"))
		if err != nil {
			go utils.Logger(ctx, mw.Logger.Error, requestDesc, serviceDesc, requestID, response.SetError(err))
			response.ToJson(w, http.StatusNotFound, response.SetError(err))
			return
		}

		userRole, err := client.GetRealmRolesByUserID(r.Context(), tokenHeader, mw.Service.ConfigInstance().GetString("api.keycloak.realm"), *userInfo.Sub)
		if err != nil {
			go utils.Logger(ctx, mw.Logger.Error, requestDesc, serviceDesc, requestID, response.SetError(err))
			response.ToJson(w, http.StatusNotFound, response.SetError(err))
			return
		}

		if len(userRole) < 1 {
			go utils.Logger(ctx, mw.Logger.Error, requestDesc, serviceDesc, requestID, response.RoleIsEmpty)
			response.ToJson(w, http.StatusNotFound, response.RoleIsEmpty)
			return
		}

		if len(userRole) > 1 {
			go utils.Logger(ctx, mw.Logger.Error, requestDesc, serviceDesc, requestID, response.TooManyRoles)
			response.ToJson(w, http.StatusNotFound, response.TooManyRoles)
			return
		}

		getRealmRoles, err := client.GetRealmRoleByID(r.Context(), tokenHeader, mw.Service.ConfigInstance().GetString("api.keycloak.realm"), *userRole[0].ID)
		if err != nil {
			go utils.Logger(ctx, mw.Logger.Error, requestDesc, serviceDesc, requestID, response.SetError(err))
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

		go utils.Logger(ctx, mw.Logger.Info, responseDesc, serviceDesc, requestID, response.AccessDenied)

		response.ToJson(w, http.StatusForbidden, response.AccessDenied)
		return

	})
}

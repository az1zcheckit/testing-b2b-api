package authentication

import (
	response "b2b-api/internal/models"
	user "b2b-api/internal/models"
	"b2b-api/internal/pkg/utils"
	"fmt"
	"github.com/Nerzal/gocloak/v9"
	"net/http"
)

// RefreshToken godoc
// @Summary      RefreshToken route
// @Description  RefreshToken
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        refreshToken   header     string  true  " "  Format(UUID)
// @Success      200  {object}   models.Token
// @Failure      401  {object}   models.ErrorResponse
// @Router       /auth/refresh-token [get]
func (ah authHandler) RefreshToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		serviceDesc := "RefreshToken"
		requestID := fmt.Sprintf("%v", r.Context().Value(utils.CTXRequestID))
		
		refreshToken := r.Header.Get("refreshToken")

		client := gocloak.NewClient(ah.svc.ConfigInstance().GetString("api.keycloak.server"))

		token, err := client.RefreshToken(ctx,refreshToken, ah.svc.ConfigInstance().GetString("api.keycloak.clientID"), ah.svc.ConfigInstance().GetString("api.keycloak.clientSecret"), ah.svc.ConfigInstance().GetString("api.keycloak.realm"))
		
		if err != nil {
			go utils.Logger(ctx, ah.Logger.Error, requestDesc, serviceDesc, requestID, err.Error())
			response.ToJson(w, http.StatusSeeOther, user.SetError(err))
			return
		}
		usr := &user.User{}
		usr.Token = token.AccessToken
		usr.RefreshToken = token.RefreshToken

		userInfo, err := client.GetUserInfo(ctx, usr.Token, ah.svc.ConfigInstance().GetString("api.keycloak.realm"))
		if err != nil{
			go utils.Logger(ctx, ah.Logger.Error, requestDesc, serviceDesc, requestID, err.Error())
			response.ToJson(w, http.StatusSeeOther, user.SetError(err))
			return
		}
		usrPhone := fmt.Sprintf("%s%s", phonePrefix, *userInfo.PreferredUsername)

		if len(usr.Token) != 0 {
			sqlResponse := ah.svc.AuthRepositoryInstance().SessionInsert(ctx, usrPhone, usr.Token, "EN")
			if sqlResponse.ErrorCode != 0 {
				go utils.Logger(ctx, ah.Logger.Error, requestDesc, serviceDesc, requestID, sqlResponse)
				response.ToJson(w, http.StatusInternalServerError, sqlResponse)
				return
			}
		}

		go utils.Logger(ctx, ah.Logger.Info, responseDesc, serviceDesc, requestID, token)
		
		response.ToJson(w, http.StatusOK, usr)
	}
}

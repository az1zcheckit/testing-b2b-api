package authentication

import (
	response "b2b-api/internal/models"
	user "b2b-api/internal/models"
	"b2b-api/internal/pkg/utils"
	"encoding/json"
	"fmt"
	"github.com/Nerzal/gocloak/v9"
	"net/http"
	"strconv"
	"time"
)

const (
	phonePrefix = "992"
	googleAuth  = "google-authenticator-se"
	smsOTP      = "sms-otp"
)

// Authentication godoc
// @Summary      Authentication route
// @Description  Authentication
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        authentication    body  models.Login  true  "authentication"  Format(json)
// @Success      200  {object}   models.LoginResponse
// @Failure     303  {object}   models.ErrorResponse
// @Router       /auth/login [post]
func (ah authHandler) Authentication() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		usr := &user.User{}

		if err := json.NewDecoder(r.Body).Decode(usr); err != nil {
			response.ToJson(w, http.StatusBadRequest, err)
			return
		}

		if err := usr.LoginValidate(); err != nil {
			response.ToJson(w, http.StatusBadRequest, err)
			return
		}

		client := gocloak.NewClient(ah.svc.ConfigInstance().GetString("api.keycloak.server"))

		token, err := client.Login(ctx, ah.svc.ConfigInstance().GetString("api.keycloak.clientID"), ah.svc.ConfigInstance().GetString("api.keycloak.clientSecret"), ah.svc.ConfigInstance().GetString("api.keycloak.realm"), usr.Phone, usr.Password)
		if err != nil {
			response.ToJson(w, http.StatusSeeOther, user.SetError(err))
			return
		}

		userInfo, err := client.GetUserInfo(ctx, token.AccessToken, ah.svc.ConfigInstance().GetString("api.keycloak.realm"))
		if err != nil {
			response.ToJson(w, http.StatusSeeOther, user.SetError(err))
			return
		}

		userRealmRoles, err := client.GetRealmRolesByUserID(ctx, token.AccessToken, ah.svc.ConfigInstance().GetString("api.keycloak.realm"), *userInfo.Sub)
		if err != nil {
			response.ToJson(w, http.StatusSeeOther, user.SetError(err))
			return
		}

		if len(userRealmRoles) < 1 {
			response.ToJson(w, http.StatusNotFound, user.RoleIsEmpty)
			return
		}

		if len(userRealmRoles) > 1 {
			response.ToJson(w, http.StatusNotFound, user.TooManyRoles)
			return
		}

		requestID := fmt.Sprintf("%v", r.Context().Value(utils.CTXRequestID))

		usr.Phone = fmt.Sprintf("%s%s", phonePrefix, usr.Phone)
		usr.Token = token.AccessToken
		usr.RefreshToken = token.RefreshToken

		userAttr, err := client.GetUserByID(ctx, token.AccessToken, ah.svc.ConfigInstance().GetString("api.keycloak.realm"), *userInfo.Sub)

		if err != nil {
			response.ToJson(w, http.StatusSeeOther, user.SetError(err))
			return
		}

		secure := &response.TwoFactorAuthentication{}

		if userAttr.Attributes != nil {
			for key, element := range *userAttr.Attributes {
				if key == googleAuth && len(element[0]) >= 16 {
					secure.TwoFA.Gauth.IsActive = true
					secure.TwoFA.Gauth.GauthSecretEndpoint = element[0]
				} else if key == smsOTP && func(str string) (b bool) {
					b, err := strconv.ParseBool(str)
					if err != nil {
						return false
					}
					return
				}(element[0]) {
					secure.TwoFA.SMSotp.IsActive = true
				}
			}
		}

		usr.Security = secure
		usr.Role = *userRealmRoles[0].Name

		if usr.Security.TwoFA.Gauth.IsActive || usr.Security.TwoFA.SMSotp.IsActive {
			usrJSON, err := json.Marshal(usr)
			if err != nil {
				response.ToJson(w, http.StatusSeeOther, err)
				return
			}
			err = ah.svc.RedisInstance().Set(ctx, requestID, string(usrJSON), ah.svc.ConfigInstance().GetDuration("api.server.sessionExpirition")*time.Second)
			if err != nil {
				response.ToJson(w, http.StatusSeeOther, err)
				return
			}
		}
		resp := user.WaitingForOTP
		secure.TwoFA.Gauth.GauthSecretEndpoint = ""
		resp.AditionalInfo = secure

		response.ToJson(w, http.StatusOK, resp)
	}
}

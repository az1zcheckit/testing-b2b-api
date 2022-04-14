package merchant

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

// TwoFAInit godoc
// @Summary      TwoFAInit route
// @Description  TwoFAInit
// @Tags         Merchant
// @Accept       json
// @Produce      json
// @Param        token    header     string  true  " "  Format(UUID)
// @Success      200  {object}   models.LoginResponse
// @Failure     303  {object}   models.ErrorResponse
// @Router       /merchant/2fa-init [get]
func (mh merchHandler) TwoFAInit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		usr := &user.User{}

		token := r.Header.Get("token")
		if len(token) == 0 {
			response.ToJson(w, http.StatusSeeOther, response.TokenIsEmpty)
			return
		}

		client := gocloak.NewClient(mh.svc.ConfigInstance().GetString("api.keycloak.server"))

		userInfo, err := client.GetUserInfo(ctx, token, mh.svc.ConfigInstance().GetString("api.keycloak.realm"))

		if err != nil {
			response.ToJson(w, http.StatusSeeOther, user.SetError(err))
			return
		}

		requestID := fmt.Sprintf("%v", r.Context().Value(utils.CTXRequestID))
		usr.Phone = fmt.Sprintf("%s%s", phonePrefix, *userInfo.PreferredUsername)

		userAttr, err := client.GetUserByID(ctx, token, mh.svc.ConfigInstance().GetString("api.keycloak.realm"), *userInfo.Sub)

		if err != nil {
			response.ToJson(w, http.StatusSeeOther, user.SetError(err))
			return
		}

		secure := &user.TwoFactorAuthentication{}

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

		if usr.Security.TwoFA.Gauth.IsActive || usr.Security.TwoFA.SMSotp.IsActive {
			usrJSON, err := json.Marshal(usr)
			if err != nil {
				response.ToJson(w, http.StatusSeeOther, err)
				return
			}
			err = mh.svc.RedisInstance().Set(ctx, requestID, string(usrJSON), mh.svc.ConfigInstance().GetDuration("api.server.sessionExpirition")*time.Second)
			if err != nil {
				response.ToJson(w, http.StatusSeeOther, err)
				return
			}
		}
		resp := response.WaitingForConfirm
		secure.TwoFA.Gauth.GauthSecretEndpoint = ""
		resp.AditionalInfo = secure
		response.ToJson(w, http.StatusOK, resp)

	}
}

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
	"github.com/gorilla/mux"
)

const (
	phonePrefix = "992"
	googleAuth  = "google-authenticator-se"
	smsOTP      = "sms-otp"
)

// ConfirmTransaction godoc
// @Summary      ConfirmTransaction route
// @Description  ConfirmTransaction
// @Tags         Merchant
// @Accept       json
// @Produce      json
// @Param        token    header     string  true  " "  Format(UUID)
// @Param        transactionId   query     string  true  " "  Format(json)
// @Success      200  {object}   models.LoginResponse
// @Failure      401  {object}   models.ErrorResponse
// @Router       /merchant/confirm-transaction [get]
func (mh merchHandler) ConfirmTransaction() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		usr := &user.User{}
		serviceDesc := "Confirm-Transaction-Handler"
		requestID := fmt.Sprintf("%v", r.Context().Value(utils.CTXRequestID))
		token := r.Header.Get("token")		
		transactionId := mux.Vars(r)["transactionId"]

		if len(token) == 0 {
			go utils.Logger(ctx, mh.Logger.Error, requestDesc, serviceDesc, requestID, response.TokenIsEmpty)
			response.ToJson(w, http.StatusSeeOther, response.TokenIsEmpty)
			return
		}
		client := gocloak.NewClient(mh.svc.ConfigInstance().GetString("api.keycloak.server"))
		userInfo, err := client.GetUserInfo(ctx, token, mh.svc.ConfigInstance().GetString("api.keycloak.realm"))
		if err != nil {
			go utils.Logger(ctx, mh.Logger.Error, requestDesc, serviceDesc, requestID, user.SetError(err))
			response.ToJson(w, http.StatusSeeOther, user.SetError(err))
			return
		}
		usr.Phone = fmt.Sprintf("%s%s", phonePrefix, *userInfo.PreferredUsername)
		usr.TransactionId = transactionId
		usr.Token = token
		go utils.Logger(ctx, mh.Logger.Info, requestDesc, serviceDesc, requestID, usr.Phone)

		userAttr, err := client.GetUserByID(ctx, token, mh.svc.ConfigInstance().GetString("api.keycloak.realm"), *userInfo.Sub)

		if err != nil {
			go utils.Logger(ctx, mh.Logger.Error, requestDesc, serviceDesc, requestID, user.SetError(err))
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
				go utils.Logger(ctx, mh.Logger.Error, requestDesc, serviceDesc, requestID, err.Error())
				response.ToJson(w, http.StatusSeeOther, err)
				return
			}
			err = mh.svc.RedisInstance().Set(ctx, requestID, string(usrJSON), mh.svc.ConfigInstance().GetDuration("api.server.sessionExpirition")*time.Second)
			if err != nil {
				go utils.Logger(ctx, mh.Logger.Error, requestDesc, serviceDesc, requestID, err.Error())
				response.ToJson(w, http.StatusSeeOther, err)
				return
			}
		}
		resp := response.WaitingForConfirm
		secure.TwoFA.Gauth.GauthSecretEndpoint = ""
		resp.AditionalInfo = secure

		go utils.Logger(ctx, mh.Logger.Info, responseDesc, serviceDesc, requestID, resp)

		response.ToJson(w, http.StatusOK, resp)

	}
}

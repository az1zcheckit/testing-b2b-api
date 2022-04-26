package authentication

import (
	models "b2b-api/internal/models"
	response "b2b-api/internal/models"
	"b2b-api/internal/pkg/utils"
	"encoding/json"
	"net/http"
	"time"
)

// SendSmsOTP godoc
// @Summary      SMS sender route
// @Description  SMS sender
// @Tags         2FA
// @Produce      json
// @Param        X-Request-Id    header     string  true  " "  Format(UUID)
// @Success      200  {object}   models.OTP
// @Failure     303  {object}   models.ErrorResponse
// @Router       /2fa/send-sms-otp [get]
func (ah authHandler) SendSmsOTP() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		serviceDesc := "Send-Sms-Otp"
		requestID := r.Header.Get("X-Request-ID")
		if len(requestID) == 0 {
			go utils.Logger(ctx, ah.Logger.Error, requestDesc, serviceDesc, requestID, response.RequestIdIsEmpty)
			response.ToJson(w, http.StatusSeeOther, response.RequestIdIsEmpty)
			return
		}
		go utils.Logger(ctx, ah.Logger.Info, requestDesc, serviceDesc, requestID, "")
		usr := &models.User{}

		userInfo, err := ah.svc.RedisInstance().Get(ctx, requestID)
		if err != nil {
			go utils.Logger(ctx, ah.Logger.Error, requestDesc, serviceDesc, requestID, response.SetError(err))
			response.ToJson(w, http.StatusSeeOther, response.SetError(err))
			return
		}

		if err := json.Unmarshal([]byte(userInfo), usr); err != nil {
			go utils.Logger(ctx, ah.Logger.Error, requestDesc, serviceDesc, requestID, response.SetError(err))
			response.ToJson(w, http.StatusBadRequest, response.SetError(err))
			return
		}

		if usr.Security.TwoFA.SMSotp.TotalSent >= ah.svc.ConfigInstance().GetInt("api.otp.otpMaxPerSession") {
			go utils.Logger(ctx, ah.Logger.Error, requestDesc, serviceDesc, requestID, response.OTPCountExhausted)
			response.ToJson(w, http.StatusSeeOther, response.OTPCountExhausted)
			return
		}

		unixTime := time.Now().Unix()
		if usr.Security.TwoFA.SMSotp.Expirition > unixTime {
			go utils.Logger(ctx, ah.Logger.Error, requestDesc, serviceDesc, requestID, response.OTPsent)
			response.ToJson(w, http.StatusSeeOther, response.OTPsent)
			return
		}

		usr.Security.TwoFA.SMSotp.Code = utils.GenerateOTP(ah.svc.ConfigInstance().GetInt("api.otp.otpDigits"))

		otpExp := time.Now().Add(time.Second * time.Duration(ah.svc.ConfigInstance().GetInt("api.otp.otpExpirition"))).Unix()

		usr.Security.TwoFA.SMSotp.Expirition = otpExp

		_, err = ah.svc.TwoFaInstance().SmsOTP(usr.Phone, usr.Security.TwoFA.SMSotp.Code)
		if err != nil {
			go utils.Logger(ctx, ah.Logger.Error, requestDesc, serviceDesc, requestID, err.Error())
			response.Json(w, http.StatusSeeOther, err.Error())
			return
		}

		usr.Security.TwoFA.SMSotp.TotalSent += 1

		usrJSON, err := json.Marshal(usr)

		if err != nil {
			go utils.Logger(ctx, ah.Logger.Error, requestDesc, serviceDesc, requestID, response.SetError(err))
			response.ToJson(w, http.StatusSeeOther, response.SetError(err))
			return
		}

		err = ah.svc.RedisInstance().Set(ctx, requestID, string(usrJSON), ah.svc.ConfigInstance().GetDuration("api.server.sessionExpirition")*time.Second)
		if err != nil {
			go utils.Logger(ctx, ah.Logger.Error, requestDesc, serviceDesc, requestID, response.SetError(err))
			response.ToJson(w, http.StatusSeeOther, response.SetError(err))
			return
		}

		usr.Security.TwoFA.SMSotp.Code = ""
		usr.Security.TwoFA.SMSotp.Expirition = otpExp - unixTime
		go utils.Logger(ctx, ah.Logger.Info, responseDesc, serviceDesc, requestID, "Total sent:", usr.Security.TwoFA.SMSotp.TotalSent)
		response.ToJson(w, http.StatusOK, usr.Security.TwoFA.SMSotp)
	}
}

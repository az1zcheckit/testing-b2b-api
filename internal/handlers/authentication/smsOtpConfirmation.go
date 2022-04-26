package authentication

import (
	models "b2b-api/internal/models"
	response "b2b-api/internal/models"
	"b2b-api/internal/pkg/utils"
	"encoding/json"
	"net/http"
)

// SmsOtpConfirmation godoc
// @Summary      SMS Confirmation route
// @Description  SMS Confirmation
// @Tags         2FA
// @Produce      json
// @Param        X-Request-Id    header     string  true  " "  Format(UUID)
// @Param        otp    header     string  true  " "  Format(uint)
// @Success      200  {object}   models.Token
// @Failure     303  {object}   models.ErrorResponse
// @Router       /2fa/confirm-sms-otp [get]
func (ah authHandler) SmsOtpConfirmation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		serviceDesc := "Sms-Otp-Confirmation"
		requestID := r.Header.Get("X-Request-ID")
		if len(requestID) == 0 {
			go utils.Logger(ctx, ah.Logger.Error, requestDesc, serviceDesc, requestID, response.RequestIdIsEmpty)
			response.ToJson(w, http.StatusSeeOther, response.RequestIdIsEmpty)
			return
		}
		go utils.Logger(ctx, ah.Logger.Info, requestDesc, serviceDesc, requestID, "")
		otp := r.Header.Get("otp")
		if len(requestID) == 0 {
			go utils.Logger(ctx, ah.Logger.Error, requestDesc, serviceDesc, requestID, response.OTPIncorrect)
			response.ToJson(w, http.StatusSeeOther, response.OTPIncorrect)
			return
		}

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

		if usr.Security.TwoFA.SMSotp.Code != otp {
			go utils.Logger(ctx, ah.Logger.Error, requestDesc, serviceDesc, requestID, response.OTPIncorrect)
			response.ToJson(w, http.StatusSeeOther, response.OTPIncorrect)
			return
		}
		tokenInfo := response.User{Token: usr.Token, RefreshToken: usr.RefreshToken, Role: usr.Role}

		if len(usr.Token) != 0 {
			sqlResponse := ah.svc.AuthRepositoryInstance().SessionInsert(ctx, usr.Phone, usr.Token, "EN")
			if sqlResponse.ErrorCode != 0 {
				go utils.Logger(ctx, ah.Logger.Error, requestDesc, serviceDesc, requestID, sqlResponse)
				response.ToJson(w, http.StatusInternalServerError, sqlResponse)
				return
			}
		}
		go utils.Logger(ctx, ah.Logger.Info, responseDesc, serviceDesc, requestID, map[string]interface{}{
			"tokenInfo": "secret",
			"msg":       response.ErrorResponse{ErrorCode: 0, ErrorDesc: "Success"},
		})
		response.ToJson(w, http.StatusOK, map[string]interface{}{
			"tokenInfo": tokenInfo,
			"msg":       response.ErrorResponse{ErrorCode: 0, ErrorDesc: "Success"},
		})
	}
}

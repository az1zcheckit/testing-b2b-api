package authentication

import (
	models "b2b-api/internal/models"
	response "b2b-api/internal/models"
	"b2b-api/internal/pkg/utils"
	"encoding/json"
	"net/http"
)

// Google-Authenticator  godoc
// @Summary      Google-Authenticator Confirmation route
// @Description  Google-Authenticator Confirmation
// @Tags         2FA
// @Produce      json
// @Param        X-Request-Id    header     string  true  " "  Format(UUID)
// @Param        otp    header     string  true  " "  Format(uint)
// @Success      200  {object}   models.Token
// @Failure      401  {object}   models.ErrorResponse
// @Router       /2fa/confirm-google-totp [get]
func (ah authHandler) GAuthenticator() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		serviceDesc := "Google-Authenticator"
		requestID := r.Header.Get("X-Request-ID")
		if len(requestID) == 0 {
			go utils.Logger(ctx, ah.Logger.Error, requestDesc, serviceDesc, requestID, response.RequestIdIsEmpty)
			response.ToJson(w, http.StatusSeeOther, response.RequestIdIsEmpty)
			return
		}
		totp := r.Header.Get("otp")
		if len(totp) == 0 {
			go utils.Logger(ctx, ah.Logger.Error, requestDesc, serviceDesc, requestID, response.OTPIncorrect)
			response.ToJson(w, http.StatusSeeOther, response.OTPIncorrect)
			return
		}

		usr := &models.User{}

		userInfo, err := ah.svc.RedisInstance().Get(ctx, requestID)
		go utils.Logger(ctx, ah.Logger.Info, requestDesc, serviceDesc, requestID, "X-Request-Id is:", requestID, "Otp is: secret")
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

		if len(usr.Security.TwoFA.Gauth.GauthSecretEndpoint) < 16 {
			go utils.Logger(ctx, ah.Logger.Error, requestDesc, serviceDesc, requestID, response.GauthIsNotActive)
			response.ToJson(w, http.StatusSeeOther, response.GauthIsNotActive)
			return
		}

		IsCorrect, err := ah.svc.TwoFaInstance().GAuthenticate(usr.Security.TwoFA.Gauth.GauthSecretEndpoint, totp)

		if !IsCorrect {
			go utils.Logger(ctx, ah.Logger.Error, requestDesc, serviceDesc, requestID, response.OTPIncorrect)
			response.ToJson(w, http.StatusSeeOther, response.OTPIncorrect)
			return
		}

		tokenInfo := response.User{Token: usr.Token, RefreshToken: usr.RefreshToken, Role: usr.Role}
		
		if len(usr.TransactionId) != 0 {
			sqlResponse := ah.svc.AuthRepositoryInstance().ConfirmTransaction(ctx, usr.Token, "EN",usr.TransactionId)
			if sqlResponse.ErrorCode != 0 {
				go utils.Logger(ctx, ah.Logger.Error, requestDesc, serviceDesc, requestID, sqlResponse)
				response.ToJson(w, http.StatusInternalServerError, sqlResponse)
				return
			}
			go utils.Logger(ctx, ah.Logger.Error, requestDesc, serviceDesc, requestID,sqlResponse)
			response.ToJson(w, http.StatusOK, sqlResponse)
			return
		}

		if len(usr.Token) != 0 {
			sqlSessionInsertResponse := ah.svc.AuthRepositoryInstance().SessionInsert(ctx, usr.Phone, usr.Token, "EN")
			if sqlSessionInsertResponse.ErrorCode != 0 {
				go utils.Logger(ctx, ah.Logger.Error, requestDesc, serviceDesc, requestID, sqlSessionInsertResponse)
				response.ToJson(w, http.StatusInternalServerError, sqlSessionInsertResponse)
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

package authentication

import (
	models "b2b-api/internal/models"
	response "b2b-api/internal/models"
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
// @Failure     303  {object}   models.ErrorResponse
// @Router       /2fa/confirm-google-totp [get]
func (ah authHandler) GAuthenticator() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		requestID := r.Header.Get("X-Request-ID")
		if len(requestID) == 0 {
			response.ToJson(w, http.StatusSeeOther, response.RequestIdIsEmpty)
			return
		}
		totp := r.Header.Get("otp")
		if len(totp) == 0 {
			response.ToJson(w, http.StatusSeeOther, response.OTPIncorrect)
			return
		}

		usr := &models.User{}

		userInfo, err := ah.svc.RedisInstance().Get(ctx, requestID)

		if err != nil {
			response.ToJson(w, http.StatusSeeOther, response.SetError(err))
			return
		}

		if err := json.Unmarshal([]byte(userInfo), usr); err != nil {
			response.ToJson(w, http.StatusBadRequest, response.SetError(err))
			return
		}

		if len(usr.Security.TwoFA.Gauth.GauthSecretEndpoint) < 16 {
			response.ToJson(w, http.StatusSeeOther, response.GauthIsNotActive)
			return
		}

		IsCorrect, err := ah.svc.TwoFaInstance().GAuthenticate(usr.Security.TwoFA.Gauth.GauthSecretEndpoint, totp)

		if !IsCorrect {
			response.ToJson(w, http.StatusSeeOther, response.OTPIncorrect)
			return
		}

		tokenInfo := response.User{Token: usr.Token, RefreshToken: usr.RefreshToken, Role: usr.Role}

		sqlSessionInsertResponse := ah.svc.AuthRepositoryInstance().SessionInsert(ctx, usr.Phone, usr.Token, "EN")
		if sqlSessionInsertResponse.ErrorCode != 0 {
			response.ToJson(w, http.StatusInternalServerError, sqlSessionInsertResponse)
			return
		}

		response.ToJson(w, http.StatusOK, map[string]interface{}{
			"tokenInfo": tokenInfo,
			"msg":       response.ErrorResponse{ErrorCode: 0, ErrorDesc: "Success"},
		})
	}
}

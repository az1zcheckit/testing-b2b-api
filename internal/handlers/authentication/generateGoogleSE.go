package authentication

import (
	response "b2b-api/internal/models"
	"net/http"
)

func (ah authHandler) GenerateGoogleSE() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		secret := ah.svc.TwoFaInstance().GauthGenerateSecretEndpoint()
		response.Json(w, http.StatusSeeOther, secret)
	}
}

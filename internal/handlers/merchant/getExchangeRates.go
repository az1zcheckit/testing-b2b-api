package merchant

import (
	response "b2b-api/internal/models"
	"net/http"
)

// GetExchangeRates    godoc
// @Summary      GetExchangeRates route
// @Description  GetExchangeRates
// @Tags         Merchant
// @Accept       json
// @Produce      json
// @Param        token   header     string  true  " "  Format(UUID)
// @Success      200  {object}   models.GetExchangeRatesResponse
// @Failure     303  {object}   models.ErrorResponse
// @Router       /merchant/get-rates [get]
func (mh merchHandler) GetExchangeRates() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		token := r.Header.Get("token")
		if len(token) == 0 {
			response.ToJson(w, http.StatusSeeOther, response.TokenIsEmpty)
			return
		}

		base, rates, err := mh.svc.MerchRepositoryInstance().GetExchangeRates(ctx, token, "EN")

		if err.ErrorCode != 0 {
			response.ToJson(w, http.StatusSeeOther, response.SetError(err))
			return
		}

		response.ToJson(w, http.StatusOK, map[string]interface{}{
			"base":  base,
			"rates": rates,
			"msg":   err,
		})
	}
}

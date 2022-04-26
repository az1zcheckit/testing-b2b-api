package merchant

import (
	response "b2b-api/internal/models"
	"b2b-api/internal/pkg/utils"
	"fmt"
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
		serviceDesc := "Get-Exchange-Rates-Handler"
		requestID := fmt.Sprintf("%v", r.Context().Value(utils.CTXRequestID))
		token := r.Header.Get("token")
		if len(token) == 0 {
			go utils.Logger(ctx, mh.Logger.Error, requestDesc, serviceDesc, requestID, response.TokenIsEmpty)
			response.ToJson(w, http.StatusSeeOther, response.TokenIsEmpty)
			return
		}
		go utils.Logger(ctx, mh.Logger.Info, requestDesc, serviceDesc, requestID, "")
		base, rates, err := mh.svc.MerchRepositoryInstance().GetExchangeRates(ctx, token, "EN")
		if err.ErrorCode != 0 {
			go utils.Logger(ctx, mh.Logger.Error, requestDesc, serviceDesc, requestID, response.SetError(err))
			response.ToJson(w, http.StatusSeeOther, response.SetError(err))
			return
		}

		go utils.Logger(ctx, mh.Logger.Info, responseDesc, serviceDesc, requestID, map[string]interface{}{
			"base":  base,
			"rates": rates,
			"msg":   err,
		})

		response.ToJson(w, http.StatusOK, map[string]interface{}{
			"base":  base,
			"rates": rates,
			"msg":   err,
		})
	}
}

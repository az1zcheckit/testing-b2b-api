package merchant

import (
	response "b2b-api/internal/models"
	"b2b-api/internal/pkg/utils"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// GetLimitConversion    godoc
// @Summary      GetLimitConversion route
// @Description  GetLimitConversion
// @Tags         Merchant
// @Accept       json
// @Produce      json
// @Param        token   header     string  true  " "  Format(UUID)
// @Param        currency   query     string  true  " "  Format(json)
// @Success      200  {object}   models.GetLimitConversion
// @Failure      401  {object}   models.ErrorResponse
// @Router       /merchant/get-limit-conversion [get]
func (mh merchHandler) GetLimitConversion() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		serviceDesc := "Get-Limit-Conversion-Handler"
		requestID := fmt.Sprintf("%v", r.Context().Value(utils.CTXRequestID))
		token := r.Header.Get("token")
		if len(token) == 0 {
			go utils.Logger(ctx, mh.Logger.Error, requestDesc, serviceDesc, requestID, response.TokenIsEmpty)
			response.ToJson(w, http.StatusSeeOther, response.TokenIsEmpty)
			return
		}
		currency := mux.Vars(r)["currency"]

		go utils.Logger(ctx, mh.Logger.Info, requestDesc, serviceDesc, requestID, currency)

		currentLimit, err := mh.svc.MerchRepositoryInstance().GetLimitConversion(ctx, token, "EN", currency)

		if err.ErrorCode != 0 {
			go utils.Logger(ctx, mh.Logger.Error, requestDesc, serviceDesc, requestID, err)
			response.ToJson(w, http.StatusSeeOther, err)
			return
		}

		go utils.Logger(ctx, mh.Logger.Info, responseDesc, serviceDesc, requestID, map[string]interface{}{
			"currentLimit": currentLimit,
			"msg":          err,
		})

		response.ToJson(w, http.StatusOK, map[string]interface{}{
			"currentLimit": currentLimit,
			"msg":          err,
		})
	}
}

package merchant

import (
	response "b2b-api/internal/models"
	"b2b-api/internal/pkg/utils"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// GetConvertedValue    godoc
// @Summary      GetConvertedValue route
// @Description  GetConvertedValue
// @Tags         Merchant
// @Accept       json
// @Produce      json
// @Param        token   header     string  true  " "  Format(UUID)
// @Param        value   query     float64  true  " "  Format(json)
// @Param        currencyFrom   query     string  true  " "  Format(json)
// @Param        currencyTo   query     string  true  " "  Format(json)
// @Success      200  {object}   models.ConvertedValueResponse
// @Failure      401  {object}   models.ErrorResponse
// @Router       /merchant/get-converted-value [get]
func (mh merchHandler) GetConvertedValue() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		serviceDesc := "Get-Converted-Value-Handler"
		requestID := fmt.Sprintf("%v", r.Context().Value(utils.CTXRequestID))
		token := r.Header.Get("token")
		value := mux.Vars(r)["value"]
		currencyFrom := mux.Vars(r)["currencyFrom"]
		currencyTo := mux.Vars(r)["currencyTo"]
		if len(token) == 0 {
			go utils.Logger(ctx, mh.Logger.Error, requestDesc, serviceDesc, requestID, response.TokenIsEmpty)
			response.ToJson(w, http.StatusSeeOther, response.TokenIsEmpty)
			return
		}
		convValue, err := mh.svc.MerchRepositoryInstance().GetConvertedValue(ctx, token, "EN",value,currencyFrom,currencyTo)
		
		if err.ErrorCode != 0 {
			go utils.Logger(ctx, mh.Logger.Error, requestDesc, serviceDesc, requestID, err)
			response.ToJson(w, http.StatusSeeOther, err)
			return
		}

		go utils.Logger(ctx, mh.Logger.Info, requestDesc, serviceDesc, requestID, convValue)
		response.ToJson(w, http.StatusOK, map[string]interface{}{
			"convertedValue": convValue,
			"msg":          err,
		})
	}
}

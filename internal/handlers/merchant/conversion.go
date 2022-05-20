package merchant

import (
	response "b2b-api/internal/models"
	"b2b-api/internal/models/merchant"
	"b2b-api/internal/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

// Conversion    godoc
// @Summary      Conversion route
// @Description  Conversion
// @Tags         Merchant
// @Accept       json
// @Produce      json
// @Param        token   header     string  true  " "  Format(UUID)
// @Param        conversion body  models.ConversionRequest  true  "conversion"  Format(json)
// @Success      200  {object}   models.ConversionResponse
// @Failure      401  {object}   models.ErrorResponse
// @Router       /merchant/conversion [post]
func (mh merchHandler) Conversion() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		token := r.Header.Get("token")
		serviceDesc := "Conversion-Handler"
		requestID := fmt.Sprintf("%v", r.Context().Value(utils.CTXRequestID))
		if len(token) == 0 {
			go utils.Logger(ctx, mh.Logger.Error, requestDesc, serviceDesc, requestID, response.TokenIsEmpty)
			response.ToJson(w, http.StatusBadRequest, response.TokenIsEmpty)
		}
		go utils.Logger(ctx, mh.Logger.Info, requestDesc, serviceDesc, requestID, "")
		converRequest := &merchant.ConversionRequest{}
		
		if err := json.NewDecoder(r.Body).Decode(converRequest); err != nil {
			go utils.Logger(ctx, mh.Logger.Error, requestDesc, serviceDesc, requestID, err.Error())
			response.ToJson(w, http.StatusBadRequest, err)
			return
		}
		
		transaction, err := mh.svc.MerchRepositoryInstance().Conversion(ctx, token, "EN", *converRequest)

		if err.ErrorCode != 0 {
			go utils.Logger(ctx, mh.Logger.Error, requestDesc, serviceDesc, requestID, err)
			response.ToJson(w, http.StatusSeeOther, err)
			return
		}

		go utils.Logger(ctx, mh.Logger.Info, responseDesc, serviceDesc, requestID, map[string]interface{}{
			"transaction": transaction,
			"msg":        err,
		})

		response.ToJson(w, http.StatusOK, map[string]interface{}{
			"transaction": transaction,
			"msg":        err,
		})
	}
}

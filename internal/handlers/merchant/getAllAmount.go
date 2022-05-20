package merchant

import (
	response "b2b-api/internal/models"
	"b2b-api/internal/pkg/utils"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// GetAllAmount    godoc
// @Summary      GetAllAmount route
// @Description  GetAllAmount
// @Tags         Merchant
// @Accept       json
// @Produce      json
// @Param        token   header     string  true  " "  Format(UUID)
// @Param        currency   query     string  true  " "  Format(json)
// @Success      200  {object}   models.GetAllAmountsResponse
// @Failure      401 {object}   models.ErrorResponse
// @Router       /merchant/get-all-amount [get]
func (mh merchHandler) GetAllAmount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		serviceDesc := "Get-All-Amount-Handler"
		requestID := fmt.Sprintf("%v", r.Context().Value(utils.CTXRequestID))
		token := r.Header.Get("token")
		if len(token) == 0 {
			go utils.Logger(ctx, mh.Logger.Error, requestDesc, serviceDesc, requestID, response.TokenIsEmpty)
			response.ToJson(w, http.StatusSeeOther, response.TokenIsEmpty)
			return
		}
		currency := mux.Vars(r)["currency"]
		go utils.Logger(ctx, mh.Logger.Info, requestDesc, serviceDesc, requestID, currency)

		allAmount, err := mh.svc.MerchRepositoryInstance().GetAllAmount(ctx, token, "EN", currency)

		if err.ErrorCode != 0 {
			go utils.Logger(ctx, mh.Logger.Error, requestDesc, serviceDesc, requestID, err)
			response.ToJson(w, http.StatusSeeOther, err)
			return
		}
		err.AditionalInfo = allAmount
		go utils.Logger(ctx, mh.Logger.Info, responseDesc, serviceDesc, requestID, err)
		response.ToJson(w, http.StatusOK, err)
	}
}

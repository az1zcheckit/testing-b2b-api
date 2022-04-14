package merchant

import (
	response "b2b-api/internal/models"
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
// @Failure     303  {object}   models.ErrorResponse
// @Router       /merchant/get-all-amount [get]
func (mh merchHandler) GetAllAmount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		token := r.Header.Get("token")
		if len(token) == 0 {
			response.ToJson(w, http.StatusSeeOther, response.TokenIsEmpty)
			return
		}

		currency := mux.Vars(r)["currency"]

		allAmount, err := mh.svc.MerchRepositoryInstance().GetAllAmount(ctx, token, "EN", currency)

		if err.ErrorCode != 0 {
			response.ToJson(w, http.StatusSeeOther, response.SetError(err))
			return
		}
		err.AditionalInfo = allAmount

		response.ToJson(w, http.StatusOK, err)
	}
}

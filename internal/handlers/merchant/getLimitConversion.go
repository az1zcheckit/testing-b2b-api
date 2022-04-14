package merchant

import (
	response "b2b-api/internal/models"
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
// @Failure      303  {object}   models.ErrorResponse
// @Router       /merchant/get-limit-conversion [get]
func (mh merchHandler) GetLimitConversion() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		token := r.Header.Get("token")
		if len(token) == 0 {
			response.ToJson(w, http.StatusSeeOther, response.TokenIsEmpty)
			return
		}
		currency := mux.Vars(r)["currency"]

		currentLimit, err := mh.svc.MerchRepositoryInstance().GetLimitConversion(ctx, token, "EN", currency)

		if err.ErrorCode != 0 {
			response.ToJson(w, http.StatusSeeOther, response.SetError(err))
			return
		}

		response.ToJson(w, http.StatusOK, map[string]interface{}{
			"currentLimit": currentLimit,
			"msg":          err,
		})
	}
}

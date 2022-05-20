package merchant

import (
	response "b2b-api/internal/models"
	"github.com/gorilla/mux"
	"net/http"
)

// GetAccountDetails    godoc
// @Summary      		GetAccountDetails route
// @Description  		GetAccountDetails
// @Tags         		Merchant
// @Accept       		json
// @Produce      		json
// @Param        		token   header     string  true  " "  Format(UUID)
// @Param        		accountNumber   query     string  true  " "  Format(json)
// @Success      		200  {object}   models.AccountDetailsForBankTrans
// @Failure     		401  {object}   models.ErrorResponse
// @Router       		/merchant/get-account-details [get]
func (mh merchHandler) GetAccountDetails() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		token := r.Header.Get("token")
		if len(token) == 0 {
			response.ToJson(w, http.StatusBadRequest, response.TokenIsEmpty)
		}
		accountNumber := mux.Vars(r)["accountNumber"]
		local, err := mh.svc.MerchRepositoryInstance().GetAccountDetails(ctx, token, "EN", accountNumber)
		if err.ErrorCode != 0 {
			response.ToJson(w, http.StatusSeeOther, response.SetError(err))
			return
		}
		response.ToJson(w, http.StatusOK, map[string]interface{}{
			"local": local,
			"msg":   err,
		})
	}
}

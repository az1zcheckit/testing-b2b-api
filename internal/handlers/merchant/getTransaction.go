package merchant

import (
	response "b2b-api/internal/models"
	"github.com/gorilla/mux"
	"net/http"
)

// GetTransactions    godoc
// @Summary      GetTransactions route
// @Description  GetTransactions
// @Tags         Merchant
// @Accept       json
// @Produce      json
// @Param        token    header     string  true  " "  Format(UUID)
// @Param        dateFrom query      string  true  " "  Format(json)
// @Param        dateTo   query      string  true  " "  Format(json)
// @Success      200  {object}   models.GetTransactionResponse
// @Failure     303  {object}   models.ErrorResponse
// @Router       /merchant/get-transactions [get]
func (mh merchHandler) GetTransactions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		token := r.Header.Get("token")
		if len(token) == 0 {
			response.ToJson(w, http.StatusSeeOther, response.TokenIsEmpty)
			return
		}

		dateFrom := mux.Vars(r)["dateFrom"]
		dateTo := mux.Vars(r)["dateTo"]

		transaction, err := mh.svc.MerchRepositoryInstance().GetTransactions(ctx, token, "EN", dateFrom, dateTo)

		if err.ErrorCode != 0 {
			response.ToJson(w, http.StatusSeeOther, response.SetError(err))
			return
		}

		response.ToJson(w, http.StatusOK, map[string]interface{}{
			"transactionList": transaction,
			"msg":             err,
		})
	}
}

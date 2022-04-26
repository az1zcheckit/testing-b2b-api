package merchant

import (
	response "b2b-api/internal/models"
	"net/http"
)

// GetTransactionByID    godoc
// @Summary      GetTransactionByID route
// @Description  GetTransactionByID
// @Tags         Merchant
// @Accept       json
// @Produce      json
// @Param        token    header     string  true  " "  Format(UUID)
// @Param        transId query      string  true  " "  Format(json)
// @Success      200  {object}   models.GetTransactionByIdResponse
// @Failure     303  {object}   models.ErrorResponse
// @Router       /merchant/get-transaction-by-id [get]
func (mh merchHandler) GetTransactionByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		token := r.Header.Get("token")
		if len(token) == 0 {
			response.ToJson(w, http.StatusSeeOther, response.TokenIsEmpty)
			return
		}
		transId := r.Header.Get("transId")
		//transById := &merchant.Transaction{}
		//if err := json.NewDecoder(r.Body).Decode(transById); err != nil {
		//	response.ToJson(w, http.StatusBadRequest, err)
		//	return
		//}
		transaction, err := mh.svc.MerchRepositoryInstance().GetTransactionsByID(ctx, token, "EN", transId)

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

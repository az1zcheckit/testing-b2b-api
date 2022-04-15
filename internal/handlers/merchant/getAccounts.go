package merchant

import (
	response "b2b-api/internal/models"
	"net/http"
)

// GetAccounts    godoc
// @Summary      GetAccounts route
// @Description  GetAccounts
// @Tags         Merchant
// @Accept       json
// @Produce      json
// @Param        token   header     string  true  " "  Format(UUID)
// @Success      200  {object}   models.GetAccountsResponse
// @Failure     303  {object}   models.ErrorResponse
// @Router       /merchant/get-accounts [get]
func (mh merchHandler) GetAccounts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		token := r.Header.Get("token")
		if len(token) == 0 {
			mh.svc.LoggerInstance().Error("%v", response.TokenIsEmpty)
			response.ToJson(w, http.StatusSeeOther, response.TokenIsEmpty)
			return
		}
		merchName, accounts, err := mh.svc.MerchRepositoryInstance().GetAccounts(ctx, token, "EN")

		if err.ErrorCode != 0 {
			response.ToJson(w, http.StatusSeeOther, response.SetError(err))
			return
		}

		response.ToJson(w, http.StatusOK, map[string]interface{}{
			"MerchantName": merchName,
			"accountsList": accounts,
			"msg":          err,
		})
	}
}

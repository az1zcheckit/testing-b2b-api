package merchant

import (
	response "b2b-api/internal/models"
	"b2b-api/internal/pkg/utils"
	"fmt"
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
// @Failure      401  {object}   models.ErrorResponse
// @Router       /merchant/get-accounts [get]
func (mh merchHandler) GetAccounts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		serviceDesc := "Get-Accounts-Handler"
		requestID := fmt.Sprintf("%v", r.Context().Value(utils.CTXRequestID))
		token := r.Header.Get("token")
		if len(token) == 0 {
			go utils.Logger(ctx, mh.Logger.Error, requestDesc, serviceDesc, requestID, response.TokenIsEmpty)
			response.ToJson(w, http.StatusSeeOther, response.TokenIsEmpty)
			return
		}
		merchName, accounts, err := mh.svc.MerchRepositoryInstance().GetAccounts(ctx, token, "EN")
		go utils.Logger(ctx, mh.Logger.Info, requestDesc, serviceDesc, requestID, merchName)
		if err.ErrorCode != 0 {
			go utils.Logger(ctx, mh.Logger.Error, requestDesc, serviceDesc, requestID, err)
			response.ToJson(w, http.StatusSeeOther, err)
			return
		}

		go utils.Logger(ctx, mh.Logger.Info, responseDesc, serviceDesc, requestID, map[string]interface{}{
			"MerchantName": merchName,
			"accountsList": accounts,
			"msg":          err,
		})

		response.ToJson(w, http.StatusOK, map[string]interface{}{
			"MerchantName": merchName,
			"accountsList": accounts,
			"msg":          err,
		})
	}
}

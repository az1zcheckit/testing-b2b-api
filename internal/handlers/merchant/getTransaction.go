package merchant

import (
	response "b2b-api/internal/models"
	"b2b-api/internal/models/merchant"
	"b2b-api/internal/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

// GetHistoryOfTransactions    godoc
// @Summary      GetHistoryOfTransactions route
// @Description GetHistoryOfTransactions
// @Tags         Merchant
// @Accept       json
// @Produce      json
// @Param        token    header     string  true  " "  Format(UUID)
// @Param        history filters body  merchant.FilterHistoryTransaction  true  "history filters"  Format(json)
// @Success      200  {object}   models.HistoryOfTransactionsResponse
// @Failure      401  {object}   models.ErrorResponse
// @Router       /merchant/transactions [post]
func (mh merchHandler) GetHistoryOfTransactions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		token := r.Header.Get("token")
		serviceDesc := "HistoryOfTransactions-Handler"
		requestID := fmt.Sprintf("%v", r.Context().Value(utils.CTXRequestID))
		if len(token) == 0 {
			go utils.Logger(ctx, mh.Logger.Error, requestDesc, serviceDesc, requestID, response.TokenIsEmpty)
			response.ToJson(w, http.StatusSeeOther, response.TokenIsEmpty)
			return
		}

		var filters merchant.FilterHistoryTransaction

		errb := json.NewDecoder(r.Body).Decode(&filters)
		if errb != nil {
			go utils.Logger(ctx, mh.Logger.Error, requestDesc, serviceDesc, requestID, errb.Error())
			//mh.svc.LoggerInstance().Error("", "GetHistoryOfTransactions error: ", errb)
			response.ToJson(w, http.StatusBadRequest, response.SetError(errb))
		}
		go utils.Logger(ctx, mh.Logger.Info, requestDesc, serviceDesc, requestID, "filters:", filters)
		//mh.svc.LoggerInstance().Info("", "GetHistoryOfTransactions filters:", filters)

		var (
			history []merchant.Transaction
			err     response.ErrorResponse
		)

		if len(filters.Account) == 0 {
			history, err = mh.svc.MerchRepositoryInstance().GetAllHistoryTransactions(ctx, token, "EN", &filters)
			if err.ErrorCode != 0 {
				go utils.Logger(ctx, mh.Logger.Error, requestDesc, serviceDesc, requestID, err)
				response.ToJson(w, http.StatusSeeOther, response.SetError(err))
				return
			}

		} else {
			history, err = mh.svc.MerchRepositoryInstance().GetAccountHistoryTransactions(ctx, token, "EN", &filters)
			if err.ErrorCode != 0 {
				go utils.Logger(ctx, mh.Logger.Error, requestDesc, serviceDesc, requestID, err)
				response.ToJson(w, http.StatusSeeOther, response.SetError(err))
				return
			}
		}

		groups := GroupTransactions(history)

		response.ToJson(w, http.StatusOK, map[string]interface{}{
			"groups": groups,
			"msg":    err,
		})
	}
}

func GroupTransactions(transactions []merchant.Transaction) (groups []merchant.GroupOfTransactions) {
	sort.Slice(transactions, func(i, j int) bool {
		return transactions[i].CreatedAt.Before(transactions[j].CreatedAt)
	})

	sortedByDate := make(map[string][]merchant.Transaction, len(transactions))

	for _, transaction := range transactions {
		date := transaction.CreatedAt.Format("2006/01/02")
		sortedByDate[date] = append(sortedByDate[date], transaction)
	}

	for date, transactions := range sortedByDate {
		var group merchant.GroupOfTransactions
		group.GroupDate = date
		group.Transactions = transactions
		groups = append(groups, group)
	}
	return groups
}

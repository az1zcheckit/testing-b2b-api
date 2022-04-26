package merchant

import (
	response "b2b-api/internal/models"
	"b2b-api/internal/models/merchant"
	"b2b-api/internal/pkg/utils"
	"database/sql"
	"fmt"
	"golang.org/x/net/context"
)

func (r repository) GetAllAmount(ctx context.Context, token, lang, currency string) (allAmount merchant.GetAllAmount, errResponse response.ErrorResponse) {
	serviceDesc := "Get-All-Amount-Repository"
	requestID := fmt.Sprintf("%v", ctx.Value(utils.CTXRequestID))
	db := r.Oracle.GetConnection()
	//stmt, err := db.PrepareContext(ctx, getAllAmountSQL)
	//if err != nil {
	//	errResponse = response.SetError(err)
	//	return
	//}
	//defer stmt.Close()

	go utils.Logger(ctx, r.Logger.Info, requestDesc, serviceDesc, requestID, currency)

	_, err := db.ExecContext(ctx, getAllAmountSQL,
		token,
		lang,
		currency,
		sql.Out{Dest: &allAmount.Currency},
		sql.Out{Dest: &allAmount.Balance},
		sql.Out{Dest: &errResponse.ErrorCode},
		sql.Out{Dest: &errResponse.ErrorDesc},
	)

	if err != nil {
		go utils.Logger(ctx, r.Logger.Error, requestDesc, serviceDesc, requestID, err.Error())
		errResponse = response.SetError(err)
		return
	}

	if errResponse.ErrorCode != 0 {
		go utils.Logger(ctx, r.Logger.Error, requestDesc, serviceDesc, requestID, err.Error())
		return
	}

	go utils.Logger(ctx, r.Logger.Info, responseDesc, serviceDesc, requestID, errResponse)

	return
}

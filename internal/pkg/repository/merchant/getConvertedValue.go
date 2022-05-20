package merchant

import (
	response "b2b-api/internal/models"
	"b2b-api/internal/models/merchant"
	"b2b-api/internal/pkg/utils"
	"database/sql"
	"fmt"
	"golang.org/x/net/context"
)

func (r repository) GetConvertedValue(ctx context.Context, token, lang,value,currencyFrom,currencyTo string) (conerValue merchant.ConvertedValue, errResponse response.ErrorResponse) {
	var rate string
	serviceDesc := "Get-Converted-Value"
	requestID := fmt.Sprintf("%v", ctx.Value(utils.CTXRequestID))
	db := r.Oracle.GetConnection()
	_, err := db.ExecContext(ctx, getConvertedValueSQL,
		token,
		lang,
		value,
		currencyFrom,
		currencyTo,
		sql.Out{Dest: &conerValue.Value},
		sql.Out{Dest: &rate},
		sql.Out{Dest: &errResponse.ErrorCode},
		sql.Out{Dest: &errResponse.ErrorDesc},
	)

	if err != nil {
		errResponse = response.DbExecContext
		go utils.Logger(ctx, r.Logger.Error, requestDesc, serviceDesc, requestID, response.SetError(err))
		return
	}

	if errResponse.ErrorCode != 0 {
		go utils.Logger(ctx, r.Logger.Error, requestDesc, serviceDesc, requestID, errResponse)
		return
	}
	
	if currencyFrom == "TJS"{
		conerValue.Rate = fmt.Sprintf("1 %s = %s %s",merchant.GetSymCurrency(currencyTo),rate,merchant.GetSymCurrency(currencyFrom))	
	}else{
		conerValue.Rate = fmt.Sprintf("1 %s = %s %s",merchant.GetSymCurrency(currencyFrom),rate,merchant.GetSymCurrency(currencyTo))
	}
	
	go utils.Logger(ctx, r.Logger.Info, responseDesc, serviceDesc, requestID, conerValue, errResponse)

	return
}

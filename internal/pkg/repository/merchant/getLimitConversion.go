package merchant

import (
	response "b2b-api/internal/models"
	"b2b-api/internal/pkg/utils"
	"database/sql"
	"fmt"
	"golang.org/x/net/context"
)

func (r repository) GetLimitConversion(ctx context.Context, token, lang, currency string) (currentLimit float64, errResponse response.ErrorResponse) {
	serviceDesc := "Get-Limit-Conversion-Repository"
	requestID := fmt.Sprintf("%v", ctx.Value(utils.CTXRequestID))
	db := r.Oracle.GetConnection()
	//stmt, err := db.PrepareContext(ctx, getLimitConversionSQL)
	//if err != nil {
	//	errResponse = response.SetError(err)
	//	return
	//}
	//defer stmt.Close()
	go utils.Logger(ctx, r.Logger.Info, requestDesc, serviceDesc, requestID, currency)

	_, err := db.ExecContext(ctx, getLimitConversionSQL,
		token,
		lang,
		currency,
		sql.Out{Dest: &currentLimit},
		sql.Out{Dest: &errResponse.ErrorCode},
		sql.Out{Dest: &errResponse.ErrorDesc},
	)
	if err != nil {
		go utils.Logger(ctx, r.Logger.Error, responseDesc, serviceDesc, requestID, err.Error())
		errResponse = response.SetError(err)
	}

	go utils.Logger(ctx, r.Logger.Info, responseDesc, serviceDesc, requestID, currentLimit, errResponse)

	return
}

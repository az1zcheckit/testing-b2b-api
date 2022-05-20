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
		errResponse = response.DbExecContext
		go utils.Logger(ctx, r.Logger.Error, requestDesc, serviceDesc, requestID, response.SetError(err))
		return
	}

	if errResponse.ErrorCode != 0 {
		go utils.Logger(ctx, r.Logger.Error, responseDesc, serviceDesc, requestID, errResponse)
		return
	}

	go utils.Logger(ctx, r.Logger.Info, responseDesc, serviceDesc, requestID, currentLimit, errResponse)

	return
}

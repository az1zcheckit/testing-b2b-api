package merchant

import (
	response "b2b-api/internal/models"
	"b2b-api/internal/models/merchant"
	"b2b-api/internal/pkg/utils"
	"database/sql"
	"fmt"
	"golang.org/x/net/context"
)

func (r repository) Conversion(ctx context.Context, token, lang string, conv merchant.Conversion) (documentId int64, errResponse response.ErrorResponse) {
	serviceDesc := "Conversion-Repository"
	requestID := fmt.Sprintf("%v", ctx.Value(utils.CTXRequestID))
	db := r.Oracle.GetConnection()
	//stmt, err := db.PrepareContext(ctx, conversionSQL)
	//if err != nil {
	//	errResponse = response.SetError(err)
	//	return
	//}
	//defer stmt.Close()

	go utils.Logger(ctx, r.Logger.Info, requestDesc, serviceDesc, requestID, "")

	_, err := db.ExecContext(ctx, conversionSQL,
		token,
		lang,
		conv.AccFrom,
		conv.AccTo,
		conv.Amount,
		conv.Dest,
		sql.Out{Dest: &documentId},
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

	go utils.Logger(ctx, r.Logger.Info, responseDesc, serviceDesc, requestID, documentId, errResponse)

	return
}

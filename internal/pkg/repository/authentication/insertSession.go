package authentication

import (
	response "b2b-api/internal/models"
	"b2b-api/internal/pkg/utils"
	"database/sql"
	"fmt"
	"golang.org/x/net/context"
)

func (r repository) SessionInsert(ctx context.Context, phone string, token string, lang string) (errResponse response.ErrorResponse) {
	serviceDesc := "Session-Insert-Repository"
	requestID := fmt.Sprintf("%v", ctx.Value(utils.CTXRequestID))
	db := r.Oracle.GetConnection()
	go utils.Logger(ctx, r.Logger.Info, requestDesc, serviceDesc, requestID, phone)
	_, err := db.ExecContext(ctx, sessionInsertSQL,
		phone,
		token,
		lang,
		sql.Out{Dest: &errResponse.ErrorCode},
		sql.Out{Dest: &errResponse.ErrorDesc},
	)
	if err != nil {
		go utils.Logger(ctx, r.Logger.Error, requestDesc, serviceDesc, requestID, response.SetError(err))
		errResponse = response.SetError(err)
		return
	}
	go utils.Logger(ctx, r.Logger.Info, responseDesc, serviceDesc, requestID, errResponse)
	return
}

package authentication

import (
	response "b2b-api/internal/models"
	"database/sql"
	"golang.org/x/net/context"
)

func (r repository) SessionInsert(ctx context.Context, phone string, token string, lang string) (errResponse response.ErrorResponse) {
	db := r.Oracle.GetConnection()
	_, err := db.ExecContext(ctx, sessionInsertSQL,
		phone,
		token,
		lang,
		sql.Out{Dest: &errResponse.ErrorCode},
		sql.Out{Dest: &errResponse.ErrorDesc},
	)
	if err != nil {
		errResponse = response.SetError(err)
		return
	}
	return
}

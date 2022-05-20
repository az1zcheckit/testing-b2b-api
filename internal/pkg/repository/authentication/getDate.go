package authentication

import (
	response "b2b-api/internal/models"
	"database/sql"
	"golang.org/x/net/context"
)

func (r repository) GetDate(ctx context.Context) (date string, errResponse response.ErrorResponse) {
	db := r.Oracle.GetConnection()

	_, err := db.ExecContext(ctx, getDateSQL,
		sql.Out{Dest: &date},
		sql.Out{Dest: &errResponse.ErrorCode},
		sql.Out{Dest: &errResponse.ErrorDesc},
	)
	if err != nil {
		errResponse = response.SetError(err)
		return
	}
	return
}
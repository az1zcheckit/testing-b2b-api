package merchant

import (
	response "b2b-api/internal/models"
	"database/sql"
	"golang.org/x/net/context"
)

func (r repository) GetLimitConversion(ctx context.Context, token, lang, currency string) (currentLimit float64, errResponse response.ErrorResponse) {
	db := r.Oracle.GetConnection()
	stmt, err := db.PrepareContext(ctx, getLimitConversionSQL)
	if err != nil {
		errResponse = response.SetError(err)
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		token,
		lang,
		currency,
		sql.Out{Dest: &currentLimit},
		sql.Out{Dest: &errResponse.ErrorCode},
		sql.Out{Dest: &errResponse.ErrorDesc},
	)
	if err != nil {
		errResponse = response.SetError(err)
	}
	return
}

package merchant

import (
	response "b2b-api/internal/models"
	"b2b-api/internal/models/merchant"
	"database/sql"
	"golang.org/x/net/context"
)

func (r repository) GetAllAmount(ctx context.Context, token, lang, currency string) (allAmount merchant.GetAllAmount, errResponse response.ErrorResponse) {
	db := r.Oracle.GetConnection()
	stmt, err := db.PrepareContext(ctx, getAllAmountSQL)
	if err != nil {
		errResponse = response.SetError(err)
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		token,
		lang,
		currency,
		sql.Out{Dest: &allAmount.Currency},
		sql.Out{Dest: &allAmount.Balance},
		sql.Out{Dest: &errResponse.ErrorCode},
		sql.Out{Dest: &errResponse.ErrorDesc},
	)

	if err != nil {
		errResponse = response.SetError(err)
		return
	}

	if errResponse.ErrorCode != 0 {
		return
	}

	return
}

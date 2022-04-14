package merchant

import (
	response "b2b-api/internal/models"
	"b2b-api/internal/models/merchant"
	"database/sql"
	"golang.org/x/net/context"
)

func (r repository) Conversion(ctx context.Context, token, lang string, conv merchant.Conversion) (documentId int64, errResponse response.ErrorResponse) {
	db := r.Oracle.GetConnection()
	stmt, err := db.PrepareContext(ctx, conversionSQL)
	if err != nil {
		errResponse = response.SetError(err)
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
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
		errResponse = response.SetError(err)
		return
	}

	if errResponse.ErrorCode != 0 {
		return
	}

	return
}

package merchant

import (
	response "b2b-api/internal/models"
	"b2b-api/internal/models/merchant"
	"context"
	"database/sql"
)

func (r repository) GetTransactionsByID(ctx context.Context, token, lang, transId string) (transById merchant.Transaction, errResponse response.ErrorResponse) {
	//var refCursor driver.Rows
	//var example merchant.Transaction
	db := r.Oracle.GetConnection()
	//stmt, err := db.PrepareContext(ctx, getTransactionsSQL)
	//
	//if err != nil {
	//	errResponse = response.SetError(err)
	//	return nil, errResponse
	//}
	//defer stmt.Close()

	_, err := db.ExecContext(ctx, getTransactionByIdSQL,
		token,
		lang,
		transId,
		sql.Out{Dest: transById.CreatedDate},
		sql.Out{Dest: transById.ProceedDate},
		sql.Out{Dest: transById.AccFrom},
		sql.Out{Dest: transById.AccTo},
		sql.Out{Dest: transById.TransId},
		sql.Out{Dest: transById.TransType},
		sql.Out{Dest: transById.PaymentPurpose},
		sql.Out{Dest: transById.Amount},
		sql.Out{Dest: &errResponse.ErrorCode},
		sql.Out{Dest: &errResponse.ErrorDesc},
	)

	if err != nil {
		errResponse = response.SetError(err)
		return
	}

	if errResponse.ErrorCode != 0 {
		errResponse = response.SetError(err)
		return
	}

	return
}

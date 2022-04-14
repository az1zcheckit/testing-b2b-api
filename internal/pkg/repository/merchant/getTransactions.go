package merchant

import (
	response "b2b-api/internal/models"
	"b2b-api/internal/models/merchant"
	"database/sql"
	"database/sql/driver"
	"golang.org/x/net/context"
	"io"
)

func (r repository) GetTransactions(ctx context.Context, token, lang, dateFrom, dateTo string) ([]merchant.Transactions, response.ErrorResponse) {
	var refCursor driver.Rows
	var errResponse response.ErrorResponse

	db := r.Oracle.GetConnection()
	stmt, err := db.PrepareContext(ctx, getTransactionsSQL)

	if err != nil {
		errResponse = response.SetError(err)
		return nil, errResponse
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		token,
		lang,
		dateFrom,
		dateTo,
		sql.Out{Dest: &refCursor},
		sql.Out{Dest: &errResponse.ErrorCode},
		sql.Out{Dest: &errResponse.ErrorDesc},
	)

	if err != nil {
		errResponse = response.SetError(err)
		return nil, errResponse
	}

	if errResponse.ErrorCode != 0 {
		errResponse = response.SetError(err)
		return nil, errResponse
	}
	defer refCursor.Close()

	trans := make([]merchant.Transactions, 0, len(refCursor.Columns()))
	transaction := &merchant.Transactions{}

	rows := make([]driver.Value, len(refCursor.Columns()))

	for {
		if err := refCursor.Next(rows); err != nil {
			if err == io.EOF {
				break
			}
			continue
		}
		transaction.AccountKt = rows[0].(string)
		transaction.AccountDt = rows[1].(string)
		transaction.DocId = rows[2].(string)
		transaction.DateProcess = rows[3].(string)
		transaction.Nazn = rows[4].(string)
		transaction.TransType = rows[5].(string)
		transaction.SenderName = rows[6].(string)
		transaction.RecipientName = rows[7].(string)
		transaction.DocState = rows[8].(string)
		transaction.Amount = rows[9].(string)

		trans = append(trans, *transaction)
	}
	return trans, errResponse
}

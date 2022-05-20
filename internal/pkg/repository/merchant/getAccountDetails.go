package merchant

import (
	response "b2b-api/internal/models"
	"b2b-api/internal/models/merchant"
	"database/sql"
	"database/sql/driver"
	"golang.org/x/net/context"
	"io"
)

func (r repository) GetAccountDetails(ctx context.Context, token, lang, accountNumber string) (accountsDetails []merchant.AccountDetailsForBankTrans, errResponse response.ErrorResponse) {
	var refCursor driver.Rows
	db := r.Oracle.GetConnection()
	_, err := db.ExecContext(ctx, getAccountDetails,
		token,
		lang,
		accountNumber,
		sql.Out{Dest: &refCursor},
		sql.Out{Dest: &errResponse.ErrorCode},
		sql.Out{Dest: &errResponse.ErrorDesc},
	)
	if err != nil {
		errResponse = response.SetError(err)
		return
	}

	if errResponse.ErrorCode != 0 {
		errResponse = response.SetError(errResponse)
		return
	}
	accountDetails := &merchant.AccountDetailsForBankTrans{}
	rows := make([]driver.Value, len(refCursor.Columns()))
	for {
		if err := refCursor.Next(rows); err != nil {
			if err == io.EOF {
				break
			}
			continue
		}
		for i := 0; i < len(rows); i += 2 {
			accountDetails.Key = rows[i].(string)
			accountDetails.Value = rows[i+1].(string)
			accountsDetails = append(accountsDetails, *accountDetails)
		}
	}
	return accountsDetails, errResponse
}

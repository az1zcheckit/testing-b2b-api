package merchant

import (
	response "b2b-api/internal/models"
	"b2b-api/internal/models/merchant"
	"b2b-api/internal/pkg/utils"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"golang.org/x/net/context"
	"io"
)

func (r repository) GetAccounts(ctx context.Context, token, lang string) (merchName string, accounts []merchant.Accounts, errResponse response.ErrorResponse) {
	var refCursor driver.Rows
	serviceDesc := "Get-Accounts-Repository"
	requestID := fmt.Sprintf("%v", ctx.Value(utils.CTXRequestID))
	db := r.Oracle.GetConnection()

	go utils.Logger(ctx, r.Logger.Info, requestDesc, serviceDesc, requestID, merchName)

	_, err := db.ExecContext(ctx, getAccountsSQL,
		token,
		lang,
		sql.Out{Dest: &merchName},
		sql.Out{Dest: &refCursor},
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
	defer refCursor.Close()

	account := &merchant.Accounts{}

	rows := make([]driver.Value, len(refCursor.Columns()))

	for {
		if err = refCursor.Next(rows); err != nil {
			if err == io.EOF {
				break
			}
			continue
		}
		account.Account = rows[0].(string)
		account.Balance = rows[1].(string)
		account.Currency = rows[2].(string)
		account.TypeAcc = rows[3].(string)

		accounts = append(accounts, *account)
	}

	go utils.Logger(ctx, r.Logger.Info, responseDesc, serviceDesc, requestID, merchName, accounts, errResponse)

	return
}

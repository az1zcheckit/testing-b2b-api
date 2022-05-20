package merchant

import (
	response "b2b-api/internal/models"
	"b2b-api/internal/models/merchant"
	"b2b-api/internal/pkg/utils"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/godror/godror"
	"golang.org/x/net/context"
)

func (r repository) GetAllHistoryTransactions(ctx context.Context, token, lang string,
	filters *merchant.FilterHistoryTransaction) (transactions []merchant.Transaction,
	errResponse response.ErrorResponse) {

	var refCursor driver.Rows

	serviceDesc := "AllHistoryTransactions-Repository"
	requestID := fmt.Sprintf("%v", ctx.Value(utils.CTXRequestID))
	db := r.Oracle.GetConnection()

	go utils.Logger(ctx, r.Logger.Info, requestDesc, serviceDesc, requestID, "")

	_, err := db.ExecContext(ctx, getTransactionsSQL,
		token,
		lang,
		filters.Pagination.Count,
		filters.Pagination.CurrentPage,
		filters.FromDate,
		filters.ToDate,
		sql.Out{Dest: &refCursor},
		sql.Out{Dest: &errResponse.ErrorCode},
		sql.Out{Dest: &errResponse.ErrorDesc},
	)

	if err != nil {
		errResponse = response.DbExecContext
		go utils.Logger(ctx, r.Logger.Error, requestDesc, serviceDesc, requestID, response.SetError(err))
		return nil, errResponse
	}

	if errResponse.ErrorCode != 0 {
		log.Println(err)
		go utils.Logger(ctx, r.Logger.Error, requestDesc, serviceDesc, requestID, errResponse)
		return nil, errResponse
	}
	defer refCursor.Close()

	trans := make([]merchant.Transaction, 0, len(refCursor.Columns()))
	transaction := &merchant.Transaction{}

	rows := make([]driver.Value, len(refCursor.Columns()))

	for {
		if err := refCursor.Next(rows); err != nil {
			if err == io.EOF {
				break
			}
			continue
		}
		transaction.FromAccount.Number = rows[0].(string)
		transaction.ToAccount.Number = rows[1].(string)
		transaction.Id = rows[2].(string)
		transaction.CreatedAtStr = rows[3].(string)
		transaction.PaymentPurpose = rows[4].(string)
		transaction.Type = rows[5].(string)
		transaction.FromAccount.Title = rows[6].(string)
		transaction.ToAccount.Title = rows[7].(string)
		transaction.Status = rows[8].(string)
		transaction.FromAccount.Amount.Value = rows[9].(string)
		transaction.FromAccount.Amount.Currency = rows[10].(string)
		transaction.DocumentNumber = rows[11].(string)
		transaction.ProceededAtStr = rows[12].(string)
		parse, err := time.Parse("02.01.2006 15:04:05", transaction.CreatedAtStr)
		if err != nil {
			go utils.Logger(ctx, r.Logger.Error, requestDesc, serviceDesc, requestID, response.SetError(err))
		}
		transaction.CreatedAt = parse

		parseProceed, err := time.Parse("02.01.2006 15:04:05", transaction.CreatedAtStr)
		if err != nil {
			go utils.Logger(ctx, r.Logger.Error, requestDesc, serviceDesc, requestID, response.SetError(err))
		}
		transaction.ProceededAt = parseProceed

		transaction.FromAccount.Amount.CurrencySymbol = merchant.GetSymCurrency(transaction.FromAccount.Amount.Currency)
		//	transaction.Id = transaction.IdGodror.String()

		trans = append(trans, *transaction)
	}
	return trans, errResponse

}

func (r repository) GetAccountHistoryTransactions(ctx context.Context, token, lang string,
	filters *merchant.FilterHistoryTransaction) (transactions []merchant.Transaction,
	errResponse response.ErrorResponse) {
	serviceDesc := "AccountHistoryTransactions-Repository"
	requestID := fmt.Sprintf("%v", ctx.Value(utils.CTXRequestID))

	var refCursor driver.Rows

	db := r.Oracle.GetConnection()

	go utils.Logger(ctx, r.Logger.Info, requestDesc, serviceDesc, requestID, "")
	_, err := db.ExecContext(ctx, getAccountTransactions,
		token,
		lang,
		godror.ArraySize(len(filters.Account)),
		godror.PlSQLArrays,
		filters.Account,
		filters.Pagination.Count,
		filters.Pagination.CurrentPage,
		filters.FromDate,
		filters.ToDate,
		sql.Out{Dest: &refCursor},
		sql.Out{Dest: &errResponse.ErrorCode},
		sql.Out{Dest: &errResponse.ErrorDesc},
	)

	if err != nil {
		errResponse = response.SetError(err)
		go utils.Logger(ctx, r.Logger.Error, requestDesc, serviceDesc, requestID, response.SetError(err))
		return nil, errResponse
	}

	if errResponse.ErrorCode != 0 {
		go utils.Logger(ctx, r.Logger.Error, requestDesc, serviceDesc, requestID, errResponse)
		return nil, errResponse
	}
	defer refCursor.Close()

	trans := make([]merchant.Transaction, 0, len(refCursor.Columns()))
	transaction := &merchant.Transaction{}

	rows := make([]driver.Value, len(refCursor.Columns()))
	for {

		if err := refCursor.Next(rows); err != nil {
			if err == io.EOF {
				break
			}
			continue
		}

		transaction.FromAccount.Number = rows[0].(string)
		transaction.ToAccount.Number = rows[1].(string)
		transaction.Id = rows[2].(string)
		transaction.CreatedAtStr = rows[3].(string)
		transaction.PaymentPurpose = rows[4].(string)
		transaction.Type = rows[5].(string)
		transaction.FromAccount.Title = rows[6].(string)
		transaction.ToAccount.Title = rows[7].(string)
		transaction.Status = rows[8].(string)
		transaction.FromAccount.Amount.Value = rows[9].(string)
		transaction.FromAccount.Amount.Currency = rows[10].(string)
		transaction.DocumentNumber = rows[12].(string)
		transaction.ProceededAtStr = rows[13].(string)
		parseCreate, err := time.Parse("02.01.2006 15:04:05", transaction.CreatedAtStr)
		if err != nil {
			go utils.Logger(ctx, r.Logger.Error, requestDesc, serviceDesc, requestID, response.SetError(err))
		}
		transaction.CreatedAt = parseCreate

		parseProceed, err := time.Parse("02.01.2006 15:04:05", transaction.CreatedAtStr)
		if err != nil {
			go utils.Logger(ctx, r.Logger.Error, requestDesc, serviceDesc, requestID, response.SetError(err))
		}
		transaction.ProceededAt = parseProceed

		trans = append(trans, *transaction)
		//	transaction.Id = transaction.IdGodror.String()
		transaction.FromAccount.Amount.CurrencySymbol = merchant.GetSymCurrency(transaction.FromAccount.Amount.Currency)
	}
	return trans, errResponse
}

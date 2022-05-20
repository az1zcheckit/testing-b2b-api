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

func (r repository) Conversion(ctx context.Context, token, lang string, convRequest merchant.ConversionRequest) (transaction merchant.ConversionResponse, errResponse response.ErrorResponse) {
	var refCursor driver.Rows
	var amountFromCur, amountToCur, rateCurr, rate string
	serviceDesc := "Conversion-Repository"
	requestID := fmt.Sprintf("%v", ctx.Value(utils.CTXRequestID))
	db := r.Oracle.GetConnection()
	go utils.Logger(ctx, r.Logger.Info, requestDesc, serviceDesc, requestID, "")

	_, err := db.ExecContext(ctx, conversionSQL,
		token,
		lang,
		convRequest.AccountFrom,
		convRequest.AccountTo,
		convRequest.Amount,
		convRequest.Footnote,
		convRequest.Rate,
		convRequest.DocNumber,
		convRequest.ExecutionDate,
		sql.Out{Dest: &transaction.Id},
		sql.Out{Dest: &transaction.DocumentNumber},
		sql.Out{Dest: &rate},
		sql.Out{Dest: &rateCurr},
		sql.Out{Dest: &transaction.Footnote},
		sql.Out{Dest: &transaction.CreateDate},
		sql.Out{Dest: &transaction.ExecutionDate},
		sql.Out{Dest: &transaction.AmountFrom.Value},
		sql.Out{Dest: &amountFromCur},
		sql.Out{Dest: &transaction.AmountTo.Value},
		sql.Out{Dest: &amountToCur},
		sql.Out{Dest: &transaction.AccountFrom.Name},
		sql.Out{Dest: &transaction.AccountFrom.Currency},
		sql.Out{Dest: &transaction.AccountTo.Name},
		sql.Out{Dest: &transaction.AccountTo.Currency},
		sql.Out{Dest: &refCursor},
		sql.Out{Dest: &errResponse.ErrorCode},
		sql.Out{Dest: &errResponse.ErrorDesc},
	)

	if err != nil {
		errResponse = response.DbExecContext
		go utils.Logger(ctx, r.Logger.Error, requestDesc, serviceDesc, requestID, response.SetError(err))
		return
	}

	if errResponse.ErrorCode != 0 {
		go utils.Logger(ctx, r.Logger.Error, requestDesc, serviceDesc, requestID, errResponse)
		return
	}

	defer refCursor.Close()
	transaction.AmountFrom.CurrencySymbol = merchant.GetSymCurrency(amountFromCur)
	transaction.AmountTo.CurrencySymbol = merchant.GetSymCurrency(amountToCur)
	transaction.Rate = fmt.Sprintf("%s c = 1 %s", rate, merchant.GetSymCurrency(rateCurr))
	signat := &merchant.Signature{}
	rows := make([]driver.Value, len(refCursor.Columns()))

	for {
		if err = refCursor.Next(rows); err != nil {
			if err == io.EOF {
				break
			}
			continue
		}
		signat.FullName = rows[0].(string)
		signat.Role = rows[1].(string)
		signat.Status = rows[2].(string)

		transaction.Signature = append(transaction.Signature, *signat)
	}

	go utils.Logger(ctx, r.Logger.Info, responseDesc, serviceDesc, requestID, transaction, errResponse)

	return
}

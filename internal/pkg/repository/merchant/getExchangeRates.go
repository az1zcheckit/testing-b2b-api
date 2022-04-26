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

func (r repository) GetExchangeRates(ctx context.Context, token, lang string) (base string, rates []merchant.Rates, errResponse response.ErrorResponse) {
	var refCursor driver.Rows
	serviceDesc := "Get-Exchange-Rates-Repository"
	requestID := fmt.Sprintf("%v", ctx.Value(utils.CTXRequestID))
	db := r.Oracle.GetConnection()

	go utils.Logger(ctx, r.Logger.Info, requestDesc, serviceDesc, requestID, "")

	_, err := db.ExecContext(ctx, getRatesSQL,
		token,
		lang,
		sql.Out{Dest: &base},
		sql.Out{Dest: &refCursor},
		sql.Out{Dest: &errResponse.ErrorCode},
		sql.Out{Dest: &errResponse.ErrorDesc},
	)

	if err != nil {
		go utils.Logger(ctx, r.Logger.Error, responseDesc, serviceDesc, requestID, err.Error())
		errResponse = response.SetError(err)
		return
	}

	if errResponse.ErrorCode != 0 {
		go utils.Logger(ctx, r.Logger.Error, responseDesc, serviceDesc, requestID, err.Error())
		return
	}
	defer refCursor.Close()

	rate := &merchant.Rates{}

	rows := make([]driver.Value, len(refCursor.Columns()))

	for {
		if err = refCursor.Next(rows); err != nil {
			if err == io.EOF {
				break
			}
			continue
		}
		rate.Currency = rows[0].(string)
		rate.BuyRate = rows[1].(string)
		rate.SellRate = rows[2].(string)

		rates = append(rates, *rate)
	}

	go utils.Logger(ctx, r.Logger.Info, responseDesc, serviceDesc, requestID, base, rates, errResponse)

	return
}

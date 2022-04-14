package merchant

import (
	response "b2b-api/internal/models"
	"b2b-api/internal/models/merchant"
	"database/sql"
	"database/sql/driver"
	"golang.org/x/net/context"
	"io"
)

func (r repository) GetExchangeRates(ctx context.Context, token, lang string) (base string, rates []merchant.Rates, errResponse response.ErrorResponse) {
	var refCursor driver.Rows

	db := r.Oracle.GetConnection()
	_, err := db.ExecContext(ctx, getRatesSQL,
		token,
		lang,
		sql.Out{Dest: &base},
		sql.Out{Dest: &refCursor},
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
	return
}

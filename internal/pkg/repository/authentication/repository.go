package authentication

import (
	response "b2b-api/internal/models"
	"b2b-api/internal/pkg/db/oracle"
	"go.uber.org/fx"
	"golang.org/x/net/context"
)

const sessionInsertSQL = "begin ibs.Z$ACCESS_TO_B2B_LIB_SERVICES.INSERTSESSION(:phone, :token, :lang, :error_code, :error_description); end;"

var NewAuthRepository = fx.Provide(newAuthRepository)

type IAuthRepository interface {
	SessionInsert(ctx context.Context, phone string, token string, lang string) (errResponse response.ErrorResponse)
}

type dependencies struct {
	fx.In
	Oracle oracle.Idb
}

type repository struct {
	Oracle oracle.Idb
}

func newAuthRepository(dp dependencies) IAuthRepository {
	return &repository{
		Oracle: dp.Oracle,
	}
}

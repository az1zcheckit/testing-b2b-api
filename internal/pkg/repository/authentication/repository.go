package authentication

import (
	response "b2b-api/internal/models"
	"b2b-api/internal/pkg/db/oracle"
	"b2b-api/internal/pkg/logger"
	"go.uber.org/fx"
	"golang.org/x/net/context"
)

const (
	requestDesc  = "request"
	responseDesc = "response"
)

const (
	sessionInsertSQL = "begin ibs.Z$ACCESS_TO_B2B_LIB_SERVICES.INSERTSESSION(:phone, :token, :lang, :error_code, :error_description); end;"
	getDateSQL       = "begin ibs.Z$B2B_LIB_SERVICES.GET_OPER_DATE(:OPERDATE, :errorCode, :errorDescription); end;"
	confirmTransaction = "begin ibs.Z$B2B_LIB_SERVICES.CONFIRMTRANSACTION(:TOKEN, :LANG, :IDDOC, :ERRORCODE, :ERRORDESCRIPTION); end;"
)

var NewAuthRepository = fx.Provide(newAuthRepository)

type IAuthRepository interface {
	SessionInsert(ctx context.Context, phone string, token string, lang string) (errResponse response.ErrorResponse)
	GetDate(ctx context.Context) (date string, errResponse response.ErrorResponse)
	ConfirmTransaction(ctx context.Context, token, lang, transactionId string) (errResponse response.ErrorResponse)
}

type dependencies struct {
	fx.In
	Oracle oracle.Idb
	Logger logger.ILogger
}

type repository struct {
	Oracle oracle.Idb
	Logger logger.ILogger
}

func newAuthRepository(dp dependencies) IAuthRepository {
	return &repository{
		Oracle: dp.Oracle,
		Logger: dp.Logger,
	}
}

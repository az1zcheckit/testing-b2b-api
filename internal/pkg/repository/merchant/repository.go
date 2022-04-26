package merchant

import (
	response "b2b-api/internal/models"
	"b2b-api/internal/models/merchant"
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
	getTransactionByIdSQL = "begin ibs.Z$B2B_TESTESTSE.GET_TRANSACTION_BY_ID(:token, :lang, :transId, :createdDate, :proceedDate, :accFrom, :accTo, :transType, :recipientName, :paymentPurpose, :amount, :errorCode, :errorDesc); end;"
	getAccountsSQL        = "begin ibs.Z$B2B_LIB_SERVICES.GET_MERCHANT_ACCOUNTS(:token, :lang, :MERCH_NAME, :ACCOUNTS, :errorCode, :errorDescription); end;"
	getTransactionsSQL    = "begin ibs.Z$B2B_LIB_SERVICES.GET_MERCHANT_TRANSACTIONS(:token, :lang, :c_From, :c_To, :MERCHTRANSACTIONS, :errorCode, :errorDescription); end;"
	getRatesSQL           = "begin ibs.Z$B2B_LIB_SERVICES.GET_EXCHANGE_RATES(:token, :lang, :BASE, :RATES, :errorCode, :errorDescription); end;"
	getAllAmountSQL       = "begin ibs.Z$B2B_LIB_SERVICES.GET_ALL_AMOUNT_MERCHANT(:token, :lang, :CURRENCY, :CURRENCYOUT, :BALANCE, :errorCode, :errorDescription); end;"
	conversionSQL         = "begin ibs.Z$B2B_LIB_SERVICES.CONVERSION(:token, :lang,:ACC_FROM, :ACC_TO, :AMOUNT, :Dest, :DOCUMENTID, :errorCode, :errorDescription); end;"
	getLimitConversionSQL = "begin ibs.Z$B2B_LIB_SERVICES.GET_CONVERSION_LIMIT(:token, :lang,:CURRENCY, :CURRENTLIMIT, :errorCode, :errorDescription); end;"
)

var NewRepositoryMerchant = fx.Provide(newRepositoryMerchant)

type IMerchantRepository interface {
	GetTransactionsByID(ctx context.Context, token, lang, transId string) (transById merchant.Transaction, errResponse response.ErrorResponse)
	GetTransactions(ctx context.Context, token, lang, dateFrom, dateTo string) (transactions []merchant.Transactions, errResponse response.ErrorResponse)
	GetAccounts(ctx context.Context, token, lang string) (merchName string, accounts []merchant.Accounts, errResponse response.ErrorResponse)
	GetExchangeRates(ctx context.Context, token, lang string) (base string, rates []merchant.Rates, errResponse response.ErrorResponse)
	GetAllAmount(ctx context.Context, token, lang, currency string) (allAmount merchant.GetAllAmount, errResponse response.ErrorResponse)
	Conversion(ctx context.Context, token, lang string, conv merchant.Conversion) (documentId int64, errResponse response.ErrorResponse)
	GetLimitConversion(ctx context.Context, token, lang, currency string) (currentLimit float64, errResponse response.ErrorResponse)
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

func newRepositoryMerchant(dp dependencies) IMerchantRepository {
	return &repository{
		Oracle: dp.Oracle,
		Logger: dp.Logger,
	}
}

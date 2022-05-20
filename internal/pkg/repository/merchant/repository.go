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
	getAccountsSQL         = "begin ibs.Z$B2B_LIB_SERVICES.GET_MERCHANT_ACCOUNTS(:token, :lang, :MERCH_NAME, :ACCOUNTS, :errorCode, :errorDescription); end;"
	getRatesSQL            = "begin ibs.Z$B2B_LIB_SERVICES.GET_EXCHANGE_RATES(:token, :lang, :BASE, :RATES, :errorCode, :errorDescription); end;"
	getAllAmountSQL        = "begin ibs.Z$B2B_LIB_SERVICES.GET_ALL_AMOUNT_MERCHANT(:token, :lang, :CURRENCY, :CURRENCYOUT, :BALANCE, :errorCode, :errorDescription); end;"
	conversionSQL          = "begin ibs.Z$B2B_LIB_SERVICES.CONVERSION(:TOKEN, :LANG, :ACC_FROM, :ACC_TO, :AMOUNT, :FOOTNOTE, :RATE, :DOCUMENTNUM, :EXECUTIONDATE, :createDateOut, :TRANID, :DOCNUMOUT, :RATEOUT, :RATECURRENCYOUT, :FOOTNOTEOUT, :EXECUTIONDATEOUT, :AMOUNTFROMVAL, :AMOUNTFROMCURSYM, :AMOUNTTOVAL, :AMOUNTTOCURSYM, :ACCFROMNAME, :ACCFROMCURALPHA, :ACCTONAME, :ACCTOCURALPHA, :SIGNATURES, :ERRORCODE, :ERRORDESCRIPTION); end;"
	getLimitConversionSQL  = "begin ibs.Z$B2B_LIB_SERVICES.GET_CONVERSION_LIMIT(:token, :lang,:CURRENCY, :CURRENTLIMIT, :errorCode, :errorDescription); end;"
	getAccountDetails      = "begin ibs.Z$B2B_LIB_SERVICES.GETACCDETAILS_FORBANKTRANS(:token,:lang,:accountNumber ,:ACCOUNTDETAILS ,:ERRORCODE ,:ERRORDESCRIPTION ); end;"
	getTransactionsSQL     = "begin ibs.Z$B2B_LIB_SERVICES.GET_MERCHANT_TRANSACTIONS(:token, :lang, :rowsPerPage,:page, :c_From, :c_To, :MERCHTRANSACTIONS, :errorCode, :errorDescription); end;"
	getAccountTransactions = "begin ibs.Z$B2B_LIB_SERVICES.GET_ACCOUNT_TRANSACTIONS(:token, :lang, :accounts, :rowsPerPage,:page, :c_From, :c_To, :MERCHTRANSACTIONS, :errorCode, :errorDescription); end;"
	getConvertedValueSQL   = "begin ibs.Z$B2B_LIB_SERVICES.GETCONVERTEDVALUE(:TOKEN, :LANG, :VALUE, :CURRENCYFROM, :CURRENCYTO, :VALUEOUT, :RATE, :ERRORCODE, :ERRORDESCRIPTION); end;"
)

var NewRepositoryMerchant = fx.Provide(newRepositoryMerchant)

type IMerchantRepository interface {
	GetAccounts(ctx context.Context, token, lang string) (merchName string, accounts []merchant.Accounts, errResponse response.ErrorResponse)
	GetExchangeRates(ctx context.Context, token, lang string) (base string, rates []merchant.Rates, errResponse response.ErrorResponse)
	GetAllAmount(ctx context.Context, token, lang, currency string) (allAmount merchant.GetAllAmount, errResponse response.ErrorResponse)
	Conversion(ctx context.Context, token, lang string, convRequest merchant.ConversionRequest) (transaction merchant.ConversionResponse, errResponse response.ErrorResponse)
	GetLimitConversion(ctx context.Context, token, lang, currency string) (currentLimit float64, errResponse response.ErrorResponse)
	GetAccountDetails(ctx context.Context, token, lang, accountNumber string) (accountsDetails []merchant.AccountDetailsForBankTrans, errResponse response.ErrorResponse)
	GetAllHistoryTransactions(ctx context.Context, token, lang string, filters *merchant.FilterHistoryTransaction) (transactions []merchant.Transaction, errResponse response.ErrorResponse)
	GetAccountHistoryTransactions(ctx context.Context, token, lang string, filters *merchant.FilterHistoryTransaction) (transactions []merchant.Transaction, errResponse response.ErrorResponse)
	GetConvertedValue(ctx context.Context, token, lang,value,currencyFrom,currencyTo string) (conerValue merchant.ConvertedValue, errResponse response.ErrorResponse)
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

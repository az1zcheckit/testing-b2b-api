package models

import "b2b-api/internal/models/merchant"

type ErrorModel struct {
	ErrorCode int    `json:"errorCode"`
	ErrorDesc string `json:"errorDescription"`
}

type Login struct {
	Phone    string `json:"phone" minLength:"9" maxLength:"9"`
	Password string `json:"password"`
}

type Token struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
	Role         string `json:"role"`
}

type OTP struct {
	IsActive   bool  `json:"isActive"`
	Expirition int64 `json:"expirition"`
	TotalSent  int   `json:"totalSent"`
}

type LoginResponse struct {
	Msg      ErrorModel `json:"msg"`
	Response twoFa      `json:"response"`
}

type twoFa struct {
	Gauth  googleAuth `json:"googleAuthenticator,omitempty"`
	SMSotp smsOTP     `json:"smsOtp,omitempty"`
}

type smsOTP struct {
	IsActive bool `json:"isActive"`
}

type googleAuth struct {
	IsActive bool `json:"isActive"`
}

type ConversionRequest struct {
	AccFrom string  `json:"accFrom"`
	AccTo   string  `json:"accTo"`
	Amount  float64 `json:"amount"`
	Dest    string  `json:"dest"`
}

type ConversionResponse struct {
	Msg   ErrorModel `json:"msg"`
	DocID string     `json:"docID"`
}

type GetAccountsResponse struct {
	Msg       ErrorModel        `json:"msg"`
	MerchName string            `json:"merchName"`
	Accounts  merchant.Accounts `json:"accounts"`
}

type GetAllAmountsResponse struct {
	MerchName string     `json:"merchName"`
	Balance   float64    `json:"balance"`
	Msg       ErrorModel `json:"msg"`
}

type GetExchangeRatesResponse struct {
	Base  string         `json:"base"`
	Rates merchant.Rates `json:"rates"`
	Msg   ErrorModel     `json:"msg"`
}

type GetTransactionResponse struct {
	MerchName string                `json:"merchName"`
	TrnList   merchant.Transactions `json:"trnList"`
	Msg       ErrorModel            `json:"msg"`
}

type GetTransactionByIdResponse struct {
	TransId string               `json:"transId"`
	TrnList merchant.Transaction `json:"trnList"`
	Msg     ErrorModel           `json:"msg"`
}

type GetLimitConversion struct {
	CurrentLimit float64    `json:"currentLimit"`
	Msg          ErrorModel `json:"msg"`
}

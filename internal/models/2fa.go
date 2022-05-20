package models

type TwoFactorAuthentication struct {
	TwoFA TwoFA `json:"2FA,omitempty"`
}

type TwoFA struct {
	Gauth  GoogleAuth `json:"googleAuthenticator,omitempty"`
	SMSotp SMSotp     `json:"smsOtp,omitempty"`
}

type SMSotp struct {
	IsActive   bool   `json:"isActive" default:"false"`
	Code       string `json:"code,omitempty"`
	Expirition int64  `json:"expirition,omitempty"`
	TotalSent  int    `json:"totalSent,omitempty"`
}

type GoogleAuth struct {
	IsActive            bool   `json:"isActive" default:"false"`
	GauthSecretEndpoint string `json:"gAuthSecret,omitempty"`
}

func (o TwoFA) Implementation()                     {}
func (tfa TwoFactorAuthentication) Implementation() {}

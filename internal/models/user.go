package models

import "github.com/pkg/errors"

type User struct {
	Phone        string                   `json:"phone,omitempty"`
	Password     string                   `json:"password,omitempty"`
	Role         string                   `json:"role,omitempty"`
	Token        string                   `json:"token,omitempty"`
	RefreshToken string                   `json:"refreshToken,omitempty"`
	Security     *TwoFactorAuthentication `json:"2FA,omitempty"`
	TransactionId string                  `json:"transactionId,omitempty"`
}

func (u *User) LoginValidate() error {

	if len(u.Phone) < 9 {
		return PhoneIncorrect
	} else if len(u.Password) == 0 {
		return OTPIncorrect
	}
	return nil
}

func (u *User) OTPValidate() error {
	if len(u.Security.TwoFA.SMSotp.Code) < 4 {
		return errors.New(OTPIncorrect.Error())
	} else if len(u.Token) == 0 {
		return errors.New(TokenIsEmpty.Error())
	}
	return nil
}

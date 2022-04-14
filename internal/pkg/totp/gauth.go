package totp

import (
	gAuth "github.com/dgryski/dgoogauth"
)

func (t totp) GAuthenticate(secretCode string, totp string) (IsValid bool, err error) {

	otp := gAuth.OTPConfig{
		Secret:      secretCode,
		WindowSize:  0,
		HotpCounter: 0,
	}

	IsValid, err = otp.Authenticate(totp)
	if err != nil {
		return
	}
	return
}

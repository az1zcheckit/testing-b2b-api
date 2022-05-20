package totp

import (
	"b2b-api/internal/config"
	"go.uber.org/fx"
)

var NewTOTP = fx.Provide(newTOTP)

type Iotp interface {
	GAuthenticate(secretCode string, totp string) (IsValid bool, err error)
	SmsOTP(phone string, msg string) (string, error)
	GauthGenerateSecretEndpoint() string
}

type dependencies struct {
	fx.In
	CFG config.IConfig
}

type totp struct {
	CFG config.IConfig
}

func newTOTP(d dependencies) Iotp {
	return &totp{CFG: d.CFG}
}

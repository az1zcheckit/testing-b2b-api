package totp

import (
	"crypto/rand"
	"encoding/base32"
)

func (t totp) GauthGenerateSecretEndpoint() string {
	random := make([]byte, 10)
	rand.Read(random)
	secret := base32.StdEncoding.EncodeToString(random)
	return secret
}

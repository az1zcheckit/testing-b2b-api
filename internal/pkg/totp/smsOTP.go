package totp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func (t totp) SmsOTP(phone string, msg string) (string, error) {

	resp, err := http.Get(fmt.Sprintf(t.CFG.GetString("api.smsc.server"), phone, strings.ReplaceAll(msg, " ", "%20")))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

package models

import (
	"encoding/json"
	"net/http"
)

const (
	// error messages
	undefinedError    = "Undefined error"
	tokenExpired      = "Token expired"
	malformedToken    = "Malformed token or token expired"
	accessDenied      = "Access denied"
	tokenIsEmpty      = "Token is empty"
	roleIsEmpty       = "You haven't got a role to login"
	tooManyRoles      = "Too many roles"
	phoneIncorrect    = "phone number is not correct"
	otpIncorrect      = "OTP is not correct"
	otpSent           = "OTP has already sent. Confirm it"
	otpCountExhausted = "A count of sms for this session has been exhausted"
	requestIdIsEmpty  = "Request-ID is empty"
	gAuthIsNotActive  = "Google authenticator is not active"
	incorrectLang     = "Incorrect language"
	// messages without error
	waitingForOTP = "Login successful. Waiting for OTP..."

	waitingForConfirm = "Success. Waiting for Confirm..."
)

var (
	// error messages
	TokenExpired      = getError(tokenExpired, -1)
	MalformedToken    = getError(malformedToken, -2)
	AccessDenied      = getError(accessDenied, -3)
	TokenIsEmpty      = getError(tokenIsEmpty, -4)
	RoleIsEmpty       = getError(roleIsEmpty, -5)
	TooManyRoles      = getError(tooManyRoles, -6)
	UndefinedError    = getError(undefinedError, -7)
	PhoneIncorrect    = getError(phoneIncorrect, -8)
	OTPIncorrect      = getError(otpIncorrect, -9)
	RequestIdIsEmpty  = getError(requestIdIsEmpty, -10)
	OTPsent           = getError(otpSent, -11)
	OTPCountExhausted = getError(otpCountExhausted, -12)
	GauthIsNotActive  = getError(gAuthIsNotActive, -13)
	IncorrectLang     = getError(incorrectLang, -14)

	// messages without error
	WaitingForOTP     = getMsg(waitingForOTP, 9)
	WaitingForConfirm = getError(waitingForConfirm, 10)
)

type ErrorResponse struct {
	ErrorCode     int    `json:"errorCode"`
	ErrorDesc     string `json:"errorDescription"`
	AditionalInfo Imodel `json:"response,omitempty"`
}

func (r ErrorResponse) Error() string {
	return r.toJson()
}

func (r ErrorResponse) toJson() string {
	reply, err := json.Marshal(r)
	if err != nil {
		return "{error:\"Marshal error\"}"
	}
	return string(reply)
}

func ToJson(w http.ResponseWriter, status int, data interface{}) {

	reply, err := json.Marshal(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	w.Write(reply)
}

func Json(w http.ResponseWriter, status int, data string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	w.Write([]byte(data))
}

func getError(err string, code int) ErrorResponse {
	return ErrorResponse{
		ErrorCode: code,
		ErrorDesc: err,
	}
}

func SetError(err error) ErrorResponse {
	return ErrorResponse{
		ErrorDesc: err.Error(),
	}
}

func getMsg(err string, code int) ErrorResponse {
	return ErrorResponse{
		ErrorCode: code,
		ErrorDesc: err,
	}
}

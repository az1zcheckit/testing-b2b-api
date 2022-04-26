package utils

import (
	"context"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"io"
)

type jwtType int8

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

const (
	JWTkey jwtType = iota
	CTXRequestID
	JWTSault string = "this-is-jwt-sault"
)

func Logger(ctx context.Context, log func(format string, v ...interface{}), method string, service string, requestID string, v ...interface{}) (info string) {
	if ctx.Err() != context.Canceled {
		logTxt := fmt.Sprintf("%s -> %s ::: %s --> %v", method, service, requestID, v)
		log("%s", logTxt)
	}
	return info
}

func GenerateSession(token string) (session string) {
	data := []byte(token + JWTSault)
	hashByte := md5.Sum(data)
	session = string(hashByte[:])
	return
}

func GenerateOTP(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

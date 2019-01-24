package errors

import (
	"net/http"
)

const (
	InvalidId = "invalid id"
	NotFound  = "not found"

	//JWT errors
	CannotFindPrivateKey   = "cannot find private key file"
	CannotFindPublicKey    = "cannot find public key file"
	TokenNotExists         = "token does not exist"
	IncorrectSigningMethod = "unexpected signing method"
	NotToken               = "not a token"
	TimeExpired            = "time expired"
	InvalidToken           = "invalid token"
	CantHandleToken        = "couldn't handle this token"

	//Field errors
	EmailExists = "email already exists"
)

var statusCode = map[string]int{
	InvalidId:              http.StatusBadRequest,
	NotFound:               http.StatusNotFound,
	CannotFindPrivateKey:   http.StatusUnauthorized,
	CannotFindPublicKey:    http.StatusUnauthorized,
	TokenNotExists:         http.StatusUnauthorized,
	IncorrectSigningMethod: http.StatusUnauthorized,
	NotToken:               http.StatusUnauthorized,
	TimeExpired:            http.StatusUnauthorized,
	InvalidToken:           http.StatusUnauthorized,
	CantHandleToken:        http.StatusUnauthorized,
	EmailExists:            http.StatusBadRequest,
}

func StatusCode(error string) int {
	if val, ok := statusCode[error]; ok {
		return val
	} else {
		return http.StatusInternalServerError
	}

}

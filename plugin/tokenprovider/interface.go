package tokenprovider

import (
	"errors"
	"nckh-BE/appCommon"
	"net/http"
	"time"
)

type Provider interface {
	Generate(data TokenPayload, expiry int) (*Token, error)
	Validate(token string) (*TokenPayload, error)
}

var (
	ErrNotFound = appCommon.NewCustomError(
		http.StatusNotFound,
		errors.New("token not found"),
		"token not found",
		"ErrNotFound",
	)

	ErrEncodingToken = appCommon.NewCustomError(
		http.StatusInternalServerError,
		errors.New("error encoding the token"),
		"error encoding the token",
		"ErrEncodingToken",
	)

	ErrInvalidToken = appCommon.NewCustomError(
		http.StatusBadRequest,
		errors.New("invalid token provided"),
		"invalid token provided",
		"ErrInvalidToken",
	)
)

type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

type TokenPayload struct {
	UserId    string `json:"user_id"`
	Role      string `json:"role"`
	SessionID string `json:"session_id"`
}

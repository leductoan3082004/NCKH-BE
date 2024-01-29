package appCommon

import (
	"errors"
	"fmt"
	"github.com/lequocbinh04/go-sdk/sdkcm"
	"net/http"
	"strings"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

func NewErrorResponse(statusCode int, root error, msg, log, key string) *sdkcm.AppError {
	return &sdkcm.AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewFullErrorResponse(statusCode int, root error, msg, log, key string) *sdkcm.AppError {
	return &sdkcm.AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewUnauthorized(root error, msg, key string) *sdkcm.AppError {
	return &sdkcm.AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		Key:        key,
	}
}

func NewCustomError(statusCode int, root error, msg string, key string) *sdkcm.AppError {
	if root != nil {
		return NewErrorResponse(statusCode, root, msg, root.Error(), key)
	}

	return NewErrorResponse(statusCode, errors.New(msg), msg, msg, key)
}

//api (e *sdkcm.AppError) RootError() error {
//	if err, ok := e.RootErr.(*sdkcm.AppError); ok {
//		return err.RootError()
//	}
//
//	return e.RootErr
//}

//api (e *sdkcm.AppError) Error() string {
//	return e.RootError().Error()
//}

func ErrDB(err error) *sdkcm.AppError {
	return NewErrorResponse(http.StatusInternalServerError, err, "something went wrong with DB", err.Error(), "DB_ERROR")
}

func ErrInvalidRequest(err error) *sdkcm.AppError {
	return NewErrorResponse(http.StatusBadRequest, err, "invalid request", err.Error(), "ErrInvalidRequest")
}

func ErrInternal(err error) *sdkcm.AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err,
		"something went wrong in the server", err.Error(), "ErrInternal")
}

func ErrCannotListEntity(entity string, err error) *sdkcm.AppError {
	return NewCustomError(
		http.StatusInternalServerError,
		err,
		fmt.Sprintf("Cannot list %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotList%s", entity),
	)
}

func ErrCannotDeleteEntity(entity string, err error) *sdkcm.AppError {
	return NewCustomError(
		http.StatusInternalServerError,
		err,
		fmt.Sprintf("Cannot delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotDelete%s", entity),
	)
}

func ErrCannotUpdateEntity(entity string, err error) *sdkcm.AppError {
	return NewCustomError(
		http.StatusInternalServerError,
		err,
		fmt.Sprintf("Cannot gin %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotUpdate%s", entity),
	)
}

func ErrCannotGetEntity(entity string, err error) *sdkcm.AppError {
	return NewCustomError(
		http.StatusInternalServerError,
		err,
		fmt.Sprintf("Cannot get %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotGet%s", entity),
	)
}

func ErrEntityDeleted(entity string, err error) *sdkcm.AppError {
	return NewCustomError(
		http.StatusInternalServerError,
		err,
		fmt.Sprintf("%s deleted", strings.ToLower(entity)),
		fmt.Sprintf("Err%sDeleted", entity),
	)
}

func ErrEntityExisted(entity string, err error) *sdkcm.AppError {
	return NewCustomError(
		http.StatusBadRequest,
		err,
		fmt.Sprintf("%s already exists", strings.ToLower(entity)),
		fmt.Sprintf("Err%sAlreadyExists", entity),
	)
}

func ErrEntityNotFound(entity string, err error) *sdkcm.AppError {
	return NewCustomError(
		http.StatusBadRequest,
		err,
		fmt.Sprintf("%s not found", strings.ToLower(entity)),
		fmt.Sprintf("Err%sNotFound", entity),
	)
}

func ErrCannotCreateEntity(entity string, err error) *sdkcm.AppError {
	return NewCustomError(
		http.StatusInternalServerError,
		err,
		fmt.Sprintf("Cannot Create %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotCreate%s", entity),
	)
}

func ErrNoPermission(err error) *sdkcm.AppError {
	return NewCustomError(
		http.StatusForbidden,
		err,
		fmt.Sprintf("You have no permission"),
		fmt.Sprintf("ErrNoPermission"),
	)
}

var ErrRecordNotFound = errors.New("record not found")
var ErrTooManyRequests = NewCustomError(
	http.StatusTooManyRequests,
	nil,
	"Too many requests",
	"ErrTooManyRequests")

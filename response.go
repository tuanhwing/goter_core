package gotercore

import "strings"

const (
	FailedToProcessRequestErrorCode = 1
	SomethingWentWrongErrorCode     = 2
	TokenNotFoundErrorCode          = 3
	UnauthorizedErrorCode           = 4
)

type ErrorReponse struct {
	ErrorCode int         `json:"error_code"`
	Errors    interface{} `json:"errors"`
}

type Response[T any] struct {
	Status  bool         `json:"status"`
	Error   ErrorReponse `json:"error"`
	Message string       `json:"message"`
	Data    T            `json:"data"`
}

type EmptyObj struct {
}

func BuildResponse[T any](status bool, message string, data T) Response[T] {
	res := Response[T]{
		Status:  status,
		Message: message,
		Data:    data,
	}
	return res
}

func BuildErrorResponse[T any](code int, message, errMessages string, data T) Response[T] {
	splittedError := strings.Split(errMessages, "\n")
	err := ErrorReponse{
		ErrorCode: code,
		Errors:    splittedError,
	}

	res := Response[T]{
		Status:  false,
		Error:   err,
		Message: message,
		Data:    data,
	}
	return res
}

var NotFoundTokenResponse = BuildErrorResponse[interface{}](
	TokenNotFoundErrorCode,
	"Token not found",
	"Failed to proccess request",
	nil)

var UnauthorizedResponse = BuildErrorResponse[interface{}](
	UnauthorizedErrorCode,
	"Token invalid or expired",
	"Failed to proccess request",
	nil)

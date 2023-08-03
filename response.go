package gotercore

import "strings"

const (
	FailedToProcessRequestErrorCode = 1
	SomethingWentWrongErrorCode     = 2
	TokenNotFoundErrorCode          = 3
	UnauthorizedErrorCode           = 4
)

type Response struct {
	Status  bool        `json:"status"`
	Error   interface{} `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct {
}

func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
	return res
}

func BuildErrorResponse(code int, message, errMessages string, data interface{}) Response {
	splittedError := strings.Split(errMessages, "\n")
	err := struct {
		ErrorCode int         `json:"error_code"`
		Errors    interface{} `json:"errors"`
	}{
		ErrorCode: code,
		Errors:    splittedError,
	}

	res := Response{
		Status:  false,
		Error:   err,
		Message: message,
		Data:    data,
	}
	return res
}

var NotFoundTokenResponse = BuildErrorResponse(
	TokenNotFoundErrorCode,
	"Token not found",
	"Failed to proccess request",
	nil)

var UnauthorizedResponse = BuildErrorResponse(
	UnauthorizedErrorCode,
	"Token invalid or expired",
	"Failed to proccess request",
	nil)

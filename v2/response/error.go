package response

import "fmt"

// Code response code
type Code int

// Error error
type Error struct {
	Code   Code   `json:"code"`
	Msg    string `json:"msg"`
	Detail string `json:"detail"`
}

func (e Error) Error() string {
	var msg = e.Msg
	if len(msg) < 1 {
		msg = "unknown error"
	}
	if len(e.Detail) < 1 {
		return msg
	}
	return fmt.Sprintf("%s|%s", msg, e.Detail)
}

// Code
const (
	CodeOK Code = 0

	CodeUnknown      Code = 10000
	CodeDBConnect    Code = 10001
	CodeDBOperation  Code = 10002
	CodeRequestParam Code = 10003

	CodeTokenInvalid     Code = 20101
	CodePermissionDenied Code = 20102
)

// System error
var (
	UnknownError = func(detail string) Error {
		return Error{Code: CodeUnknown, Msg: "unknown error", Detail: detail}
	}

	DBConnectError = func(detail string) Error {
		return Error{Code: CodeDBConnect, Msg: "database connection error", Detail: detail}
	}

	DBOperationError = func(detail string) Error {
		return Error{Code: CodeDBOperation, Msg: "database operation error", Detail: detail}
	}

	RequestParamError = func(detail string) Error {
		return Error{Code: CodeRequestParam, Msg: "request param error", Detail: detail}
	}

	TokenInvalidError = func(detail string) Error {
		return Error{Code: CodeTokenInvalid, Msg: "token invalid error", Detail: detail}
	}

	PermissionDeniedError = func(detail string) Error {
		return Error{Code: CodePermissionDenied, Msg: "permission denied error", Detail: detail}
	}
)

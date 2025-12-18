package http_response

import (
	"net/http"
	_ "net/http"
)

const (
	ResponseOK                   = "OK"
	ErrLoginPasswordIncorrect    = "A0001"
	ResponseWarning              = "WARNING"
	ErrInternalServerError       = "ISE"
	ErrUserTokenInvalidOrExpired = "TOKERR"
	ErrCannotGenerateToken       = "TOKCG"
	ErrPermissionDenied          = "PERMISSION DENIED"
	ErrNotFillAllFields          = "DEV0001"
	ErrDataFormatWrong           = "DEV0002"
	ErrStrconv                   = "DEV0004"
	ErrUserAlreadyExists         = "U0002"
	ErrUserNotFound              = "U0003"
	ErrStartAfterEnd             = "E0001"
)

type ResponseInfo struct {
	Message    string
	HttpStatus int
}

var ResponseMap = map[string]ResponseInfo{
	ResponseOK: {
		Message:    "SUS0000",
		HttpStatus: http.StatusOK,
	},
	ResponseWarning: {
		Message:    "SUS0000",
		HttpStatus: http.StatusOK,
	},
	ErrInternalServerError: {
		Message:    "Internal Server Error",
		HttpStatus: http.StatusInternalServerError,
	},
	ErrUserTokenInvalidOrExpired: {
		Message:    "Token is invalid or already expired",
		HttpStatus: http.StatusUnauthorized,
	},
	ErrUserAlreadyExists: {
		Message:    "Username already exists",
		HttpStatus: http.StatusBadRequest,
	},
	ErrNotFillAllFields: {
		Message:    "Not fill all fields",
		HttpStatus: http.StatusBadRequest,
	},
	ErrUserNotFound: {
		Message:    "User not found",
		HttpStatus: http.StatusNotFound,
	},

	ErrDataFormatWrong: {
		Message:    "Data format wrong",
		HttpStatus: http.StatusBadRequest,
	},

	ErrPermissionDenied: {
		Message:    "Permission denied",
		HttpStatus: http.StatusForbidden,
	},
	ErrStrconv: {
		Message:    "fail to convert string to int",
		HttpStatus: http.StatusBadRequest,
	},
	ErrLoginPasswordIncorrect: {
		Message:    "login password incorrect",
		HttpStatus: http.StatusUnauthorized,
	},
	ErrCannotGenerateToken: {
		Message:    "Error in generating user token",
		HttpStatus: http.StatusInternalServerError,
	},
	ErrStartAfterEnd: {
		Message:    "Error in start after end",
		HttpStatus: http.StatusBadRequest,
	},
}

func GetMessage(code string) string {
	if val, ok := ResponseMap[code]; ok {
		return val.Message
	} else {
		return ""
	}
}

func GetHttpResponseStatus(code string) int {
	if val, ok := ResponseMap[code]; ok {
		return val.HttpStatus
	} else {
		return 555
	}
}

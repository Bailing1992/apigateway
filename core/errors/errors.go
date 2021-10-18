package errors

import (
	"fmt"
)

type ErrorType uint8

const (
	UnknownErrorType ErrorType = 1 + iota
	InternalErrorType
	DependenceErrorType
	RequestErrorType
	ProxyErrorType
	BackendErrorType
	AsyncErrorType
)

const (
	SuccessErrorCode int = 0 - iota
	UnknownErrorCode
	InternalErrorCode
	DependenceErrorCode
	RequestErrorCode
	ProxyErrorCode
	BackendErrorCode
	AsyncErrorCode
)

var (
	ErrorTypeMap = map[ErrorType]string{
		InternalErrorType:   "INTERNAL_ERROR",
		DependenceErrorType: "DEPENDENCE_ERROR",
		RequestErrorType:    "REQUEST_ERROR",
		ProxyErrorType:      "PROXY_ERROR",
		BackendErrorType:    "BACKEND_ERROR",
		AsyncErrorType:      "ASYNC_ERROR",
	}
)

type Error struct {
	errorType ErrorType
	code      int
	message   string
}

func (e Error) New(errorType ErrorType, errorCode int, errorMessage string) Error {
	return Error{
		errorType: errorType,
		code:      errorCode,
		message:   errorMessage,
	}
}

func NewInternalError(format string, v ...interface{}) Error {
	errorMessage := fmt.Sprintf(format, v...)
	return Error{
		errorType: InternalErrorType,
		code:      InternalErrorCode,
		message:   errorMessage,
	}
}

func NewCoreError(format string, v ...interface{}) Error {
	errorMessage := fmt.Sprintf(format, v...)
	return Error{
		errorType: InternalErrorType,
		code:      InternalErrorCode,
		message:   errorMessage,
	}
}

func NewRequestError(format string, v ...interface{}) Error {
	errorMessage := fmt.Sprintf(format, v...)
	return Error{
		errorType: RequestErrorType,
		code:      RequestErrorCode,
		message:   errorMessage,
	}
}

func NewProxyError(format string, v ...interface{}) Error {
	errorMessage := fmt.Sprintf(format, v...)
	return Error{
		errorType: ProxyErrorType,
		code:      ProxyErrorCode,
		message:   errorMessage,
	}
}

func NewDependenceError(format string, v ...interface{}) Error {
	errorMessage := fmt.Sprintf(format, v...)
	return Error{
		errorType: DependenceErrorType,
		code:      DependenceErrorCode,
		message:   errorMessage,
	}
}

func NewFromErr(err error, errorType ErrorType, errorCode int) Error {
	errorMessage := "Unknown error info"
	if err != nil {
		errorMessage = err.Error()
	}

	return Error{
		errorType: errorType,
		code:      errorCode,
		message:   errorMessage,
	}
}

func NewAsyncError(format string, v ...interface{}) Error {
	errorMessage := fmt.Sprintf(format, v...)
	return Error{
		errorType: AsyncErrorType,
		code:      AsyncErrorCode,
		message:   errorMessage,
	}
}

func (e Error) String() string {
	errorString := fmt.Sprintf("ErrorType: %s, ErrorCode: %d, ErrorMessage: %s", ErrorTypeMap[e.errorType], e.code, e.message)
	return errorString
}

func (e Error) Error() string {
	return e.String()
}

package errors

import "errors"

type ErrorCode int

const (
	ErrCodeNotFound ErrorCode = iota + 1
	ErrCodeBadRequest
	ErrCodeUnauthorized
	ErrCodeForbidden
	ErrCodeInternal
)

type AppError struct {
	Message string    `json:"message"`
	Code    int       `json:"code"` // HTTP status
	ErrCode ErrorCode `json:"-"`    // internal code
	Err     error     `json:"-"`    // optional root cause
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(code int, errCode ErrorCode, msg string, err error) error {
	return &AppError{
		Message: msg,
		Code:    code,
		ErrCode: errCode,
		Err:     err,
	}
}

func NewNotFoundError(msg string) error {
	return NewAppError(404, ErrCodeNotFound, msg, nil)
}

func NewBadRequestError(msg string) error {
	return NewAppError(400, ErrCodeBadRequest, msg, nil)
}

func NewUnauthorizedError(msg string) error {
	return NewAppError(401, ErrCodeUnauthorized, msg, nil)
}

func NewForbiddenError(msg string) error {
	return NewAppError(403, ErrCodeForbidden, msg, nil)
}

func NewInternalError(msg string) error {
	return NewAppError(500, ErrCodeInternal, msg, nil)
}

func IsErrorCode(err error, code ErrorCode) bool {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr.ErrCode == code
	}
	return false
}

func IsNotFoundError(err error) bool {
	return IsErrorCode(err, ErrCodeNotFound)
}

func IsBadRequestError(err error) bool {
	return IsErrorCode(err, ErrCodeBadRequest)
}

func IsUnauthorizedError(err error) bool {
	return IsErrorCode(err, ErrCodeUnauthorized)
}

func IsForbiddenError(err error) bool {
	return IsErrorCode(err, ErrCodeForbidden)
}

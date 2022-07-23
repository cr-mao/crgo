package bizerror

import (
	"fmt"

	"github.com/pkg/errors"
)

var ErrBiz = errors.New("biz")

// BizError 代表业务上捕捉到的错误/异常。
type BizError struct {
	Code    int    // ErrorResponse.ErrCode, return to gRPC/HTTP Client
	Msg     string // ErrorResponse.ErrMsg,  return to gRPC/HTTP Client
	wrapped error  // Underlying error, from 3rd API/Library or errors.Wrap/New/WithMessage
}

func (e *BizError) Error() string {
	return fmt.Sprintf("code: %d, msg: %s, error: %s", e.Code, e.Msg, e.wrapped)
}

func (e *BizError) Unwrap() error {
	return e.wrapped
}

func (e *BizError) Is(target error) bool {
	return target == ErrBiz
}

var NewError = Newf

// New construct BizError with code and msg
func New(code int, msg string) error {
	return &BizError{Code: code, Msg: msg}
}

// Newf construct BizError with code, msg and extra message
func Newf(code int, msg string, format string, args ...interface{}) error {
	return &BizError{Code: code, Msg: msg, wrapped: errors.Errorf(format, args...)}
}

// Wrap construct BizError with code, msg and underlying error
func Wrap(code int, msg string, err error) error {
	return &BizError{Code: code, Msg: msg, wrapped: err}
}

// Wrapf construct BizError with code, msg, underlying error and extra message
func Wrapf(code int, msg string, err error, format string, args ...interface{}) error {
	return &BizError{Code: code, Msg: msg, wrapped: errors.Wrapf(err, format, args...)}
}

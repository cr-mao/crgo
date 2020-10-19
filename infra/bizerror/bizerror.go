package bizerror

import (
	"fmt"
)

// BizError ...
type BizError struct {
	Code   int
	Msg    string
	Detail string
}

func (e *BizError) Error() string {
	return fmt.Sprintf("code = %d ; msg = %s", e.Code, e.Msg)
}

// NewError 声明一个错误
func NewError(code int, msg string, detail string) *BizError {
	return &BizError{Code: code, Msg: msg, Detail: detail}
}

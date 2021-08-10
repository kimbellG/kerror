package kerror

import (
	"errors"
	"fmt"
)

type Error struct {
	err  error
	code StatusCode
}

func New(err error, code StatusCode) error {
	return Error{
		err:  err,
		code: code,
	}
}

func Newf(code StatusCode, format string, args ...interface{}) error {
	return New(fmt.Errorf(format, args...), code)
}

func (te Error) Error() string {
	return te.err.Error()
}

func (te Error) StatusCode() StatusCode {
	return te.code
}

func Errorf(err error, format string, a ...interface{}) error {
	a = append(a, err)
	newErr := fmt.Errorf(format+": %v", a...)

	if trnt := (Error{}); errors.As(err, &trnt) {
		trnt.err = newErr
		return trnt
	}

	return newErr
}

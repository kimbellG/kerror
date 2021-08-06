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

func (te Error) Error() string {
	return te.err.Error()
}

func (te Error) StatusCode() StatusCode {
	return te.code
}

func (te Error) StatusMessage() string {
	return statusMessage[te.code]
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

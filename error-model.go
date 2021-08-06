package trnmnterr

import (
	"errors"
	"fmt"
)

type TrnmntError struct {
	err  error
	code StatusCode
}

func New(err error, code StatusCode) error {
	return TrnmntError{
		err:  err,
		code: code,
	}
}

func (te TrnmntError) Error() string {
	return te.err.Error()
}

func (te TrnmntError) StatusCode() StatusCode {
	return te.code
}

func (te TrnmntError) StatusMessage() string {
	return statusMessage[te.code]
}

func Errorf(err error, format string, a ...interface{}) error {
	a = append(a, err)
	newErr := fmt.Errorf(format+": %v", a...)

	if trnt := (TrnmntError{}); errors.As(err, &trnt) {
		trnt.err = newErr
		return trnt
	}

	return newErr
}

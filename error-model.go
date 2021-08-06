package trnmnterr

import (
	"errors"
	"fmt"
)

type trnmntError struct {
	err  error
	code StatusCode
}

func New(err error, code StatusCode) trnmntError {
	return trnmntError{
		err:  err,
		code: code,
	}
}

func (te trnmntError) Error() string {
	return te.err.Error()
}

func (te trnmntError) StatusCode() StatusCode {
	return te.code
}

func (te trnmntError) StatusMessage() string {
	switch te.code {
	case InvalidId:
		return "invalid element's id"
	case NotFound:
		return "element not found"
	}

	panic(fmt.Sprintf("incorrect error status code: %d, err: %v", te.code, te.err))
}

func Errorf(err error, format string, a ...interface{}) error {
	a = append(a, err)
	newErr := fmt.Errorf(format+": %v", a...)

	if trnt := (trnmntError{}); errors.As(err, &trnt) {
		trnt.err = newErr
		return trnt
	}

	return newErr
}

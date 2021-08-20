package kerror

import (
	"errors"

	log "github.com/sirupsen/logrus"
)

func ErrorLog(entry *log.Entry, err error, msg string) {
	errorEntry := entry.WithField("status", "error")

	var code StatusCode
	if trnt := (Error{}); errors.As(err, &trnt) {
		code = trnt.StatusCode()
	} else {
		code = Unknown
	}

	errorEntry.WithFields(log.Fields{
		"code": code.Message(),
		"msg":  err.Error(),
	}).Debug(msg)
}

func ErrorLogf(entry *log.Entry, err error, format string, args ...interface{}) {
	errorEntry := entry.WithField("status", "error")

	var code StatusCode
	if trnt := (Error{}); errors.As(err, &trnt) {
		code = trnt.StatusCode()
	} else {
		code = Unknown
	}

	errorEntry.WithFields(log.Fields{
		"code": code.Message(),
		"msg":  err.Error(),
	}).Debugf(format, args...)
}

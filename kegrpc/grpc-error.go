package kegrpc

import (
	"errors"

	"github.com/kimbellG/kerror"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var keToGrpcDict = map[kerror.StatusCode]codes.Code{
	kerror.InvalidID:                   codes.InvalidArgument,
	kerror.BadRequest:                  codes.InvalidArgument,
	kerror.NotFound:                    codes.NotFound,
	kerror.TournamentDoesntExists:      codes.NotFound,
	kerror.UserDoesntExists:            codes.NotFound,
	kerror.SQLConstraintError:          codes.FailedPrecondition,
	kerror.SQLQueryError:               codes.Internal,
	kerror.SQLPrepareStatementError:    codes.Internal,
	kerror.SQLScanError:                codes.Internal,
	kerror.SQLExecutionError:           codes.Internal,
	kerror.SQLTransactionError:         codes.Aborted,
	kerror.SQLTransactionBeginError:    codes.Aborted,
	kerror.SQLTransactionRoolbackError: codes.Aborted,
	kerror.SQLTransactionCommitError:   codes.Aborted,
}

func Newf(code kerror.StatusCode, format string, args ...interface{}) error {
	return status.Errorf(MarshalStatusCode(code), format, args...)
}

func Errorf(err error, format string, args ...interface{}) error {
	args = append(args, err)

	if trnt := (kerror.Error{}); errors.As(err, &trnt) {
		return Newf(trnt.StatusCode(), format+": %v", args...)
	}

	return status.Errorf(codes.Unknown, format+": %v", args...)
}

func MarshalStatusCode(code kerror.StatusCode) codes.Code {
	return grpcDict[code]
}

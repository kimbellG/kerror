package kerror

type StatusCode int

const (
	InvalidID StatusCode = iota
	NotFound
	BadRequest
	InternalServerError

	TournamentDoesntExists
	UserDoesntExists

	SQLConstraintError

	SQLQueryError
	SQLPrepareStatementError
	SQLScanError
	SQLExecutionError

	SQLTransactionError
	SQLTransactionBeginError
	SQLTransactionRoolbackError
	SQLTransactionCommitError

	Unknown
)

var MessageForCode = map[StatusCode]string{
	InvalidID:           "InvalidID",
	NotFound:            "Not Found",
	BadRequest:          "BadRequest",
	InternalServerError: "InternalServerError",

	TournamentDoesntExists: "TouranamentDoesntExists",
	UserDoesntExists:       "UserDoesntExists",

	SQLConstraintError: "SQLConstraintError",

	SQLQueryError:            "SQLQueryError",
	SQLPrepareStatementError: "SQLPrepareStatementError",
	SQLScanError:             "SQLScanError",
	SQLExecutionError:        "SQLExecutionError",

	SQLTransactionError:         "SQLTransactionError",
	SQLTransactionBeginError:    "SQLTransactionError",
	SQLTransactionRoolbackError: "SQLTransactionRoolbackError",
	SQLTransactionCommitError:   "SQLTransactionCommitError",

	Unknown: "Unknown",
}

func (s StatusCode) Message() string {
	return MessageForCode[s]
}

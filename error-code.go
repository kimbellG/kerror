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

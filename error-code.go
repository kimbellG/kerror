package kerror

type StatusCode int

const (
	InvalidID StatusCode = iota
	NotFound
	BadRequest
	IntervalServerError

	TournamentDoesntExists
	UserDoesntExists

	SQLPrepareStatementError
	SQLConstraintError
	SQLScanError
	SQLExecutionError

	SQLTransactionBeginError
	SQLTransactionRoolbackError
	SQLTransactionCommitError
)

package kerror

type StatusCode int

const (
	InvalidID StatusCode = iota
	NotFound
	BadRequest
	IntervalServerError
	TournamentDoesntExists

	SQLPrepareStatementError
	SQLConstraintError
	SQLScanError
	SQLExecutionError

	SQLTransactionBeginError
	SQLTransactionRoolbackError
	SQLTransactionCommitError
)

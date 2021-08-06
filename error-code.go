package kerror

type StatusCode int

const (
	InvalidID StatusCode = iota
	NotFound
)

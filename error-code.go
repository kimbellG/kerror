package trnmnterr

type StatusCode int

const (
	InvalidId StatusCode = iota
	NotFound
)

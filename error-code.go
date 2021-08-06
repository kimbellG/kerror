package trnmnterr

type StatusCode int

const (
	InvalidID StatusCode = iota
	NotFound
)

var statusMessage = map[StatusCode]string{
	InvalidID: "invalid element's id",
	NotFound:  "element not found",
}

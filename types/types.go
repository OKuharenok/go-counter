package types

// Reader is the interface that wraps Read method.
type Reader interface {
	Read() ([]byte, error)
}

// Result stores source path, count substring and error text
type Result struct {
	Path  string
	Count int
	Error error
}

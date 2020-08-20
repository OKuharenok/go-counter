package filereader

import (
	"fmt"
	"io/ioutil"

	"github.com/OKuharenok/go-counter/types"
)

type reader struct {
	Path string
}

// NewReader creates reader for file source
func NewReader(path string) types.Reader {
	return &reader{Path: path}
}

// Read is method for reading data from file
func (fr *reader) Read() ([]byte, error) {
	file, err := ioutil.ReadFile(fr.Path)
	if err != nil {
		return nil, fmt.Errorf("Reading file failed: %s", err)
	}

	return file, nil
}

package filereader

import (
	"fmt"
	"io/ioutil"
)

// Reader stores path of the data source file
type Reader struct {
	Path string
}

// NewReader creates reader for file source
func NewReader(path string) *Reader {
	return &Reader{Path: path}
}

// Read is method for reading data from file
func (fr *Reader) Read() ([]byte, error) {
	file, err := ioutil.ReadFile(fr.Path)
	if err != nil {
		return nil, fmt.Errorf("Reading file failed: %s", err)
	}

	return file, nil
}

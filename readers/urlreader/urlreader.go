package urlreader

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Reader stores URL of the data source
type Reader struct {
	Path string
}

// NewReader creates reader for URL source
func NewReader(path string) *Reader {
	return &Reader{Path: path}
}

// Read is method for reading data from URL
func (ur *Reader) Read() ([]byte, error) {
	resp, err := http.Get(ur.Path)
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return nil, fmt.Errorf("Request failed: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Get error http status: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Reading body failed: %s", err)
	}

	return body, nil
}

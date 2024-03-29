package compress

import (
	"bytes"
	"compress/gzip"
)

// GzBytes takes a given byte array and returns a compressed version using gz
func GzBytes(data []byte) ([]byte, error) {
	var renderedGzBuffer bytes.Buffer
	gzw := gzip.NewWriter(&renderedGzBuffer)
	if _, err := gzw.Write(data); err != nil {
		return nil, err
	}
	gzd := renderedGzBuffer.Bytes()
	renderedGzBuffer.Reset()

	return gzd, nil
}

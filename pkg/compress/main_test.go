package compress

import (
	"bytes"
	"compress/gzip"
	"testing"
)

func TestGzBytes(t *testing.T) {
	// Test data
	inputData := []byte("hello, world")

	// Call the function
	_, err := GzBytes(inputData)
	if err != nil {
		t.Fatalf("Expected gz compress to work for buffer, but failed: %v", err)
	}
}

// Helper function to compress data using gzip
func compressData(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	_, err := gz.Write(data)
	if err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

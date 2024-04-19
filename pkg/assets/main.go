package assets

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func loadFileOrUrl(path string) ([]byte, error) {
	var data []byte
	// Check if the file exists locally
	if _, err := os.Stat(path); err == nil {
		// File exists locally, read it
		data, err = os.ReadFile(path)
		if err != nil {
			return nil, err
		}
	} else if strings.HasPrefix(path, "http") {
		// File does not exist locally, download it from the internet
		resp, err := http.Get(path)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf(path + " could not be found and is no resolvable URL")
	}
	return data, nil
}

// Must return valid result
func Must(s string, err error) string {
	if err != nil {
		panic(err)
	}
	return s
}

func shouldBeSvgMimeType(mimeType string) bool {
	return mimeType == xmlMimeType || mimeType == plainTextMimeType
}

// InlineIcon takes a local path or URL and returns a base64-encoded data URI
func InlineIcon(path string) (string, error) {
	data, err := loadFileOrUrl(path)
	if err != nil {
		return "", err
	}

	mimeType := http.DetectContentType(data)
	if shouldBeSvgMimeType(mimeType) {
		mimeType = svgMimeType
	}

	return fmt.Sprintf("data:%s;base64,%s", mimeType, base64.StdEncoding.EncodeToString(data)), nil
}

// LookupImgAnyExt tries to load a file with any of the supported given image file types until one is found
// in case there is no such file, it returns an error
func LookupImgAnyExt(baseName string) (string, error) {
	for _, extension := range supportedImageExtensions {
		filePath := baseName + "." + extension
		if _, err := os.Stat(filePath); err != nil {
			continue
		}
		return filePath, nil
	}

	return "", fmt.Errorf("no supported image file format found for " + baseName)
}

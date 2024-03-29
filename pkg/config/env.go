package config

import (
	"github.com/timo-reymann/yal/pkg/assets"
	"os"
	path2 "path"
	"strings"
)

func envWithDefault(variable string, fallback string) string {
	val := os.Getenv("YAL_" + variable)
	if strings.TrimSpace(val) == "" {
		val = fallback
	}

	return val
}

func imageFromEnv(variable string, fallbackFile string) (string, error) {
	path := envWithDefault(variable, fallbackFile)
	resolvedPath, err := assets.LookupImgAnyExt(path2.Join(ImageFolder(), path))
	// in case file can not be found locally try to load from URL
	// or search for full file name
	if err != nil {
		return assets.InlineIcon(path)
	}

	return assets.InlineIcon(resolvedPath)
}

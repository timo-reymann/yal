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
	if err != nil {
		return "", err
	}

	return assets.InlineIcon(resolvedPath)
}

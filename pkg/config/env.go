package config

import (
	"os"
	"strings"
)

func envWithDefault(variable string, fallback string) string {
	val := os.Getenv("YAL_" + variable)
	if strings.TrimSpace(val) == "" {
		val = fallback
	}

	return val
}

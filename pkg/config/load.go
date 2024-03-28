package config

import (
	"encoding/json"
	"github.com/timo-reymann/yal/pkg/assets"
	"github.com/timo-reymann/yal/pkg/model"
	"os"
	"strings"
)

type Assets struct {
	Logo       string
	Mascot     string
	Background string
	Favicon    string
}

type Configuration struct {
	Sections      []model.Section
	SearchEngines []model.SearchEngine
	Assets        Assets
}

func envWithDefault(variable string, fallback string) string {
	val := os.Getenv("YAL_" + variable)
	if strings.TrimSpace(val) == "" {
		val = fallback
	}

	return val
}

func imageFromEnv(variable string, fallbackFile string) (string, error) {
	path := envWithDefault(variable, "images/"+fallbackFile)
	resolvedPath, err := assets.LookupImgAnyExt(path)
	if err != nil {
		return "", err
	}

	return assets.InlineIcon(resolvedPath)
}

func Port() string {
	return envWithDefault("PORT", "2024")
}

// Load config based on file system and environment variables
func Load() (*Configuration, error) {
	sections, err := loadSections()
	if err != nil {
		return nil, err
	}

	searchEngines, err := loadSearchEngines()
	if err != nil {
		return nil, err
	}

	logo, err := imageFromEnv("LOGO", "logo")
	if err != nil {
		return nil, err
	}

	mascot, err := imageFromEnv("MASCOT", "mascot")
	if err != nil {
		return nil, err
	}

	background, err := imageFromEnv("BACKGROUND", "background")
	if err != nil {
		return nil, err
	}

	favicon, err := imageFromEnv("FAVICON", "favicon")
	if err != nil {
		return nil, err
	}

	return &Configuration{
		Sections:      sections,
		SearchEngines: searchEngines,
		Assets: Assets{
			Logo:       logo,
			Mascot:     mascot,
			Background: background,
			Favicon:    favicon,
		},
	}, nil
}

func loadConfig(baseName string, v interface{}) error {
	data, err := os.ReadFile("config/" + baseName + ".json")
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}

func loadSearchEngines() ([]model.SearchEngine, error) {
	var searchEngines []model.SearchEngine
	err := loadConfig("searchEngines", &searchEngines)
	return searchEngines, err
}

func loadSections() ([]model.Section, error) {
	var sections []model.Section
	err := loadConfig("items", &sections)
	return sections, err
}

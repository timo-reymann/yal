package config

import (
	"encoding/json"
	"github.com/timo-reymann/yal/pkg/model"
	"os"
	path2 "path"
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

// Port returns the HTTP port the service will be available on
func Port() string {
	return envWithDefault("PORT", "2024")
}

// ImageFolder returns the folder where images will be searched in
func ImageFolder() string {
	return envWithDefault("IMAGES_FOLDER", "images")
}

// ConfigFolder returns the folder where config files are located
func ConfigFolder() string {
	return envWithDefault("CONFIG_FOLDER", "config")
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
	data, err := os.ReadFile(path2.Join(ConfigFolder(), baseName+".json"))
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

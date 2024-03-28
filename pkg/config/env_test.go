package config

import (
	"os"
	"testing"
)

func TestEnvWithDefault_EnvSet(t *testing.T) {
	os.Setenv("YAL_VARIABLE", "value")

	result := envWithDefault("VARIABLE", "fallback")
	expected := "value"

	if result != expected {
		t.Errorf("Expected result %s, got %s", expected, result)
	}
}

func TestEnvWithDefault_EnvNotSet(t *testing.T) {
	os.Unsetenv("YAL_VARIABLE")

	result := envWithDefault("VARIABLE", "fallback")
	expected := "fallback"

	if result != expected {
		t.Errorf("Expected result %s, got %s", expected, result)
	}
}

func TestEnvWithDefault_EnvEmpty(t *testing.T) {
	os.Setenv("YAL_VARIABLE", "")

	result := envWithDefault("VARIABLE", "fallback")
	expected := "fallback"

	if result != expected {
		t.Errorf("Expected result %s, got %s", expected, result)
	}
}

func TestImageFromEnv_ExistingFile(t *testing.T) {
	os.Setenv("YAL_IMAGE_VARIABLE", "testdata/icon")
	_, err := imageFromEnv("IMAGE_VARIABLE", "testdata/fallback.png")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestImageFromEnv_NonExistingFile(t *testing.T) {
	os.Setenv("YAL_IMAGE_VARIABLE", "nonexistent.png")

	_, err := imageFromEnv("IMAGE_VARIABLE", "fallback.png")
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}

func TestPort_NotSet(t *testing.T) {
	if Port() != "2024" {
		t.Error("Expected default port to be set")
	}
}

func TestConfigFolder_NotSet(t *testing.T) {
	if ConfigFolder() != "config" {
		t.Error("Expected default config folder to be set")
	}
}

func TestLoad_Default(t *testing.T) {
	os.Setenv("YAL_CONFIG_FOLDER", "../../config")
	os.Setenv("YAL_IMAGES_FOLDER", "../../images")

	c, err := Load()
	if err != nil {
		t.Fatal(err)
	}

	if len(c.Sections) == 0 {
		t.Error("Default sections should be populated")
	}

	if len(c.SearchEngines) == 0 {
		t.Error("Default search engines should be populated")
	}
}

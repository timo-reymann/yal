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

func TestLoad_Default(t *testing.T) {
	os.Setenv("YAL_CONFIG_FOLDER", "../../config")
	os.Setenv("YAL_IMAGES_FOLDER", "../../images")

	c := Get()
	err := c.Load()
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

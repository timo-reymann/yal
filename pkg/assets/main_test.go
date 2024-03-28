package assets

import (
	"errors"
	"fmt"
	"testing"
)

func TestMust_ValidInput(t *testing.T) {
	// Call Must with valid input
	result := Must("valid", nil)

	// Check if the result matches the expected value
	expected := "valid"
	if result != expected {
		t.Errorf("Must returned incorrect result, got: %s, want: %s", result, expected)
	}
}

func TestMust_PanicOnInvalidInput(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Must did not panic with invalid input")
		}
	}()

	// Call Must with invalid input
	Must("", errors.New("test"))
}

func TestLookupImgAnyExt_Exists(t *testing.T) {
	baseName := "testdata/logo"
	expectedPath := baseName + ".png"

	// Call LookupImgAnyExt with an existing image file
	result, err := LookupImgAnyExt(baseName)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expectedPath {
		t.Errorf("Expected path %s, got %s", expectedPath, result)
	}
}

func TestLookupImgAnyExt_NotExists(t *testing.T) {
	baseName := "nonexistent"

	// Call LookupImgAnyExt with a non-existing image file
	result, err := LookupImgAnyExt(baseName)
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
	if result != "" {
		t.Errorf("Expected empty result, got %s", result)
	}
	expectedError := fmt.Sprintf("no supported image file format found for %s", baseName)
	if err.Error() != expectedError {
		t.Errorf("Expected error message %s, got %s", expectedError, err.Error())
	}
}

func TestInlineIcon_LoadFile_Success(t *testing.T) {
	result, err := InlineIcon("testdata/logo.png")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result == "" {
		t.Error("Expected result to be not empty")
	}
}

func TestInlineIcon_LoadUrl_Error(t *testing.T) {
	// Call InlineIcon with a mock URL
	result, err := InlineIcon("https://exasmple.abc/logo.png")
	if err == nil {
		t.Fatalf("Expected error, but got nil")
	}

	if result != "" {
		t.Errorf("Expected empty result, got %s", result)
	}
	expectedError := "Get \"https://exasmple.abc/logo.png\": dial tcp: lookup exasmple.abc: no such host"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %s, got %s", expectedError, err.Error())
	}
}

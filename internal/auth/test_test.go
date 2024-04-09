package auth

import (
	"net/http"
	"testing"
)

// Success case
func TestGetAPIKeySuccess(t *testing.T) {
	headers := http.Header{}
	expectedKey := "yourActualApiKey"
	headers.Set("Authorization", "ApiKey "+expectedKey)

	apiKey, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if apiKey != expectedKey {
		t.Fatalf("Expected API key %s, got %s", expectedKey, apiKey)
	}
}

// No Authorization header case
func TestGetAPIKeyNoAuthHeader(t *testing.T) {
	headers := http.Header{}

	_, err := GetAPIKey(headers)
	if err == nil {
		t.Fatalf("Expected an error when no Authorization header is included, got none")
	}

	// You could also test for the specific error message if needed
}

// Malformed Authorization header case
func TestGetAPIKeyMalformedAuthHeader(t *testing.T) {
	headers := http.Header{}
	// Example of a malformed header
	headers.Set("Authorization", "BadFormatApiKey")

	_, err := GetAPIKey(headers)
	if err == nil {
		t.Fatalf("Expected an error for malformed Authorization header, got none")
	}

	// Optional: check if the error is the specific one you expected.
}

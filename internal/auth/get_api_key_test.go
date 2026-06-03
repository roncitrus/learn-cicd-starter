package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestApiKey(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey 123456")

	result, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "123456"
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}

}

func TestEmptyHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "")

	_, err := GetAPIKey(headers)
	if err == nil {
		t.Fatalf("expected error, didn't find one")
	}

	if !errors.Is(err, ErrNoAuthHeaderIncluded) {
		t.Errorf("expected error %v, got %v", ErrNoAuthHeaderIncluded, err)
	}
}

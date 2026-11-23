package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	h := http.Header{}
	h.Set("Authorization", "ApiKey 1234567")

	got, err := GetAPIKey(h)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", "1234567", err)
	}
	if got != "1234567" {
		t.Fatalf("expected: %v, got: %s", "1234567", got)
	}
}

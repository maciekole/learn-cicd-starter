package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKeyEmptyHeader(t *testing.T) {
	header := http.Header{}
	got, gotErr := GetAPIKey(header)
	want := ""
	if got != want && gotErr == nil {
		t.Fatalf("expected: %v, %v, got: %v, %v", want, ErrNoAuthHeaderIncluded, got, gotErr)
	}
}

func TestGetAPIKeyInvalidAuthHeader(t *testing.T) {
	header := http.Header{
		"Authorization": {},
	}
	got, gotErr := GetAPIKey(header)
	want := ""
	wantErr := errors.New("malformed authorization header")
	if got != want && gotErr == nil {
		t.Fatalf("expected: %v, %v, got: %v, %v", want, wantErr, got, gotErr)
	}
}

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthAndReadyEndpoints(t *testing.T) {
	handler := newHandler()

	for _, path := range []string{"/health", "/ready"} {
		req := httptest.NewRequest(http.MethodGet, path, nil)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("%s: expected 200, got %d", path, rec.Code)
		}
		if rec.Header().Get("X-Request-ID") == "" {
			t.Fatalf("%s: missing request id", path)
		}
		if rec.Header().Get("X-Content-Type-Options") != "nosniff" {
			t.Fatalf("%s: missing content type protection", path)
		}
	}
}

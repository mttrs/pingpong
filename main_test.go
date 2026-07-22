package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	handler(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("status code = %d, want %d", res.StatusCode, http.StatusOK)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}

	if got, want := string(body), "pong\n"; got != want {
		t.Errorf("body = %q, want %q", got, want)
	}
}

func TestHandlerAnyPath(t *testing.T) {
	// handler responds with "pong" regardless of the request path.
	paths := []string{"/", "/ping", "/anything/else"}
	for _, p := range paths {
		req := httptest.NewRequest(http.MethodGet, p, nil)
		rec := httptest.NewRecorder()

		handler(rec, req)

		if got, want := rec.Body.String(), "pong\n"; got != want {
			t.Errorf("path %q: body = %q, want %q", p, got, want)
		}
	}
}

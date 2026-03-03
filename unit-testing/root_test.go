package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoot(t *testing.T) {
	mux := NewMux()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expect %d result %d", http.StatusOK, rec.Code)
	}

	if rec.Body.String() != "Hello Golang Ghana" {
		t.Fatalf("Unexpected body %v", rec.Body.String())
	}
}

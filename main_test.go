package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(t *testing.T) {
	req, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	index(w, req)
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		t.Errorf("expected resp.StatusCode: %d, got: %d", 200, resp.StatusCode)
	}

	if string(body) != "<html><body>Hello World!</body></html>" {
		t.Errorf("expected body: %s, got: %s", "<html><body>Hello World!</body></html>", body)
	}
}

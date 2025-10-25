package util

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestJsonRequest runs a JsonHandler with a given body and returns the recorder
func TestJsonRequest(t *testing.T, handler http.Handler, method, url string, body any) *httptest.ResponseRecorder {
	t.Helper()

	var buf bytes.Buffer
	if body != nil {
		if err := json.NewEncoder(&buf).Encode(body); err != nil {
			t.Fatalf("failed to encode body: %v", err)
		}
	}

	req := httptest.NewRequest(method, url, &buf)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	return w
}

func TestJsonRequestAndDecode[T any](t *testing.T, handler http.Handler, method, url string, body any) (T, int) {
	t.Helper()
	var parsed T

	resp := TestJsonRequest(t, handler, method, url, body)

	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	return parsed, resp.Code
}

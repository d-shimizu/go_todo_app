package testutil

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func AssertJson(t *testing.T, want, got []byte) {
	t.Helper()

	var jw, jq any
	if err := json.Unmarshal(want, &jw); err != nil {
		t.Fatalf("failed to unmarshal want: %v", err)
	}
	if err := json.Unmarshal(got, &jq); err != nil {
		t.Fatalf("failed to unmarshal got: %v", err)
	}
	if diff := cmp.Diff(jw, jq); diff != "" {
		t.Errorf("JSON mismatch (-want +got):\n%s", diff)
	}
}

func AssertResponse(t *testing.T, got *http.Response, status int, body []byte) {
	t.Helper()

	t.Cleanup(func() { _ = got.Body.Close() })
	gb, err := io.ReadAll(got.Body)
	if err != nil {
		t.Fatal(err)
	}
	if got.StatusCode != status {
		t.Fatalf("want status %d, but got %d, body: %q", status, got.StatusCode, gb)
	}

	if len(gb) == 0 && len(body) == 0 {
		return
	}

	AssertJson(t, body, gb)
}

func LoadFile(t *testing.T, path string) []byte {
	t.Helper()

	bt, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read file %s: %v", path, err)
	}
	return bt
}

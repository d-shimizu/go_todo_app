package testutil

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func AssertJSON(t *testing.T, want, got []byte) {
	t.Helper()

	var jw, jq any
	if err := json.Unmarshal(want, &jw); err != nil {
		t.Fatalf("cannot unmarshal want: %q: %v", want, err)
	}
	if err := json.Unmarshal(got, &jq); err != nil {
		t.Fatalf("cannnot unmarshal got: %q: %v", got, err)
	}
	if diff := cmp.Diff(jq, jw); diff != "" {
		t.Errorf("(-got +want)\n%s", diff)
	}
}

func AssertResponse(t *testing.T, got *http.Response, status int, body []byte) {
	t.Helper()

	if got.StatusCode != status {
		t.Errorf("status code: got %d, want %d", got.StatusCode, status)
	}
	if got.Body == nil {
		t.Fatalf("body is nil")
	}
	defer got.Body.Close()

	gotBody, err := io.ReadAll(got.Body)
	if err != nil {
		t.Fatalf("cannot read body: %v", err)
	}
	AssertJSON(t, body, gotBody)
}

func LoadFile(t *testing.T, path string) []byte {
	t.Helper()

	bt, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("cannnot read from %q: %v", path, err)
	}
	return bt
}

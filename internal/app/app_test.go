package app

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunWithOutput(t *testing.T) {
	var buf bytes.Buffer

	if err := RunWithOutput(nil, &buf); err != nil {
		t.Fatalf("RunWithOutput returned error: %v", err)
	}

	if got := buf.String(); !strings.Contains(got, Version) {
		t.Fatalf("output %q does not contain version %q", got, Version)
	}
}

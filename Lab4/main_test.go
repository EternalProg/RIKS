package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("pipe error: %v", err)
	}
	defer r.Close()

	os.Stdout = w
	main()
	_ = w.Close()
	os.Stdout = oldStdout

	out, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("read error: %v", err)
	}

	text := string(out)
	if !strings.Contains(text, "Int sequence:") || !strings.Contains(text, "Average of float sequence:") {
		t.Fatalf("unexpected output: %q", text)
	}
}

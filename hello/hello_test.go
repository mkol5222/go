package hello

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestHello(t *testing.T) {
	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the function
	Hello()

	// Restore stdout
	w.Close()
	os.Stdout = oldStdout

	// Read the captured output
	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// Verify the output
	expected := "Hello, World!\n"
	if output != expected {
		t.Errorf("Hello() output = %q, want %q", output, expected)
	}
}

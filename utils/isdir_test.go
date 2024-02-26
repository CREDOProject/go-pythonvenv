package utils

import (
	"os"
	"testing"
)

func TestIsDir(t *testing.T) {
	tempdir := t.TempDir()
	if !IsDir(tempdir) {
		t.Errorf("IsDir(%q) = false, want true", tempdir)
	}

	tempfile, err := os.CreateTemp(tempdir, "file")
	if err != nil {
		t.Fatal(err)
	}
	defer tempfile.Close()

	if IsDir(tempfile.Name()) {
		t.Errorf("IsDir(%q) = true, want false", tempfile.Name())
	}
}

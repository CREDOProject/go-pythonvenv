package gopythonvenv

import (
	"os"
	"path"
	"testing"
)

func Test_CreateVenv(t *testing.T) {
	// expectedDir := []string{"bin", "include", "lib"}
	tmp := t.TempDir()
	v, e := Create(path.Join(tmp, "venv"))
	if e != nil {
		t.Error(e)
	}
	l, e := os.ReadDir(v.Path)
	if len(l) != 4 {
		t.Error("Some files are not expected.")
	}
	// Create once again
	v, e = Create(path.Join(tmp, "venv"))
	if e != nil {
		t.Error(e)
	}
	l, e = os.ReadDir(v.Path)
	if len(l) != 4 {
		t.Error("Some files are not expected.")
	}
}

package providers

import (
	"runtime"
	"testing"
)

func TestLooksLikePython(t *testing.T) {
	var tests map[string]bool

	switch runtime.GOOS {
	case "windows":
		tests = map[string]bool{
			"python.exe":         true,
			"python3.exe":        true,
			"python39.exe":       true,
			"python310.exe":      true,
			"python-build":       false,
			"python-python3.exe": false,
		}
	default:
		tests = map[string]bool{
			"python":         true,
			"python3":        true,
			"python3.9":      true,
			"python3.10":     true,
			"python-build":   false,
			"python-python3": false,
		}
	}

	for name, want := range tests {
		t.Run(name, func(t *testing.T) {
			got := looksLikePython(name)
			if got != want {
				t.Errorf("looksLikePython(%q) = %v, want %v", name, got, want)
			}
		})
	}
}

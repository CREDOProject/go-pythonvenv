package providers

import (
	"fmt"
	"os/exec"
	"strings"

	pep440Version "github.com/aquasecurity/go-pep440-version"
)

// PythonExecutable contains information about a Python executable.
type PythonExecutable struct {
	// Version is the parsed Python version.
	Version *pep440Version.Version

	// Path is the absolute path to the Python executable.
	Path string
}

// NewPythonExecutable creates a new PythonExecutable from a Python path.
func NewPythonExecutable(executable string) (*PythonExecutable, error) {
	versionInfo, err := getPythonVersion(executable)
	if err != nil {
		return nil, err
	}
	return &PythonExecutable{Version: versionInfo, Path: executable}, nil
}

// Returns a string representation of a PythonExecutable
func (v *PythonExecutable) String() string {
	return fmt.Sprintf("%s (%s)", v.Version, v.Path)
}

// getPythonVersion returns the version information for the given Python
// executable.
func getPythonVersion(executable string) (*pep440Version.Version, error) {
	cmd := exec.Command(executable, "--version")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	// Note that the output follows this example scheme:
	// Python 3.12.2
	// Python <version><LF/CRLF>
	_, version, found := strings.Cut(string(output), " ")
	if !found {
		return nil, fmt.Errorf("Unable to parse Python version: %q", output)
	}
	version = strings.TrimRight(version, "\r\n")

	versionInfo, err := pep440Version.Parse(version)
	if err != nil {
		return nil, err
	}
	return &versionInfo, nil
}

// Returns a pep440Version.Specifier.
func Specifier(versionInfo *pep440Version.Version) (*pep440Version.Specifiers, error) {
	s, err := pep440Version.NewSpecifiers("== " + getGlobVersion(versionInfo))

	if err != nil {
		return nil, err
	}
	return &s, nil
}

// getGlobVersion returns the glob version for the given version if it's not
// a complete version, otherwise it returns the given version.
//
// This function assumes that the given version is a final release.
func getGlobVersion(version *pep440Version.Version) string {
	v := version.String()
	switch len(strings.Split(v, ".")) {
	case 1, 2:
		return fmt.Sprintf("%s.*", v)
	default:
		return v
	}
}

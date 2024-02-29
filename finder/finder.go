package finder

import (
	"github.com/CREDOProject/go-pythonvenv/providers"
	version "github.com/aquasecurity/go-pep440-version"
)

// finder is a Python version finder.
type finder struct {
	providers []providers.Provider
}

// New returns a new Python version finder.
func New() *finder {
	f := &finder{}
	f.setupProviders()
	return f
}

func (f *finder) setupProviders() {
	f.providers = append(f.providers, providers.NewPathProvider())
}

// Find returns the latest Python version installed on the system.
func (f *finder) Find() (*providers.PythonExecutable, error) {
	var maxVersion *providers.PythonExecutable

	baseversion, err := version.Parse("3.0.0") // Uses python 3
	if err != nil {
		return nil, err
	}
	specifier, err := providers.Specifier(&baseversion)
	if err != nil {
		return nil, err
	}

	// seen is a set of Python executables which were already seen by the
	// providers. This is used to avoid returning duplicate Python versions.
	// This contains the absolute path to the Python executable.
	seen := make(map[string]struct{})

	for _, p := range f.providers {
		executables, err := p.Executables()
		if err != nil {
			return nil, err
		}
		for _, executable := range executables {
			if _, ok := seen[executable]; ok {
				continue
			}
			seen[executable] = struct{}{}

			pythonExecutable, err := providers.NewPythonExecutable(executable)
			if err != nil {
				return nil, err
			}
			if maxVersion == nil {
				maxVersion = pythonExecutable
			}
			if specifier.Check(*pythonExecutable.Version) {
				if pythonExecutable.Version.GreaterThan(*maxVersion.Version) {
					maxVersion = pythonExecutable
				}
			}
		}
	}
	return maxVersion, nil
}

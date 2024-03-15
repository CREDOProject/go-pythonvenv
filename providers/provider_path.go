package providers

import (
	"os"
	"strings"

	"github.com/CREDOProject/sharedutils/files"
)

// pathProvider is a Provider that finds Python executables in the PATH
// environment variable.
type pathProvider struct {
	paths []string
}

// newPathProvider returns a new pathProvider.
func NewPathProvider() *pathProvider {
	return &pathProvider{
		paths: strings.Split(os.Getenv("PATH"), string(os.PathListSeparator)),
	}
}

func (p *pathProvider) Executables() ([]string, error) {
	var executables []string
	for _, path := range p.paths {
		if !files.IsDir(path) {
			continue
		}
		execs, err := files.ExecsInPath(path, looksLikePython)
		if err != nil {
			return nil, err
		}
		executables = append(executables, execs...)
	}
	return executables, nil
}

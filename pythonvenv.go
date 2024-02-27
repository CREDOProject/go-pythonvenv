package gopythonvenv

import (
	"errors"
	"os"
	"os/exec"

	"github.com/CREDOProject/go-pythonvenv/finder"
	"github.com/CREDOProject/go-pythonvenv/utils"
	"github.com/CREDOProject/sharedutils/files"
)

var (
	ErrAlreadyPresent = errors.New("Virtual environment already exists.")
)

// Structure representing a PythonVenv
type PythonVenv struct {
	Path string
}

// Creates a virtual environment in the specified path using the latest
// Python version available in the system.
func Create(path string) (*PythonVenv, error) {
	err := createVenv(path)

	if err != nil && err != ErrAlreadyPresent {
		return nil, err
	}

	return &PythonVenv{
		Path: path,
	}, nil
}

func createVenv(path string) error {
	if files.IsDir(path) {
		return ErrAlreadyPresent
	}
	v, err := finder.New().Find()
	if err != nil {
		return err
	}
	cmd := exec.Command(v.Path, "-m", "venv", path)
	if err = cmd.Start(); err != nil {
		return err
	}
	return cmd.Wait()
}

// Activates the virtual environment in the path.
func (g *PythonVenv) Activate() []string {
	environment := utils.Env()

	environment["VIRTUAL_ENV"] = g.Path
	environment["PATH"] = g.Path + "/bin:" + environment["PATH"]
	delete(environment, "PYTHONHOME")

	return utils.Roll(environment)
}

// Removes the virtual environment in the specified path.
func (g *PythonVenv) Destroy() error {
	if err := os.RemoveAll(g.Path); err != nil {
		return err
	}
	return nil
}

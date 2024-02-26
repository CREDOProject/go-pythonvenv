package gopythonvenv

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/CREDOProject/go-pythonvenv/finder"
	"github.com/CREDOProject/go-pythonvenv/utils"
)

// Structure representing a PythonVenv
type GoPythonVenv struct {
	Path string
}

// Creates a virtual environment in the specified path using the latest
// Python version available in the system.
func Create(path string) (*GoPythonVenv, error) {
	err := createVenv(path)

	if err != nil {
		return nil, err
	}

	return &GoPythonVenv{
		Path: path,
	}, nil
}

func createVenv(path string) error {
	if utils.IsDir(path) {
		return fmt.Errorf("Virtual environment already exists.")
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

// Activates the virtual environment in the specified path.
func (g *GoPythonVenv) Activate() []string {
	environment := utils.Env()

	environment["VIRTUAL_ENV"] = g.Path
	environment["PATH"] = g.Path + "/bin:" + environment["PATH"]
	delete(environment, "PYTHONHOME")

	return utils.Roll(environment)
}

// Removes the virtual environment in the specified path.
func (g *GoPythonVenv) Destroy() error {
	if err := os.RemoveAll(g.Path); err != nil {
		return err
	}
	return nil
}

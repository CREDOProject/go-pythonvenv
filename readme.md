# go-pythonvenv

A bridge helping manage python virtual environments from go.

---

Based on: https://github.com/dhruvmanila/pie

## Usage:

```go

env, err := pythonvenv.Create("path")
if err != nil {
  // Handle error
}

// You can use this environment for further operation.
venv := env.Activate()

// This removes all the venv files.
env.Destroy()
```

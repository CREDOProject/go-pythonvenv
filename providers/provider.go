package providers

// A provider is a type of object that provides Python Executables.
type Provider interface {
	Executables() ([]string, error)
}

// looksLikePython returns true if the given filename looks like a Python
// executable.
func looksLikePython(name string) bool {
	return pythonFileRegex.MatchString(name)
}

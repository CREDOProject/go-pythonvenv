package providers

import "regexp"

var pythonFileRegex = regexp.MustCompile(`^python(\d(\d\d?)?)?\.exe$`)

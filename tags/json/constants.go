package json

import (
	"regexp"
)

var (
	overwriteTagsRegExp = regexp.MustCompile(`json:"[^"]*"`)
)

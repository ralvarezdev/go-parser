package go_parser

import (
	"errors"
)

var (
	NilFileSetError = errors.New("file set is nil")
	NilASTNodeError = errors.New("ast node is nil")
)

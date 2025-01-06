package go_parser

import (
	"errors"
)

var (
	ErrNilParser  = errors.New("go file parser cannot be nil")
	ErrNilFileSet = errors.New("file set cannot be nil")
	ErrNilASTNode = errors.New("ast node cannot be nil")
)

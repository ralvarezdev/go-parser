package json

import (
	"errors"
)

var (
	ErrNilStructsFieldsMapper      = errors.New("structs fields mapper cannot be nil")
	ErrNilFilesStructsFieldsMapper = errors.New("files structs fields mapper cannot be nil")
	ErrNilStructsTagsMapper        = errors.New("structs tags mapper cannot be nil")
)

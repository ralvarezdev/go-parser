package json

import (
	"errors"
)

var (
	NilStructsFieldsMapperError  = errors.New("nil structs fields mapper")
	NilGoStructFieldsMapperError = errors.New("nil go struct fields mapper")
	NilStructJSONTagMapperError  = errors.New("nil struct json tag mapper")
)

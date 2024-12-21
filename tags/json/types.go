package json

type (
	// FieldJSONTagMapper map the field name to the new JSON tag
	FieldJSONTagMapper map[string]string

	// StructJSONTagMapper map the struct name to the field JSON tag
	StructJSONTagMapper map[string]FieldJSONTagMapper

	// Fields is a slice of struct fields
	Fields []string

	// StructsFieldsMapper map the struct name to the fields
	StructsFieldsMapper map[string]Fields

	// GoFileStructFieldsMapper map the Go file path to the struct fields
	GoFileStructFieldsMapper map[string]*StructsFieldsMapper
)

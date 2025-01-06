package json

type (
	// FieldsTagsMapper map the field name to the new JSON tag
	FieldsTagsMapper map[string]string

	// StructsTagsMapper map the struct name to the field JSON tag
	StructsTagsMapper map[string]FieldsTagsMapper

	// Fields is a slice of struct fields
	Fields []string

	// StructsFieldsMapper map the struct name to the fields
	StructsFieldsMapper map[string]Fields

	// FilesStructsFieldsMapper map the Go file path to the struct fields
	FilesStructsFieldsMapper map[string]*StructsFieldsMapper
)

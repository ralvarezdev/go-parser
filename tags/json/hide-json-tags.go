package json

import (
	goflagsmode "github.com/ralvarezdev/go-flags/mode"
)

// HideFileJSONTags hides the JSON tags in the specified file
func HideFileJSONTags(
	filePath string,
	structsFieldsMapper *StructsFieldsMapper,
	mode *goflagsmode.Flag,
) error {
	// Check if the structs fields mapper is nil
	if structsFieldsMapper == nil {
		return NilStructsFieldsMapperError
	}

	// Generate the StructJSONTag
	structJSONTagMapper := StructJSONTagMapper{}

	// Loop through the struct fields
	for structName, fields := range *structsFieldsMapper {
		fieldJSONTagMapper := FieldJSONTagMapper{}
		for _, field := range fields {
			fieldJSONTagMapper[field] = "-"
		}
		structJSONTagMapper[structName] = fieldJSONTagMapper
	}
	// Overwrite the JSON tags
	return OverwriteJSONTags(filePath, &structJSONTagMapper, mode)
}

// HideFilesJSONTags hides the JSON tags in the specified files
func HideFilesJSONTags(
	goFileStructFieldsMapper *GoFileStructFieldsMapper,
	mode *goflagsmode.Flag,
) error {
	// Check if the GoFileStructFieldsMapper is nil
	if goFileStructFieldsMapper == nil {
		return NilGoStructFieldsMapperError
	}

	// Loop through the file paths
	for filePath, structsFieldsMapper := range *goFileStructFieldsMapper {
		// Hide the JSON tags in the file
		if err := HideFileJSONTags(
			filePath,
			structsFieldsMapper,
			mode,
		); err != nil {
			return err
		}
	}
	return nil
}

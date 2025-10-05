package json

import (
	"log/slog"
)

// DetectedField prints a detected field
//
// Parameters:
//
//   - structTypeName: the name of the struct type
//   - fieldName: the name of the field
//   - logger: the logger to use, if nil, no logging is done
func DetectedField(
	structTypeName string,
	fieldName string,
	logger *slog.Logger,
) {
	if logger != nil {
		logger.Debug(
			"detected field on struct type",
			slog.String("struct_type_name", structTypeName),
			slog.String("field_name", fieldName),
		)
	}
}

// FieldsNotUpdated prints the fields that were not updated
//
// Parameters:
//
//   - structJSONTagMapper: the mapper of struct names to their fields and JSON tags
//   - logger: the logger to use, if nil, no logging is done
func FieldsNotUpdated(
	structJSONTagMapper StructsTagsMapper,
	logger *slog.Logger,
) {
	// Check if the structJSONTagMapper is nil
	if structJSONTagMapper == nil {
		return
	}

	// Iterate over the fields that haven't been updated
	for structTypeName := range structJSONTagMapper {
		var structFields []string

		// Get the struct fields
		for fieldName := range structJSONTagMapper[structTypeName] {
			structFields = append(
				structFields,
				fieldName,
			)
		}

		if logger != nil {
			logger.Debug(
				"some fields haven't been updated on struct",
				slog.String("struct_type_name", structTypeName),
				slog.Any("struct_fields", structFields),
			)
		}
	}
}

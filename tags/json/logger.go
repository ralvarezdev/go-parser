package json

import (
	gologgermode "github.com/ralvarezdev/go-logger/mode"
	gologgermodenamed "github.com/ralvarezdev/go-logger/mode/named"
)

// Logger is the JWT validator logger
type Logger struct {
	logger gologgermodenamed.Logger
}

// NewLogger creates a new JWT validator logger
func NewLogger(header string, modeLogger gologgermode.Logger) (*Logger, error) {
	// Initialize the mode named logger
	namedLogger, err := gologgermodenamed.NewDefaultLogger(header, modeLogger)
	if err != nil {
		return nil, err
	}

	return &Logger{logger: namedLogger}, nil
}

// DetectedField prints a detected field
func (l *Logger) DetectedField(
	structTypeName string,
	fieldName string,
) {
	l.logger.Debug(
		"detected field on struct type: "+structTypeName,
		"field name: "+fieldName,
	)
}

// FieldsNotUpdated prints the fields that were not updated
func (l *Logger) FieldsNotUpdated(
	structJSONTagMapper *StructsTagsMapper,
) {
	// Check if the structJSONTagMapper is nil
	if structJSONTagMapper == nil {
		return
	}

	// Iterate over the fields that haven't been updated
	for structTypeName := range *structJSONTagMapper {
		var structFields []string

		// Get the struct fields
		for fieldName := range (*structJSONTagMapper)[structTypeName] {
			structFields = append(
				structFields,
				"field name: "+fieldName,
			)
		}

		l.logger.Debug(
			"some fields haven't been updated on struct: "+structTypeName,
			structFields...,
		)
	}
}

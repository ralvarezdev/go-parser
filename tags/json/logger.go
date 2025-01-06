package json

import (
	"fmt"
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

// PrintField prints a field
func (l *Logger) PrintField(
	structName string,
	fieldName string,
) {
	l.logger.Debug(
		fmt.Sprintf("detected field on '%v'", structName),
		fmt.Sprintf("field '%v'", fieldName),
	)
}

// PrintFieldsNotUpdated prints the fields that were not updated
func (l *Logger) PrintFieldsNotUpdated(
	structJSONTagMapper *StructsTagsMapper,
) {
	// Check if the structJSONTagMapper is nil
	if structJSONTagMapper == nil {
		return
	}

	// Iterate over the fields that haven't been updated
	var structFields []string
	for structName := range *structJSONTagMapper {
		for fieldName := range (*structJSONTagMapper)[structName] {
			structFields = append(
				structFields,
				fmt.Sprintf("field: '%s'", fieldName),
			)
		}
	}

	l.logger.Debug(
		fmt.Sprintf("some fields haven't been updated"),
		structFields...,
	)
}

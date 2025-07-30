package json

import (
	"fmt"
	"go/ast"
	"go/token"
	"regexp"

	goparser "github.com/ralvarezdev/go-parser"
)

type (
	// Writer is the interface for writing JSON tags
	Writer interface {
		OverwriteTags(
			filePath string,
			structJSONTagMapper *StructsTagsMapper,
		) error
		HideStructsTags(
			filePath string,
			structsFieldsMapper *StructsFieldsMapper,
		) error
		HideFilesTags(
			goFileStructFieldsMapper *FilesStructsFieldsMapper,
		) error
	}

	// DefaultWriter is the struct for the default JSON tag writer
	DefaultWriter struct {
		logger *Logger
		parser goparser.Parser
	}
)

// NewDefaultWriter creates a new DefaultWriter struct
func NewDefaultWriter(
	parser goparser.Parser,
	logger *Logger,
) (*DefaultWriter, error) {
	// Check if the parser is nil
	if parser == nil {
		return nil, goparser.ErrNilParser
	}

	return &DefaultWriter{logger, parser}, nil
}

// OverwriteTags overwrite the given structs fields JSON tags from the given Go file
func (d *DefaultWriter) OverwriteTags(
	filePath string,
	structsTagsMapper *StructsTagsMapper,
) error {
	// Check if the structs tags mapper is nil
	if structsTagsMapper == nil {
		return ErrNilStructsTagsMapper
	}

	// Compile the regex pattern to match any content inside JSON tag
	regex, err := regexp.Compile(`json:"[^"]*"`)
	if err != nil {
		return fmt.Errorf("failed to compile regex: %w", err)
	}

	// Parse the Go file
	fileSet, node, err := d.parser.ParseFile(filePath)
	if err != nil {
		return err
	}

	// Traverse the AST to find the struct and field
	err = d.parser.TraverseAST(
		node, func(n ast.Node) bool {
			// Check if the node is a type spec
			ts, ok := n.(*ast.TypeSpec)
			if !ok {
				return true
			}

			// Check if the node is a struct
			st, ok := ts.Type.(*ast.StructType)
			if !ok {
				return true
			}

			// Get the struct type name
			structTypeName := ts.Name.Name

			// Check if the struct name is in the map
			fieldJSONTag, ok := (*structsTagsMapper)[structTypeName]
			if !ok {
				return true
			}

			// Iterate the struct fields
			for _, field := range st.Fields.List {
				for _, name := range field.Names {
					// Check if the field name is in the map and get the new JSON tag
					fieldName := name.Name
					newJSONTag, ok := fieldJSONTag[fieldName]
					if !ok {
						continue
					}

					// Print the struct and field name
					if d.logger != nil {
						d.logger.DetectedField(
							structTypeName,
							fieldName,
						)
					}

					// Modify the JSON tag
					if field.Tag != nil {
						tag := field.Tag.Value

						// Replace the matched content
						newTag := regex.ReplaceAllString(
							tag,
							fmt.Sprintf(`json:"%s"`, newJSONTag),
						)

						field.Tag.Value = newTag
					} else {
						field.Tag = &ast.BasicLit{
							Kind:  token.STRING,
							Value: fmt.Sprintf("`json:\"%s\"`", newJSONTag),
						}
					}

					// Remove the field from the map
					delete(fieldJSONTag, fieldName)
				}
			}

			// Check if the struct has fields to update
			numFields := len(fieldJSONTag)
			if numFields == 0 {
				delete(*structsTagsMapper, ts.Name.Name)
				return false
			}
			return true
		},
	)
	if err != nil {
		return err
	}

	// Check if all structs have been updated
	if len(*structsTagsMapper) > 0 {
		// Print the structs fields that haven't been updated
		if d.logger != nil {
			d.logger.FieldsNotUpdated(structsTagsMapper)
		}
		return fmt.Errorf("failed to update all structs")
	}

	// Write the modified AST back to the file
	if err := d.parser.WriteFile(filePath, fileSet, node); err != nil {
		return err
	}

	return nil
}

// HideStructsTags hides the JSON tags from the structs in the specified file
func (d *DefaultWriter) HideStructsTags(
	filePath string,
	structsFieldsMapper *StructsFieldsMapper,
) error {
	// Check if the structs fields mapper is nil
	if structsFieldsMapper == nil {
		return ErrNilStructsFieldsMapper
	}

	// Generate the StructJSONTag
	structJSONTagMapper := StructsTagsMapper{}

	// Loop through the struct fields
	for structTypeName, fields := range *structsFieldsMapper {
		fieldJSONTagMapper := FieldsTagsMapper{}
		for _, field := range fields {
			fieldJSONTagMapper[field] = "-"
		}
		structJSONTagMapper[structTypeName] = fieldJSONTagMapper
	}
	// Overwrite the JSON tags
	return d.OverwriteTags(filePath, &structJSONTagMapper)
}

// HideFilesTags hides the JSON tags from the structs in the specified files
func (d *DefaultWriter) HideFilesTags(
	filesStructsFieldsMapper *FilesStructsFieldsMapper,
) error {
	// Check if the FilesStructsFieldsMapper is nil
	if filesStructsFieldsMapper == nil {
		return ErrNilFilesStructsFieldsMapper
	}

	// Loop through the file paths
	for filePath, structsFieldsMapper := range *filesStructsFieldsMapper {
		// Hide the JSON tags from the structs in the given file
		if err := d.HideStructsTags(
			filePath,
			structsFieldsMapper,
		); err != nil {
			return err
		}
	}
	return nil
}

package json

import (
	"fmt"
	"go/ast"
	"go/token"
	"log/slog"

	goparser "github.com/ralvarezdev/go-parser"
)

type (
	// DefaultWriter is the struct for the default JSON tag writer
	DefaultWriter struct {
		logger *slog.Logger
		parser goparser.Parser
	}
)

// NewDefaultWriter creates a new DefaultWriter struct
//
// Parameters:
//
//   - parser goparser.Parser: the Go parser to use
//   - logger *slog.Logger: the logger to use, if nil, no logging is done
//
// Returns:
//
//   - *DefaultWriter: the default JSON tag writer
//   - error: if any error occurs
func NewDefaultWriter(
	parser goparser.Parser,
	logger *slog.Logger,
) (*DefaultWriter, error) {
	// Check if the parser is nil
	if parser == nil {
		return nil, goparser.ErrNilParser
	}

	if logger != nil {
		logger = logger.With(
			slog.String("component", "tags_json_writer"),
		)
	}

	return &DefaultWriter{logger, parser}, nil
}

// OverwriteTags overwrite the given structs fields JSON tags from the given Go file
//
// Parameters:
//
//   - filePath string: the path to the Go file
//   - structsTagsMapper StructsTagsMapper: the mapper of struct names to their fields and new JSON tags
//
// Returns:
//
//   - error: if any error occurs
func (d DefaultWriter) OverwriteTags(
	filePath string,
	structsTagsMapper StructsTagsMapper,
) error {
	// Check if the structs tags mapper is nil
	if structsTagsMapper == nil {
		return ErrNilStructsTagsMapper
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
			fieldJSONTag, ok := structsTagsMapper[structTypeName]
			if !ok {
				return true
			}

			// Iterate the struct fields
			for _, field := range st.Fields.List {
				for _, name := range field.Names {
					// Check if the field name is in the map and get the new JSON tag
					fieldName := name.Name
					newJSONTag, tagOk := fieldJSONTag[fieldName]
					if !tagOk {
						continue
					}

					// Print the struct and field name
					DetectedField(
						structTypeName,
						fieldName,
						d.logger,
					)

					// Modify the JSON tag
					if field.Tag != nil {
						tag := field.Tag.Value

						// Replace the matched content
						newTag := overwriteTagsRegExp.ReplaceAllString(
							tag,
							fmt.Sprintf(`json:"%q"`, newJSONTag),
						)

						field.Tag.Value = newTag
					} else {
						field.Tag = &ast.BasicLit{
							Kind:  token.STRING,
							Value: fmt.Sprintf("`json:\"%q\"`", newJSONTag),
						}
					}

					// Remove the field from the map
					delete(fieldJSONTag, fieldName)
				}
			}

			// Check if the struct has fields to update
			numFields := len(fieldJSONTag)
			if numFields == 0 {
				delete(structsTagsMapper, ts.Name.Name)
				return false
			}
			return true
		},
	)
	if err != nil {
		return err
	}

	// Check if all structs have been updated
	if len(structsTagsMapper) > 0 {
		FieldsNotUpdated(structsTagsMapper, d.logger)
		return fmt.Errorf("failed to update all structs")
	}

	// Write the modified AST back to the file
	return d.parser.WriteFile(filePath, fileSet, node)
}

// HideStructsTags hides the JSON tags from the structs in the specified file
//
// Parameters:
//
//   - filePath string: the path to the Go file
//   - structsFieldsMapper StructsFieldsMapper: the mapper of struct names to their fields
//
// Returns:
//
//   - error: if any error occurs
func (d DefaultWriter) HideStructsTags(
	filePath string,
	structsFieldsMapper StructsFieldsMapper,
) error {
	// Check if the structs fields mapper is nil
	if structsFieldsMapper == nil {
		return ErrNilStructsFieldsMapper
	}

	// Generate the StructJSONTag
	structJSONTagMapper := StructsTagsMapper{}

	// Loop through the struct fields
	for structTypeName, fields := range structsFieldsMapper {
		fieldJSONTagMapper := FieldsTagsMapper{}
		for _, field := range fields {
			fieldJSONTagMapper[field] = "-"
		}
		structJSONTagMapper[structTypeName] = fieldJSONTagMapper
	}
	// Overwrite the JSON tags
	return d.OverwriteTags(filePath, structJSONTagMapper)
}

// HideFilesTags hides the JSON tags from the structs in the specified files
//
// Parameters:
//
//   - filesStructsFieldsMapper FilesStructsFieldsMapper: the mapper of Go file paths to their struct fields
//
// Returns:
//
//   - error: if any error occurs
func (d DefaultWriter) HideFilesTags(
	filesStructsFieldsMapper FilesStructsFieldsMapper,
) error {
	// Check if the FilesStructsFieldsMapper is nil
	if filesStructsFieldsMapper == nil {
		return ErrNilFilesStructsFieldsMapper
	}

	// Loop through the file paths
	for filePath, structsFieldsMapper := range filesStructsFieldsMapper {
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

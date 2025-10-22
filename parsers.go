package goparser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
)

type (
	// DefaultParser is the struct for the default Go parser
	DefaultParser struct{}
)

// NewDefaultParser creates a new default Go parser
//
// Returns:
//
//   - *DefaultParser: the default Go parser
func NewDefaultParser() *DefaultParser {
	return &DefaultParser{}
}

// ParseFile parse the given Go file and return the file set and the AST node
//
// Parameters:
//
//   - filePath string: the path to the Go file
//
// Returns:
//
//   - *token.FileSet: the file set
//   - *ast.File: the AST node
//   - error: if any error occurs
func (d DefaultParser) ParseFile(filePath string) (
	*token.FileSet,
	*ast.File,
	error,
) {
	// Parse the Go file
	fileSet := token.NewFileSet()
	node, err := parser.ParseFile(fileSet, filePath, nil, parser.ParseComments)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse file: %w", err)
	}

	return fileSet, node, nil
}

// WriteFile write the given AST node to the given file path
//
// Parameters:
//
//   - filePath string: the path to the Go file
//   - fileSet *token.FileSet: the file set
//   - node *ast.File: the AST node
//
// Returns:
//
//   - error: if any error occurs
func (d DefaultParser) WriteFile(
	filePath string,
	fileSet *token.FileSet,
	node *ast.File,
) error {
	// Check the file set and the AST node
	if fileSet == nil {
		return ErrNilFileSet
	}
	if node == nil {
		return ErrNilASTNode
	}

	// Write the modified AST back to the file
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	if err = printer.Fprint(file, fileSet, node); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	return nil
}

// TraverseAST traverse the given AST node and call the given function for each node
//
// Parameters:
//
//   - node *ast.File: the AST node
//   - fn func(ast.Node) bool: the function to call for each node
//
// Returns:
//
//   - error: if any error occurs
func (d DefaultParser) TraverseAST(
	node *ast.File,
	fn func(ast.Node) bool,
) error {
	// Check if the node is nil
	if node == nil {
		return ErrNilFileSet
	}

	// Traverse the AST to find the struct and field
	ast.Inspect(
		node, fn,
	)
	return nil
}

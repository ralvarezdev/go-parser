package go_parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
)

// ParseGoFile parse the given Go file and return the file set and the AST node
func ParseGoFile(filePath string) (*token.FileSet, *ast.File, error) {
	// Parse the Go file
	fileSet := token.NewFileSet()
	node, err := parser.ParseFile(fileSet, filePath, nil, parser.ParseComments)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse file: %w", err)
	}

	return fileSet, node, nil
}

// WriteGoFile write the given AST node to the given file path
func WriteGoFile(filePath string, fileSet *token.FileSet, node *ast.File) error {
	// Check the file set and the AST node
	if fileSet == nil {
		return NilFileSetError
	}
	if node == nil {
		return NilASTNodeError
	}

	// Write the modified AST back to the file
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	if err := printer.Fprint(file, fileSet, node); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

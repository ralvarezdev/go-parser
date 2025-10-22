package goparser

import (
	"go/ast"
	"go/token"
)

type (
	// Parser is an interface to parse Go files
	Parser interface {
		ParseFile(filePath string) (*token.FileSet, *ast.File, error)
		WriteFile(filePath string, fileSet *token.FileSet, node *ast.File) error
		TraverseAST(node *ast.File, fn func(ast.Node) bool) error
	}
)

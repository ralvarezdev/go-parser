package go_parser

import (
	"go/ast"
)

// TraverseAST traverse the given AST node and call the given function for each node
func TraverseAST(node *ast.File, fn func(ast.Node) bool) error {
	// Check if the node is nil
	if node == nil {
		return NilFileSetError
	}

	// Traverse the AST to find the struct and field
	ast.Inspect(
		node, fn,
	)
	return nil
}

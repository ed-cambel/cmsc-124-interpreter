package interpreter

import "arkana/pkg/ast"

// Evaluate walks the tree post-order and returns the node's value
func Evaluate(e ast.Expr) any {
	switch n := e.(type) {
	case ast.Literal:
		return n.Value
	case ast.Grouping:
		return Evaluate(n.Expression)
		// TODO: add cases for ast.Unary (-, !) and ast.Binary
	}
	return nil
}

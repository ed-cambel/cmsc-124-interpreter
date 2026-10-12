package interpreter

import "arkana/pkg/ast"

// Evaluate walks the tree post-order and returns the node's value
func Evaluate(e ast.Expr) any {
	switch n := e.(type) {
	case ast.Literal:
		return n.Value
	case ast.Grouping:
		return Evaluate(n.Expression)

	case ast.Unary:
		right := Evaluate(n.Right)

		switch n.Operator.Type {
		case scanner.MINUS:
			checkNumberOperand(n.Operator, right)
			switch x := right.(type) {
			case int:
				return -x
			case float64:
				return -x
			}

		case scanner.NOT:
			return !isTruthy(right)
		}

		panic("Unknown unary operator: " + n.Operator.Lexeme)

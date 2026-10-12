package interpreter

import (
	"arkana/pkg/ast"
	"arkana/pkg/scanner"
)

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

	case ast.Binary:
		left := Evaluate(n.Left)
		right := Evaluate(n.Right)

		switch n.Operator.Type {
		case scanner.OR:
			if isTruthy(left) {
				return left
			}
			return Evaluate(n.Right)
		case scanner.AND:
			if !isTruthy(left) {
				return left
			}
			return Evaluate(n.Right)
		}

		switch n.Operator.Type {
		case scanner.PLUS:
			checkNumberOperands(n.Operator, left, right)
			switch x := right.(type) {
			case int:
				return left.(int) + x
			case float64:
				return left.(float64) + x	// TODO: string concatenation
			}
		case scanner.MINUS:
			checkNumberOperands(n.Operator, left, right)
			switch x := right.(type) {
			case int:
				return left.(int) - x
			case float64:
				return left.(float64) - x
			}
		case scanner.STAR:
			checkNumberOperands(n.Operator, left, right)
			switch x := right.(type) {
			case int:
				return left.(int) * x
			case float64:
				return left.(float64) * x
			}
		case scanner.SLASH: 
			checkNumberOperands(n.Operator, left, right)
			switch x := right.(type) {
			case int:
				if x == 0 {
					panic("Division by zero.") // TODO: handle this later
				}
				return left.(int) / x
			case float64:
				if x == 0 {
					panic("Division by zero.")
				}
				return left.(float64) / x
			}
		case scanner.PERCENT:
			checkNumberOperands(n.Operator, left, right) 
				switch x := right.(type) {
				case int:
					if x == 0 {
						panic("Modulo by zero.") // TODO: handle this later
					}
					return left.(int) % x
				}
			

		case scanner.LESS:
			checkNumberOperands(n.Operator, left, right)
			switch x := right.(type) {
			case int:
				return left.(int) < x
			case float64:
				return left.(float64) < x
			}
		case scanner.EQUAL_EQUAL:
			return isEqual(left, right)
		
		default:
			panic("Unknown binary operator: " + n.Operator.Lexeme)
		}
	
		panic("Unknown expression type.")
	} 
	panic("Unknown expression type.")
}
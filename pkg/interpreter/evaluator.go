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

		// && and || check since it decides from left value before touchign the right
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

		right := Evaluate(n.Right) // every other operator needs both sides

		leftInt, leftIsInt := left.(int)
		rightInt, rightIsInt := right.(int)
		bothInt := leftIsInt && rightIsInt
		leftFloat, isLeftNum := toFloat(left)
		rightFloat, isRightNum := toFloat(right)

		switch n.Operator.Type {
		case scanner.PLUS:
			leftStringOp, isTrueLeft := left.(string)
			rightStringOp, isTrueRight := right.(string)
			
			// returns concatenated string if both are string 
			if isTrueLeft && isTrueRight {
				return  leftStringOp + rightStringOp
			} 

			if !isLeftNum || !isRightNum {
				panic(RuntimeError{Token: n.Operator, Message: "Operands must be two numbers or two strings."})
			}

			// int + int stays int
			if bothInt {
				return leftInt + rightInt
			}
			// any float involved: both sides as float
			return leftFloat + rightFloat

		case scanner.MINUS:
			checkNumberOperands(n.Operator, left, right)

			if bothInt {
				return leftInt - rightInt
			} 

			return leftFloat - rightFloat

		case scanner.STAR:
			checkNumberOperands(n.Operator, left, right)
		
			if bothInt {
				return leftInt * rightInt
			}
			return leftFloat * rightFloat

		case scanner.SLASH:

			checkNumberOperands(n.Operator, left, right)
			if rightFloat == 0 {
				panic(RuntimeError{Token: n.Operator, Message: "Division by zero."})
			}
			return leftFloat / rightFloat // int / int promotes to float

		case scanner.PERCENT:
			checkNumberOperands(n.Operator, left, right)
			switch x := right.(type) {
			case int:
				if x == 0 {
					panic(RuntimeError{Token: n.Operator, Message: "Modulo by zero."})
				}
				return left.(int) % x
			}

		case scanner.LESS, scanner.LESS_EQUAL, scanner.GREATER, scanner.GREATER_EQUAL:
			checkNumberOperands(n.Operator, left, right)
			leftOp, _ := toFloat(left)
			rightOp, _ := toFloat(right)

			switch n.Operator.Type {
			case scanner.LESS: 
				return leftOp < rightOp
			case scanner.LESS_EQUAL:
				return leftOp <= rightOp
			case scanner.GREATER:
				return leftOp > rightOp
			default:
				return leftOp >= rightOp
			}
		case scanner.EQUAL_EQUAL:
			return isEqual(left, right)
		case scanner.NOT_EQUAL:
			return !isEqual(left, right)
		default:
			panic("Unknown binary operator: " + n.Operator.Lexeme)
		}

		panic("Unknown expression type.")
	}
	panic("Unknown expression type.")
}

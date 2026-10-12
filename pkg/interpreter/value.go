package interpreter

import (
	"arkana/pkg/ast"
	"arkana/pkg/scanner"
)

type Evaluator struct {
	tokens  []scanner.Token
	message string
}

// Helper Methods

// Stringify is the one place a runtime value becomes text
func Stringify(v any) string {
	return ast.FormatLiteral(v)
}

// isTruthy is the only truthiness rule: bane and nil are falsey
func isTruthy(v any) bool {
	switch x := v.(type) {
	case nil:	
		return false
	case bool:
		return x
	}
	return true
}

// toFloat is used to convert int to float when necessary
func toFloat(v any) (float64, bool) {
	switch x := v.(type) {
	case int:
		return float64(x), true
	case float64:
		return x, true
	}
	// if value isn't a supported number, report
	return 0, false
}

// isEqual works on any pair of types
// Numbers compare by value across int and float; values of different kinds are never equal.

func isEqual(a, b any) bool {
	if af, ok := toFloat(a); ok {
		bf, ok := toFloat(b)
		return ok && af == bf // makes 5 == 5.0 true since it converts 5 to 5.0
	}
	return a == b
}

// for unary function
func checkNumberOperand(op scanner.Token, n any) {
	if _, ok := toFloat(n); !ok {
		panic("Operand must be a number.")
	}
}

func checkNumberOperands(op scanner.Token, left, right any) {
	if _, ok := toFloat(left); !ok {
		panic("Left operand must be a number.")
	}
	if _, ok := toFloat(right); !ok {
		panic("Right operand must be a number.")
	}
}


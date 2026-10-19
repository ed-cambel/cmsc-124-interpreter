package interpreter

import (
	"arkana/pkg/scanner"
	"fmt"
	"strconv"
	"strings"
)

type Evaluator struct {
	tokens  []scanner.Token
	message string
}

// Helper Methods

// a runtime value becomes text
func Stringify(v any) string {
	switch x := v.(type) {
	case nil:
		return "nil"
	case bool:
		if x {
			return "boon"
		}
		return "bane"
	case float64:
		s := strconv.FormatFloat(x, 'f', -1, 64) // shortest form that roundtrips (at most 17 digits)
		if !strings.Contains(s, ".") {
			s += ".0" // numbers never trim .0
		}
		return s
	case int:
		return strconv.Itoa(x)
	case string:
		return x // no quotes
	}
	return fmt.Sprint(v)
}

// isTruthy is the only truthiness rule: bane, 0, and nil are falsey
func isTruthy(v any) bool {
	switch x := v.(type) {
	case nil:
		return false
	case bool:
		return x
	case int:
		return x != 0
	case float64:
		return x != 0 // -0.0 == 0.0 in Go, so both are falsey
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
// numbers compare by value across int and float; values of different kinds are never equal.
func isEqual(a, b any) bool {
	if ai, ok := a.(int); ok {
		if bi, ok := b.(int); ok {
			return ai == bi
		}
	}

	if af, ok := toFloat(a); ok {
		bf, ok := toFloat(b)
		return ok && af == bf // makes 5 == 5.0 true since it converts 5 to 5.0; a number vs a non-number is false
	}
	return a == b // for non-numbers (nil, bool, string) compares by type and value, not by reference
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

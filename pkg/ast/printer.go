package ast

import (
	"fmt"
	"strconv"
	"strings"
)

// switch-case is used for the "pattern match over node types" option, which replaces the visitor pattern
// print returns the prefix parenthesized form of a tree, e.g. (+ 2.0 (* 3.0 4.0))
func Print(e Expr) string {
	switch n := e.(type) {
	case Literal:
		return formatLiteral(n.Value)
	case Grouping:
		return "(group " + Print(n.Expression) + ")"
	case Unary:
		return "(" + n.Operator.Lexeme + " " + Print(n.Right) + ")"
	case Binary:
		return "(" + n.Operator.Lexeme + " " + Print(n.Left) + " " + Print(n.Right) + ")"
	}

	return "<unknown>"
}

// numbers always print with a decimal point so 2 and 2.0 are one string
func formatLiteral(v any) string {
	switch x := v.(type) {
	case nil:
		return "nil"
	case float64:
		s := strconv.FormatFloat(x, 'f', -1, 64)
		if !strings.Contains(s, ".") {
			s += ".0"
		}
		return s
	case int:
		// the scanner returns int for whole numbers
		return fmt.Sprintf("%d.0", x)
	default:
		// bool and string
		return fmt.Sprint(x)
	}
}

package ast

import (
	"fmt"
	"strconv"
	"strings"
)

// switch-case is used for the "pattern match over node types" option, which replaces the visitor pattern
// Print returns the prefix parenthesized form of a tree, e.g. (+ 2 (* 3.0 4)) for the source 2 + 3.0 * 4
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
	case Variable:
		return n.Name.Lexeme
	case Assign:
		return "(-> " + n.Name.Lexeme + " " + Print(n.Value) + ")"
	case CallExpr:
		s := "(call " + Print(n.Callee)
		for _, arg := range n.Args {
			s += " " + Print(arg)
		}
		return s + ")"
	}

	return "<unknown>"
}

func formatLiteral(v any) string {
	switch x := v.(type) {
	case nil:
		return "nil"
	case bool:
		if x {
			return "boon"
		}
		return "bane"
	case float64:
		s := strconv.FormatFloat(x, 'f', -1, 64)
		if !strings.Contains(s, ".") {
			s += ".0"
		}
		return s
	case int:
		// the scanner returns int for whole numbers
		return strconv.Itoa(x)
	default:
		// string
		return fmt.Sprint(x)
	}
}

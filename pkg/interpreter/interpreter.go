package interpreter

import (
	"arkana/pkg/ast"
	"fmt"
)

// evaluates one expression, used by --eval and the REPL
func Interpret(e ast.Expr) (value any, err error) {
	defer catch(&err)
	return Evaluate(e), nil
}

// runs one statement, used by plain ./run
func Execute(s ast.Stmt) (err error) {
	defer catch(&err)

	switch n := s.(type) {
	case ast.ExprStmt:
		Evaluate(n.Expression) // value is discarded
	case ast.RevealStmt:
		fmt.Println(Stringify(Evaluate(n.Expression)))
	default:
		panic("Unknown statement type.")
	}
	return nil
}

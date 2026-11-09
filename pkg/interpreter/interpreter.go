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
	case ast.IfStmt:
		if isTruthy(Evaluate(n.Condition)) {
			return Execute(n.Then)
		} else if n.Else != nil {
			return Execute(n.Else)
		}
	case ast.BlockStmt:
		for _, stmt := range n.Statements {
			if err := Execute(stmt); err != nil {
				return err
			}
		}
	default:
		panic("Unknown statement type.")
	}
	return nil
}

package interpreter

import (
	"arkana/pkg/ast"
	"fmt"
)

type Interpreter struct {
	env *Environment
}

func New() *Interpreter {
	return &Interpreter{env: NewEnvironment(nil)}
}

// evaluates one expression, used by --eval and the REPL
func (in *Interpreter) Interpret(e ast.Expr) (value any, err error) {
	defer catch(&err)
	return Evaluate(e), nil
}

// runs one statement, used by plain ./run
func (in *Interpreter)Execute(s ast.Stmt) (err error) {
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

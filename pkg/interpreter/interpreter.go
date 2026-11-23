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
	return in.Evaluate(e), nil
}

// runs one statement, used by plain ./run
func (in *Interpreter) Execute(s ast.Stmt) (err error) {
	defer catch(&err)
	in.exec(s)
	return nil
}

func (in *Interpreter) exec(s ast.Stmt) {
	switch n := s.(type) {
	case ast.ExprStmt:
		in.Evaluate(n.Expression) // value is discarded
	case ast.RevealStmt:
		fmt.Println(Stringify(in.Evaluate(n.Expression)))
	case ast.IfStmt:
		if isTruthy(in.Evaluate(n.Condition)) {
			in.exec(n.Then)
		} else if n.Else != nil {
			in.exec(n.Else)
		}
	case ast.SpellStmt:
		in.env.Define(n.Name, kindSpell, true, &Spell{decl: &n, closure: in.env})
	case ast.VarStmt:
		var value any
		if n.Initializer != nil {
			value = in.Evaluate(n.Initializer)
		}
		in.env.Define(n.Name, declaredKind(n.Type.Type), n.IsSeal, value)
	case ast.BlockStmt:
		in.executeBlock(n.Statements, NewEnvironment(in.env))
	case ast.RitualStmt:
		for isTruthy(in.Evaluate(n.Condition)) {
			in.exec(n.Body)
		}
	default:
		panic("Unknown statement type.")
	}
}

func (in *Interpreter) executeBlock(stmts []ast.Stmt, env *Environment) {
	previous := in.env
	in.env = env
	defer func() { in.env = previous }()

	for _, s := range stmts {
		in.exec(s)
	}
}

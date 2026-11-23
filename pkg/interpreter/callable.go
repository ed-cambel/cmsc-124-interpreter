package interpreter

import (
	"arkana/pkg/ast"
	
)

type Callable interface {
	Arity() int                           // number of parameters it expects
	Call(in *Interpreter, args []any) any // runs the body and returns its value (nil if none)
}

// user-defined functions
type Spell struct {
	decl 		*ast.SpellStmt
	closure		*Environment
}

func (f *Spell) Arity() int { return len(f.decl.Params) }
func (f *Spell) Call(in *Interpreter, args []any) any {
    env := NewEnvironment(f.closure)

	for i, p := range f.decl.Params {
        env.Define(p.Name, declaredKind(p.Type.Type), false, args[i])
    }

    in.executeBlock(f.decl.Body, env)
    return nil
} 
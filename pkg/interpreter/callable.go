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
	decl    *ast.SpellStmt
	closure *Environment
}

func (f *Spell) Arity() int { return len(f.decl.Params) }
func (f *Spell) Call(in *Interpreter, args []any) (result any) {
	env := NewEnvironment(f.closure)

	for i, p := range f.decl.Params {
		env.Define(p.Name, declaredKind(p.Type.Type), false, args[i])
	}

	defer func() {
		r := recover()
		if r == nil {
			return // fell off the end of the body, so the result stays nil
		}

		signal, isReturn := r.(YieldSignal)
		if !isReturn {
			panic(r) // a RuntimeError or a bug, not ours to handle; keep unwinding
		}

		result = signal.Value
	}()

	in.executeBlock(f.decl.Body, env)
	return nil
}

// carries a function's result up through every enclosing statement to the call running it
type YieldSignal struct {
	Value any
}

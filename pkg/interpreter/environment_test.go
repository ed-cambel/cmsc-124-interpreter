package interpreter

import (
	"testing"

	"arkana/pkg/ast"
	"arkana/pkg/scanner"
)

func tok(name string) scanner.Token {
	return scanner.Token{Lexeme: name, Line: 1}
}

// fails unless f panics with a RuntimeError carrying exactly want
func expectRuntimeError(t *testing.T, want string, f func()) {
	t.Helper()
	defer func() {
		r := recover()
		re, ok := r.(RuntimeError)
		if !ok {
			t.Fatalf("expected RuntimeError, got %v", r)
		}
		if re.Message != want {
			t.Errorf("message: got %q, want %q", re.Message, want)
		}
	}()
	f()
}

func TestGetWalksOutward(t *testing.T) {
	outer := NewEnvironment(nil)
	outer.Define(tok("y"), kindGrain, false, 1)
	inner := NewEnvironment(outer)

	if got := inner.Get(tok("y")); got != 1 {
		t.Errorf("got %v, want 1", got)
	}
}

func TestInnerShadowsOuter(t *testing.T) {
	outer := NewEnvironment(nil)
	outer.Define(tok("x"), kindGrain, false, 1)
	inner := NewEnvironment(outer)
	inner.Define(tok("x"), kindGrain, false, 2)

	if got := inner.Get(tok("x")); got != 2 {
		t.Errorf("inner: got %v, want 2", got)
	}
	if got := outer.Get(tok("x")); got != 1 {
		t.Errorf("outer: got %v, want 1", got)
	}
}

func TestAssignWalksOutward(t *testing.T) {
	outer := NewEnvironment(nil)
	outer.Define(tok("x"), kindGrain, false, 1)
	inner := NewEnvironment(outer)

	if got := inner.Assign(tok("x"), 5); got != 5 {
		t.Errorf("assign result: got %v, want 5", got)
	}
	if got := outer.Get(tok("x")); got != 5 {
		t.Errorf("outer: got %v, want 5", got)
	}
}

func TestAssignNeverCreates(t *testing.T) {
	env := NewEnvironment(nil)

	expectRuntimeError(t, "Undefined variable 'x'.", func() { env.Assign(tok("x"), 1) })
	expectRuntimeError(t, "Undefined variable 'x'.", func() { env.Get(tok("x")) })
}

func TestNilBindingIsFound(t *testing.T) {
	env := NewEnvironment(nil)
	env.Define(tok("x"), kindGrain, false, nil)

	if got := env.Get(tok("x")); got != nil {
		t.Errorf("got %v, want nil", got)
	}
}

func TestRedeclareInSameScopeFails(t *testing.T) {
	env := NewEnvironment(nil)
	env.Define(tok("x"), kindGrain, false, 1)

	expectRuntimeError(t, "Variable 'x' is already declared in this scope.", func() { env.Define(tok("x"), kindGrain, false, 2) })
	if got := env.Get(tok("x")); got != 1 {
		t.Errorf("failed redeclaration changed the value: got %v, want 1", got)
	}
}

func TestSameNameInInnerScopeIsAllowed(t *testing.T) {
	outer := NewEnvironment(nil)
	outer.Define(tok("x"), kindGrain, false, 1)

	NewEnvironment(outer).Define(tok("x"), kindGrain, false, 2) // must not panic
}

func TestGrainPromotesToDram(t *testing.T) {
	env := NewEnvironment(nil)
	env.Define(tok("d"), kindDram, false, 1) // initializer promotes

	if got := env.Get(tok("d")); got != float64(1) {
		t.Errorf("define: got %v (%T), want float64 1", got, got)
	}
	if got := env.Assign(tok("d"), 2); got != float64(2) {
		t.Errorf("assign result: got %v (%T), want float64 2", got, got)
	}
}

func TestDramToGrainFails(t *testing.T) {
	env := NewEnvironment(nil)
	env.Define(tok("g"), kindGrain, false, 1)

	expectRuntimeError(t, "Type mismatch for variable 'g'.", func() { env.Assign(tok("g"), 1.5) })
}

func TestMismatchIsAnError(t *testing.T) {
	env := NewEnvironment(nil)
	env.Define(tok("x"), kindGrain, false, 1)

	expectRuntimeError(t, "Type mismatch for variable 'x'.", func() { env.Assign(tok("x"), "a") })
	if got := env.Get(tok("x")); got != 1 {
		t.Errorf("failed assign changed the value: got %v, want 1", got)
	}
}

func TestNilFitsAnyType(t *testing.T) {
	kinds := map[string]kind{"g": kindGrain, "d": kindDram, "s": kindScroll}

	for name, k := range kinds {
		env := NewEnvironment(nil)
		env.Define(tok(name), k, false, nil) // nil initializer

		if got := env.Assign(tok(name), nil); got != nil {
			t.Errorf("%s: got %v, want nil", name, got)
		}
	}
}

func TestSealCannotBeReassigned(t *testing.T) {
	env := NewEnvironment(nil)
	env.Define(tok("s"), kindNil, true, 1)

	expectRuntimeError(t, "Cannot reassign seal 's'.", func() { env.Assign(tok("s"), 2) })
	if got := env.Get(tok("s")); got != 1 {
		t.Errorf("failed reassign changed the value: got %v, want 1", got)
	}
}

func TestRestoreAfterError(t *testing.T) {
	in := New()
	global := in.env

	block := ast.BlockStmt{Statements: []ast.Stmt{
		ast.ExprStmt{Expression: ast.Variable{Name: tok("missing")}},
	}}

	if err := in.Execute(block); err == nil {
		t.Fatal("expected a runtime error")
	}
	if in.env != global {
		t.Error("environment was not restored after the error")
	}

}
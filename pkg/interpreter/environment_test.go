package interpreter

import (
	"testing"

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
	outer.Define(tok("y"), 1)
	inner := NewEnvironment(outer)

	if got := inner.Get(tok("y")); got != 1 {
		t.Errorf("got %v, want 1", got)
	}
}

func TestInnerShadowsOuter(t *testing.T) {
	outer := NewEnvironment(nil)
	outer.Define(tok("x"), 1)
	inner := NewEnvironment(outer)
	inner.Define(tok("x"), 2)

	if got := inner.Get(tok("x")); got != 2 {
		t.Errorf("inner: got %v, want 2", got)
	}
	if got := outer.Get(tok("x")); got != 1 {
		t.Errorf("outer: got %v, want 1", got)
	}
}

func TestAssignWalksOutward(t *testing.T) {
	outer := NewEnvironment(nil)
	outer.Define(tok("x"), 1)
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
	env.Define(tok("x"), nil)

	if got := env.Get(tok("x")); got != nil {
		t.Errorf("got %v, want nil", got)
	}
}

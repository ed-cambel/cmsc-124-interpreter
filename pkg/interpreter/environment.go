package interpreter

import "arkana/pkg/scanner"

type binding struct {
	value  any
	typ    kind
	isSeal bool
}

// Environment holds the bindings of one scope and a link to the scope enclosing it
type Environment struct {
	values    map[string]*binding
	enclosing *Environment // global environment has enclosing == nil
}

func NewEnvironment(enclosing *Environment) *Environment {
	return &Environment{values: make(map[string]*binding), enclosing: enclosing}
}

// creates a binding in this scope only. it never walks outward
// redeclaring in the same scope is an error, while an inner scope may reuse the name (shadowing)
func (e *Environment) Define(name scanner.Token, typ kind, isSeal bool, value any) {
	if _, declared := e.values[name.Lexeme]; declared {
		panic(RuntimeError{Token: name, Message: "Variable '" + name.Lexeme + "' is already declared in this scope."})
	}

	if isSeal {
		typ = kindOf(value)
	} else {
		v, ok := coerceKind(typ, value)
		if !ok {
			panic(typeMismatch(name))
		}

		value = v
	}

	e.values[name.Lexeme] = &binding{value: value, typ: typ, isSeal: isSeal}
}

// walks outward until some scope has the name. the innermost match wins
func (e *Environment) Get(name scanner.Token) any {
	// the comma-ok form matters: a binding holding nil must still count as found
	b, ok := e.values[name.Lexeme]
	if !ok {
		if e.enclosing == nil {
			panic(undefinedVariable(name))
		}
		return e.enclosing.Get(name) // not in this scope, keep walking outward
	}
	return b.value
}

// updates an existing binding, searching outward. it never creates one
// returns the assigned value, since assignment is an expression
func (e *Environment) Assign(name scanner.Token, value any) any {
	b, ok := e.values[name.Lexeme]
	if !ok {
		if e.enclosing == nil {
			panic(undefinedVariable(name))
		}
		return e.enclosing.Assign(name, value) // not in this scope, keep walking outward
	}

	if b.isSeal {
		panic(RuntimeError{Token: name, Message: "Cannot reassign seal '" + name.Lexeme + "'."})
	}

	v, ok := coerceKind(b.typ, value)
	if !ok {
		panic(typeMismatch(name))
	}

	b.value = v
	return v
}

func undefinedVariable(name scanner.Token) RuntimeError {
	return RuntimeError{Token: name, Message: "Undefined variable '" + name.Lexeme + "'."}
}

func typeMismatch(name scanner.Token) RuntimeError {
	return RuntimeError{Token: name, Message: "Type mismatch for variable '" + name.Lexeme + "'."}
}

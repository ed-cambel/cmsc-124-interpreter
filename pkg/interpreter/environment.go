package interpreter

import "arkana/pkg/scanner"

// Environment holds the bindings of one scope and a link to the scope enclosing it
type Environment struct {
	values    map[string]any
	enclosing *Environment // global environment has enclosing == nil
}

func NewEnvironment(enclosing *Environment) *Environment {
	return &Environment{values: make(map[string]any), enclosing: enclosing}
}

// creates a binding in this scope only. it never walks outward
func (e *Environment) Define(name scanner.Token, value any) {
	e.values[name.Lexeme] = value
}

// walks outward until some scope has the name. the innermost match wins
func (e *Environment) Get(name scanner.Token) any {
	// the comma-ok form matters: a binding holding nil must still count as found
	if value, ok := e.values[name.Lexeme]; ok {
		return value
	}
	if e.enclosing != nil {
		return e.enclosing.Get(name)
	}
	panic(undefinedVariable(name))
}

// updates an existing binding, searching outward. it never creates one
// returns the assigned value, since assignment is an expression
func (e *Environment) Assign(name scanner.Token, value any) any {
	if _, ok := e.values[name.Lexeme]; ok {
		e.values[name.Lexeme] = value
		return value
	}
	if e.enclosing != nil {
		return e.enclosing.Assign(name, value)
	}
	panic(undefinedVariable(name))
}

func undefinedVariable(name scanner.Token) RuntimeError {
	return RuntimeError{Token: name, Message: "Undefined variable '" + name.Lexeme + "'."}
}

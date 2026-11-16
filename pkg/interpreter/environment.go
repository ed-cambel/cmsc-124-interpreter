package interpreter

import "arkana/pkg/scanner"

type binding struct {
	value 	any
	typ   	kind
	isSeal	bool
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
	if _, ok := e.values[name.Lexeme]; ok {
		panic(RuntimeError{Token: name, Message: "Variable '" + name.Lexeme + "' is already declared in this scope."})
	}
	
	if isSeal {
		typ = kindOf(value)
	} else {
		v, ok := coerceKind(typ, kindOf(value))
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
	if b, ok := e.values[name.Lexeme]; ok {
		if b.isSeal {
			panic(RuntimeError{Token:name, Message: "Cannot reassign seal '" + name.Lexeme + "."})
		}

		v, ok := coerceKind(b.typ, kindOf(value))
		if !ok{
			panic(typeMismatch(name))
		}
		b.value = v
		return v
	}
	if e.enclosing != nil {
		return e.enclosing.Assign(name, value)
	}
	panic(undefinedVariable(name))
}

func undefinedVariable(name scanner.Token) RuntimeError {
	return RuntimeError{Token: name, Message: "Undefined variable '" + name.Lexeme + "'."}
}

func typeMismatch(name scanner.Token) RuntimeError {
	return RuntimeError{Token: name, Message: "Type mismatch for variable '" + name.Lexeme + "."}
}
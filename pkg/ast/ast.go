// defines the structure of AST

package ast

import "arkana/pkg/scanner"

// ast expression interface: common type for all expression nodes
// Literal, Unary, Binary, Grouping, Variable, Assign
type Expr interface {
	expr()
}

type Literal struct {
	Value any
}

func (Literal) expr() {}

type Unary struct {
	Operator scanner.Token
	Right    Expr
}

func (Unary) expr() {}

type Binary struct {
	Left     Expr
	Operator scanner.Token
	Right    Expr
}

func (Binary) expr() {}

type Grouping struct {
	Expression Expr
}

func (Grouping) expr() {}

// Variable is a read of a named binding, e.g. the x in `reveal x;`
// Name is kept as a token (not a string) so runtime errors can report the line
type Variable struct {
	Name scanner.Token
}

func (Variable) expr() {}

// Assign writes to an existing binding, e.g. `x -> 5`. It is an expression, so it produces the assigned value.
// Name is the target's identifier token; Value is the right-hand side.
type Assign struct {
	Name  scanner.Token
	Value Expr
}

func (Assign) expr() {}

// Tree is the root of a whole file, its expressions in source order
// not an Expr, so no expr() method
type Tree struct {
	ASTs []Expr // the root's children, one AST per valid statement
}

type CallExpr struct {
	Callee Expr
	Paren  scanner.Token // closing ')', used for error line numbers
	Args   []Expr
}

func (CallExpr) expr() {}
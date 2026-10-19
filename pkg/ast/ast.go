// defines the structure of AST

package ast

import "arkana/pkg/scanner"

// ast expression interface: common type for all expression nodes (Literal, Unary, Binary, Grouping)
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

// Tree is the root of a whole file, its expressions in source order
// not an Expr, so no expr() method
type Tree struct {
	ASTs []Expr // the root's children, one AST per valid statement
}

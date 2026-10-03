// defines the structure of AST

package ast

import "arkana/pkg/scanner"

// ast expression interface: common type for all future expression nodes (Literal, Unary, Binary, Grouping)
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

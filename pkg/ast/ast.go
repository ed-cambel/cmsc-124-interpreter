// defines the structure of AST

package ast

// ast expression interface: common type for all future expression nodes (Literal, Unary, Binary, Grouping)
type Expr interface {
	expr()
}

type Literal struct {
	Value any
}

func (Literal) expr() {}

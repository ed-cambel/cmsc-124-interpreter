package ast

// Stmt is a node that does something instead of producing a value
// (ExprStmt, RevealStmt). Separate from Expr so the two can't be mixed up.
type Stmt interface {
	stmt()
}

// ExprStmt evaluates an expression and discards the result
type ExprStmt struct {
	Expression Expr
}

func (ExprStmt) stmt() {}

// RevealStmt evaluates an expression and prints the result
type RevealStmt struct {
	Expression Expr
}

func (RevealStmt) stmt() {}
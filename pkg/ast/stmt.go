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

// IfStmt evaluates a condition and executes one of two statements
type IfStmt struct {
	Condition Expr
	Then      Stmt
	Else      Stmt
}

func (IfStmt) stmt() {}

// BlockStmt groups multiple statements together inside curly braces
type BlockStmt struct {
	Statements []Stmt
}

func (BlockStmt) stmt() {}

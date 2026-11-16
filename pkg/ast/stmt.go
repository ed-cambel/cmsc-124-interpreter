package ast

import "arkana/pkg/scanner"

// Stmt is a node that does something instead of producing a value
// VarStmt, ExprStmt, RevealStmt, BlockStmt. Separate from Expr
type Stmt interface {
	stmt()
}

// VarStmt introduces a binding in the current scope: `grain x -> 5;`, `dram y;`, `seal z -> 1;`
type VarStmt struct {
	Name        scanner.Token
	Type        scanner.Token // type keyword (GRAIN, DRAM, SCROLL, RUNE); meaningless when IsSeal is true
	Initializer Expr          // nil when there is no initializer; the binding then holds nil
	IsSeal      bool          // true for a constant declaration; its type is inferred from the initializer
}

func (VarStmt) stmt() {}

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

// BlockStmt groups declarations and statements into their own nested scope
type BlockStmt struct {
	Statements []Stmt // declarations are also Stmts (VarStmt), in source order
}

func (BlockStmt) stmt() {}

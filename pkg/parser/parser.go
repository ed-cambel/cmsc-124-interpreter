package parser

import (
	"arkana/pkg/ast"
	"arkana/pkg/scanner"
)

type Parser struct {
	tokens  []scanner.Token
	current int
}

// Constructor
func NewParser(tokens []scanner.Token) *Parser {
	return &Parser{
		tokens: tokens,
	}
}

func (p *Parser) Parse() ast.Expr {
	return p.primary()
}

func (p *Parser) primary() ast.Expr {
	token := p.tokens[p.current] // token := p.advance()

	if token.Type == scanner.BOON {
		p.current++ // remove this in the future
		return ast.Literal{
			Value: token.Literal,
		}
	}

	return nil
}

// Helper Methods

//peek()
//advance()
//check()
//match()
//consume()
//previous()

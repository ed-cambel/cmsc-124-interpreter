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
	return p.factor()
}

func (p *Parser) factor() ast.Expr {
	// TODO: Add unary parsing if unary ops are included
	return p.primary()
}

func (p *Parser) primary() ast.Expr {
	token := p.advance()

	if token.Type == scanner.BOON {
		return ast.Literal{
			Value: token.Literal,
		}
	}

	if token.Type == scanner.NUMBER {
		return ast.Literal{
			Value: token.Literal,
		}
	}

	return nil
}

// Helper Methods

func (p *Parser) advance() scanner.Token {
	token := p.tokens[p.current]

	if !p.isAtEnd() {
		p.current++
	}

	return token
}

func (p *Parser) peek() scanner.Token {
	return p.tokens[p.current]
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == scanner.EOF
}

// TODO: implement other helper methods
// match()
// previous()
// check()
// consume()

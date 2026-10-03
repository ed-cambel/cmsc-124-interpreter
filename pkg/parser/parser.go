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

func (p *Parser) unary() ast.Expr {
	if p.match(scanner.NOT, scanner.MINUS) {
		operator := p.previous()
		right := p.unary()
		return ast.Unary{
			Operator: operator,
			Right:    right,
		}
	}

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

func (p *Parser) check(expected scanner.TokenType) bool {
	if p.isAtEnd() {
		return false
	}

	return p.peek().Type == expected
}

func (p *Parser) previous() scanner.Token {
	return p.tokens[p.current-1]
}

func (p *Parser) match(expected ...scanner.TokenType) bool {
	for _, tokenType := range expected {
		if p.check(tokenType) {
			p.advance()
			return true
		}
	}
	return false
}


func (p *Parser) consume(expected scanner.TokenType, message string) scanner.Token {
	if p.check(expected) {
		return p.advance()
	} 
	
	// TODO: implement error handling rawr
	panic(message)
}
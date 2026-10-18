package parser

import (
	"fmt"
	"os"

	"arkana/pkg/ast"
	"arkana/pkg/scanner"
)

type parseAbort struct{}

type Parser struct {
	tokens  []scanner.Token
	current int
	IsError bool
}

// Constructor
func NewParser(tokens []scanner.Token) *Parser {
	return &Parser{
		tokens: tokens,
	}
}

// reads the whole token list as a sequence of expressions each ending in semicolon
func (p *Parser) Parse() []ast.Expr {
	exprs := []ast.Expr{}

	for !p.isAtEnd() {
		if expr := p.parseExpression(); expr != nil {
			exprs = append(exprs, expr)
		}
	}

	return exprs
}

func (p *Parser) expression() ast.Expr {
	return p.or()
}

func (p *Parser) parseExpression() (expr ast.Expr) {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(parseAbort); !ok {
				panic(r) 
			}
			p.synchronize()
			expr = nil
		}
	}()

	expr = p.expression()
	p.consume(scanner.SEMICOLON, "Expect ';' after expression.")
	return expr
}

func (p *Parser) or() ast.Expr {
	expr := p.and()

	for p.match(scanner.OR) {
		operator := p.previous()
		right := p.and()
		expr = ast.Binary{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}

	return expr
}

func (p *Parser) and() ast.Expr {
	expr := p.equality()

	for p.match(scanner.AND) {
		operator := p.previous()
		right := p.equality()
		expr = ast.Binary{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}

	return expr
}

func (p *Parser) equality() ast.Expr {
	expr := p.comparison()

	for p.match(scanner.EQUAL_EQUAL, scanner.NOT_EQUAL) {
		operator := p.previous()
		right := p.comparison()
		expr = ast.Binary{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}

	return expr
}
func (p *Parser) comparison() ast.Expr {
	expr := p.term()

	for p.match(scanner.LESS, scanner.LESS_EQUAL, scanner.GREATER, scanner.GREATER_EQUAL) {
		operator := p.previous()
		right := p.term()
		expr = ast.Binary{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}

	return expr
}

func (p *Parser) term() ast.Expr {
	expr := p.factor()

	for p.match(scanner.PLUS, scanner.MINUS) {
		operator := p.previous()
		right := p.factor()
		expr = ast.Binary{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}

	return expr
}

func (p *Parser) factor() ast.Expr {
	expr := p.unary() // left operand

	for p.match(scanner.STAR, scanner.SLASH, scanner.PERCENT) {
		operator := p.previous()
		right := p.unary()
		expr = ast.Binary{
			Left:     expr, //the old tree becomes the left child
			Operator: operator,
			Right:    right,
		}
	}

	return expr
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

	if p.match(scanner.NUMBER, scanner.STRING, scanner.BOON, scanner.BANE, scanner.NIL) {
		return ast.Literal{
			Value: p.previous().Literal,
		}
	}

	if p.match(scanner.LEFT_PAREN) {
		expr := p.or() //change to p.expression() once that exists
		p.consume(scanner.RIGHT_PAREN, "Expect ')' after expression.")
		return ast.Grouping{Expression: expr}
	}

	p.ParseError(p.peek(), "Expect expression")
	panic(parseAbort{})
}

// Helper Methods

// returns current token and moves forward one position (consumes it) but never moves past EOF
func (p *Parser) advance() scanner.Token {
	token := p.tokens[p.current]

	if !p.isAtEnd() {
		p.current++
	}

	return token
}

// returns current token without consuming (lookahead only)
func (p *Parser) peek() scanner.Token {
	return p.tokens[p.current]
}

// reports whether current token is EOF
func (p *Parser) isAtEnd() bool {
	return p.peek().Type == scanner.EOF
}

// reports whether current token is of the expected type, w/o consuming (false at EOF)
func (p *Parser) check(expected scanner.TokenType) bool {
	if p.isAtEnd() {
		return false
	}

	return p.peek().Type == expected
}

// returns the most recently consumed token (the one just before current)
func (p *Parser) previous() scanner.Token {
	return p.tokens[p.current-1]
}

// consumes the current token if it matches any of the given types and returns true; otherwise, consumes nothing and returns false
func (p *Parser) match(expected ...scanner.TokenType) bool {
	for _, tokenType := range expected {
		if p.check(tokenType) {
			p.advance()
			return true
		}
	}
	return false
}

// requires current token to be the expected type: it consumes and returns if it so; otherwise, reports an error
func (p *Parser) consume(expected scanner.TokenType, message string) scanner.Token {
	if p.check(expected) {
		return p.advance()
	}

	p.ParseError(p.peek(), message)
	panic(parseAbort{})
}

func (p *Parser) ParseError(token scanner.Token, message string) {
    if p.isAtEnd() {
		p.errorMessage(token.Line, " at the end", message)
	} else {
		p.errorMessage(token.Line,  " at '"+ token.Lexeme +"'", message)
	}
}

// prints error message and sets p.IsError 
func (p *Parser) ParseError(token scanner.Token, message string) {
    where := " at '" + token.Lexeme + "'"
	if p.isAtEnd() {
		where = " at the end"
	}
	fmt.Fprintf(os.Stderr, "[line %d] Error%s: %s\n", token.Line, where, message)
	p.IsError = true
}

func (p *Parser) synchronize() {
	if !p.isAtEnd() {
		p.advance() // always make progress
	}
	for !p.isAtEnd() {
		if p.previous().Type == scanner.SEMICOLON {
			return
		}
		p.advance()
	}
}

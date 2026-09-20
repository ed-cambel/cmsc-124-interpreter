package scanner

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Scanner struct {
	source  string
	start   int
	current int
	line    int
	IsError	bool
}

// Constructor
func NewScanner(source string) *Scanner {
	return &Scanner{
		source: source,
		line:   1,
	}
}

// Helper for formatting literal values for printing
func StringLiteral(literal any) any {
	if literal == nil {
		return "null"
	}

	return literal
}

// Main scanning function
func (s *Scanner) ScanTokens() []Token {
	tokens := []Token{}

	//TODO: Continue implementing the scanning process for other token types
	for s.current < len(s.source) {
		s.start = s.current
		char := s.advance()

		switch char {
		case '(':
			tokens = append(tokens, Token{
				Type:   LEFT_PAREN,
				Lexeme: string(char),
				Line:   s.line,
			})
		case ')':
			tokens = append(tokens, Token{
				Type:   RIGHT_PAREN,
				Lexeme: string(char),
				Line:   s.line,
			})
		case '{':
			tokens = append(tokens, Token{
				Type:   LEFT_BRACE,
				Lexeme: string(char),
				Line:   s.line,
			})
		case '}':
			tokens = append(tokens, Token{
				Type:   RIGHT_BRACE,
				Lexeme: string(char),
				Line:   s.line,
			})
		case ',':
			tokens = append(tokens, Token{
				Type:   COMMA,
				Lexeme: string(char),
				Line:   s.line,
			})
		case '.':
			if s.current < len(s.source) && s.peek() >= '0' && s.peek() <= '9' {
				lexeme := s.scanNumber()

				literal, _ := strconv.ParseFloat(lexeme, 64)

				tokens = append(tokens, Token{
					Type:    NUMBER,
					Lexeme:  lexeme,
					Literal: literal,
					Line:    s.line,
				})
			} else {
				tokens = append(tokens, Token{
					Type:   DOT,
					Lexeme: string(char),
					Line:   s.line,
				})
			}
		case ';':
			tokens = append(tokens, Token{
				Type:   SEMICOLON,
				Lexeme: string(char),
				Line:   s.line,
			})
		case '+':
			tokens = append(tokens, Token{
				Type:   PLUS,
				Lexeme: string(char),
				Line:   s.line,
			})
		case '-':
			tokens = append(tokens, Token{
				Type:   MINUS,
				Lexeme: string(char),
				Line:   s.line,
			})
		case '*':
			tokens = append(tokens, Token{
				Type:   STAR,
				Lexeme: string(char),
				Line:   s.line,
			})
		case '/':
			tempType := SLASH

			if s.match('>') {
				for {
					if s.peek() == '\x00' {
						s.printError("Unterminated comment block.")
						break // eof check
					}

					if s.peek() == '\n' {
						s.line++ // include newline in commenr
					}

					if s.peek() == '<' {
						s.advance() // consume <

						if s.match('/') {
							break
						}

						continue
					}

					s.advance()
				}

				break
			}

			tokens = append(tokens, Token{
				Type:   tempType,
				Lexeme: string(char),
				Line:   s.line,
			})
		case '%':
			tokens = append(tokens, Token{
				Type:   PERCENT,
				Lexeme: string(char),
				Line:   s.line,
			})
		case '=':
			tempType := EQUAL
			if s.match('=') {
				tempType = EQUAL_EQUAL
			}
			tokens = append(tokens, Token{
				Type:   tempType,
				Lexeme: s.source[s.start:s.current],
				Line:   s.line,
			})
		case '<':
			tempType := LESS
			if s.match('=') {
				tempType = LESS_EQUAL
			}
			tokens = append(tokens, Token{
				Type:   tempType,
				Lexeme: s.source[s.start:s.current],
				Line:   s.line,
			})
		case '>':
			tempType := GREATER
			if s.match('=') {
				tempType = GREATER_EQUAL
			}
			tokens = append(tokens, Token{
				Type:   tempType,
				Lexeme: s.source[s.start:s.current],
				Line:   s.line,
			})
		case '!':
			tempType := NOT
			if s.match('=') {
				tempType = NOT_EQUAL
			}
			tokens = append(tokens, Token{
				Type:   tempType,
				Lexeme: s.source[s.start:s.current],
				Line:   s.line,
			})
		case '&':
			if s.match('&') {
				tokens = append(tokens, Token{
					Type:   AND,
					Lexeme: s.source[s.start:s.current],
					Line:   s.line,
				})
			}
		case '|':
			if s.match('|') {
				tokens = append(tokens, Token{
					Type:   OR,
					Lexeme: s.source[s.start:s.current],
					Line:   s.line,
				})
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			lexeme := s.scanNumber()

			var literal any

			if strings.Contains(lexeme, ".") {
				literal, _ = strconv.ParseFloat(lexeme, 64)
			} else {
				literal, _ = strconv.Atoi(lexeme)
			}

			tokens = append(tokens, Token{
				Type:    NUMBER,
				Lexeme:  lexeme,
				Literal: literal,
				Line:    s.line,
			})
		case '"': // indicates a string
			tokenLine := s.line
			literal, isValidString := s.scanString()
			
			if !isValidString {
				s.printError("Unterminated string.")
			} else {
			tokens = append(tokens, Token{
				Type:    STRING,
				Lexeme:  s.source[s.start:s.current],
				Literal: literal,
				Line:    tokenLine,
			})}
		case ' ', '\t', '\r':
			// ignore whitespaces, tabs, and carriage return	
		case '\n':
			s.line++
		default:
			s.printError(fmt.Sprintf("Unexpected character '%c'", char))
		}

	}

	tokens = append(tokens, Token{
		Type:    EOF,
		Lexeme:  "",
		Literal: nil,
		Line:    s.line,
	})

	return tokens
}

// Helper methods
func (s *Scanner) advance() byte {
	char := s.source[s.current]
	s.current++
	return char
}

func (s *Scanner) peek() byte {
	if s.current >= len(s.source) {
		return '\x00'
	}
	return s.source[s.current]
}

func (s *Scanner) match(expected byte) bool {
	if s.current >= len(s.source) {
		return false
	}

	if s.source[s.current] != expected {
		return false
	}

	s.current++
	return true
}

func (s *Scanner) scanNumber() string {
	for s.peek() >= '0' && s.peek() <= '9' {
		s.advance()
	}

	if s.peek() == '.' && s.current+1 < len(s.source) && s.source[s.current+1] >= '0' && s.source[s.current+1] <= '9' {
		s.advance()

		for s.peek() >= '0' && s.peek() <= '9' {
			s.advance()
		}
	}

	return s.source[s.start:s.current]
}

func (s *Scanner) scanString() (string, bool) {
	for s.peek() != '"' && s.current < len(s.source) {
		if s.peek() == '\\' { // basically checks for backlash
			s.advance() // consumes the backlash

			switch s.peek() {
			case 'n', 't', '"', '\\':
				s.advance()
			default:
				// invalid escape sequence
			}

			continue
		}

		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}

	if s.current >= len(s.source) {
		return "", false
	}
	s.advance() // consume closing 
	return s.source[s.start+1 : s.current-1], true // to not include double quotation marks
}

func (s *Scanner) printError(errMessage string) {
	fmt.Fprintf(os.Stderr, "[Line %d] Error: %s\n", s.line, errMessage)
	s.IsError = true
}
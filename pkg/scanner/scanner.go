package scanner

type Scanner struct {
	source  string
	start   int
	current int
	line    int
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
			tokens = append(tokens, Token{
				Type:   DOT,
				Lexeme: string(char),
				Line:   s.line,
			})
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
			tokens = append(tokens, Token{
				Type:   SLASH,
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

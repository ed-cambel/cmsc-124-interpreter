// defines the types and structure of tokens produced by the scanner

package scanner

type TokenType int

const (
	// Single-character tokens
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	SEMICOLON

	// Arithmetic operators
	PLUS
	MINUS
	STAR
	SLASH
	PERCENT

	// Logical operators
	AND
	OR
	NOT

	// Assignment & Comparison operators
	EQUAL
	LESS
	GREATER
	EQUAL_EQUAL
	NOT_EQUAL
	LESS_EQUAL
	GREATER_EQUAL

	// Identifier & Literals
	IDENTIFIER
	STRING
	NUMBER
	TRUE
	FALSE
	NIL

	// Flow control keywords
	IF
	ELSE

	// TODO: More Witchcraft-themed keywords
	BREW
	SPELL
	CAST

	EOF
)

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal any
	Line    int
}

// method to print the name of TokenType instead of its number
func (t TokenType) String() string {
	switch t {
	case LEFT_PAREN:
		return "LEFT_PAREN"
	case RIGHT_PAREN:
		return "LEFT_PAREN"
	case LEFT_BRACE:
		return "LEFT_BRACE"
	case RIGHT_BRACE:
		return "RIGHT_BRACE"
	case COMMA:
		return "COMMA"
	case DOT:
		return "DOT"
	case SEMICOLON:
		return "SEMICOLON"
	case PLUS:
		return "PLUS"
	case MINUS:
		return "MINUS"
	case STAR:
		return "STAR"
	case SLASH:
		return "SLASH"
	case PERCENT:
		return "PERCENT"
	case EOF:
		return "EOF"
	case EQUAL:
		return "EQUAL"
	case LESS:
		return "LESS"
	case GREATER:
		return "GREATER"
	case NOT:
		return "NOT"
	default:
		return "UNKNOWN"
	}
}
